# 03. Tecnologías unificadas

La librería utiliza varias tecnologías para cumplir sus funciones.  La tabla siguiente resume su propósito, problemas detectados, mejoras sugeridas y alternativas en Go (en caso de una reescritura).

| Tecnología                | Propósito en BaileysCSharp                                        | Problemas / Consideraciones              | Mejora o alternativa (.NET)                        | Alternativa en Go                       |
|--------------------------|-------------------------------------------------------------------|-----------------------------------------|----------------------------------------------------|-----------------------------------------|
| **SQLite**               | Persistir credenciales, sesiones y estados de conexión            | Faltan migraciones; bloqueo concurrente; credenciales en claro | Implementar esquema con migraciones y WAL; cifrar archivos usando `Microsoft.Data.Sqlite` con cifrado; evaluar LiteDB | Uso de [bbolt](https://pkg.go.dev/go.etcd.io/bbolt) o `modernc.org/sqlite`; cifrado con `sqlcipher` |
| **WebSocket**            | Mantener conexión con WhatsApp Web                                | Reconexión básica; no gestiona backpressure | Crear cliente con `System.Net.WebSockets` o `Websocket.Client` implementando pings, heartbeats, reintentos exponenciales y circuit breaker (`Polly`) | Librerías `nhooyr.io/websocket` o `gorilla/websocket`, con contexto y cancelación |
| **LibSignal**            | Implementar cifrado de extremo a extremo y ratcheting (Signal Protocol) | Gestión de claves poco documentada | Encapsular en un servicio para generar, almacenar y rotar claves; proteger claves en reposo (DPAPI/KeyVault) | Uso de `github.com/whisperfish/libsignal-protocol-c` o `signal-cli` con bindings |
| **Protobuf (`WAProto`)** | Serializar y deserializar mensajes según el protocolo de WhatsApp | Archivo generado gigante; compilación lenta | Generar código en proyecto aparte; considerar `Google.Protobuf` y usar streams en vez de cargar todo en memoria | `google.golang.org/protobuf` con generación y módulos separados |
| **Crypto / Hashing**     | Hash HMAC‑SHA256, AES‑CMAC y HKDF para autenticación             | Uso disperso de helpers; riesgo de errores | Centralizar en un módulo `CryptoService`; añadir tests de vectors conocidos | Uso de `crypto/hmac`, `crypto/aes` y `golang.org/x/crypto/hkdf` |
| **Servicios de Google**  | En algunos proyectos se usa Google API para traducción o almacenamiento de archivos (detectado en DocumentaciónJules) | Dependencia externa sin abstracción | Encapsular integraciones en adaptadores; evaluar alternativas locales | Uso de clients Go oficiales (`cloud.google.com/go`) |

La elección de tecnologías debe alinearse con los objetivos del proyecto.  En una reescritura en Go es importante valorar la madurez de las bibliotecas (especialmente para LibSignal) y el soporte multiplataforma.

Proveniencia: Codex, Jules, Copilot y análisis propio.
