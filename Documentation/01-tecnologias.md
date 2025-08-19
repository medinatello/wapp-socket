# 01 - Decisiones Tecnológicas

Este documento resume las tecnologías seleccionadas para el proyecto `wapp-socket`.

| Área                | Tecnología Objetivo (Futuro) | Interfaz de Abstracción        | Adaptador Sprint 1 (Fake)         |
| ------------------- | ---------------------------- | ------------------------------ | --------------------------------- |
| **WebSocket**       | `nhooyr/websocket`           | `port.WebSocketDialer`         | `adapter.ws.fake.FakeWebSocket`   |
| **Almacén Sesiones**| `modernc.org/sqlite`         | `port.SessionStore`            | `adapter.store.fake.FakeStore`    |
| **Almacén Media**   | Almacenamiento local/S3      | `port.MediaStore`              | `adapter.media.fake.FakeMedia`    |
| **Criptografía**    | Noise Protocol (`noise`)     | `port.Crypto`                  | `adapter.crypto.fake.FakeCrypto`  |
| **Codificador Proto** | `google.golang.org/protobuf` | `port.ProtoCodec`              | `adapter.proto.fake.FakeCodec`    |
| **Logging**         | `log/slog` o `zerolog`       | `port.Logger`                  | `adapter.log.slog.SlogLogger`     |
| **CLI**             | `spf13/cobra`                | `interface.cli` (subcomandos)  | Implementación directa en `main`    |
| **Telemetría**      | OpenTelemetry (`otel`)       | `telemetry.Telemetry`          | `telemetry.OtelNoop`              |
| **Reintentos**      | `cenkalti/backoff`           | (Integrado en casos de uso)    | Lógica de reintento no implementada |
| **Mocks (Testing)** | `vektra/mockery`             | (Se aplica a todas las interfaces) | `scripts/gen-mocks.sh` (placeholder)|
