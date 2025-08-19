# Codex Agents - TemplateGo

Este archivo documenta los agentes y asistentes recomendados para automatizar tareas en TemplateGo, siguiendo el esquema de prompts y documentación del proyecto.

Comunicate siempre en español, incluso los logs que emites cuando vas realizando las tareas.

## Índice de Agentes

- [Gemini](#gemini)
- [Claude](#claude)
- [Copilot](#copilot)
- [Codex](#codex)

## Agentes Disponibles

### Gemini
**Descripción**: Agente de Google Gemini, especializado en prompts para Go y análisis de arquitectura, migraciones y automatización de sprints.
- Usa el archivo `GEMINI.md` para prompts y ejemplos.
- Recomendado para: migraciones, análisis de arquitectura, generación de documentación técnica.

### Claude
**Descripción**: Agente de Anthropic Claude, enfocado en prompts de revisión, ejecución de sprints y análisis de código Go.
- Usa el archivo `Claude.md` para prompts y ejemplos.
- Recomendado para: ejecución de tareas de MVP, code review, generación de reportes de calidad.

### Copilot
**Descripción**: GitHub Copilot, asistente de autocompletado y generación de código Go en tiempo real.
- Integrado en VS Code y otros IDEs.
- Recomendado para: generación de código, refactorización rápida, sugerencias de tests y documentación.

### Codex
**Descripción**: OpenAI Codex, agente para generación de código, automatización de scripts y análisis de dependencias en Go.
- Usa prompts similares a los de Gemini y Claude.
- Recomendado para: generación de scripts, automatización de pipelines, análisis de dependencias y refactorizaciones masivas.

## Uso de los Agentes

- Selecciona el agente según la tarea (ver recomendaciones arriba).
- Utiliza el archivo de prompts correspondiente (`Claude.md`, `GEMINI.md`).
- Documenta los resultados y decisiones en los archivos de resultado de cada MVP.

## Ejemplo de flujo de trabajo

1. Ejecuta el prompt de sprint en `Claude.md` para iniciar el trabajo de un MVP.
2. Usa Gemini para análisis de arquitectura y migraciones.
3. Apóyate en Copilot para generación de código y tests.
4. Utiliza Codex para automatización de scripts y refactorizaciones.
5. Documenta todo en los archivos de resultado y documentación principal.

## Notas

- Todos los agentes deben seguir los estándares definidos en `Cases/Documentation Main/`.
- Mantén la trazabilidad de decisiones y cambios en los archivos de resultado.
- Actualiza este archivo si se integran nuevos agentes o asistentes.
