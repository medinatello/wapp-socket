# 9. Análisis de Riesgos, Costos y Deuda Técnica

Este documento consolida los riesgos, la deuda técnica acumulada y una estimación de alto nivel de los costos asociados a cada variante estratégica.

## 9.1. Deuda Técnica Priorizada

La deuda técnica es el costo implícito de retrabajo causado por elegir una solución fácil ahora en lugar de usar un mejor enfoque que tomaría más tiempo. La librería actual tiene una deuda técnica significativa:

1.  **Deuda Arquitectónica (Crítica)**:
    -   **God Object (`WASocket`)**: Centraliza demasiadas responsabilidades, violando el Principio de Responsabilidad Única.
    -   **Alto Acoplamiento**: La lógica de negocio depende directamente de la infraestructura (LiteDB, WebSockets).
    -   **Falta de Inyección de Dependencias**: Dificulta la sustitución de componentes y las pruebas.
    -   **Impacto**: Dificultad extrema para testear, mantener y extender la librería.

2.  **Deuda de Seguridad (Alta)**:
    -   **Credenciales en Texto Plano**: Las claves de sesión se guardan en un archivo JSON sin cifrar.
    -   **Impacto**: Riesgo muy alto de compromiso de la sesión si se accede al sistema de archivos.

3.  **Deuda de Robustez (Alta)**:
    -   **Sin Reconexión Automática**: La librería delega la lógica de resiliencia al cliente.
    -   **Manejo de Errores Básico**: No hay una estrategia clara para diferenciar errores transitorios de fatales.
    -   **Impacto**: La librería no es robusta por sí misma y requiere código "boilerplate" en cada implementación.

4.  **Deuda de Observabilidad (Media)**:
    -   **Logging No Estructurado**: Los logs son `string` a la consola, inútiles para análisis automáticos.
    -   **Sin Telemetría**: Imposible medir el rendimiento o diagnosticar cuellos de botella sin modificar el código.
    -   **Impacto**: Dificultad para operar y depurar la librería en un entorno de producción.

## 9.2. Matriz de Riesgos

Se identifican los siguientes riesgos, priorizados por Impacto y Probabilidad.

| Riesgo                                                | Impacto  | Probabilidad | Mitigación                                                                                                             |
| ----------------------------------------------------- | -------- | ------------ | ---------------------------------------------------------------------------------------------------------------------- |
| **Seguridad: Robo de credenciales de sesión**         | Crítico  | Media        | **Prioridad #1**: Implementar el cifrado en reposo de `creds.json` usando `ProtectedData` (Sprint 5 del plan de adaptación). |
| **Externo: WhatsApp cambia el protocolo**             | Crítico  | Alta         | Mantener una vigilancia activa del proyecto original de Baileys (JS) y estar preparado para portar los cambios rápidamente. |
| **Técnico: La arquitectura actual impide la evolución** | Alto     | Alta         | **Prioridad #2**: Ejecutar los Sprints 1-2 del plan de adaptación para refactorizar a una Arquitectura Hexagonal.      |
| **Técnico: Cascada de fallos por falta de resiliencia** | Alto     | Alta         | Implementar reconexión automática con Polly (Sprint 3 del plan de adaptación).                                       |
| **Operacional: Incapacidad de depurar en producción** | Medio    | Alta         | Implementar logging estructurado y telemetría (Sprint 4 del plan de adaptación).                                      |
| **Proyecto: La reescritura en Go excede el presupuesto/tiempo** | Medio    | Media        | Empezar con un MVP claro, reutilizar la librería actual como referencia funcional y tener un equipo con experiencia en Go. |
| **Legal: El fork original de Baileys es abandonado**  | Bajo     | Media        | La comunidad suele tomar el relevo de forks populares. El riesgo es la velocidad a la que se adaptan a los cambios de WA. |

## 9.3. Análisis de Costos (Cualitativo)

Este es un análisis de alto nivel, sin cifras monetarias.

### Variante A: Mejorar/Adaptar el Proyecto .NET

-   **Inversión Inicial**: **Media**. Requiere un esfuerzo de refactorización significativo (estimado en ~3 meses/equipo).
-   **Costo de Mantenimiento a Largo Plazo**: **Medio**. La base de código será mucho más limpia, pero seguirá siendo una base de código heredada con sus complejidades.
-   **Costo de Talento**: **Bajo**. Se aprovecha el conocimiento del equipo actual.
-   **Riesgo de Ejecución**: **Bajo**. El plan es incremental y cada sprint aporta valor tangible.

### Variante B: Reescribir el Proyecto en Go

-   **Inversión Inicial**: **Alta**. Una reescritura desde cero es un proyecto completo (estimado en 3-6 meses/equipo).
-   **Costo de Mantenimiento a Largo Plazo**: **Bajo**. El resultado será un producto más simple, eficiente y fácil de desplegar, lo que reduce los costos operativos y de mantenimiento.
-   **Costo de Talento**: **Alto**. Requiere desarrolladores de Go, que pueden ser más caros o requerir formación.
-   **Riesgo de Ejecución**: **Medio-Alto**. Riesgo de que surjan complejidades imprevistas del protocolo no documentado.

**Conclusión**: La variante de **adaptación de .NET es una inversión más segura y rápida a corto plazo**, mientras que la **reescritura en Go es una inversión estratégica a largo plazo** que podría resultar en un TCO (Costo Total de Propiedad) más bajo si el rendimiento y la eficiencia operativa son críticos.
