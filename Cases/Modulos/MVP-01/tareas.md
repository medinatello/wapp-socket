# Sprint 13: Testing y Deployment

## Objetivos
- Crear suite completa de tests (unitarios, integración, E2E)
- Implementar CI/CD pipeline con GitHub Actions
- Configurar Docker y Docker Compose
- Establecer deployment automatizado
- Implementar monitoring y health checks

## Tareas Técnicas

### 1. Test Suite Comprehensiva
```go
// test/integration/api_test.go
func TestAPIIntegration(t *testing.T) {
    // Setup test database
    testDB := setupTestDatabase(t)
    defer cleanupTestDatabase(t, testDB)
    
    // Setup test server
    server := setupTestServer(t, testDB)
    defer server.Close()
    
    // Test user registration flow
    t.Run("User Registration Flow", func(t *testing.T) {
        // Register user
        registerReq := RegisterRequest{
            Username:  "testuser",
            Email:     "test@example.com",
            Password:  "SecurePassword123!",
            FirstName: "Test",
            LastName:  "User",
        }
        
        resp := makeRequest(t, server, "POST", "/auth/register", registerReq)
        assert.Equal(t, http.StatusCreated, resp.Code)
        
        var registerResp UserResponse
        err := json.Unmarshal(resp.Body.Bytes(), &registerResp)
        require.NoError(t, err)
        assert.Equal(t, registerReq.Username, registerResp.Username)
        
        // Login with registered user
        loginReq := LoginRequest{
            Username: registerReq.Username,
            Password: registerReq.Password,
        }
        
        resp = makeRequest(t, server, "POST", "/auth/login", loginReq)
        assert.Equal(t, http.StatusOK, resp.Code)
        
        var loginResp LoginResponse
        err = json.Unmarshal(resp.Body.Bytes(), &loginResp)
        require.NoError(t, err)
        assert.NotEmpty(t, loginResp.AccessToken)
        
        // Test protected endpoint
        req := httptest.NewRequest("GET", "/users/me", nil)
        req.Header.Set("Authorization", "Bearer "+loginResp.AccessToken)
        
        w := httptest.NewRecorder()
        server.Handler.ServeHTTP(w, req)
        
        assert.Equal(t, http.StatusOK, w.Code)
    })
}

// test/load/load_test.go
func TestLoadPerformance(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping load test in short mode")
    }
    
    server := setupTestServer(t, nil)
    defer server.Close()
    
    // Concurrent requests test
    concurrency := 50
    totalRequests := 1000
    requestsPerWorker := totalRequests / concurrency
    
    var wg sync.WaitGroup
    results := make(chan time.Duration, totalRequests)
    
    for i := 0; i < concurrency; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            
            for j := 0; j < requestsPerWorker; j++ {
                start := time.Now()
                resp := makeRequest(t, server, "GET", "/health", nil)
                duration := time.Since(start)
                
                assert.Equal(t, http.StatusOK, resp.Code)
                results <- duration
            }
        }()
    }
    
    wg.Wait()
    close(results)
    
    // Calculate statistics
    var durations []time.Duration
    var total time.Duration
    
    for duration := range results {
        durations = append(durations, duration)
        total += duration
    }
    
    average := total / time.Duration(len(durations))
    
    sort.Slice(durations, func(i, j int) bool {
        return durations[i] < durations[j]
    })
    
    p95 := durations[int(float64(len(durations))*0.95)]
    p99 := durations[int(float64(len(durations))*0.99)]
    
    t.Logf("Load test results:")
    t.Logf("Total requests: %d", len(durations))
    t.Logf("Average response time: %v", average)
    t.Logf("95th percentile: %v", p95)
    t.Logf("99th percentile: %v", p99)
    
    // Assert performance requirements
    assert.Less(t, average, 100*time.Millisecond, "Average response time should be under 100ms")
    assert.Less(t, p95, 200*time.Millisecond, "95th percentile should be under 200ms")
}
```

### 2. Docker Configuration
```dockerfile
# Dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git ca-certificates

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/api/main.go

# Final stage
FROM gcr.io/distroless/static:nonroot

WORKDIR /

# Copy the binary from builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs

# Use nonroot user
USER nonroot:nonroot

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/main", "--health-check"]

EXPOSE 8080

ENTRYPOINT ["/main"]
```

```yaml
# docker-compose.yml
version: '3.8'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - APP_ENVIRONMENT=development
      - APP_DATABASE_HOST=postgres
      - APP_DATABASE_PASSWORD=password
      - APP_CACHE_REDIS_HOST=redis
      - APP_SECURITY_JWT_SECRET=your-secret-key-here
    depends_on:
      postgres:
        condition: service_healthy
      redis:
        condition: service_healthy
    healthcheck:
      test: ["CMD", "/main", "--health-check"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s

  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: go_api_dev
      POSTGRES_USER: dev_user
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U dev_user -d go_api_dev"]
      interval: 10s
      timeout: 5s
      retries: 5

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 3s
      retries: 3

  prometheus:
    image: prom/prometheus:latest
    ports:
      - "9090:9090"
    volumes:
      - ./monitoring/prometheus.yml:/etc/prometheus/prometheus.yml
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'

  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
    volumes:
      - grafana_data:/var/lib/grafana
      - ./monitoring/grafana/provisioning:/etc/grafana/provisioning

volumes:
  postgres_data:
  redis_data:
  grafana_data:
```

### 3. CI/CD Pipeline
```yaml
# .github/workflows/ci.yml
name: CI/CD Pipeline

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

env:
  GO_VERSION: '1.21'
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: postgres
          POSTGRES_DB: test_db
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432
      
      redis:
        image: redis:7
        options: >-
          --health-cmd "redis-cli ping"
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 6379:6379

    steps:
    - uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: ${{ env.GO_VERSION }}

    - name: Cache Go modules
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
        restore-keys: |
          ${{ runner.os }}-go-

    - name: Install dependencies
      run: go mod download

    - name: Run linter
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest

    - name: Generate Swagger docs
      run: |
        go install github.com/swaggo/swag/cmd/swag@latest
        swag init -g cmd/api/main.go -o docs/

    - name: Run unit tests
      run: go test -v -race -coverprofile=coverage.out ./...

    - name: Run integration tests
      env:
        APP_DATABASE_HOST: localhost
        APP_DATABASE_PASSWORD: postgres
        APP_CACHE_REDIS_HOST: localhost
        APP_SECURITY_JWT_SECRET: test-secret-key
      run: go test -v -tags integration ./test/integration/...

    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out

  security:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4

    - name: Run Gosec Security Scanner
      uses: securecodewarrior/github-action-gosec@master
      with:
        args: './...'

    - name: Run Trivy vulnerability scanner
      uses: aquasecurity/trivy-action@master
      with:
        scan-type: 'fs'
        scan-ref: '.'

  build:
    needs: [test, security]
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v4

    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v3

    - name: Log in to Container Registry
      uses: docker/login-action@v3
      with:
        registry: ${{ env.REGISTRY }}
        username: ${{ github.actor }}
        password: ${{ secrets.GITHUB_TOKEN }}

    - name: Extract metadata
      id: meta
      uses: docker/metadata-action@v5
      with:
        images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
        tags: |
          type=ref,event=branch
          type=ref,event=pr
          type=sha,prefix={{branch}}-
          type=raw,value=latest,enable={{is_default_branch}}

    - name: Build and push Docker image
      uses: docker/build-push-action@v5
      with:
        context: .
        push: true
        tags: ${{ steps.meta.outputs.tags }}
        labels: ${{ steps.meta.outputs.labels }}
        cache-from: type=gha
        cache-to: type=gha,mode=max

  deploy:
    needs: build
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    environment: production
    
    steps:
    - uses: actions/checkout@v4

    - name: Deploy to production
      run: |
        echo "Deploying to production..."
        # Add your deployment logic here
        # Examples: kubectl apply, helm upgrade, etc.
```

### 4. Monitoring Configuration
```yaml
# monitoring/prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'go-api'
    static_configs:
      - targets: ['api:8080']
    metrics_path: '/metrics'
    scrape_interval: 5s

  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']

rule_files:
  - "alert_rules.yml"

alerting:
  alertmanagers:
    - static_configs:
        - targets:
          - alertmanager:9093
```

### 5. Health Check Enhancement
```go
// cmd/api/health_check.go
func runHealthCheck() error {
    config, err := loadConfig()
    if err != nil {
        return fmt.Errorf("failed to load config: %w", err)
    }
    
    // Check database
    if err := checkDatabase(config.Database); err != nil {
        return fmt.Errorf("database health check failed: %w", err)
    }
    
    // Check Redis
    if err := checkRedis(config.Cache.Redis); err != nil {
        return fmt.Errorf("redis health check failed: %w", err)
    }
    
    // Check external dependencies
    if err := checkExternalServices(); err != nil {
        return fmt.Errorf("external services health check failed: %w", err)
    }
    
    fmt.Println("All health checks passed")
    return nil
}

func main() {
    if len(os.Args) > 1 && os.Args[1] == "--health-check" {
        if err := runHealthCheck(); err != nil {
            fmt.Fprintf(os.Stderr, "Health check failed: %v\n", err)
            os.Exit(1)
        }
        os.Exit(0)
    }
    
    // Normal application startup
    startApplication()
}
```

### 6. Production Configuration
```yaml
# configs/config.production.yaml
app:
  name: "go-api-template"
  environment: "production"
  debug: false

server:
  port: 8080
  host: "0.0.0.0"
  read_timeout: "30s"
  write_timeout: "30s"
  idle_timeout: "120s"

database:
  host: "${DATABASE_HOST}"
  port: 5432
  username: "${DATABASE_USERNAME}"
  password: "${DATABASE_PASSWORD}"
  database: "${DATABASE_NAME}"
  ssl_mode: "require"
  max_open_conns: 50
  max_idle_conns: 25
  max_lifetime: "1h"

logging:
  level: "info"
  format: "json"
  output: "stdout"

metrics:
  enabled: true
  port: 9090
  path: "/metrics"

security:
  jwt_secret: "${JWT_SECRET}"
  jwt_expiration: "4h"
  rate_limit:
    enabled: true
    requests: 100
    window: "1m"
```

### 7. Makefile for Development
```makefile
# Makefile
.PHONY: build test run docker-build docker-run clean deps lint

# Variables
APP_NAME=go-api-template
VERSION?=latest
DOCKER_IMAGE=$(APP_NAME):$(VERSION)

# Build the application
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$(APP_NAME) cmd/api/main.go

# Run tests
test:
	go test -v -race -coverprofile=coverage.out ./...

# Run integration tests
test-integration:
	go test -v -tags integration ./test/integration/...

# Run the application
run:
	go run cmd/api/main.go

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run linter
lint:
	golangci-lint run

# Generate swagger docs
docs:
	swag init -g cmd/api/main.go -o docs/

# Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run with Docker Compose
docker-run:
	docker-compose up -d

# Stop Docker Compose
docker-stop:
	docker-compose down

# Clean build artifacts
clean:
	rm -rf bin/ docs/swagger.json coverage.out

# Setup development environment
setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	cp .env.example .env

# Deploy to staging
deploy-staging:
	echo "Deploying to staging..."

# Deploy to production
deploy-production:
	echo "Deploying to production..."

# Health check
health-check:
	./bin/$(APP_NAME) --health-check
```

## Criterios de Aceptación ✅

- [ ] Suite completa de tests (unit, integration, E2E)
- [ ] CI/CD pipeline con GitHub Actions
- [ ] Docker containerization con multi-stage builds
- [ ] Docker Compose para desarrollo local
- [ ] Health checks comprehensivos
- [ ] Monitoring con Prometheus y Grafana
- [ ] Security scanning en pipeline
- [ ] Performance testing automatizado
- [ ] Production-ready configuration
- [ ] Automated deployment pipeline

## Conclusion

Sprint 13 completa la plantilla con:
- ✅ Testing framework robusto con alta cobertura
- ✅ CI/CD pipeline automatizado
- ✅ Containerization optimizada para producción
- ✅ Monitoring y observabilidad completa
- ✅ Security scanning y hardening
- ✅ Performance testing y optimization
- ✅ Production deployment readiness

La plantilla está ahora lista para uso en producción con todas las mejores prácticas de DevOps implementadas.