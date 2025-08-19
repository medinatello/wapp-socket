# BaileysCSharp - Documentación Técnica del Proyecto

## 🌟 Descripción General

**BaileysCSharp** es una implementación en C# del protocolo de WhatsApp Web, inspirada en la biblioteca Baileys de Node.js. Este proyecto permite crear bots y aplicaciones de WhatsApp utilizando el protocolo oficial sin necesidad de usar WhatsApp Business API.

### 📋 Información Básica
- **Lenguaje**: C# (.NET 9.0)
- **Tipo de Proyecto**: Biblioteca de cliente de WhatsApp Web
- **Arquitectura**: Multi-dispositivo compatible
- **Protocolo**: WebSocket con cifrado end-to-end

## 🏗️ Estructura del Proyecto

```
BaileysCSharp/
├── BaileysCSharp/              # 📚 Biblioteca principal
├── BaileysCSharp.Tests/        # 🧪 Pruebas unitarias (NUnit)
├── WhatsSocketConsole/         # 🎯 Aplicación de ejemplo/demostración
└── BaileysCSharp.sln          # 📁 Solución de Visual Studio
```

### 📦 Proyectos en la Solución

1. **BaileysCSharp** - Biblioteca principal (Library)
2. **BaileysCSharp.Tests** - Proyecto de pruebas (Test Project)
3. **WhatsSocketConsole** - Aplicación de consola de ejemplo (Console App)

## 🔧 Arquitectura del Sistema

### 🎨 Patrones de Diseño Utilizados

1. **Event-Driven Architecture**: Sistema basado en eventos para manejo de mensajes y estados
2. **Repository Pattern**: Para almacenamiento de datos y sesiones
3. **Socket Pattern**: Comunicación WebSocket con WhatsApp
4. **Builder Pattern**: Configuración de sockets y autenticación

### 🗂️ Organización del Código (BaileysCSharp/)

```
Core/
├── Converters/         # 🔄 Convertidores de datos (JSON, Buffer)
├── Events/            # 📡 Sistema de eventos y stores
├── Exceptions/        # ⚠️ Excepciones personalizadas
├── Extensions/        # 🔧 Métodos de extensión
├── Helper/           # 🛠️ Utilidades y helpers
├── Logging/          # 📝 Sistema de logging
├── Models/           # 📊 Modelos de datos
├── NoSQL/            # 💾 Almacenamiento (LiteDB)
├── Signal/           # 🔐 Protocolo Signal (cifrado)
├── Sockets/          # 🌐 Implementación de sockets
├── Stores/           # 🗄️ Almacenes de datos
├── Types/            # 📋 Definiciones de tipos
├── Utils/            # ⚙️ Utilidades generales
└── WABinary/         # 📦 Protocolo binario de WhatsApp

LibSignal/            # 🔒 Implementación del protocolo Signal
Proto/                # 📜 Archivos Protocol Buffers
Resources/            # 📁 Recursos (mimeTypes.json)
```

### 🏛️ Componentes Principales

#### 1. **WASocket** (`Core/Sockets/WASocket.cs:8`)
- **Función**: Clase principal para conexión con WhatsApp
- **Herencia**: `WASocket` → `NewsletterSocket` → `GroupSocket` → ... → `BaseSocket`
- **Responsabilidad**: Punto de entrada principal para todas las operaciones

#### 2. **EventEmitter** (`Core/Events/EventEmitter.cs:24`)
- **Función**: Sistema central de eventos
- **Eventos Disponibles**:
  - `Connection` - Estado de conexión
  - `Auth` - Autenticación
  - `Message` - Mensajes entrantes/salientes
  - `Pressence` - Estado de presencia
  - `Group` - Eventos de grupos
  - `Newsletter` - Newsletters/Canales
  - Y más...

#### 3. **SocketConfig** (`Core/Types/SocketConfig.cs:18`)
- **Función**: Configuración del socket
- **Parámetros clave**:
  - `Version`: Versión del protocolo WhatsApp
  - `Browser`: Información del navegador simulado
  - `Auth`: Estado de autenticación
  - `SessionName`: Nombre de la sesión
  - `Logger`: Configuración de logging

#### 4. **NoiseHandler** (`Core/Utils/NoiseHandler.cs:13`)
- **Función**: Manejo del protocolo Noise para cifrado
- **Responsabilidad**: Establecer conexión segura con WhatsApp

## 🔌 Dependencias Principales

### 📚 Librerías de Producción (BaileysCSharp.csproj)

| Paquete | Versión | Propósito |
|---------|---------|-----------|
| **Google.Protobuf** | 3.27.0 | Serialización de mensajes Protocol Buffers |
| **Grpc.Tools** | 2.66.0 | Herramientas para compilar archivos .proto |
| **LiteDB** | 5.0.20 | Base de datos NoSQL para almacenamiento local |
| **Portable.BouncyCastle** | 1.9.0 | Criptografía y cifrado (Signal Protocol) |
| **SkiaSharp** | 2.88.8 | Procesamiento de imágenes |
| **FFMpegCore** | 5.1.0 | Procesamiento de audio/video |

### 🧪 Librerías de Desarrollo y Testing

| Paquete | Propósito |
|---------|-----------|
| **NUnit** | Framework de pruebas unitarias |
| **Microsoft.NET.Test.Sdk** | SDK de testing de .NET |
| **QRCoder** (Console) | Generación de códigos QR |

## 🔒 Sistema de Autenticación y Seguridad

### 🗝️ Componentes de Seguridad

1. **AuthenticationCreds**: Credenciales de autenticación
2. **SignalRepository**: Implementación del protocolo Signal
3. **KeyStore**: Almacenamiento seguro de claves
4. **NoiseHandler**: Protocolo Noise para handshake inicial

### 💾 Almacenamiento de Sesión

- **Archivo**: `{SessionName}/creds.json`
- **Base de Datos**: LiteDB para almacenar mensajes y metadatos
- **Claves**: Sistema de claves distribuido para cifrado E2E

## 📨 Sistema de Mensajes

### 📝 Tipos de Mensajes Soportados

#### Mensajes de Texto
- `TextMessageContent` - Mensajes de texto básicos
- `ExtendedTextMessage` - Texto con funciones extendidas

#### Mensajes Multimedia
- `ImageMessageContent` - Imágenes (JPEG, PNG, WebP)
- `VideoMessageContent` - Videos (MP4, GIF)
- `AudioMessageContent` - Audio (MP3, OGG, AAC)
- `DocumentMessageContent` - Documentos (PDF, DOC, etc.)

#### Mensajes Especiales
- `ContactMessageContent` - Compartir contactos
- `LocationMessageContent` - Ubicaciones
- `ReactMessageContent` - Reacciones a mensajes
- `DeleteMessageContent` - Eliminar mensajes

### 🎯 Eventos de Mensajes

```csharp
// Ejemplo de manejo de eventos
socket.EV.Message.Upsert += (sender, e) => {
    if (e.Type == MessageEventType.Notify) {
        // Nuevos mensajes
    }
    if (e.Type == MessageEventType.Append) {
        // Mensajes sincronizados offline
    }
};
```

## 👥 Funcionalidades de Grupos

### 🔧 Operaciones de Grupo

- **Creación**: `GroupCreate(nombre, participantes)`
- **Metadatos**: `GroupUpdateSubject`, `GroupUpdateDescription`
- **Participantes**: `GroupParticipantsUpdate` (agregar/remover/promover)
- **Configuración**: `GroupSettingUpdate` (anuncios, permisos)
- **Invitaciones**: `GroupInviteCode`, `GroupGetInviteInfo`

## 📢 Sistema de Newsletters/Canales

- **Consulta**: `QueryRecommendedNewsletters`
- **Suscripción**: `NewsletterFollow`/`NewsletterUnFollow`
- **Gestión**: `NewsletterCreate`, `NewsletterDelete`
- **Envío**: `SendNewsletterMessage`

## 🔧 Configuración y Uso

### ⚙️ Configuración Básica

```csharp
var config = new SocketConfig() {
    SessionName = "mi-sesion",
    Version = [2, 3000, 1023223821],
    Logger = { Level = LogLevel.Info }
};

var socket = new WASocket(config);
```

### 🚀 Inicialización

```csharp
// Configurar autenticación
config.Auth = new AuthenticationState() {
    Creds = authentication,
    Keys = new FileKeyStore(config.CacheRoot)
};

// Suscribirse a eventos
socket.EV.Connection.Update += Connection_Update;
socket.EV.Message.Upsert += Message_Upsert;

// Iniciar conexión
socket.MakeSocket();
```

## 🐛 Manejo de Errores

### ⚠️ Excepciones Personalizadas

- `Boom` - Error de conexión con WhatsApp
- `SessionException` - Errores de sesión
- `GroupCipherException` - Errores de cifrado de grupo

### 🔄 Reconexión Automática

```csharp
if (connection.Connection == WAConnectionState.Close) {
    if (connection.LastDisconnect.Error is Boom boom && 
        boom.Data?.StatusCode != (int)DisconnectReason.LoggedOut) {
        // Reconectar automáticamente
        socket.MakeSocket();
    }
}
```

## 🛠️ Herramientas de Desarrollo

### 📋 Comandos Útiles

```bash
# Compilar el proyecto
dotnet build

# Ejecutar pruebas
dotnet test

# Ejecutar aplicación de ejemplo
dotnet run --project WhatsSocketConsole
```

### 🔍 Debugging y Logging

- **LogLevel.Raw**: Todos los mensajes (muy verbose)
- **LogLevel.Trace**: Información detallada
- **LogLevel.Info**: Información general
- **LogLevel.Warn**: Advertencias
- **LogLevel.Error**: Solo errores

## 🚨 Consideraciones Especiales

### 🍎 Configuración para macOS ARM64

El proyecto incluye configuración especial para Apple Silicon:
- Generación manual de archivos Protocol Buffers
- Uso de `protoc` del sistema en lugar del paquete NuGet

### 🔐 Seguridad

- **No commitear** archivos de credenciales (`creds.json`, `session/`)
- Las claves se almacenan cifradas localmente
- Implementación completa del protocolo Signal para E2E encryption

### ⚡ Rendimiento

- Uso de `ConcurrentDictionary` para thread safety
- Manejo eficiente de memoria con `ArrayPool`
- Almacenamiento optimizado con LiteDB

## 🔄 Flujo de Trabajo de Desarrollo

### 🎯 Para Agregar Nueva Funcionalidad

1. **Analizar** el protocolo de WhatsApp requerido
2. **Implementar** en la capa de Socket correspondiente
3. **Agregar** eventos necesarios en EventEmitter
4. **Crear** modelos de datos en `Core/Models/`
5. **Añadir** pruebas en `BaileysCSharp.Tests/`
6. **Documentar** en este archivo

### 🔧 Para Corregir Bugs

1. **Reproducir** el error en `WhatsSocketConsole`
2. **Escribir** prueba que falle
3. **Implementar** la corrección
4. **Verificar** que la prueba pase
5. **Probar** en la aplicación de consola

## 📚 Recursos Adicionales

- **Protocol Buffers**: `Proto/WAProto.proto` - Definiciones del protocolo
- **Tipos MIME**: `Resources/mimeTypes.json` - Mapeo de tipos de archivo
- **Navegadores**: `Core/Utils/Browsers.cs` - User agents simulados

---

*Este documento debe mantenerse actualizado conforme evolucione el proyecto. Para contribuir, asegúrate de entender la arquitectura descrita y seguir los patrones establecidos.*