# 1. Resumen Ejecutivo

## 1.1. Estado Actual de la Librería `BaileysCSharp`

**¿Qué hace hoy la librería?**
`BaileysCSharp` es un fork de la popular librería Baileys (JavaScript) que implementa el protocolo de WhatsApp Web. Permite a una aplicación .NET actuar como un cliente de WhatsApp, capaz de conectar, autenticar, enviar y recibir mensajes de texto y multimedia.

**¿En qué estado está?**
La librería es **funcional a nivel de prototipo**, pero presenta **deficiencias críticas** en áreas clave que la hacen inadecuada para un uso en producción sin una intervención significativa. Funciona, pero es frágil, insegura y difícil de mantener.

**Problemas Críticos Identificados**:
1.  **Riesgo de Seguridad Crítico**: Las credenciales de sesión, incluidas las claves privadas maestras, se almacenan en el disco **sin cifrar**. Un acceso al sistema de archivos compromete la sesión por completo.
2.  **Arquitectura Rígida y Acoplada**: La lógica de negocio está fuertemente acoplada a la infraestructura (base de datos, red). Esto hace que la librería sea extremadamente difícil de probar unitariamente, mantener y extender.
3.  **Falta de Robustez**: La librería no gestiona su propio ciclo de vida de conexión. La lógica de reconexión automática en caso de fallo de red es responsabilidad del cliente que la consume, lo que lleva a implementaciones inconsistentes y frágiles.

## 1.2. Los Dos Caminos: Adaptar vs. Reescribir

Se presentan dos caminos estratégicos para abordar la deuda técnica y los riesgos identificados.

### Camino A: Mejorar y Adaptar el Proyecto .NET
Consiste en una refactorización profunda de la base de código actual para introducir una arquitectura moderna (Hexagonal), mejorar la seguridad, añadir resiliencia (Polly) y observabilidad (OpenTelemetry).
-   **Ventajas**: Menor costo y tiempo inicial, aprovecha el código existente, riesgo de ejecución controlado.
-   **Desventajas**: No alcanza el máximo rendimiento teórico, arrastra parte de la complejidad del fork original.

### Camino B: Reescribir el Proyecto desde Cero en Go
Consiste en desarrollar una nueva librería en Go, utilizando la implementación actual como referencia funcional.
-   **Ventajas**: Producto final técnicamente superior, con mejor rendimiento de concurrencia y despliegue trivial (un solo binario). Código limpio sin deuda técnica.
-   **Desventajas**: Mayor costo y tiempo, requiere talento especializado en Go, riesgo de implementación al tratarse de un protocolo no documentado.

## 1.3. Recomendación Inicial y Semáforo

**Semáforo de Decisión: 🟡 Amarillo**

El estado actual del proyecto requiere precaución. No es apto para producción (`Rojo`), pero es salvable con un esfuerzo de ingeniería bien dirigido (`Verde`). La recomendación es proceder, pero con una refactorización planificada.

---

**Recomendación: Iniciar con el Camino A (Mejorar y Adaptar .NET)**

**Razones Clave**:
1.  **Pragmatismo y Velocidad**: El plan de adaptación (detallado en `08-plan-sprints-y-roadmap.md`) permite mitigar los riesgos más críticos (seguridad, robustez) en los primeros sprints, aportando valor de forma incremental y rápida.
2.  **Inversión Reutilizable**: La principal mejora (refactor a Arquitectura Hexagonal) consiste en desacoplar la lógica de negocio de la infraestructura. Este trabajo **no se desperdicia**. El "Core" de negocio resultante serviría como una especificación ejecutable perfecta si en el futuro se decide reescribir la librería en Go, reduciendo drásticamente el riesgo de ese segundo paso.
3.  **Menor Riesgo de Ejecución**: Refactorizar es inherentemente menos arriesgado que reescribir un protocolo complejo y no documentado desde cero. Permite a la organización construir sobre lo que ya funciona mientras se pagan las deudas técnicas más urgentes.

Se recomienda seguir el plan de adaptación de .NET y reevaluar la necesidad de una reescritura en Go en 6-12 meses, una vez que se tenga una base de código estable, segura y robusta.
