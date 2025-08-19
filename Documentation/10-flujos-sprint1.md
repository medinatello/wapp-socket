# 10 - Flujos Simulados (Sprint 1)

Este documento describe el comportamiento esperado y los logs generados por cada caso de uso simulado a través de la CLI. El comportamiento aleatorio es reproducible usando el flag `--seed`.

## 1. `whats-cli connect`

Este comando simula el proceso de conexión al servidor de WhatsApp.

- **Flujo**:
  1. El `container` DI inyecta el `FakeWebSocketDialer` en el caso de uso `ConnectUseCase`.
  2. Se llama a `ConnectUseCase.Execute()`.
  3. El caso de uso invoca a `FakeWebSocketDialer.Dial()`.
  4. Basado en el `seed`, el `Dialer` simula uno de tres resultados:
     - **70% Éxito**:
       - Log: `INFO: [FakeWS] Connection successful.`
       - Retorna una `FakeWebSocketConn` y `nil` error.
       - El caso de uso guarda la sesión "conectada" en el `FakeSessionStore`.
       - Log: `INFO: [ConnectUseCase] Session stored successfully for JID <fake-jid>.`
       - CLI Output: `Connection successful. Session ready.`
     - **20% Timeout**:
       - Log: `WARN: [FakeWS] Connection attempt timed out.`
       - Retorna `nil` y un error `context.DeadlineExceeded`.
       - El caso de uso propaga el error.
       - CLI Output: `Error: connection timed out.`
     - **10% Credenciales Inválidas**:
       - Log: `ERROR: [FakeWS] Invalid credentials provided.`
       - Retorna `nil` y un error `domain.ErrInvalidCredentials`.
       - El caso de uso propaga el error.
       - CLI Output: `Error: invalid credentials.`

## 2. `whats-cli send --to <jid> --text "..."`

Simula el envío de un mensaje de texto.

- **Flujo**:
  1. El `container` inyecta la `FakeWebSocketConn` (si existe) y el `FakeSessionStore` en `SendMessageUseCase`.
  2. `SendMessageUseCase.Execute()` es llamado.
  3. El caso de uso primero verifica si hay una sesión "conectada" en `FakeSessionStore`.
     - **Sin Sesión**:
       - Log: `ERROR: [SendMessageUseCase] No active session found. Cannot send message.`
       - Retorna `domain.ErrSessionNotFound`.
       - CLI Output: `Error: not connected. Please run 'connect' first.`
     - **Con Sesión**:
       - El caso de uso invoca a `FakeWebSocketConn.Send()` con un frame de texto simulado.
       - Log: `INFO: [FakeWS] Sending frame: <frame-details>`
       - El método `Send()` simula una latencia aleatoria (p. ej., 50-200ms).
       - Log: `INFO: [SendMessageUseCase] Message sent to <jid>. Waiting for ACK.`
       - El sistema simula la recepción de un ACK a través del flujo de `receive`.
       - CLI Output: `Message sent successfully.`

## 3. `whats-cli stream --duration 10s`

Simula la recepción de eventos del servidor durante un período de tiempo.

- **Flujo**:
  1. El `container` obtiene el canal de eventos de la `FakeWebSocketConn` (o se suscribe al `EventBus`).
  2. El caso de uso `ReceiveUseCase` comienza a leer los eventos.
  3. La `FakeWebSocketConn` tiene un bucle interno (goroutine) que, basado en el `seed`, genera eventos cada N milisegundos (p. ej., 500-2000ms).
  4. Tipos de eventos generados aleatoriamente:
     - `message_received`: Un mensaje de texto de un contacto fake.
     - `presence_update`: El estado de un contacto cambia a "composing" o "online".
     - `qr_refresh`: Simula la necesidad de escanear un nuevo código QR.
     - `disconnected`: El servidor cierra la conexión.
  5. Por cada evento, el `ReceiveUseCase` logea la información.
     - Log: `INFO: [ReceiveUseCase] Event received: <event-type> | Data: <event-data>`
  6. El comando termina cuando el `duration` expira o se recibe un evento `disconnected`.

## 4. `whats-cli groups list`

Simula la obtención de una lista de grupos.

- **Flujo**:
  1. `GroupsUseCase.List()` es invocado.
  2. El caso de uso devuelve una lista de 3-5 grupos dummy con nombres y JIDs generados de forma predecible a partir del `seed`.
  3. Log: `INFO: [GroupsUseCase] Returning 4 dummy groups.`
  4. CLI Output: Muestra una tabla con los nombres e IDs de los grupos.
