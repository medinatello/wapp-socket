# 07. Evaluación .NET vs Go

## Matriz comparativa
| Criterio | Peso | .NET 9 | Go 1.22 |
|---------|-----|-------|-------|
| Rendimiento | 0.2 | 7 | 8 |
| Concurrencia | 0.2 | 6 | 9 |
| Ecosistema/paquetes | 0.15 | 8 | 7 |
| Despliegue (single binary) | 0.1 | 6 | 9 |
| Latencia | 0.1 | 7 | 8 |
| Soporte/comunidad | 0.1 | 8 | 8 |
| Costo de talento | 0.1 | 7 | 6 |
| Testabilidad | 0.05 | 8 | 7 |

**Puntaje ponderado**
- .NET: `7.3`
- Go: `8.2`

## Mapeo de equivalencias
| Función | .NET | Go |
|--------|------|----|
| WebSocket | `System.Net.WebSockets` | `nhooyr/websocket` o `gorilla/websocket` |
| Base de datos | `Microsoft.Data.Sqlite` / LiteDB | `mattn/modernc sqlite`, `bbolt` |
| Logging | `Serilog` / MEL | `slog`, `zerolog`, `log` |
| Resiliencia | `Polly` | `cenkalti/backoff`, `retry` |
| DI/IoC | `Microsoft.Extensions.DependencyInjection` | `wire`, `fx`, patrones simples |

## Recomendación final
- **.NET** es adecuado si se busca reuso del código actual y compatibilidad con ecosistemas Windows/Azure.
- **Go** ofrece mejor modelo de concurrencia y despliegue ligero, ideal para servicios de alto volumen.

### Riesgos y mitigaciones
- Migrar a Go implica reescritura completa del protocolo: *mitigar* usando especificación clara y pruebas de interoperabilidad.
- Mantener .NET requiere refactor significativo: *mitigar* estableciendo capas y tests desde el inicio.
