# Bloqueos, Supuestos y Decisiones

Este documento registra los puntos que requieren clarificación, los supuestos que se han tomado para poder avanzar y las decisiones clave que podrían necesitar revisión.

## Supuestos

- **Seed Aleatorio**: Se asume que el `SEED_ALEATORIO` provisto como variable de entorno será un entero válido. Si no se provee, se usará el tiempo actual como semilla, lo cual es no-reproducible.
- **Configuración Local**: Para el desarrollo inicial, se usará `configs/config.local.yaml` que no será versionado. El archivo `config.example.yaml` servirá de plantilla.
- **Feature Flags**: El mapa de `FEATURE_FLAGS` se leerá de la configuración, pero para el Sprint 1, el único flag relevante es `FAKE_MODE: true`, que se asumirá como activado por defecto en el código si no se especifica lo contrario.
- **Estado de la CLI**: Los comandos de la CLI (`whats-cli`) son sin estado. Cada comando (`connect`, `send`) se ejecuta en un proceso separado y crea un nuevo contenedor de dependencias. Para que `send` funcione, se ha tomado el atajo de que `send` llame implícitamente a la lógica de `connect` antes de enviar el mensaje. Esto no es ideal, pero es una solución pragmática para el Sprint 1. Un enfoque futuro podría implicar un daemon con el que la CLI se comunique, o la serialización del estado de la sesión en disco.

## Decisiones

- **Makefile**: Se ha creado un `Makefile` con objetivos estándar. Las rutas y nombres de binarios están hardcodeados por simplicidad (`bin/wapp-socket-cli`, `bin/wapp-socket-daemon`).
- **Logging**: Se usará `slog` como librería de logging subyacente, encapsulada por nuestra propia interfaz `Logger` para desacoplar.
- **Telemetría Prometheus**: La tarea original mencionaba un `prometheus_noop.go`. Dado que la interfaz `Telemetry` actual solo tiene `StartSpan` y `RecordCounter`, una implementación no-op separada para Prometheus sería redundante con `otel_noop.go`. Se ha decidido omitir la creación de `prometheus_noop.go` en el Sprint 1 y reconsiderarlo si se añaden métricas más específicas de Prometheus (Gauges, Histograms) a la interfaz.

## Faltantes (Sprint 1)

- No se implementará la lógica de reintentos con backoff (p. ej. `cenkalti/backoff`) en los fakes, solo se simularán los resultados de éxito/fallo.
- La telemetría será completamente no-op. Los puntos de instrumentación se añadirán en el código, pero no emitirán datos.
- Los tests de contrato e integración estarán vacíos (`t.Skip()`) en este sprint.
