# 02. Arquitectura actual

## Diagrama general
```mermaid
flowchart LR
    A[Program/Cliente] -->|SocketConfig| B[WASocket]
    B --> C[WebSocketClient]
    B --> D[EventEmitter]
    B --> E[SignalRepository]
    B --> F[MemoryStore (LiteDB)]
    B --> G[LibSignal/Criptografía]
    C -->|Frames binarios| H[WhatsApp Web]
```

## Dependencias internas
- `WASocket` hereda de una cadena de sockets (NewsletterSocket → BusinessSocket → ChatSocket → BaseSocket).
- `BaseSocket` gestiona eventos, reintentos y cifrado de mensajes.
- `MemoryStore` persiste credenciales y claves usando LiteDB.

## Dependencias externas
| Tecnología | Uso actual |
|------------|-----------|
| System.Net.WebSockets | Conexión al servidor de WhatsApp.
| LiteDB | Persistencia embebida de chats, claves y metadatos.
| BouncyCastle | Operaciones de cifrado/descifrado y curvas elípticas.
| Google.Protobuf & Grpc.Tools | Generación de clases de protocolo WA.
| FFMpegCore & SkiaSharp | Procesamiento de multimedia (audio/video/imágenes).

## Puntos de acoplamiento fuerte
1. **`BaseSocket`** mezcla lógica de transporte, negocio y almacenamiento; difícil de testear.
2. **Persistencia** acoplada a LiteDB sin interfaz; sustituirla implica tocar numerosas clases.
3. **Eventos**: `EventEmitter` usa delegados sin contratos claros, lo que dificulta aislamiento.
