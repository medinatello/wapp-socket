# 06. Mejoras: errores, logs y telemetría

## Línea base
- Manejo de errores se limita a `try/catch` con mensajes en consola.
- `DefaultLogger` no diferencia niveles ni estructura.
- Eventos no registran correlación ni IDs de operación.

## Propuesta de manejo de errores
1. Definir jerarquía de excepciones (`Transient`, `Fatal`, `AuthError`).
2. Usar `Polly` para retires/backoff en operaciones de red.
3. Propagar `CancellationToken` en todas las tareas.

## Logging estructurado
- Adoptar `Microsoft.Extensions.Logging` + `Serilog`.
- Formato JSON con campos: `traceId`, `jid`, `event`.
- Niveles: `Trace`, `Debug`, `Info`, `Warn`, `Error`.

## Telemetría
- Instrumentar con **OpenTelemetry**: traces y métricas (latencia, reconexiones, tamaño de mensaje).
- Exportadores configurables: Jaeger, Prometheus, Application Insights.
- Política de retención: 30 días, anonimización de datos sensibles.

## Scrub de datos
- Filtrar números telefónicos y texto sensible antes de loguear.
- Almacenar hashes de JIDs en lugar de valores completos.
