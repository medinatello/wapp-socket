# 03. Tecnologías usadas

| Tecnología | Propósito | Uso actual | Problemas detectados | Mejoras sugeridas | Detalle |
|------------|-----------|------------|---------------------|------------------|--------|
| LiteDB | Almacenamiento embebido | Guardado de chats, claves y sesiones | Sin cifrado, posible corrupción en acceso concurrente | Abstraer mediante interfaz; evaluar SQLite o base cifrada | [03a-sqlite.md](03a-sqlite.md) |
| System.Net.WebSockets | Transporte | Cliente WebSocket hacia `wss://web.whatsapp.com/ws/chat` | Sin manejo robusto de reconexión/backoff | Usar política de reconexión exponencial, heartbeat y métricas | [03b-websocket.md](03b-websocket.md) |
| BouncyCastle / System.Security.Cryptography | Criptografía | ECDH, AES, firmas | Almacenamiento de claves en texto plano | Uso de KDF, cifrado en reposo, rotación de claves | [03c-criptografia-y-claves.md](03c-criptografia-y-claves.md) |
| Google.Protobuf / Grpc.Tools | Serialización protobuf | Generación de WAProto | Código generado mezclado con fuente, difícil de actualizar | Automatizar build; usar versiones actualizadas | [03d-servicios-google-o-otros.md](03d-servicios-google-o-otros.md) |
| FFMpegCore & SkiaSharp | Procesamiento multimedia | Transcodificación de audio/video, manipulación de imágenes | Dependencias nativas pesadas, sin control de recursos | Pooling de procesos, validación de formatos | [03e-procesamiento-media.md](03e-procesamiento-media.md) |
| Logging propio (`DefaultLogger`) | Telemetría básica | Escribe en consola | No soporta niveles estructurados ni sinks | Migrar a `Microsoft.Extensions.Logging`/Serilog | [06-mejoras-errores-logs-telemetria.md](06-mejoras-errores-logs-telemetria.md) |
