# 10. Glosario y Referencias

Este documento proporciona definiciones para términos técnicos y de dominio utilizados a lo largo de la auditoría, junto con enlaces a recursos externos para una mayor profundización.

## 10.1. Glosario de Términos

-   **Adaptador (Adapter)**: En la Arquitectura Hexagonal, es un componente de infraestructura que implementa un Puerto. Se encarga de la comunicación entre el núcleo de la aplicación y el mundo exterior.

-   **Arquitectura Hexagonal (Hexagonal Architecture)**: Un patrón de diseño de software también conocido como Puertos y Adaptadores. Su objetivo es aislar la lógica de negocio de las dependencias de infraestructura.

-   **Backpressure**: Una estrategia para manejar la sobrecarga en sistemas de streaming de datos. Consiste en que un consumidor lento pueda notificar al productor que reduzca la velocidad de envío de datos para evitar ser desbordado.

-   **BouncyCastle**: Una librería de criptografía muy popular y completa para Java y C# (.NET).

-   **Deuda Técnica (Technical Debt)**: El costo implícito de retrabajo futuro causado por elegir una solución fácil o rápida en el presente en lugar de una mejor solución que llevaría más tiempo.

-   **DI (Dependency Injection / Inyección de Dependencias)**: Un patrón de diseño en el que un objeto recibe sus dependencias de una fuente externa en lugar de crearlas él mismo. Facilita el desacoplamiento y la testabilidad.

-   **Idempotencia**: La propiedad de una operación por la cual si se realiza múltiples veces, el resultado es el mismo que si se hubiera realizado una sola vez.

-   **JID (Jabber ID)**: El formato de dirección utilizado en el protocolo XMPP, sobre el que se construyó originalmente WhatsApp. Se sigue utilizando para identificar a usuarios (`<numero>@s.whatsapp.net`), grupos (`<id>@g.us`), etc.

-   **LiteDB**: Una base de datos NoSQL, embebida, de código abierto y escrita en .NET. Utiliza un único archivo para almacenar los datos.

-   **Noise Protocol Framework**: Un framework para construir protocolos criptográficos. El handshake de WhatsApp Web se basa en Noise para establecer un canal de comunicación seguro y cifrado.

-   **Observabilidad (Observability)**: La capacidad de un sistema para ser entendido desde el exterior a partir de los datos que genera (logs, métricas y trazas).

-   **OpenTelemetry (OTel)**: Un estándar y conjunto de herramientas de código abierto para la instrumentación, generación y recolección de datos de telemetría.

-   **Polly**: Una librería de resiliencia y manejo de fallos transitorios para .NET. Permite definir políticas de reintentos, circuit breaker, timeouts, etc.

-   **Puerto (Port)**: En la Arquitectura Hexagonal, es una interfaz definida en el núcleo de la aplicación que establece un contrato para la comunicación con la infraestructura externa.

-   **Protobuf (Protocol Buffers)**: Un formato de serialización de datos binarios desarrollado por Google. Es eficiente en tamaño y velocidad, y es utilizado por WhatsApp para la comunicación entre cliente y servidor.

-   **Signal Protocol**: El protocolo de cifrado de extremo a extremo desarrollado por Open Whisper Systems. Es el que utiliza WhatsApp para asegurar la confidencialidad e integridad de los mensajes.

## 10.2. Referencias y Enlaces

-   **Tecnologías Principales**:
    -   [.NET](https://dotnet.microsoft.com/)
    -   [Go](https://go.dev/)
    -   [LiteDB](https://www.litedb.org/)
    -   [BouncyCastle for C#](https://www.bouncycastle.org/csharp/)
    -   [Polly](http://www.thepollyproject.org/)
    -   [Serilog](https://serilog.net/)
    -   [OpenTelemetry](https://opentelemetry.io/)
    -   [BenchmarkDotNet](https://benchmarkdotnet.org/)

-   **Arquitectura y Patrones**:
    -   [Arquitectura Hexagonal (Alistair Cockburn)](https://alistair.cockburn.us/hexagonal-architecture/)
    -   [Martin Fowler - Inyección de Dependencias](https://martinfowler.com/articles/injection.html)

-   **Protocolos**:
    -   [Noise Protocol Framework](https://noiseprotocol.org/)
    -   [Signal Protocol](https://signal.org/docs/)
