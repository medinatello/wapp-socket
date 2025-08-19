# 6c. Propuestas de Mejora: Eficiencia y Rendimiento

Una librería de comunicación en tiempo real debe ser altamente eficiente para manejar un flujo constante de datos sin agotar los recursos del sistema (CPU, memoria, I/O). Esta sección identifica posibles cuellos de botella y propone mejoras de rendimiento.

## 6c.1. Hotspots de CPU y Memoria Identificados

Basado en el análisis del código, los siguientes son los hotspots de rendimiento más probables:

1.  **Operaciones Criptográficas (`BouncyCastle`)**:
    -   El establecimiento de nuevas sesiones de Signal (intercambio de claves Diffie-Hellman) y el cifrado/descifrado de cada mensaje son operaciones CPU-intensivas.
    -   **Riesgo**: Picos de CPU al recibir una gran cantidad de mensajes nuevos o al iniciar muchas conversaciones nuevas simultáneamente.

2.  **Serialización/Deserialización (`Google.Protobuf`)**:
    -   Aunque Protobuf es eficiente, procesar miles de mensajes durante la sincronización del historial puede generar una presión significativa sobre la CPU y el Garbage Collector (GC) debido a las asignaciones de memoria de los objetos `WebMessageInfo`.

3.  **Procesamiento de Medios (`FFMpegCore`, `SkiaSharp`)**:
    -   La generación de miniaturas o la conversión de formatos de audio/video son operaciones muy pesadas. `FFMpegCore` a menudo funciona lanzando un proceso `ffmpeg` externo, lo cual tiene una sobrecarga considerable.
    -   **Riesgo**: Un único envío de medio puede bloquear un hilo o consumir una cantidad desproporcionada de CPU y memoria.

4.  **Asignaciones de Buffer de Red**:
    -   El `WebSocketClient` utiliza `new byte[8192]` para cada operación de lectura. En un flujo constante de datos, esto genera una presión innecesaria sobre el GC.

## 6c.2. Mejoras de Rendimiento Sugeridas

### Optimización de Buffers con `ArrayPool`

-   **Propuesta**: En lugar de `new byte[]`, utilizar `System.Buffers.ArrayPool<byte>.Shared.Rent()` para obtener buffers de un pool compartido y `Return()` para devolverlos después de su uso.
-   **Beneficio**: Reduce drásticamente las asignaciones de memoria y la frecuencia de las recolecciones de basura (GC), mejorando la latencia y el rendimiento general.
-   **Aplicación**: En el bucle de lectura de `WebSocketClient` y en cualquier lugar donde se manejen buffers de bytes para criptografía o I/O.

### Gestión de Backpressure

-   **Problema**: Actualmente, si el servidor envía datos más rápido de lo que la aplicación puede procesarlos (especialmente durante la sincronización del historial), los mensajes se acumularán en memoria sin límite, lo que puede llevar a un `OutOfMemoryException`.
-   **Propuesta**: Introducir un mecanismo de backpressure. Una excelente opción en .NET es usar **`System.Threading.Tasks.Dataflow`** o **`System.Threading.Channels`**.
    -   Se puede crear un `BufferBlock<T>` o un `Channel<T>` con una capacidad limitada (ej. 1000 mensajes).
    -   El hilo del WebSocket escribe en el bloque/canal.
    -   Uno o más hilos consumidores leen del bloque/canal para procesar (descifrar, guardar en DB).
-   **Beneficio**: Si el buffer se llena, el hilo del WebSocket se bloqueará de forma natural (o asíncrona), lo que ralentiza la lectura del socket. Esto ejerce "presión" hacia atrás hasta el servidor TCP, evitando que la aplicación se sobrecargue.

### Offloading de Procesamiento de Medios

-   **Propuesta**: El procesamiento de medios no debe realizarse en el mismo hilo que gestiona la lógica de mensajería.
-   **Solución**: Crear una cola de trabajo dedicada para estas tareas. Cuando se necesita enviar un medio, se encola una tarea de "procesamiento" y, una vez completada, el resultado (el medio procesado) se pasa al flujo de envío principal.
-   **Beneficio**: Aísla las operaciones pesadas, evitando que bloqueen o degraden el rendimiento de las operaciones en tiempo real.

## 6c.3. Benchmarking

Para validar las mejoras y encontrar nuevos cuellos de botella, es fundamental establecer una suite de benchmarks.

-   **Herramienta Sugerida**: **BenchmarkDotNet**. Es el estándar de facto para el benchmarking en .NET.
-   **Escenarios a Medir**:
    1.  **Cifrado/Descifrado**: Medir el rendimiento de `SessionCipher.Encrypt/Decrypt` para mensajes de diferentes tamaños.
    2.  **Serialización**: Medir el throughput de serialización y deserialización de `WebMessageInfo`.
    3.  **Conexión**: Medir el tiempo que se tarda en establecer una sesión (nueva y reanudada).
    4.  **End-to-End (E2E)**: Medir la latencia completa de envío y recepción de mensajes bajo diferentes cargas.

Estos benchmarks proporcionarían datos concretos para guiar los esfuerzos de optimización y demostrar el impacto de los cambios propuestos.
