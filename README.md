# wapp-socket

Un cliente de WhatsApp implementado en Go siguiendo arquitectura limpia y principios de Clean Architecture.

## 🚀 Características

- **Arquitectura Hexagonal**: Separación clara entre dominio, aplicación e infraestructura
- **Adaptadores Fake**: Desarrollo y testing sin dependencias externas
- **CLI Interactiva**: Comandos para conectar, enviar mensajes y escuchar eventos
- **Daemon HTTP**: Servicio en background con health checks
- **Logging Estructurado**: Usando log/slog nativo de Go
- **Testing Completo**: 100% cobertura en dominio, 80%+ en aplicación

## 📦 Instalación

### Prerrequisitos
- Go 1.22 o superior

### Compilación
```bash
# Clonar repositorio
git clone https://github.com/medinatello/wapp-socket.git
cd wapp-socket

# Instalar dependencias
make deps

# Compilar binarios
make build
```

## 🎯 Uso Rápido

### CLI Interactiva
```bash
# Conectar al servicio
./bin/wapp-socket-cli connect

# Enviar mensaje
./bin/wapp-socket-cli send --to "1234567890@s.whatsapp.net" --text "¡Hola mundo!"

# Escuchar eventos (10 segundos)
./bin/wapp-socket-cli stream --duration 10s

# Listar grupos
./bin/wapp-socket-cli groups list
```

### Daemon de Servicio
```bash
# Iniciar daemon
./bin/wapp-socket-daemon

# Health check
curl http://localhost:8081/health
```

## 🏗️ Arquitectura

```
wapp-socket/
├── cmd/                    # Puntos de entrada
│   ├── whats-cli/         # CLI interactiva
│   └── whatsd/            # Daemon de servicio
├── internal/              # Código privado
│   ├── domain/           # Entidades de negocio
│   ├── usecase/          # Casos de uso
│   ├── port/             # Interfaces (puertos)
│   └── adapter/          # Implementaciones
└── interface/            # Interfaces externas
    ├── cli/             # Comandos CLI
    └── http/            # Servidor HTTP
```

### Principios de Diseño

1. **Clean Architecture**: Dependencias apuntan hacia adentro
2. **Dependency Inversion**: Interfaces definidas en capas internas
3. **Single Responsibility**: Cada componente tiene una responsabilidad
4. **Testabilidad**: Adaptadores fake para desarrollo y testing

## 🧪 Testing

```bash
# Ejecutar todos los tests
make test

# Tests con reporte de cobertura
make test-coverage

# Solo tests del dominio
make test-domain

# Verificar estándares del proyecto
make check-standards
```

### Métricas de Calidad
- **Dominio**: 100% cobertura (obligatorio)
- **Aplicación**: 89.5% cobertura (objetivo: 80%+)
- **Logger**: 100% cobertura (objetivo: 70%+)

## 🔧 Desarrollo

### Comandos Make Disponibles
```bash
make help           # Mostrar ayuda
make all            # Build completo (deps, fmt, lint, test, build)
make build          # Compilar binarios
make test           # Ejecutar tests
make test-coverage  # Tests con cobertura
make lint           # Análisis estático
make fmt            # Formatear código
make clean          # Limpiar artefactos
make demo           # Demo funcional
```

### Configuración

El proyecto usa Viper para configuración flexible:

```yaml
# configs/config.yaml
app:
  name: "wapp-socket"
  log_level: "info"

fakes:
  seed: 0                      # 0 = random, otro = reproducible
  connect_timeout_chance: 0.2   # 20% probabilidad de timeout
  connect_fail_chance: 0.1      # 10% probabilidad de fallo
  receive_interval_ms: 1000     # Intervalo entre eventos fake

features:
  fake_mode: true              # Usar adaptadores fake
  eventing_model: "bus"        # Modelo de eventos
```

Variables de entorno (prefijo `WAPP_`):
```bash
export WAPP_APP_LOG_LEVEL=debug
export WAPP_FEATURES_FAKE_MODE=false
```

## 🔌 Adaptadores

### Implementados (Fake)
- **WebSocket**: Simula conexiones con eventos aleatorios
- **SessionStore**: Almacenamiento en memoria
- **Logger**: Wrapper sobre log/slog

### Futuros (Reales)
- **WebSocket**: nhooyr/websocket
- **SessionStore**: SQLite con modernc.org/sqlite
- **Crypto**: Noise Protocol
- **ProtoCodec**: google.golang.org/protobuf

## 📊 Estado del Proyecto

| MVP | Estado | Descripción |
|-----|--------|-------------|
| MVP-01 | ✅ Completo | Arquitectura base y adaptadores fake |
| MVP-02 | 📋 Planificado | Adaptadores reales (WebSocket) |
| MVP-03 | 📋 Planificado | Persistencia (SQLite) |
| MVP-04 | 📋 Planificado | Criptografía (Noise Protocol) |

## 🤝 Contribuir

1. Fork el proyecto
2. Crear rama feature (`git checkout -b feature/nueva-funcionalidad`)
3. Commit cambios (`git commit -am 'Agregar nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Crear Pull Request

### Estándares de Contribución
- Seguir [Cases/Documentation Main/MEJORES_PRACTICAS.md](Cases/Documentation%20Main/MEJORES_PRACTICAS.md)
- Mantener 100% cobertura en dominio
- Documentar funciones exportadas
- Tests para nueva funcionalidad

## 📝 Documentación

- [Decisiones Tecnológicas](Documentation/01-tecnologias.md)
- [Interfaces y Puertos](Documentation/02-interfaces.md)
- [Flujos del Sprint 1](Documentation/10-flujos-sprint1.md)
- [Mejores Prácticas](Cases/Documentation%20Main/MEJORES_PRACTICAS.md)

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver `LICENSE` para más detalles.

## 🔍 Arquitectura en Detalle

### Dominio del Negocio
```go
// JID - Identificador de usuario WhatsApp
type JID string

// Message - Mensaje básico
type Message struct {
    ID        string
    From      JID
    To        JID
    Text      string
    Timestamp time.Time
}

// Event - Evento del sistema
type Event struct {
    Type    string // message_received, presence_update, etc.
    Payload any
}
```

### Casos de Uso
- **ConnectUseCase**: Gestiona conexiones WebSocket
- **SendMessageUseCase**: Envío de mensajes
- **ReceiveUseCase**: Procesamiento de eventos
- **GroupsUseCase**: Gestión de grupos

### Inyección de Dependencias
```go
type Container struct {
    Config    *Config
    Logger    outbound.Logger
    
    // Casos de uso
    ConnectUseCase     *usecase.ConnectUseCase
    SendMessageUseCase *usecase.SendMessageUseCase
    ReceiveUseCase     *usecase.ReceiveUseCase
    GroupsUseCase      *usecase.GroupsUseCase
}
```

---

**Desarrollado con ❤️ usando Go y Clean Architecture**