# 03b. WebSocket

## Cliente actual
- Implementación propia `WebSocketClient` sobre `System.Net.WebSockets`.
- Conexión a `wss://web.whatsapp.com/ws/chat` con cabeceras personalizadas.
- Lectura manual de frames binarios con prefijo de tamaño.

## Heartbeat y reconexión
- `BaseSocket` lanza un hilo de *keep alive* que envía `iq` cada 30s.
- Si no hay respuesta en ~35s se cierra la conexión.
- No existe política de backoff ni reconexión automática; la lógica depende del consumidor.

## Problemas detectados
1. Lectura de frames asume tamaños pequeños y bloquea en memoria.
2. Falta de cancelación por `CancellationToken` en todas las operaciones.
3. No hay métricas de latencia ni buffers.

## Mejoras sugeridas en .NET
- Utilizar `System.Net.WebSockets.ClientWebSocket` con `ReceiveAsync` en bucle, usando `ArrayPool` para buffers.
- Implementar reconexión exponencial (`Polly` o lógica propia) y notificación de estado.
- Separar el framing binario en una capa dedicada.

## Alternativas en Go
| Librería | Pros | Contras |
|---------|------|--------|
| `nhooyr/websocket` | API moderna, contexto con cancelación | No soporta algunas extensiones antiguas |
| `gorilla/websocket` | Amplio soporte y ejemplos | Proyecto en mantenimiento limitado |

## Recomendación
1. Encapsular WebSocket en una interfaz `ITransport` para permitir pruebas y reemplazos.
2. Implementar ping/pong formal y reconexión con backoff.
3. Medir buffers y throughput para dimensionar recursos.
