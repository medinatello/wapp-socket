# 08. Plan de trabajo y sprints unificado

A continuación se presenta un plan tentativo dividido en sprints de dos semanas, considerando tanto la opción de mejora en .NET como la exploración en Go.

## Variante A: Mejorar la librería actual en .NET

| Sprint | Objetivos principales | Entregables |
|------:|-----------------------|-------------|
| 1 | Definir arquitectura hexagonal, extraer cliente WebSocket en interfaz y módulo, configurar DI | Documento de arquitectura; módulo `IConnectionTransport` y una implementación inicial con `WebSocket` |
| 2 | Refactorizar persistencia: crear repositorio para sesiones y mensajes; añadir cifrado SQLite | Interfaces `ISessionStore` y `IMessageStore`; migraciones y cifrado con `sqlcipher` |
| 3 | Implementar logging estructurado y telemetría; integrar OpenTelemetry | Configuración de `Serilog` u `ILogger`; exportación de trazas a consola |
| 4 | Incrementar cobertura: escribir tests unitarios para módulos refactorizados; añadir pruebas de integración con servidor simulado | Suite de tests (>60 % cobertura) y pipeline de CI |
| 5 | Implementar mejora de manejo de errores: retries con `Polly`, backpressure y métricas | Módulos de resiliencia; benchmarks de throughput |
| 6 | Documentar API pública y crear paquetes NuGet; preparar guía de migración | Documentación actualizada y paquete pre‑release en NuGet |

## Variante B: Reescribir en Go

| Sprint | Objetivos principales | Entregables |
|------:|-----------------------|-------------|
| 1 | Crear prototipo de cliente WebSocket en Go que se conecte a WhatsApp Web y reciba el QR | Prototipo funcional; reporte de viabilidad |
| 2 | Integrar librería de cifrado Signal (explorar wrappers C o implementaciones Go) | API básica para cifrar/descifrar mensajes |
| 3 | Modelar persistencia usando bbolt o SQLite; definir interfaz para sesión | Módulo de almacenamiento; pruebas unitarias |
| 4 | Implementar envío y recepción de mensajes; flujo completo de texto | CLI de demostración; mediciones de rendimiento |
| 5 | Evaluar compatibilidad multiplataforma y empaquetado; generar binarios | Binarios para Windows, Linux y macOS; documentación |
| 6 | Decidir si continuar migración total o mantener en .NET | Informe de decisión y, según el caso, plan de siguiente fase |

Este roadmap puede ajustarse según la carga del equipo y los descubrimientos durante la ejecución.

Proveniencia: Codex, Jules, Copilot y análisis propio.
