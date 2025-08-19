# 2. Análisis de la Arquitectura Actual

## 2.1. Descripción General

La arquitectura de `BaileysCSharp` se centra en la clase `WASocket`, que actúa como una fachada (Facade) para orquestar todas las operaciones. El diseño es fuertemente orientado a eventos, utilizando un `EventEmitter` para comunicar los cambios de estado y los datos recibidos (mensajes, recibos, etc.) al código cliente.

El sistema se puede dividir en las siguientes capas lógicas:

1.  **Capa de Aplicación (`WhatsSocketConsole`)**: El consumidor final de la librería. Es responsable de instanciar, configurar y suscribirse a los eventos de `WASocket`.
2.  **Capa de Orquestación (`WASocket`)**: El punto de entrada principal y fachada de la librería. Gestiona el ciclo de vida de la conexión, la autenticación y expone métodos para las acciones del usuario (enviar mensaje, descargar medios, etc.).
3.  **Capa de Comunicación (WebSocket)**: Gestiona la conexión en tiempo real con los servidores de WhatsApp. Utiliza la implementación nativa de .NET (`System.Net.WebSockets`). El manejo del protocolo de WhatsApp (Noise Protocol para el handshake, frames binarios, etc.) se realiza sobre esta capa.
4.  **Capa de Lógica de Negocio (Handlers y Utils)**: Un conjunto de clases de ayuda y manejadores que procesan los nodos binarios recibidos del WebSocket, los decodifican y los convierten en eventos de dominio.
5.  **Capa de Persistencia (`FileKeyStore`, `LiteDB`)**: Responsable de almacenar el estado de la sesión, incluyendo las credenciales de autenticación y las claves de cifrado. Utiliza `LiteDB` como motor de base de datos NoSQL embebido y archivos planos para otros datos.
6.  **Capa de Criptografía (`BouncyCastle`, `LibSignal`)**: Implementa el protocolo Signal para el cifrado de extremo a extremo de los mensajes.

## 2.2. Diagrama de Componentes (Mermaid)

```mermaid
graph TD
    subgraph "Capa de Aplicación"
        A[WhatsSocketConsole]
    end

    subgraph "Librería BaileysCSharp"
        B(WASocket - Fachada)
        C{EventEmitter (EV)}
        D[Cliente WebSocket]
        E[Handlers & Utils]
        F[Capa de Persistencia]
        G[Capa de Criptografía]
    end

    subgraph "Dependencias Externas"
        H[Servidores WhatsApp]
        I[LiteDB]
        J[BouncyCastle]
        K[Sistema de Archivos]
    end

    A -->|1. Configura e Instancia| B
    B -->|2. Se suscribe a| C
    B -->|3. Inicia Conexión| D
    D -->|4. Comunica con| H
    H -->|5. Envía datos binarios| D
    D -->|6. Pasa datos a| E
    E -->|7. Decodifica y Procesa| B
    E -->|8. Usa Cripto| G
    B -->|9. Emite Evento| C
    C -->|10. Notifica a| A

    B -->|Almacena/Recupera Sesión| F
    F -->|Usa| I
    F -->|Usa| K
    G -->|Usa| J
```

## 2.3. Mapa de Dependencias y Acoplamiento

### Dependencias Internas

-   `WhatsSocketConsole` depende directamente de `BaileysCSharp`.
-   `WASocket` es el componente más acoplado. Depende directamente de la capa de sockets, la de persistencia, la de criptografía y la de eventos.
-   Los `Handlers` y `Utils` tienen dependencias cruzadas, lo que puede dificultar el mantenimiento.
-   La capa de eventos (`EventEmitter`) está diseñada para ser un punto de desacoplamiento, pero su implementación monolítica (una única fuente de eventos `EV`) puede convertirse en un cuello de botella.

### Dependencias Externas

-   **`LiteDB`**: Acoplamiento fuerte en la capa de persistencia (`BaseKeyStore`, `FileKeyStore`). Cambiar a otra base de datos requeriría implementar una nueva serie de clases de almacenamiento.
-   **`BouncyCastle`**: Acoplamiento fuerte en toda la lógica de criptografía. Es una dependencia fundamental para el protocolo Signal.
-   **`Google.Protobuf`**: Esencial para la serialización/deserialización de los mensajes definidos en los archivos `.proto`.
-   **`System.Net.WebSockets`**: Dependencia a nivel de framework para la comunicación. Es una abstracción sólida, por lo que el riesgo es bajo.

## 2.4. Puntos Críticos y Dificultades de Cambio

1.  **Gestión de Estado en `WASocket`**: La clase `WASocket` parece ser un "God Object" que centraliza demasiada responsabilidad (estado de la conexión, configuración, lógica de envío, etc.). Esto la hace difícil de testear y modificar.
2.  **Lógica de Persistencia**: El acoplamiento directo con `LiteDB` y el sistema de archivos en la lógica de negocio dificulta la prueba unitaria de componentes que necesitan estado. No parece haber una interfaz clara que abstraiga el origen de los datos.
3.  **Manejo de Errores**: El manejo de errores parece disperso. En `Program.cs` se ve un `try-catch` para la reconexión, pero no está claro si existe una estrategia de resiliencia (reintentos, backoff exponencial) centralizada y configurable.
4.  **Flujo de Procesamiento de Mensajes**: El flujo desde la recepción de bytes en el socket hasta la emisión de un evento de "mensaje recibido" atraviesa múltiples clases y métodos estáticos (`MessageDecoder`, `ProcessMessageUtil`). Trazar y depurar este flujo es complejo.
5.  **Testabilidad**: La falta de inversión de dependencias (DI) y el uso de componentes estáticos hacen que la librería sea muy difícil de probar unitariamente. Los tests existentes parecen ser más de integración o de extremo a extremo, dependiendo de credenciales reales.
