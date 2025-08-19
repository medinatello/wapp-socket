# 06c. Eficiencia y rendimiento

## Hotspots actuales
- Lectura completa de frames en memoria (`ReadBytes`) sin `ArrayPool`.
- Procesamiento de media sin streaming; carga completa de archivos.
- Hilos dedicados para *keep alive* y recepción sin control de CPU.

## Backpressure y throughput
- No existen colas para limitar consumo; un aluvión de mensajes podría saturar memoria.
- Envío de media se realiza de uno en uno; no hay manejo de lotes.

## Estrategias propuestas
1. Usar `Channel<T>` o `System.Threading.Channels` para procesar mensajes entrantes con backpressure.
2. Implementar `SemaphoreSlim` para limitar descargas concurrentes de media.
3. Integrar `ArrayPool<byte>` y `MemoryStream` reutilizable para reducir GC.
4. Medir CPU/memoria con `dotnet-counters` y planificar *benchmarks* usando `BenchmarkDotNet`.

## Benchmarks sugeridos
- Escenario 1: 1000 mensajes de texto consecutivos.
- Escenario 2: Transferencia de 10 archivos de 10MB en paralelo.
- Métricas: tiempo total, memoria pico, uso de CPU.
