# 3b. Análisis de Comunicaciones con WebSocket

## 3b.1. Cliente WebSocket Actual

La comunicación con los servidores de WhatsApp se realiza a través de una conexión WebSocket persistente.

**Implementación**:
-   La librería utiliza el cliente nativo de .NET, `System.Net.WebSockets.ClientWebSocket`.
-   La clase `BaileysCSharp.Core.Sockets.Client.WebSocketClient` actúa como un wrapper delgado alrededor de `ClientWebSocket`. Es responsable de:
    1.  Establecer las cabeceras HTTP necesarias para la conexión (`Origin`, `Host`, etc.).
    2.  Conectar al endpoint `wss://web.whatsapp.com/ws/chat`.
    3.  Mantener un bucle de lectura (`ReceiveAsync`) para procesar los frames binarios entrantes.
    4.  Exponer un método `Send` para enviar datos binarios al servidor.

**Gestión del Ciclo de Vida**:
-   **Heartbeat (Ping/Pong)**: No se ha identificado una lógica explícita de `ping/pong` a nivel de aplicación para mantener la conexión viva. Es posible que se confíe en el `keep-alive` de TCP o que el servidor de WhatsApp gestione la vitalidad de la conexión de otra manera (ej. enviando frames de `noop`).
-   **Pérdida de Mensajes**: Si la conexión se cae entre el envío y la recepción de un mensaje, el sistema depende de la sincronización de historial de WhatsApp para recuperar los mensajes perdidos en la siguiente conexión. No hay un buffer de salida persistente a nivel de cliente.
-   **Límites de Buffer**: Se utiliza un `ArraySegment` de 8192 bytes para leer del socket. Si un frame de WhatsApp excede este tamaño, el código está preparado para leer en fragmentos hasta recibir el mensaje completo (`result.EndOfMessage`).

## 3b.2. Política de Reconexión y Estado de Sesión

Este es uno de los puntos más débiles de la arquitectura actual.

**La librería no gestiona su propia reconexión.**

El `WebSocketClient` o `WASocket` no intentan restablecer la conexión de forma autónoma en caso de fallo. La responsabilidad recae enteramente en el código cliente que consume la librería.

Como se observa en `WhatsSocketConsole/Program.cs`, el cliente debe:
1.  Suscribirse al evento `Connection.Update`.
2.  Detectar un estado de `WAConnectionState.Close`.
3.  Inspeccionar el error (`LastDisconnect.Error`). Si no es un cierre de sesión voluntario (`DisconnectReason.LoggedOut`), el cliente debe esperar un tiempo y volver a llamar a `socket.MakeSocket()`.

**Problemas de este enfoque**:
-   **Lógica Duplicada**: Cada aplicación que use la librería debe implementar su propia lógica de reconexión.
-   **Falta de Estrategia de Backoff**: El ejemplo de la consola usa un `Thread.Sleep(1000)` simple. Una estrategia robusta requeriría un backoff exponencial para no sobrecargar al servidor en caso de fallos repetidos.
-   **Gestión de Estado Compleja**: El estado de la conexión está disperso entre la excepción `Boom` y el objeto `ConnectionState`, lo que complica la toma de decisiones.

## 3b.3. Alternativas y Mejoras

### En .NET

La solución ideal es integrar una librería de resiliencia y manejo de fallos transitorios.

| Alternativa             | Ventajas                                                              | Desventajas                                                                 |
| ----------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| **Polly**               | El estándar de facto en .NET para resiliencia. Permite definir políticas complejas de reintento, backoff, circuit breaker, etc. | Añade una dependencia. Requiere refactorizar la lógica de conexión para envolverla en una política de Polly. |
| **`Websocket.Client`**    | Una librería cliente de WebSocket que ya incluye lógica de reconexión automática y personalizable. | Reemplaza el cliente nativo. Puede ser menos flexible si se necesita control de bajo nivel sobre las cabeceras o el handshake. |

**Recomendación para .NET**:
Integrar **Polly** dentro de `WASocket`. La llamada a `ConnectAsync` debería estar envuelta en una `RetryPolicy` de Polly que:
1.  Reintente la conexión en caso de `WebSocketException` o errores de red.
2.  Utilice una estrategia de backoff exponencial (ej. esperar 1s, 2s, 4s, 8s...).
3.  Permita configurar el número máximo de reintentos.
Esto haría la librería mucho más robusta y autónoma.

### En Go

Go es excelente para manejar este tipo de lógica de red concurrente.

| Alternativa             | Equivalencia .NET       | Ventajas en Go                                                                   | Desventajas                                                              |
| ----------------------- | ----------------------- | -------------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
| **`nhooyr/websocket`**  | `System.Net.WebSockets` | Considerado un cliente de WebSocket moderno y robusto para Go. API limpia.      | No incluye lógica de reconexión por defecto, se debe implementar manualmente. |
| **`gorilla/websocket`** | `Websocket.Client`      | El cliente de WebSocket más popular y antiguo. Muy probado en batalla.          | La API es un poco más verbosa que la de `nhooyr`.                        |
| **`cenkalti/backoff`**  | Polly                   | Una implementación popular del algoritmo de backoff exponencial. Simple de usar para los reintentos. | Solo gestiona el backoff, no otras estrategias como circuit breaker.     |

**Recomendación para Go**:
Una combinación de **`nhooyr/websocket`** para la conexión y **`cenkalti/backoff`** para gestionar la lógica de reintentos sería una solución idiomática y robusta en Go. La lógica de reconexión se implementaría dentro de un goroutine que gestiona el ciclo de vida de la conexión.
