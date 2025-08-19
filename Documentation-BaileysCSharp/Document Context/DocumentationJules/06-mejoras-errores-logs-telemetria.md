# 6. Propuestas de Mejora: Errores, Logs y Telemetría

La robustez y la observabilidad son cruciales para una librería que gestiona comunicaciones en tiempo real. El estado actual de la librería en estas áreas es muy básico y puede mejorarse significativamente.

## 6.1. Manejo de Errores

**Línea Base Actual**:
-   La librería utiliza una clase de excepción personalizada, `Boom`, para señalar errores de conexión.
-   El código cliente es responsable de capturar esta excepción y determinar la causa analizando su `StatusCode`.
-   No hay una distinción clara a nivel de tipo entre errores reintentables (ej. un timeout de red) y errores fatales (ej. credenciales inválidas).

**Anti-Patrones Detectados**:
-   **Excepciones Genéricas**: El cliente debe inspeccionar propiedades de la excepción para entender el error, en lugar de capturar tipos de excepción específicos.
-   **Responsabilidad Delegada**: La librería delega toda la lógica de recuperación (reintentos) al cliente.

**Propuesta de Mejora**:
1.  **Jerarquía de Excepciones Específicas**:
    -   Crear una excepción base, `BaileysException`.
    -   Derivar excepciones específicas como:
        -   `TransientNetworkException`: Para errores de red que pueden reintentarse.
        -   `AuthenticationException`: Para fallos de login (fatal).
        -   `ProtocolException`: Para errores de parseo o cifrado (usualmente fatal).
        -   `RateLimitException`: Si se detecta que el servidor está limitando las peticiones.
2.  **Patrón de Resultado (Alternativa)**: En lugar de excepciones, los métodos podrían devolver un objeto `Result<T>` que contenga el valor exitoso o un objeto de error detallado. Esto hace que el manejo de errores sea más explícito.
3.  **Resiliencia Integrada**: Como se mencionó en la sección de WebSocket, la librería debería gestionar sus propios reintentos para errores transitorios usando una política de Polly.

## 6.2. Logging Estructurado

**Línea Base Actual**:
-   Se utiliza un `ILogger` personalizado muy simple (`DefaultLogger`) que escribe mensajes de texto plano a la consola.
-   Los niveles de log son `Info`, `Warning`, `Error`, `Debug`, `Trace`, `Raw`.
-   No hay contexto en los logs (ej. un ID de correlación para seguir una operación).

**Propuesta de Mejora**:
Adoptar una librería de logging estructurado como **Serilog**.

**Beneficios**:
-   **Logs como Datos**: Los logs se escriben en formato JSON, con propiedades que se pueden buscar y filtrar fácilmente (ej. `{"userJid": "...", "messageId": "...", "latencyMs": 55}`).
-   **Enriquecimiento**: Se puede añadir contexto automáticamente a todos los logs, como un `ConnectionId` o `TraceId`.
-   **Sinks Configurables**: Los logs se pueden enviar a múltiples destinos (consola, archivo, Seq, Application Insights, etc.) con solo cambiar la configuración.

**Plan de Implementación**:
1.  Reemplazar `ILogger` por la interfaz `Microsoft.Extensions.Logging.ILogger<T>`.
2.  Configurar Serilog como el proveedor de logging.
3.  Enricher los logs con propiedades relevantes en cada punto del código.

## 6.3. Telemetría con OpenTelemetry (OTel)

**Línea Base Actual**:
-   No existe ninguna forma de telemetría. Es imposible medir el rendimiento o trazar el flujo de una operación sin modificar el código.

**Propuesta de Mejora**:
Integrar **OpenTelemetry**, el estándar de la industria para la instrumentación de software. OTel unifica tres pilares de la observabilidad:

1.  **Trazas (Traces)**:
    -   Permiten visualizar el ciclo de vida completo de una operación a través de diferentes componentes.
    -   **Ejemplo**: Una traza para "enviar mensaje" podría tener "spans" (intervalos de tiempo) para: `Validación -> Cifrado -> Serialización -> Envío por Socket -> ACK del Servidor`. Esto permite identificar cuellos de botella al instante.
2.  **Métricas (Metrics)**:
    -   Permiten medir y agregar datos numéricos.
    -   **Ejemplos**:
        -   `baileys.messages.sent.count` (Contador): Número de mensajes enviados.
        -   `baileys.encryption.duration.milliseconds` (Histograma): Latencia de la operación de cifrado.
        -   `baileys.connected.clients` (Gauge): Número de clientes actualmente conectados.
3.  **Logs**:
    -   OTel también puede exportar los logs (de Serilog), correlacionándolos automáticamente con las trazas.

**Plan de Implementación**:
1.  Añadir los paquetes de OpenTelemetry (`OpenTelemetry.Api`, `OpenTelemetry.Exporter.Console`, etc.).
2.  Crear una `ActivitySource` para generar las trazas y sus spans en puntos clave.
3.  Crear un `Meter` para registrar las métricas.
4.  Configurar un exportador (ej. a la consola para empezar, luego a Jaeger para trazas o Prometheus para métricas).

## 6.4. Políticas de Retención y Scrubbing

-   **Scrubbing de Datos Sensibles**: Al usar logging estructurado, es vital asegurarse de que los datos personales (números de teléfono, contenido de mensajes) no se escriban en los logs por defecto. Se debe implementar una política de "scrubbing" o anonimización para los logs que se envíen a sistemas externos.
-   **Retención de Datos**: La política de retención de logs y datos de la base de datos local (`store.db`) debe ser configurable por el cliente de la librería para cumplir con normativas de privacidad.
