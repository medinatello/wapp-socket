# 08. Plan de sprints y roadmap

## Variante A: Mejorar/Adaptar (.NET)
| Sprint | Objetivo | Criterios de aceptación | Entregables |
|-------|----------|------------------------|-------------|
| 1 | Configurar CI, remover credenciales, pruebas básicas | Pipeline ejecuta `dotnet build` y tests vacíos | Repo limpio, GitHub Actions |
| 2 | Abstraer persistencia `IStateStore` | Pruebas unitarias con store en memoria | Interfaz + implementación LiteDB |
| 3 | Implementar manejo de errores/logging | Logs estructurados y retries | Integración Serilog + Polly |
| 4 | Refactor WebSocket con reconexión | Simulación de caída y reconexión automática | `ITransport` con backoff |
| 5 | Flujos de envío/recepción modularizados | Tests de integración con mensajes ficticios | Servicios `IMessageService` |
| 6 | Cobertura de media y grupos | Soporte básico y pruebas | Casos de uso completos |

## Variante B: Reescritura (Go)
| Sprint | Objetivo | Criterios | Entregables |
|-------|----------|----------|-------------|
| 1 | Especificar protocolo y estructuras | Documento técnico aprobado | Repo base en Go |
| 2 | Implementar WebSocket + handshake | Test contra servidor real/simulado | Módulo transporte |
| 3 | Criptografía y almacenamiento ligero (`bbolt`) | Pruebas de cifrado | Módulo crypto + store |
| 4 | Enviar/recibir texto | Tests unitarios | API básica |
| 5 | Media y grupos | Subida/descarga controlada | Módulos adicionales |
| 6 | Integración con sistemas externos y despliegue | Docker image | MVP funcional |

## Roadmap general
1. Evaluar resultados de la Variante A tras Sprint 3.
2. Si la deuda persiste o el rendimiento es insuficiente, pivotar a Variante B.
3. Mantener paridad de funcionalidades durante la transición.
