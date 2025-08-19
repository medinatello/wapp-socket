# 02 - Interfaces (Puertos)

Este documento define las interfaces (puertos) clave en la arquitectura hexagonal de la aplicación.

## Puertos de Salida (Outbound)

Los puertos de salida son interfaces que definen cómo la aplicación se comunica con el mundo exterior (bases de datos, servicios de red, sistemas de archivos, etc.).

- **`WebSocketDialer` / `WebSocketConn`**:
  - **Responsabilidad**: Abstraer la creación y gestión de una conexión WebSocket.
  - **Métodos `WebSocketDialer`**: `Dial(ctx, url)`
  - **Métodos `WebSocketConn`**: `Send(ctx, frame)`, `Receive(ctx)` (devuelve `<-chan Frame`), `Close()`.
  - **Implementación Sprint 1**: `adapter/ws/fake` (simula conexión y eventos).

- **`SessionStore`**:
  - **Responsabilidad**: Persistir y recuperar datos de sesión (claves, JID, etc.).
  - **Métodos**: `Get(ctx, jid)`, `Put(ctx, session)`, `Delete(ctx, jid)`.
  - **Implementación Sprint 1**: `adapter/store/fake` (mapa en memoria).

- **`MediaStore`**:
  - **Responsabilidad**: Guardar y recuperar archivos multimedia.
  - **Métodos**: `Save(ctx, mediaData)` (devuelve URL/handle), `Get(ctx, url)`.
  - **Implementación Sprint 1**: `adapter/media/fake` (no-op, devuelve handles predecibles).

- **`Crypto`**:
  - **Responsabilidad**: Encapsular las operaciones criptográficas (cifrado/descifrado de mensajes, handshake).
  - **Métodos**: `Encrypt(ctx, plaintext)`, `Decrypt(ctx, ciphertext)`.
  - **Implementación Sprint 1**: `adapter/crypto/fake` (devuelve datos fijos).

- **`ProtoCodec`**:
  - **Responsabilidad**: Codificar y decodificar los mensajes en formato Protobuf.
  - **Métodos**: `Encode(ctx, object)`, `Decode(ctx, data, object)`.
  - **Implementación Sprint 1**: `adapter/proto/fake` (simula la operación).

- **`Logger`**:
  - **Responsabilidad**: Proporcionar un logging estructurado y nivelado.
  - **Métodos**: `Debug(msg, args...)`, `Info(msg, args...)`, `Warn(msg, args...)`, `Error(msg, args...)`.
  - **Implementación Sprint 1**: `adapter/log/slog` (wrapper sobre `log/slog`).

## Puertos de Entrada (Inbound)

Los puertos de entrada definen cómo los actores externos (UI, CLI, API HTTP) interactúan con la aplicación.

- **`CommandBus` (Alternativa A)**:
  - **Responsabilidad**: Aceptar comandos de los casos de uso y ejecutarlos.
  - **Métodos**: `Register(commandType, handler)`, `Dispatch(ctx, command)`.
  - **Sprint 1**: Se evaluará su implementación, pero el enfoque principal será más directo.

- **`EventBus` (Alternativa A)**:
  - **Responsabilidad**: Publicar eventos del sistema (p.ej., `message_received`, `disconnected`) para que otros componentes puedan reaccionar.
  - **Métodos**: `Publish(ctx, event)`, `Subscribe(ctx, topic)`.
  - **Sprint 1**: Implementación fake en memoria.

- **Canales Go directos (Alternativa B)**:
  - **Responsabilidad**: Exponer canales de Go directamente desde los componentes (p.ej., `WebSocketConn`) para consumir eventos.
  - **Sprint 1**: Implementación detrás de un feature flag o build tag.
