# 04. Flujos técnicos

## a) Conectar y autenticar
```mermaid
sequenceDiagram
    participant C as Cliente
    participant WS as WASocket
    participant S as Servidor WA
    C->>WS: Configura SocketConfig
    WS->>S: Conecta WebSocket
    S-->>WS: Solicita clave/QR
    WS-->>C: Evento QR generado
    C->>WS: Escanea QR / envía creds
    WS->>S: Pair-Device / Login
    S-->>WS: success
    WS-->>C: Evento Connection.Update
```

**Errores**: expiración de QR, conexión rechazada. **Retries**: regenerar QR, reconectar.

## b) Enviar mensaje (texto/media)
```mermaid
sequenceDiagram
    participant C
    participant WS
    participant S
    C->>WS: SendMessage(contenido)
    WS->>S: encripta y envía nodo
    S-->>WS: ack / messageID
    WS-->>C: Evento MessageStatus
```
Errores: archivo demasiado grande, sesión inválida. Retries: reintento con backoff exponencial.

## c) Recibir mensaje/eventos
```mermaid
sequenceDiagram
    S-->>WS: frame binario
    WS->>WS: decrypt + parse
    WS-->>C: Evento Message.Upsert
    C->>WS: opcional ack de lectura
```
Errores: frame corrupto, clave rota. Retries: solicitar reenvío.

## d) Crear/gestionar grupos
```mermaid
sequenceDiagram
    C->>WS: CreateGroup(usuarios)
    WS->>S: iq set group
    S-->>WS: group id / estado
    WS-->>C: Evento Group.Update
```
Errores: usuario no válido. Retries: validar miembros.

## e) Cerrar conexión
```mermaid
sequenceDiagram
    C->>WS: EndConnection()
    WS->>S: Close frame
    WS-->>C: Evento Connection.Update(disconnected)
```
Errores: desconexión abrupta; se intenta `Abort`.
