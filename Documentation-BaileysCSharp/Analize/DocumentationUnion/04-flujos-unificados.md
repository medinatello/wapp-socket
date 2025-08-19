# 04. Flujos unificados

A continuación se describen los flujos principales del sistema, integrando aportes de las distintas auditorías y el análisis del código.

## Conectar y autenticar

```mermaid
sequenceDiagram
    participant Cliente as Cliente (.NET)
    participant API as BaileysCSharp API
    participant WS as WebSocket WhatsApp
    Cliente->>API: Solicita iniciar conexión
    API->>WS: Abre WebSocket (wss://web.whatsapp.com)
    WS-->>API: Responde con identificador e indica necesidad de QR
    API-->>Cliente: Emite evento "QR generado"
    Cliente->>Usuario: Muestra QR
    Usuario-->>WhatsApp: Escanea QR
    WS-->>API: Envía evento de autenticación y claves
    API->>LibSignal: Genera y almacena claves de sesión
    API-->>Cliente: Emite evento "conectado"
```

## Enviar mensaje

```mermaid
sequenceDiagram
    participant Cliente
    participant API
    participant WS
    Cliente->>API: Enviar mensaje (texto/archivo)
    API->>LibSignal: Cifra mensaje usando claves de sesión
    API->>WS: Envía paquete en formato Protobuf
    WS-->>API: Ack recibido / error
    API-->>Cliente: Emite evento de confirmación / error
    API->>SQLite: Guarda estado y mensaje
```

## Recibir mensaje

```mermaid
sequenceDiagram
    participant WS
    participant API
    participant Cliente
    WS-->>API: Recibe paquete Protobuf
    API->>LibSignal: Descifra mensaje
    API->>SQLite: Guarda mensaje
    API-->>Cliente: Emite evento "mensaje recibido"
```

## Cerrar conexión

```mermaid
sequenceDiagram
    participant Cliente
    participant API
    participant WS
    Cliente->>API: Solicita cerrar sesión
    API->>WS: Envía cierre de WebSocket
    WS-->>API: Confirma desconexión
    API->>SQLite: Actualiza estado
    API-->>Cliente: Emite evento "desconectado"
```

Estos flujos permiten entender cómo se propagan las llamadas y eventos.  Se recomienda extraer estas secuencias en métodos bien definidos y añadir políticas de reintento y cancelación.

Proveniencia: Codex, Jules, Copilot y análisis propio.
