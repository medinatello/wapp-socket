# 09. Riesgos y deuda técnica unificados

A continuación se resumen los riesgos identificados y las deudas técnicas existentes, junto con su impacto, probabilidad y estrategias de mitigación.

| Riesgo / Deuda                      | Impacto | Probabilidad | Mitigación |
|------------------------------------|--------:|------------:|-----------|
| **Pérdida de sesión por fallos de red** | Alto | Alta | Implementar reconexión automática con backoff; persistir estado para reusar QR |
| **Fuga de credenciales**           | Alto | Media | Cifrar base de datos; utilizar secretos en servicios externos; controles de acceso |
| **Dependencia de LibSignal**       | Medio | Media | Investigar alternativas mantenidas; encapsular el uso para facilitar sustitución |
| **Escalabilidad limitada**         | Medio | Alta | Separar responsabilidades y aplicar patrones de concurrencia; diseñar para manejar altos volúmenes |
| **Actualizaciones de protocolo de WhatsApp** | Alto | Alta | Monitorizar cambios en el protocolo; abstraer parseo de mensajes para facilitar actualizaciones |
| **Licencia y cumplimiento**        | Medio | Media | Revisar licencias de Baileys original y respetar atribuciones; documentar modificaciones |
| **Cobertura de pruebas insuficiente** | Alto | Alta | Integrar pruebas continuas desde el inicio de cada sprint; objetivos de cobertura creciente |
| **Curva de aprendizaje de Go** (si se migra) | Medio | Media | Plan de formación y capacitación; iterar con prototipos antes de comprometerse |

Además de estos riesgos, existen numerosas **deudas técnicas**: tamaño excesivo del archivo `WAProto.cs`, acoplamiento entre módulos, falta de comentarios y documentación en el código.  Se recomienda asignar esfuerzos específicos en cada sprint para reducir esta deuda.

Proveniencia: Codex, Jules, Copilot y análisis propio.
