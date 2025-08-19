# MVP-01 — Arquitectura Base y Esqueleto del Proyecto

Este módulo corresponde al Sprint 1 del proyecto wapp-socket y tiene como objetivo establecer la arquitectura base y el esqueleto funcional del sistema.

## Objetivos

- Implementar la arquitectura hexagonal con puertos y adaptadores
- Crear adaptadores fake para todas las dependencias externas (WebSocket, Storage, Crypto, etc.)
- Establecer el patrón de inyección de dependencias con Container
- Implementar la CLI básica con comandos principales (connect, send, stream, groups)
- Configurar el sistema de logging estructurado con slog
- Crear la base de telemetría con implementación no-op

## Versiones y Dependencias Clave

- Go 1.22
- github.com/spf13/cobra v1.8.1 (CLI framework)
- github.com/spf13/viper v1.20.1 (Configuración)
- log/slog (Logging estructurado nativo)

## Contexto del Proyecto

wapp-socket es un cliente de WhatsApp implementado en Go que simula la funcionalidad de comunicación mediante WebSocket. El proyecto utiliza arquitectura limpia con implementaciones fake para el desarrollo inicial.

### Estructura del Proyecto

```
wapp-socket/
├── cmd/
│   ├── whats-cli/     # CLI para interacción manual
│   └── whatsd/        # Daemon para servicio en background
├── internal/
│   ├── domain/        # Entidades del dominio (Message, Session, JID)
│   ├── app/           # Configuración y Container DI
│   ├── usecase/       # Casos de uso (Connect, SendMessage, Receive, Groups)
│   ├── port/          # Interfaces (puertos)
│   │   ├── inbound/   # CommandBus, EventBus
│   │   └── outbound/  # Logger, WebSocket, SessionStore, etc.
│   └── adapter/       # Implementaciones (adaptadores)
│       ├── ws/fake/   # WebSocket fake
│       ├── store/fake/# Session store fake
│       ├── log/slog/  # Logger con slog
│       └── ...
├── interface/
│   ├── cli/           # Comandos CLI
│   └── http/          # Servidor HTTP básico
└── configs/           # Archivos de configuración
```

## Criterios de Aceptación

- ✅ Arquitectura hexagonal implementada con interfaces claras
- ✅ Todos los adaptadores externos tienen implementaciones fake funcionales
- ✅ CLI funcional con comandos básicos (connect, send, stream, groups)
- ✅ Sistema de configuración flexible con Viper
- ✅ Logging estructurado con slog
- ✅ Container de DI implementado y funcional
- ✅ Cobertura de tests: 100% en dominio, 80%+ en aplicación
- ✅ Compilación sin errores y warnings

## Características Implementadas

### 1. Dominio del Negocio
- **JID**: Identificadores de usuarios de WhatsApp
- **Message**: Mensajes de texto y multimedia
- **Session**: Estado de conexión
- **Event**: Eventos del sistema (message_received, presence_update, etc.)

### 2. Casos de Uso
- **ConnectUseCase**: Manejo de conexiones WebSocket
- **SendMessageUseCase**: Envío de mensajes
- **ReceiveUseCase**: Recepción de eventos
- **GroupsUseCase**: Manejo de grupos

### 3. Adaptadores Fake
- **FakeWebSocketDialer**: Simula conexiones WebSocket con probabilidades configurables
- **FakeStore**: Almacenamiento en memoria para sesiones
- **SlogLogger**: Logging estructurado

## Riesgos y Consideraciones

- Los adaptadores fake no validan la lógica real de WhatsApp
- La configuración permite modos de fallo para testing
- El estado es volátil (no persistente) en implementaciones fake
- Arquitectura preparada para migración a implementaciones reales