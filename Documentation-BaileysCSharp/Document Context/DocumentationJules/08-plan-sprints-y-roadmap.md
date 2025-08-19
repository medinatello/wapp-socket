# 8. Roadmap y Plan de Sprints

Este documento presenta dos posibles roadmaps de alto nivel, uno para cada variante estratégica discutida en la evaluación: (A) mejorar la base de código .NET existente, y (B) reescribir el proyecto en Go.

---

## Variante A: Mejorar y Adaptar el Proyecto .NET

Este plan se enfoca en refactorizar incrementalmente la librería actual para mejorar su robustez, testabilidad y mantenibilidad. Se estima una duración de 6 sprints de 2 semanas cada uno.

### Roadmap General (Adaptar .NET)
-   **Q1**: Fundamentos de Arquitectura y Resiliencia.
-   **Q2**: Observabilidad, Seguridad y Optimización.

### Plan de Sprints (Adaptar .NET)

**Sprint 1: Fundamentos de Arquitectura Hexagonal**
-   **Objetivo**: Introducir las interfaces (Puertos) para los servicios clave.
-   **Entregables**:
    -   Crear proyecto `BaileysCSharp.Core` con las interfaces: `IMessageRepository`, `IAuthCredentialStore`, `IWebSocketTransport`, `ICryptoService`.
    -   Refactorizar `WASocket` para que dependa de estas interfaces (sin implementación aún).
-   **Criterio de Aceptación**: El proyecto compila y las dependencias del núcleo apuntan hacia adentro.

**Sprint 2: Implementación de Adaptadores**
-   **Objetivo**: Mover la lógica de infraestructura a clases de Adaptadores.
-   **Entregables**:
    -   Crear adaptadores `LiteDBAdapter`, `WebSocketAdapter`, `BouncyCastleAdapter`.
    -   Configurar Inyección de Dependencias para ensamblar la aplicación.
-   **Criterio de Aceptación**: La aplicación es funcional usando la nueva arquitectura. Los tests de integración existentes pasan.

**Sprint 3: Resiliencia y Reconexión Automática**
-   **Objetivo**: Hacer la librería autónoma en su gestión de conexión.
-   **Entregables**:
    -   Integrar `Polly` en el `WebSocketAdapter`.
    -   Implementar una política de reintentos con backoff exponencial para la conexión.
    -   Eliminar la lógica de reconexión del código cliente (`WhatsSocketConsole`).
-   **Criterio de Aceptación**: La librería se reconecta automáticamente tras una pérdida de red.

**Sprint 4: Observabilidad (Logging y Telemetría)**
-   **Objetivo**: Instrumentar la librería para facilitar la depuración y el monitoreo.
-   **Entregables**:
    -   Reemplazar el logger personalizado por `Serilog` y `Microsoft.Extensions.Logging`.
    -   Integrar `OpenTelemetry` para trazas y métricas clave (conexión, envío/recepción de mensajes).
-   **Criterio de Aceptación**: Se pueden visualizar trazas en Jaeger/consola y los logs son estructurados.

**Sprint 5: Endurecimiento de la Seguridad**
-   **Objetivo**: Eliminar el riesgo de almacenamiento de credenciales en texto plano.
-   **Entregables**:
    -   Refactorizar el `AuthCredentialStore` para usar `ProtectedData` (DPAPI) para cifrar y descifrar `creds.json`.
-   **Criterio de Aceptación**: El archivo `creds.json` en disco está cifrado. La reanudación de sesión sigue funcionando.

**Sprint 6: Benchmarking y Optimización**
-   **Objetivo**: Medir y optimizar los hotspots de rendimiento.
-   **Entregables**:
    -   Crear una suite de benchmarks con `BenchmarkDotNet`.
    -   Implementar `ArrayPool` para la gestión de buffers.
    -   Realizar una ronda de optimizaciones basada en los resultados.
-   **Criterio de Aceptación**: Se entregan resultados de benchmarks que demuestran una mejora en el rendimiento.

---

## Variante B: Reescritura en Go

Este plan se enfoca en crear una nueva librería en Go desde cero, con un MVP funcional como primer hito.

### Roadmap General (Reescribir en Go)
-   **Mes 1**: Conexión y Autenticación.
-   **Mes 2**: MVP Funcional (Mensajería de Texto E2E).
-   **Mes 3**: Paridad de Funcionalidades y Hardening.

### Plan de Sprints (Reescribir en Go)

**Sprint 1: Conexión y Handshake Criptográfico**
-   **Objetivo**: Establecer una conexión WebSocket y completar el handshake de Noise.
-   **Entregables**:
    -   Cliente WebSocket básico (`nhooyr/websocket`).
    -   Implementación del handshake con la librería `crypto` nativa.
-   **Criterio de Aceptación**: Se establece un canal de comunicación cifrado con el servidor.

**Sprint 2: Autenticación con QR y Gestión de Sesión**
-   **Objetivo**: Implementar el flujo de login con QR y persistir la sesión.
-   **Entregables**:
    -   Lógica para solicitar y procesar el QR.
    -   Almacén de credenciales con `bbolt` o `modernc.org/sqlite`, cifrado con una clave maestra.
-   **Criterio de Aceptación**: Un nuevo cliente puede registrarse y un cliente existente puede reanudar la sesión.

**Sprint 3: MVP - Mensajería de Texto E2E**
-   **Objetivo**: Enviar y recibir mensajes de texto cifrados.
-   **Entregables**:
    -   Implementación del protocolo Signal (cifrado/descifrado de sesión).
    -   API pública para `SendMessage` y `OnMessageReceived`.
-   **Criterio de Aceptación**: Dos instancias de la librería pueden intercambiar mensajes de texto.

**Sprint 4: Sincronización de Historial y Eventos**
-   **Objetivo**: Procesar correctamente los eventos de sincronización y recibos.
-   **Entregables**:
    -   Lógica para manejar el historial de chats, recibos de entrega/lectura.
    -   API de eventos robusta.
-   **Criterio de Aceptación**: El estado de los mensajes (enviado, entregado, leído) se actualiza correctamente.

**Sprint 5: Soporte para Medios**
-   **Objetivo**: Enviar y recibir archivos multimedia (imágenes, audio).
-   **Entregables**:
    -   Lógica para cifrar y subir el blob del medio al servidor de WhatsApp.
    -   Lógica para descargar y descifrar medios.
-   **Criterio de Aceptación**: Se pueden enviar y recibir imágenes y documentos.

**Sprint 6: Pruebas de Integración, Benchmarking y Hardening**
-   **Objetivo**: Asegurar la calidad y robustez de la librería.
-   **Entregables**:
    -   Suite de pruebas de integración completa.
    -   Benchmarks de rendimiento.
    -   Revisión de seguridad y documentación final.
-   **Criterio de Aceptación**: La librería es estable, documentada y tiene una cobertura de pruebas > 80%.
