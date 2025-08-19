# Gemini: Guía del Desarrollador para BaileysCSharp

## 1. Descripción General

Este proyecto, `BaileysCSharp`, es una biblioteca en C# (.NET 9) diseñada para interactuar con la API web no oficial de WhatsApp. Su nombre sugiere que es un port o está fuertemente inspirado en la popular biblioteca de JavaScript [Baileys](https://github.com/WhiskeySockets/Baileys).

El objetivo principal es proporcionar una interfaz programática para automatizar acciones de WhatsApp, como enviar y recibir mensajes, gestionar grupos y manejar el estado de la conexión, todo desde un entorno .NET.

## 2. Arquitectura y Diseño

La arquitectura es asincrónica y orientada a eventos, clave para manejar las comunicaciones en tiempo real con los servidores de WhatsApp.

### 2.1. Arquitectura Orientada a Eventos
El núcleo (`BaileysCSharp/Core/Events`) utiliza un `EventEmitter`. En lugar de sondear en busca de nuevos mensajes, la biblioteca emite eventos a los que el código cliente se puede suscribir.

**Ejemplo conceptual:**
```csharp
// Suscribirse a un evento de nuevo mensaje
socket.Events.Message.OnNewMessage += (message) => {
    Console.WriteLine($"Nuevo mensaje de {message.Key.RemoteJid}: {message.Message.Conversation}");
};

// Suscribirse a un evento de conexión exitosa
socket.Events.Connection.OnOpen += (args) => {
    Console.WriteLine("¡Conexión establecida!");
};
```

### 2.2. Comunicación y Protocolo
- **WebSockets:** La comunicación se mantiene a través de una conexión WebSocket persistente (`BaileysCSharp/Core/Sockets/WASocket.cs`).
- **Protocolo Binario (`WABinary`):** WhatsApp utiliza un formato binario personalizado, no JSON/XML. La clase `BinaryNode.cs` es crucial, ya que representa los "nodos" de datos intercambiados.

### 2.3. Cifrado de Extremo a Extremo (E2EE)
- **Protocolo Signal:** El cifrado se basa en el Protocolo Signal.
- **`LibSignal/` y `Core/Signal/`:** Contienen la implementación criptográfica. `LibSignal` tiene las primitivas del protocolo, y `Core/Signal` lo integra con el resto de la biblioteca.
- **Gestión de Sesiones:** La biblioteca gestiona de forma transparente las claves de sesión, pre-claves y el cifrado/descifrado de mensajes (`SessionCipher.cs`, `GroupCipher.cs`).

### 2.4. Gestión de Estado y Autenticación
- **`Core/NoSQL/`:** La biblioteca necesita persistir el estado de la autenticación (credenciales, claves) para reanudar sesiones sin escanear el código QR cada vez.
- **`FileKeyStore.cs`:** Implementación que guarda el estado en archivos locales (JSON).
- **`MemoryStore.cs`:** Implementación en memoria para sesiones efímeras o pruebas.

## 3. Dependencias Clave y Tecnologías

- **.NET 9:** El proyecto está construido sobre la última versión de .NET.
- **Google.Protobuf:** WhatsApp utiliza Protocol Buffers para la serialización de datos. Los archivos `.proto` en `BaileysCSharp/Proto/` definen las estructuras de los mensajes.
- **LiteDB:** Para el almacenamiento de estado en `FileKeyStore`, se usa esta base de datos NoSQL embebida.
- **Portable.BouncyCastle:** Proporciona las primitivas criptográficas necesarias para el Protocolo Signal.
- **QRCoder:** Utilizado en la aplicación de ejemplo (`WhatsSocketConsole`) para generar el código QR para la autenticación.
- **NUnit:** El framework utilizado para las pruebas unitarias en `BaileysCSharp.Tests`.

## 4. Estructura del Proyecto

- **`BaileysCSharp.sln`**: La solución de Visual Studio que agrupa todos los proyectos.
- **`BaileysCSharp/`**: El proyecto principal de la biblioteca.
  - `Core/`: El corazón de la biblioteca (sockets, eventos, cifrado, etc.).
  - `Proto/`: Definiciones de Protocol Buffers. `WAProto.cs` y `WhisperTextProtocol.cs` son **archivos generados** a partir de los `.proto`.
- **`BaileysCSharp.Tests/`**: Proyecto de pruebas unitarias. **Esencial para entender el uso práctico de la biblioteca.**
- **`WhatsSocketConsole/`**: Una aplicación de consola de ejemplo. **El mejor punto de partida para un nuevo desarrollador.**

## 5. Cómo Compilar y Ejecutar

### 5.1. Prerrequisitos
- [.NET 9 SDK](https://dotnet.microsoft.com/download/dotnet/9.0)
- (Opcional, para macOS en ARM64) `protoc` (compilador de Protocol Buffers). Se puede instalar con Homebrew: `brew install protobuf`.

### 5.2. Compilar la Solución
Puedes compilar todo el proyecto usando el siguiente comando desde el directorio raíz:
```bash
dotnet build
```
**Nota sobre Protobuf en macOS (ARM64):** El archivo `BaileysCSharp.csproj` tiene una lógica de compilación condicional. Si estás en macOS con un chip Apple Silicon (ARM64), intentará usar el `protoc` del sistema para generar los archivos de C# a partir de los `.proto`. En otros sistemas operativos, usará el paquete `Grpc.Tools`.

### 5.3. Ejecutar la Aplicación de Ejemplo
La aplicación de consola es la mejor forma de ver la biblioteca en acción.
1.  Navega al directorio del proyecto de consola: `cd WhatsSocketConsole`
2.  Ejecuta la aplicación: `dotnet run`

La primera vez, generará un código QR en la consola. Escanéalo con tu teléfono desde `WhatsApp > Dispositivos Vinculados > Vincular un dispositivo`. Esto creará un archivo `creds.json` en `WhatsSocketConsole/CreateSession`, que se usará para las siguientes sesiones.

### 5.4. Ejecutar las Pruebas
Las pruebas unitarias son cruciales para verificar la funcionalidad.
```bash
dotnet test
```
Esto ejecutará todas las pruebas definidas en el proyecto `BaileysCSharp.Tests`.

## 6. Guía de Inicio Rápido para Desarrolladores

1.  **Analiza `WhatsSocketConsole/Program.cs`:** Este archivo es tu mejor amigo. Muestra el ciclo de vida completo:
    -   Instanciación y configuración del `WASocket`.
    -   Carga o creación de una sesión de autenticación.
    -   Suscripción a los eventos clave (`OnOpen`, `OnNewMessage`, etc.).
    -   Conexión al servidor de WhatsApp.
    -   Manejo de la lógica de la aplicación una vez conectado.
2.  **Explora las Pruebas (`BaileysCSharp.Tests/`):** Las pruebas ofrecen ejemplos concisos y aislados de funcionalidades específicas, como el envío de diferentes tipos de mensajes o la gestión de grupos.
3.  **Modifica y Experimenta:** Intenta modificar `Program.cs` para enviar un mensaje a tu propio número o para reaccionar a mensajes entrantes.

## 7. Puntos Clave para Contribuir o Corregir Errores

- **El Protocolo de WhatsApp Cambia:** WhatsApp actualiza su API web con frecuencia y sin previo aviso. Una tarea de mantenimiento común es adaptar la biblioteca a estos cambios, a menudo analizando la biblioteca original de Baileys JS.
- **La Criptografía es Delicada:** Cambios en `LibSignal/` o `Core/Signal/` deben hacerse con extremo cuidado y un profundo conocimiento del Protocolo Signal.
- **Depuración de Nodos Binarios:** Si un mensaje no se envía o recibe, el problema probablemente esté en la construcción/interpretación de un `BinaryNode`.
- **Gestión de Estado:** Asegúrate de que cualquier nueva información persistente se añada correctamente al `AuthenticationState` y se guarde a través del sistema `NoSQL`.