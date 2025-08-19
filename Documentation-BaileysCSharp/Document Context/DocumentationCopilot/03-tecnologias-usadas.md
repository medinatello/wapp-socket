# Tecnologías Utilizadas - BaileysCSharp

## 📋 Stack Tecnológico Principal

| Tecnología | Versión | Propósito | Estado | Archivo Detallado |
|------------|---------|-----------|--------|-------------------|
| **.NET** | 9.0 | Runtime y Framework base | ✅ Actual | - |
| **LiteDB** | 5.0.20 | Base de datos NoSQL embebida | ✅ Bueno | [03a-sqlite.md](./03a-sqlite.md) |
| **WebSocket** | Nativo .NET | Comunicación en tiempo real | ⚠️ Básico | [03b-websocket.md](./03b-websocket.md) |
| **BouncyCastle** | 1.9.0 | Criptografía y protocolo Signal | ⚠️ Legacy | [03c-criptografia-y-claves.md](./03c-criptografia-y-claves.md) |
| **Protocol Buffers** | 3.27.0 | Serialización de mensajes | ✅ Bueno | [03d-protobuf.md](./03d-protobuf.md) |
| **FFMpegCore** | 5.1.0 | Procesamiento multimedia | ✅ Bueno | [03e-ffmpeg.md](./03e-ffmpeg.md) |
| **SkiaSharp** | 2.88.8 | Manipulación de imágenes | ✅ Bueno | [03f-skia.md](./03f-skia.md) |

## 🔍 Análisis por Tecnología

### 1. **.NET 9.0 Framework**

**Propósito**: Runtime base y framework de desarrollo

**Uso Actual**:
- Target framework principal
- Nullability warnings suprimidos masivamente
- Configuración específica para ARM64 macOS

**Problemas Detectados**:
- Supresión excesiva de warnings: `CS8632;CS8618;CS8600;CS8601;CS8602;CS8603;CS8604;CS8625;CS8629;CS8765`
- No aprovecha características modernas de .NET 8+

**Mejoras Sugeridas**:
- Habilitar nullable reference types correctamente
- Migrar a patrones modernos de async/await
- Usar System.Text.Json en lugar de Newtonsoft.Json

---

### 2. **LiteDB (NoSQL Database)**

**Propósito**: Persistencia de mensajes, contactos y metadatos

**Uso Actual**:
```csharp
public class MemoryStore {
    LiteDB.LiteDatabase database;
    Store<ChatModel> chats;
    Store<MessageModel> messages;
    Store<ContactModel> contacts;
    Store<GroupMetadataModel> groupMetaData;
}
```

**Ventajas**:
- Base de datos embebida sin servidor externo
- Soporte para BSON y consultas LINQ
- Transacciones ACID

**Ver detalles completos**: [03a-sqlite.md](./03a-sqlite.md)

---

### 3. **WebSocket Cliente**

**Propósito**: Conexión en tiempo real con servidores WhatsApp

**Implementación Actual**:
```csharp
public class WebSocketClient : AbstractSocketClient {
    ClientWebSocket WebSocket;
    // Framing personalizado para protocolo WhatsApp
    var messageSize = sizeBuffer[0] >> 16 | BitConverter.ToUInt16(sizeBuffer.Skip(1).Reverse().ToArray());
}
```

**Características**:
- Framing binario custom sobre WebSocket estándar
- Reconexión automática
- Keep-alive cada 30 segundos

**Ver detalles completos**: [03b-websocket.md](./03b-websocket.md)

---

### 4. **Portable.BouncyCastle (Criptografía)**

**Propósito**: Implementación del protocolo Signal para E2E encryption

**Componentes Criptográficos**:
- Curve25519 para intercambio de claves
- AES-256-CBC para cifrado simétrico
- HMAC-SHA256 para integridad
- HKDF para derivación de claves

**Problemas**:
- Dependencia legacy vs System.Security.Cryptography moderno
- Implementación custom de Curve25519

**Ver detalles completos**: [03c-criptografia-y-claves.md](./03c-criptografia-y-claves.md)

---

### 5. **Google.Protobuf**

**Propósito**: Serialización de mensajes del protocolo WhatsApp

**Archivos Proto**:
- `WAProto.proto` - Mensajes principales
- `WhisperTextProtocol.proto` - Protocolo Signal

**Configuración Especial**:
```xml
<!-- Protobuf generation for ARM64 macOS (.NET 9 compatible) -->
<Target Name="GenerateProtobufForArm64" BeforeTargets="BeforeBuild" 
        Condition="$([MSBuild]::IsOSPlatform('OSX')) AND '$([System.Runtime.InteropServices.RuntimeInformation]::ProcessArchitecture)' == 'Arm64'">
```

**Ver detalles completos**: [03d-protobuf.md](./03d-protobuf.md)

---

## 🔧 Dependencias de Desarrollo

### Testing Framework
```xml
<PackageReference Include="Microsoft.NET.Test.Sdk" Version="17.7.1" />
<PackageReference Include="NUnit" Version="3.13.3" />
<PackageReference Include="NUnit3TestAdapter" Version="4.4.2" />
<PackageReference Include="coverlet.collector" Version="3.2.0" />
```

### Herramientas de Build
```xml
<PackageReference Include="Grpc.Tools" Version="2.66.0">
  <PrivateAssets>all</PrivateAssets>
  <IncludeAssets>runtime; build; native; contentfiles; analyzers; buildtransitive</IncludeAssets>
</PackageReference>
```

## 📊 Matriz de Evaluación Tecnológica

| Tecnología | Rendimiento | Mantenibilidad | Comunidad | Futuro | Score Total |
|------------|-------------|----------------|-----------|---------|-------------|
| LiteDB | 8/10 | 9/10 | 7/10 | 8/10 | **32/40** |
| BouncyCastle | 7/10 | 6/10 | 8/10 | 6/10 | **27/40** |
| Protobuf | 9/10 | 8/10 | 9/10 | 9/10 | **35/40** |
| SkiaSharp | 8/10 | 8/10 | 8/10 | 8/10 | **32/40** |
| FFMpegCore | 9/10 | 7/10 | 7/10 | 8/10 | **31/40** |
| WebSocket (.NET) | 8/10 | 9/10 | 9/10 | 9/10 | **35/40** |

## 🚨 Dependencias Problemáticas

### 1. **Portable.BouncyCastle**
- **Problema**: Librería legacy, .NET moderno tiene cryptografía nativa
- **Impacto**: Rendimiento subóptimo, mayor superficie de ataque
- **Solución**: Migrar a System.Security.Cryptography

### 2. **Configuración de Protobuf Compleja**
- **Problema**: Lógica específica para ARM64 macOS
- **Impacto**: Build frágil en ciertas plataformas
- **Solución**: Simplificar usando herramientas modernas

### 3. **Nullability Warnings Suprimidos**
- **Problema**: `<NoWarn>CS8632;CS8618;CS8600;...`
- **Impacto**: Potenciales NullReferenceException en runtime
- **Solución**: Habilitar nullable reference types progresivamente

## 🔄 Plan de Modernización Tecnológica

### Fase 1: Seguridad (Prioridad Alta)
1. **Migrar BouncyCastle** → System.Security.Cryptography
2. **Habilitar nullable reference types**
3. **Actualizar patrones de async/await**

### Fase 2: Herramientas (Prioridad Media)
1. **Simplificar generación de Protobuf**
2. **Migrar a System.Text.Json**
3. **Actualizar dependencias menores**

### Fase 3: Optimización (Prioridad Baja)
1. **Evaluar alternativas a LiteDB** (si hay problemas de rendimiento)
2. **Optimizar uso de SkiaSharp**
3. **Considerar alternativas a FFMpegCore**

## 📈 Equivalencias .NET vs Go

Para referencia en evaluación de migración:

| .NET | Go | Observaciones |
|------|-----|---------------|
| LiteDB | bbolt/BadgerDB | Go tiene opciones más ligeras |
| System.Net.WebSockets | nhooyr.io/websocket | Go más simple |
| System.Security.Cryptography | crypto/* | Go crypto nativo excelente |
| Google.Protobuf | protobuf/proto | Soporte similar |
| SkiaSharp | No equivalente directo | Ventaja para .NET |
| FFMpegCore | ffmpeg-go | Funcionalidad similar |

## 🎯 Recomendaciones Específicas

### Para Mejora del Proyecto Actual:
1. **Priorizar migración de BouncyCastle** - Mayor impacto en seguridad
2. **Mantener LiteDB** - Funciona bien para el caso de uso
3. **Modernizar WebSocket handling** - Mejor manejo de errores y reconexión

### Para Evaluación de Reescritura:
1. **Go sería más ligero** - Especialmente para deployment
2. **.NET mantiene ventajas** - En multimedia y tooling
3. **Protobuf es neutro** - Funciona igual en ambos

---

**Conclusión**: El stack actual es funcional pero necesita modernización selectiva, especialmente en el área de criptografía y manejo de errores.
