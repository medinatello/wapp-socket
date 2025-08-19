# 09. Riesgos, costos y deuda técnica

## Riesgos principales
| Riesgo | Impacto | Probabilidad | Plan de contingencia |
|-------|---------|--------------|----------------------|
| Exposición de credenciales | Alto | Alta | Eliminar archivos sensibles, cifrar store |
| Cambios de protocolo WA | Alto | Media | Mantener seguimiento de repositorio oficial Baileys |
| Falta de tests al refactor | Medio | Alta | Adoptar TDD y coverage mínimo 70% |
| Rendimiento insuficiente en .NET | Medio | Media | Evaluar Go tras sprint 3 |
| Dependencia de libs nativas (ffmpeg) | Medio | Baja | Contener en contenedor Docker |

## Estimación de costos (aprox.)
- **Variante A (.NET refactor)**: 6 sprints × 2 devs = ~12 semanas. Costo estimado: 12 × 2 × 40h × tarifa.
- **Variante B (Go reescritura)**: 6 sprints × 3 devs = ~18 semanas.

## Deuda técnica existente
- Ausencia de pruebas automatizadas.
- Mezcla de responsabilidades en `BaseSocket`.
- Persistencia sin cifrado ni abstracción.
- Código heredado sin documentación.
