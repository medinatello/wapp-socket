# Documentación de BaileysCSharp

Este directorio contiene una auditoría del proyecto **BaileysCSharp** y un plan de acción para su evolución.

## Índice
1. [Resumen ejecutivo](01-resumen-ejecutivo.md)
2. [Arquitectura actual](02-arquitectura-actual.md)
3. [Tecnologías usadas](03-tecnologias-usadas.md)
   - [SQLite y almacenamiento](03a-sqlite.md)
   - [WebSocket](03b-websocket.md)
   - [Criptografía y claves](03c-criptografia-y-claves.md)
   - [Servicios de Google u otros](03d-servicios-google-o-otros.md)
   - [Procesamiento de media](03e-procesamiento-media.md)
4. [Flujos técnicos](04-flujos-tecnicos.md)
5. [Casos de uso](05-casos-de-uso.md)
6. [Errores, logs y telemetría](06-mejoras-errores-logs-telemetria.md)
   - [Refactor de arquitectura](06b-refactor-arquitectura.md)
   - [Eficiencia y rendimiento](06c-eficiencia-rendimiento.md)
7. [Evaluación .NET vs Go](07-evaluacion-dotnet-vs-go.md)
8. [Plan de sprints y roadmap](08-plan-sprints-y-roadmap.md)
9. [Riesgos, costos y deuda técnica](09-riesgos-costos-y-deuda-tecnica.md)
10. [Glosario y referencias](10-glosario-y-referencias.md)

## Cómo regenerar esta documentación
1. Clonar el repositorio.
2. Instalar .NET 9 SDK.
3. Ejecutar scripts de análisis (por definir) o volver a compilar la solución para obtener datos actualizados.

## Bloqueos y cómo resolverlos
- La solución incluye archivos `creds.json` con credenciales reales; deben ser reemplazados por datos ficticios antes de compartir el repositorio.
- El repositorio no tiene `remote`; para sincronizar con GitHub agregarlo manualmente: `git remote add origin <url>`.
