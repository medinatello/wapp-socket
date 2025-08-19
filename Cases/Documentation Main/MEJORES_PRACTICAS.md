# Mejores Prácticas para Desarrollo en Go

Esta guía contiene las mejores prácticas para el desarrollo con Go, siguiendo los estándares de Google y la comunidad.

## 📋 Tabla de Contenidos

1. [Estructura de Proyecto](#estructura-de-proyecto)
2. [Convenciones de Nombres](#convenciones-de-nombres)
3. [Clean Architecture](#clean-architecture)
4. [Manejo de Errores](#manejo-de-errores)
5. [Testing](#testing)
6. [Logging](#logging)
7. [Configuración](#configuración)
8. [Seguridad](#seguridad)
9. [Performance](#performance)
10. [Documentación](#documentación)

## 🏗️ Estructura de Proyecto

### Organización de Directorios

```
project/
├── cmd/                    # Aplicaciones principales
│   └── api/               # Servidor API REST
│       └── main.go        # Punto de entrada
├── internal/              # Código privado de la aplicación
│   ├── app/              # Servicios de aplicación
│   │   ├── services/     # Lógica de negocio
│   │   └── usecases/     # Casos de uso
│   ├── domain/           # Entidades y reglas de negocio
│   │   ├── entities/     # Entidades del dominio
│   │   ├── repositories/ # Interfaces de repositorios
│   │   └── services/     # Servicios del dominio
│   ├── infrastructure/   # Dependencias externas
│   │   ├── database/     # Implementaciones de DB
│   │   ├── cache/        # Implementaciones de cache
│   │   └── http/         # Clientes HTTP
│   └── interfaces/       # Adaptadores de interfaz
│       ├── handlers/     # Handlers HTTP
│       ├── middleware/   # Middleware HTTP
│       └── routes/       # Definición de rutas
├── pkg/                  # Código público reutilizable
│   ├── logger/           # Utilidades de logging
│   ├── validator/        # Validadores personalizados
│   └── utils/            # Utilidades generales
├── api/                  # Especificaciones OpenAPI
├── configs/              # Archivos de configuración
├── docs/                 # Documentación
├── scripts/              # Scripts de utilidad
└── test/                 # Tests adicionales
```

### Principios de Organización

1. **Separación de responsabilidades**: Cada paquete tiene una responsabilidad específica
2. **Dependencias hacia adentro**: Los paquetes internos no deben depender de externos
3. **Interfaces en el dominio**: Define interfaces donde las usas, no donde las implementas
4. **Código público en pkg/**: Solo código reutilizable entre proyectos

## 🏷️ Convenciones de Nombres

### Variables y Funciones

```go
// ✅ Correcto - camelCase
var userID int
var httpClient *http.Client

// ❌ Incorrecto - snake_case
var user_id int
var http_client *http.Client

// ✅ Correcto - nombres descriptivos
func CreateUserAccount(user *User) error

// ❌ Incorrecto - nombres no descriptivos
func Create(u *User) error
```

### Constantes

```go
// ✅ Correcto - PascalCase para exportadas
const MaxRetryAttempts = 3
const DefaultTimeout = 30 * time.Second

// ✅ Correcto - camelCase para privadas
const maxRetryAttempts = 3
const defaultTimeout = 30 * time.Second
```

### Tipos y Estructuras

```go
// ✅ Correcto - PascalCase para exportados
type UserService struct {
    repository UserRepository
    logger     *zap.Logger
}

// ✅ Correcto - Interfaces terminan en -er cuando es posible
type UserRepository interface {
    Save(user *User) error
    FindByID(id int) (*User, error)
}
```

### Archivos y Paquetes

```go
// ✅ Correcto - nombres de paquetes en minúsculas
package userservice

// ✅ Correcto - nombres de archivos descriptivos
// user_service.go
// user_repository.go
// user_handler.go
```

## 🏛️ Clean Architecture

### Capas de la Arquitectura

#### 1. Domain Layer (internal/domain/)

```go
// entities/user.go
package entities

import "time"

type User struct {
    ID        int       `json:"id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Validate() error {
    if u.Email == "" {
        return errors.New("email is required")
    }
    return nil
}
```

```go
// repositories/user_repository.go
package repositories

import "github.com/your-project/internal/domain/entities"

type UserRepository interface {
    Save(user *entities.User) error
    FindByID(id int) (*entities.User, error)
    FindByEmail(email string) (*entities.User, error)
    Delete(id int) error
}
```

#### 2. Application Layer (internal/app/)

```go
// services/user_service.go
package services

import (
    "github.com/your-project/internal/domain/entities"
    "github.com/your-project/internal/domain/repositories"
)

type UserService struct {
    userRepo repositories.UserRepository
    logger   Logger
}

func NewUserService(userRepo repositories.UserRepository, logger Logger) *UserService {
    return &UserService{
        userRepo: userRepo,
        logger:   logger,
    }
}

func (s *UserService) CreateUser(name, email string) (*entities.User, error) {
    user := &entities.User{
        Name:  name,
        Email: email,
    }
    
    if err := user.Validate(); err != nil {
        return nil, err
    }
    
    return s.userRepo.Save(user)
}
```

#### 3. Infrastructure Layer (internal/infrastructure/)

```go
// database/user_repository.go
package database

import (
    "gorm.io/gorm"
    "github.com/your-project/internal/domain/entities"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Save(user *entities.User) error {
    return r.db.Save(user).Error
}

func (r *userRepository) FindByID(id int) (*entities.User, error) {
    var user entities.User
    err := r.db.First(&user, id).Error
    return &user, err
}
```

#### 4. Interface Layer (internal/interfaces/)

```go
// handlers/user_handler.go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/your-project/internal/app/services"
)

type UserHandler struct {
    userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    user, err := h.userService.CreateUser(req.Name, req.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, user)
}
```

## ⚠️ Manejo de Errores

### Creación de Errores Personalizados

```go
// ✅ Correcto - Errores con contexto
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field %s: %s", e.Field, e.Message)
}

// ✅ Correcto - Wrapping de errores
func ProcessUser(user *User) error {
    if err := user.Validate(); err != nil {
        return fmt.Errorf("failed to validate user: %w", err)
    }
    
    if err := saveUser(user); err != nil {
        return fmt.Errorf("failed to save user: %w", err)
    }
    
    return nil
}
```

### Manejo en Handlers

```go
// ✅ Correcto - Manejo de errores con códigos HTTP apropiados
func (h *UserHandler) GetUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{
            Code:    "INVALID_ID",
            Message: "Invalid user ID format",
        })
        return
    }
    
    user, err := h.userService.GetUser(id)
    if err != nil {
        if errors.Is(err, ErrUserNotFound) {
            c.JSON(http.StatusNotFound, ErrorResponse{
                Code:    "USER_NOT_FOUND",
                Message: "User not found",
            })
            return
        }
        
        h.logger.Error("Failed to get user", zap.Error(err), zap.Int("user_id", id))
        c.JSON(http.StatusInternalServerError, ErrorResponse{
            Code:    "INTERNAL_ERROR",
            Message: "Internal server error",
        })
        return
    }
    
    c.JSON(http.StatusOK, user)
}
```

## 🧪 Testing

### Cobertura Mínima Requerida

**Cobertura objetivo**:
- Paquetes del dominio (`internal/domain`): **100%**
- Servicios de aplicación (`internal/app`): **80%**
- Adaptadores críticos (`internal/adapter`): **70%**
- Casos de uso (`internal/usecase`): **80%**

### Tests Unitarios

```go
// user_service_test.go
package services_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) Save(user *entities.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func TestUserService_CreateUser(t *testing.T) {
    // Arrange
    mockRepo := new(MockUserRepository)
    mockLogger := new(MockLogger)
    service := NewUserService(mockRepo, mockLogger)
    
    user := &entities.User{Name: "John", Email: "john@example.com"}
    mockRepo.On("Save", mock.AnythingOfType("*entities.User")).Return(nil)
    
    // Act
    result, err := service.CreateUser("John", "john@example.com")
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "John", result.Name)
    mockRepo.AssertExpectations(t)
}
```

### Tests de Integración

```go
// user_handler_integration_test.go
func TestUserHandler_Integration(t *testing.T) {
    // Setup
    db := setupTestDB(t)
    defer cleanupTestDB(t, db)
    
    router := setupTestRouter(db)
    
    // Test
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/users", strings.NewReader(`{
        "name": "John Doe",
        "email": "john@example.com"
    }`))
    req.Header.Set("Content-Type", "application/json")
    
    router.ServeHTTP(w, req)
    
    // Assert
    assert.Equal(t, http.StatusCreated, w.Code)
    
    var response map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.Equal(t, "John Doe", response["name"])
}
```

## 📝 Logging

### Configuración de Zap

```go
// pkg/logger/logger.go
package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func NewLogger(level string, env string) (*zap.Logger, error) {
    var config zap.Config
    
    if env == "production" {
        config = zap.NewProductionConfig()
    } else {
        config = zap.NewDevelopmentConfig()
    }
    
    // Set log level
    switch level {
    case "debug":
        config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
    case "info":
        config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    case "warn":
        config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
    case "error":
        config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
    default:
        config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    }
    
    return config.Build()
}
```

### Uso de Logs Estructurados

```go
// ✅ Correcto - Logs estructurados
logger.Info("User created successfully",
    zap.Int("user_id", user.ID),
    zap.String("email", user.Email),
    zap.Duration("processing_time", time.Since(start)),
)

// ✅ Correcto - Logs de error con contexto
logger.Error("Failed to save user to database",
    zap.Error(err),
    zap.String("email", user.Email),
    zap.String("operation", "create_user"),
)

// ❌ Incorrecto - Logs no estructurados
logger.Info(fmt.Sprintf("User %d created with email %s", user.ID, user.Email))
```

## ⚙️ Configuración

### Uso de Viper

```go
// configs/config.go
package configs

import (
    "github.com/spf13/viper"
    "time"
)

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    JWT      JWTConfig      `mapstructure:"jwt"`
    Logger   LoggerConfig   `mapstructure:"logger"`
}

type ServerConfig struct {
    Port         int           `mapstructure:"port"`
    ReadTimeout  time.Duration `mapstructure:"read_timeout"`
    WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

func LoadConfig(path string) (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(path)
    viper.AddConfigPath(".")
    
    // Environment variables
    viper.AutomaticEnv()
    viper.SetEnvPrefix("APP")
    
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }
    
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }
    
    return &config, nil
}
```

## 🔒 Seguridad

### Validación de Input

```go
// ✅ Correcto - Validación con struct tags
type CreateUserRequest struct {
    Name     string `json:"name" validate:"required,min=2,max=50"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}

// ✅ Correcto - Sanitización de input
func sanitizeString(input string) string {
    return html.EscapeString(strings.TrimSpace(input))
}
```

### Manejo de Secrets

```go
// ✅ Correcto - No hardcodear secrets
type Config struct {
    JWTSecret string `mapstructure:"jwt_secret" validate:"required"`
    DBPassword string `mapstructure:"db_password" validate:"required"`
}

// ❌ Incorrecto - Secrets hardcodeados
const JWTSecret = "my-secret-key" // NUNCA hacer esto
```

## 🚀 Performance

### Pool de Conexiones

```go
// ✅ Correcto - Pool de conexiones configurado
func setupDB(config DatabaseConfig) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }
    
    sqlDB.SetMaxIdleConns(config.MaxIdleConns)
    sqlDB.SetMaxOpenConns(config.MaxOpenConns)
    sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
    
    return db, nil
}
```

### Context y Timeouts

```go
// ✅ Correcto - Uso de context con timeout
func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    return s.userRepo.FindByID(ctx, id)
}
```

## 📖 Documentación

### Comentarios en Go

```go
// ✅ Correcto - Comentarios para funciones exportadas
// CreateUser creates a new user with the provided name and email.
// It validates the input and returns an error if validation fails.
func CreateUser(name, email string) (*User, error) {
    // Implementation...
}

// ✅ Correcto - Comentarios para tipos exportados
// UserService handles all user-related business logic.
// It coordinates between the domain layer and infrastructure layer.
type UserService struct {
    userRepo repositories.UserRepository
    logger   *zap.Logger
}
```

### Documentación API - Filosofía Nativa First

#### Principio Fundamental: Priorizar Soluciones Nativas

**REGLA DE ORO**: Siempre preferir implementaciones nativas de Go sobre dependencias externas (Java, Node.js, Python, etc.)

#### OpenAPI 3.0 - Estrategia Nativa

**✅ RECOMENDADO: kin-openapi (100% Go nativo)**

```go
// Generación nativa OpenAPI 3.0 sin dependencias externas
import (
    "github.com/getkin/kin-openapi/openapi3"
)

// Ejemplo de generación programática de especificación
func generateOpenAPISpec() *openapi3.T {
    doc := &openapi3.T{
        OpenAPI: "3.0.1",
        Info: &openapi3.Info{
            Title:       "API Nativa Go",
            Description: "API documentada con kin-openapi nativo",
            Version:     "1.0.0",
        },
        Paths: openapi3.Paths{},
    }
    
    // Agregar endpoints programáticamente
    doc.Paths["/users"] = &openapi3.PathItem{
        Post: &openapi3.Operation{
            Summary:     "Crear usuario",
            Description: "Crea un nuevo usuario en el sistema",
            Tags:        []string{"usuarios"},
            RequestBody: &openapi3.RequestBodyRef{
                Value: &openapi3.RequestBody{
                    Content: openapi3.Content{
                        "application/json": &openapi3.MediaType{
                            Schema: openapi3.NewSchemaRef("",
                                openapi3.NewObjectSchema().
                                    WithProperty("name", openapi3.NewStringSchema()).
                                    WithProperty("email", openapi3.NewStringSchema()),
                            ),
                        },
                    },
                },
            },
            Responses: openapi3.Responses{
                "201": &openapi3.ResponseRef{
                    Value: &openapi3.Response{
                        Description: stringPtr("Usuario creado exitosamente"),
                        Content: openapi3.Content{
                            "application/json": &openapi3.MediaType{
                                Schema: openapi3.NewSchemaRef("#/components/schemas/User", nil),
                            },
                        },
                    },
                },
            },
        },
    }
    
    return doc
}

// Integración con Swagger UI (embebido)
//go:embed swagger-ui/*
var swaggerUI embed.FS

func setupSwaggerUI(router *gin.Engine, spec *openapi3.T) {
    // Servir especificación OpenAPI
    router.GET("/api/openapi.json", func(c *gin.Context) {
        specJSON, _ := json.Marshal(spec)
        c.Data(200, "application/json", specJSON)
    })
    
    // Servir Swagger UI embebido
    router.StaticFS("/swagger", http.FS(swaggerUI))
}
```

**❌ EVITAR: Dependencias externas pesadas**

```go
// ❌ Dependencias de Java (OpenAPI Generator)
// ❌ Dependencias de Node.js (swagger-ui-dist npm)
// ❌ Herramientas que requieren runtime externos
```

#### Comparación de Enfoques

| Enfoque | Nativo Go | Dependencias | Swagger UI | OpenAPI 3.0 | Compilación |
|---------|-----------|--------------|------------|-------------|-------------|
| **kin-openapi** | ✅ 100% | ✅ Solo Go | ✅ Embebido | ✅ Nativo | ✅ Auto |
| **swaggo/swag** | ⚠️ 80% | ⚠️ Herramienta | ✅ Incluido | ⚠️ Conversión | ✅ Manual |
| **oapi-codegen** | ✅ 90% | ✅ Solo Go | ❌ Separado | ✅ Nativo | ⚠️ Separado |
| **OpenAPI Generator** | ❌ 0% | ❌ Java + deps | ✅ Incluido | ✅ Nativo | ❌ Complejo |

#### Implementación Recomendada

```go
// pkg/openapi/generator.go - Generador nativo
package openapi

import (
    "encoding/json"
    "github.com/getkin/kin-openapi/openapi3"
    "github.com/gin-gonic/gin"
)

type APIDocumentator struct {
    spec *openapi3.T
}

func NewAPIDocumentator(info *openapi3.Info) *APIDocumentator {
    return &APIDocumentator{
        spec: &openapi3.T{
            OpenAPI: "3.0.1",
            Info:    info,
            Paths:   make(openapi3.Paths),
            Components: &openapi3.Components{
                Schemas: make(openapi3.Schemas),
            },
        },
    }
}

// RegisterEndpoint registra un endpoint con documentación
func (d *APIDocumentator) RegisterEndpoint(method, path string, operation *openapi3.Operation) {
    if d.spec.Paths[path] == nil {
        d.spec.Paths[path] = &openapi3.PathItem{}
    }
    
    switch method {
    case "GET":
        d.spec.Paths[path].Get = operation
    case "POST":
        d.spec.Paths[path].Post = operation
    // ... otros métodos
    }
}

// RegisterSchema registra un schema de componente
func (d *APIDocumentator) RegisterSchema(name string, schema *openapi3.Schema) {
    d.spec.Components.Schemas[name] = openapi3.NewSchemaRef("", schema)
}

// SetupRoutes configura las rutas de documentación
func (d *APIDocumentator) SetupRoutes(router *gin.Engine) {
    router.GET("/api/openapi.json", d.serveSpec)
    router.StaticFS("/swagger/", http.FS(swaggerUIFiles))
}

func (d *APIDocumentator) serveSpec(c *gin.Context) {
    c.JSON(200, d.spec)
}
```

#### Ventajas del Enfoque Nativo

1. **🔒 Sin dependencias externas**: Solo Go stdlib + kin-openapi
2. **⚡ Compilación limpia**: Todo embebido en el binario
3. **🔄 Sincronización automática**: Spec generado desde código
4. **🚀 Despliegue simple**: Un solo binario sin runtime externos
5. **🛡️ Mantenimiento reducido**: Sin actualizaciones de herramientas Java/Node
6. **📦 Distribución simple**: Sin instalación de herramientas adicionales

## 🛠️ Herramientas Recomendadas

### Linting y Formatting

```bash
# golangci-lint - Linter completo
golangci-lint run

# gofmt - Formateo estándar
gofmt -w .

# goimports - Organización de imports
goimports -w .
```

### Testing

```bash
# Tests con cobertura
go test -cover ./...

# Tests con reporte HTML
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Profiling

```go
// ✅ Correcto - Profiling en desarrollo
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Rest of application...
}
```

## 🚀 Integración Continua y Calidad de Código

Para asegurar la calidad del proyecto y minimizar regresiones, se deben configurar y seguir las siguientes prácticas:

- **Pipeline de Integración Continua (CI)**: Configurar un pipeline (por ejemplo, GitHub Actions o Jenkins) que ejecute pruebas unitarias, pruebas de integración, análisis estático de código y builds de forma automática en cada pull request.
- **Cobertura de Pruebas**: Mantener una cobertura mínima del 80 % en los paquetes de dominio y lógica de negocio utilizando `go test -cover`.
- **Análisis Estático y Formateo**: Usar herramientas como `go vet`, `golangci-lint` y `golines` para detectar errores y mantener un estilo consistente. Estos deben ejecutarse como parte del pipeline.
- **Revisión de Dependencias**: Automatizar la actualización y verificación de dependencias con herramientas como `dependabot` o `renovate`.
- **Reportes de Calidad**: Generar y publicar reportes de pruebas y cobertura como artefactos del pipeline para facilitar su revisión.

## 🔍 Code Review Checklist

- [ ] ¿Sigue las convenciones de nombres de Go?
- [ ] ¿Los errores se manejan apropiadamente?
- [ ] ¿Se usan interfaces donde es apropiado?
- [ ] ¿El código es testeable?
- [ ] ¿Los tests tienen buena cobertura?
- [ ] ¿La documentación está actualizada?
- [ ] ¿No hay secrets hardcodeados?
- [ ] ¿Se usan contexts para cancelación/timeout?
- [ ] ¿Los logs son estructurados y útiles?
- [ ] ¿Sigue los principios de Clean Architecture?

## 📊 Estándares Específicos para wapp-socket

### Patrones de Interfaces (Puertos y Adaptadores)

```go
// ✅ Correcto - Interfaces en el paquete outbound
package outbound

// Logger define interfaz estándar para logging estructurado
type Logger interface {
    Debug(msg string, args ...any)
    Info(msg string, args ...any)
    Warn(msg string, args ...any)
    Error(msg string, err error, args ...any)
}

// WebSocketDialer define interfaz para conexiones WebSocket
type WebSocketDialer interface {
    Dial(ctx context.Context, url string) (WebSocketConn, error)
}
```

### Adaptadores Fake para Testing

**Obligatorio**: Todos los adaptadores externos deben tener implementaciones fake para testing:

```go
// ✅ Correcto - Adaptador fake con configuración
type FakeWebSocketDialer struct {
    logger            outbound.Logger
    seed              int64
    timeoutChance     float64
    failChance        float64
}

func NewFakeWebSocketDialer(logger outbound.Logger, seed int64, timeoutChance, failChance float64, interval int) outbound.WebSocketDialer {
    return &FakeWebSocketDialer{
        logger:            logger,
        seed:              seed,
        timeoutChance:     timeoutChance,
        failChance:        failChance,
    }
}
```

### Inyección de Dependencias

**Patrón requerido**: Usar el patrón Container para inyección de dependencias:

```go
// ✅ Correcto - Container con todas las dependencias
type Container struct {
    Config    *Config
    Logger    outbound.Logger
    Telemetry telemetry.Telemetry
    
    // Use Cases
    ConnectUseCase     *usecase.ConnectUseCase
    SendMessageUseCase *usecase.SendMessageUseCase
}

func NewContainer() (*Container, error) {
    // Inicialización de dependencias...
    return container, nil
}
```

### Métricas de Calidad Requeridas

**Cobertura de tests mínima**:
- **Dominio** (`internal/domain`): 100%
- **Aplicación** (`internal/app`): 80%
- **Casos de uso** (`internal/usecase`): 80%
- **Adaptadores** (`internal/adapter`): 70%

**Comandos de verificación**:
```bash
# Verificar cobertura por paquete
go test -cover ./internal/domain
go test -cover ./internal/app
go test -cover ./internal/usecase

# Verificar compilación sin errores
go build ./...

# Ejecutar linting
golangci-lint run
```

---

Esta guía debe ser un documento vivo que se actualiza conforme el proyecto evoluciona y se aprenden nuevas mejores prácticas.