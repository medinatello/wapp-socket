# 3. Análisis de Tecnologías y Dependencias

## 3.1. Resumen de Tecnologías

La siguiente tabla resume las principales librerías y tecnologías de terceros utilizadas en el proyecto `BaileysCSharp`. Cada tecnología es un pilar fundamental para el funcionamiento de la librería.

| Tecnología / Librería       | Propósito en el Proyecto                                     | Problemas Detectados / Riesgos                                                              | Mejoras Sugeridas                                                                      | Análisis Detallado                                           |
| --------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| **LiteDB**                  | Base de datos NoSQL embebida para la persistencia de la sesión (claves, credenciales, etc.). | Acoplamiento directo, ausencia de un repositorio abstracto, posible contención en escrituras. | Abstraer la persistencia detrás de una interfaz (Repositorio), evaluar modo WAL si aplica. | [Ver Análisis de LiteDB](./03a-litedb.md)                    |
| **System.Net.WebSockets**   | Cliente nativo de .NET para la comunicación WebSocket con los servidores de WhatsApp. | La lógica de reconexión y gestión de estado está implementada manualmente en `WASocket`.     | Usar una librería de resiliencia como Polly para gestionar reintentos con backoff.     | [Ver Análisis de WebSocket](./03b-websocket.md)              |
| **BouncyCastle**            | Librería de criptografía para implementar el protocolo Signal (cifrado E2E). | Complejidad inherente de la API, dependencia crítica y monolítica.                        | Envolver la lógica cripto en una interfaz clara para aislar su complejidad.            | [Ver Análisis de Criptografía](./03c-criptografia-y-claves.md) |
| **Google.Protobuf**         | Serialización y deserialización de mensajes usando Protocol Buffers. | Los archivos `.proto` son la "verdad" del contrato. Cambios en el protocolo de WA rompen la librería. | Mantener un proceso claro para actualizar los `.proto` y regenerar el código.           | No requiere análisis separado.                               |
| **FFMpegCore** / **SkiaSharp** | Procesamiento de archivos multimedia (generación de miniaturas, conversión de formatos). | Dependencias externas pesadas (requieren FFMpeg en el sistema). Complejidad en el manejo de errores. | Aislar la lógica de medios en un servicio aparte. Implementar colas y reintentos para fallos. | No requiere análisis separado por ahora.                     |
| **QRCoder** (en Consola)    | Generación de códigos QR para el proceso de login.           | Usado solo en la app de consola para visualización. No es una dependencia del núcleo.       | Ninguna. Su uso es adecuado para el propósito.                                         | No requiere análisis separado.                               |

## 3.2. Dependencias del Framework .NET

-   **Target Framework**: `net9.0`. Es una versión muy reciente, lo que implica que se beneficia de las últimas mejoras de rendimiento y características del lenguaje C#. Sin embargo, podría presentar problemas de compatibilidad si se necesita ejecutar en entornos más antiguos.
-   **ASP.NET Core / Hosting**: No se utilizan. La librería es autocontenida y no depende del stack de ASP.NET, lo que la hace ligera.
-   **Inyección de Dependencias**: No se utiliza `Microsoft.Extensions.DependencyInjection`. La creación de objetos es manual, lo que lleva a un alto acoplamiento como se describió en la sección de arquitectura.

## 3.3. Consideraciones sobre Licenciamiento

-   **Baileys (Original)**: Licencia MIT. Permisiva.
-   **LiteDB**: Licencia MIT. Permisiva.
-   **BouncyCastle**: Licencia MIT. Permisiva.
-   **Google.Protobuf**: Licencia BSD-3-Clause. Permisiva.
-   **FFMpegCore**: Licencia LGPL. Requiere que `FFMpegCore` se use como una librería dinámica y no se modifique su código fuente si el proyecto se distribuye.
-   **SkiaSharp**: Licencia MIT. Permisiva.

**Conclusión de Licenciamiento**: No se han detectado conflictos de licencia graves ni dependencias "copyleft" que obliguen a abrir el código fuente del proyecto que la utilice. El uso de FFMpeg a través de `FFMpegCore` es el único punto a tener en cuenta, pero su uso como librería dinámica cumple con los requisitos de la LGPL.
