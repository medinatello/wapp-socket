# 07. Evaluación .NET vs Go

Para decidir si se mantiene la librería en .NET o se reescribe en Go se analiza cada criterio relevante.  La tabla asigna un valor subjetivo (1–5) donde 5 es mejor.  Los pesos se basan en la importancia para el proyecto (mayor concurrencia, madurez de bibliotecas, etc.).

| Criterio                       | Peso | .NET (puntuación) | Go (puntuación) | Comentarios |
|-------------------------------|-----:|------------------:|----------------:|-------------|
| Rendimiento y concurrencia    |   3  | 3                | 5              | Go cuenta con goroutines ligeras y canales, ofreciendo mejor escala para I/O concurrente; .NET dispone de async/await pero el modelo de hilos es más pesado. |
| Ecosistema de librerías       |   2  | 4                | 3              | .NET posee bibliotecas maduras para WebSocket, SQLite y DI; en Go algunas áreas (Signal Protocol) son menos maduras. |
| Facilidades de despliegue     |   2  | 4                | 5              | Go genera binarios estáticos sencillos de distribuir; .NET necesita runtime aunque .NET 9 mejora la AOT. |
| Curva de aprendizaje          |   1  | 4                | 3              | El equipo actual domina C#; migrar a Go requerirá capacitación. |
| Soporte multiplataforma       |   1  | 4                | 5              | Ambos soportan Linux, Windows y macOS; Go suele tener menos dependencias. |
| Mantenimiento a largo plazo   |   2  | 3                | 4              | Go ofrece un lenguaje sencillo con menor “magia”; .NET evoluciona rápido pero puede exigir actualizaciones frecuentes. |
| Comunidades y soporte         |   1  | 4                | 4              | Amplias comunidades en ambas; .NET en entornos corporativos y Go en herramientas de red. |
| Librerías de cifrado Signal   |   2  | 4                | 2              | En .NET existe LibSignal; en Go hay proyectos, pero algunos no están 100 % mantenidos. |

### Conclusión

- **Mantener en .NET** es viable y reduce costos de migración, aprovechando los conocimientos del equipo y la disponibilidad de librerías maduras.  
- **Migrar a Go** ofrece mejor rendimiento y simplicidad en la distribución, pero la falta de una librería madura para el protocolo Signal y la necesidad de capacitación añaden riesgos.

Se sugiere iniciar mejoras en .NET mientras se investiga un prototipo en Go para validar la factibilidad de migrar.

Proveniencia: Codex, Jules, Copilot y análisis propio.
