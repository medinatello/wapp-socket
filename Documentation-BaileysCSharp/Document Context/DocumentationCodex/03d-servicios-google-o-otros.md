# 03d. Servicios Google u otros

## SDKs y APIs
- **Google.Protobuf / Grpc.Tools**: utilizados para generar clases de protocolo (`WAProto.cs`, `WhisperTextProtocol.cs`).
- No se detectan otras APIs de Google ni llamadas a servicios externos (Firebase, Maps, etc.).

## Problemas
- Generación de código mezclada en el repositorio, lo que dificulta actualizaciones.
- Dependencia de `protoc` en máquinas ARM64 con pasos manuales.

## Recomendaciones
1. Automatizar la generación de protobuf en CI y excluir archivos generados del control de versiones.
2. Centralizar versiones de `Google.Protobuf` para evitar incompatibilidades.
3. Considerar migrar a `protobuf-net` si se busca integración más directa con .NET.
