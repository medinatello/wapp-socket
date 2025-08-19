# 3c. Análisis de Criptografía y Gestión de Claves

## 3c.1. Generación y Almacenamiento de Claves

La librería implementa el protocolo Signal para el cifrado de extremo a extremo (E2E), dependiendo en gran medida de la librería `BouncyCastle` para las operaciones criptográficas primitivas.

**Flujo de Generación y Almacenamiento**:
1.  **Generación de Identidad**: La primera vez que se conecta un cliente, se genera un par de claves de identidad (pública y privada) de largo plazo, junto con un ID de registro y un conjunto de pre-claves firmadas (`SignedPreKey`). Esto ocurre dentro de la lógica de `AuthenticationUtils.InitAuthCreds()` y las clases del directorio `LibSignal`.
2.  **Uso de `BouncyCastle`**: Se utiliza `BouncyCastle` para todas las operaciones de curva elíptica (Curve25519), AES-GCM para el cifrado simétrico, y HMAC-SHA256 para la autenticación de mensajes.
3.  **Almacenamiento**: Todas las claves generadas (identidad, pre-claves, claves de sesión, etc.) se agrupan en el objeto `AuthenticationCreds`.
4.  **Persistencia Insegura**: Este objeto `AuthenticationCreds` es serializado a un archivo JSON llamado `creds.json` por la clase `FileKeyStore`. **Las claves privadas se guardan en texto plano dentro de este archivo.**

**Ruta del Archivo de Credenciales**: `<config.CacheRoot>\\creds.json`

## 3c.2. Cifrado en Tránsito y en Reposo

-   **Cifrado en Tránsito**: La comunicación con los servidores de WhatsApp está doblemente cifrada.
    1.  **Capa de Transporte**: La conexión WebSocket (`wss://`) está protegida por TLS, cifrando todo el tráfico entre el cliente y el servidor de WhatsApp.
    2.  **Capa de Aplicación**: Los mensajes se cifran de extremo a extremo utilizando el protocolo Signal. Esto significa que ni siquiera WhatsApp puede leer el contenido de los mensajes. La librería implementa correctamente este cifrado.

-   **Cifrado en Reposo**: **No existe.** Este es el principal riesgo de seguridad. Las claves maestras de la sesión se almacenan sin cifrar en el disco. Si un atacante obtiene acceso al archivo `creds.json`, puede suplantar la identidad del usuario, descifrar mensajes futuros y potencialmente acceder al historial si también obtiene la base de datos `store.db`.

## 3c.3. Riesgos y Endurecimiento (Hardening)

**Riesgo Principal**: Compromiso de la sesión por acceso no autorizado al sistema de archivos.

### Propuestas de Endurecimiento en .NET

La solución es cifrar el archivo `creds.json` utilizando una clave derivada de secretos del sistema operativo o una contraseña proporcionada por el usuario.

| Alternativa                               | Ventajas                                                                               | Desventajas                                                              |
| ----------------------------------------- | -------------------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
| **`ProtectedData` (DPAPI)**               | Utiliza la API de Protección de Datos de Windows (o el Keyring en Linux/macOS). Cifra los datos usando una clave ligada al usuario o a la máquina. No requiere gestionar contraseñas. | Los datos solo pueden ser descifrados en la misma máquina (o por el mismo usuario). |
| **Derivación de Clave (PBKDF2)**          | Se solicita una contraseña al usuario y se deriva una clave de cifrado con un KDF (Key Derivation Function) como PBKDF2. | Requiere que la aplicación gestione de forma segura una contraseña.      |
| **Usar un `SecureString`**                | Para mantener la contraseña en memoria de forma más segura.                            | `SecureString` tiene limitaciones y su uso puede ser complejo.            |

**Recomendación para .NET**:
Utilizar **`System.Security.Cryptography.ProtectedData`**. Es la solución más robusta y transparente para el usuario final.
1.  Al guardar `creds.json`, cifrar el contenido serializado con `ProtectedData.Protect()`.
2.  Al cargar `creds.json`, descifrarlo con `ProtectedData.Unprotect()`.
Esto liga la seguridad del archivo de credenciales a la del sistema operativo.

### Equivalentes y Propuestas en Go

Go ofrece primitivas excelentes para implementar esto de forma manual.

| Alternativa                      | Equivalencia .NET       | Ventajas en Go                                                                | Desventajas                                                              |
| -------------------------------- | ----------------------- | ----------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
| **`crypto/aes`, `crypto/cipher`**| Cifrado manual con AES  | Permite implementar cifrado AES-GCM de forma directa y eficiente.             | Se debe gestionar la generación y el almacenamiento de la clave de cifrado. |
| **`golang.org/x/crypto/scrypt`** | PBKDF2                  | Implementación de Scrypt, un KDF más moderno y resistente que PBKDF2.         | Requiere una contraseña de entrada.                                      |
| **Librerías de SO (ej. `keyring`)** | `ProtectedData` (DPAPI) | Existen librerías como `99designs/keyring` que abstraen el acceso a los almacenes de secretos nativos del SO (Keychain en macOS, Credential Manager en Windows, etc.). | Añade una dependencia de terceros para una funcionalidad crítica.        |

**Recomendación para Go**:
Para una solución de servidor, donde no hay un "usuario" interactivo, la mejor opción es **cifrar el archivo de credenciales con una clave maestra que se proporciona a la aplicación a través de una variable de entorno o un sistema de gestión de secretos** (como HashiCorp Vault o AWS KMS). La librería `crypto/aes` es perfecta para esto. Para aplicaciones de escritorio, usar una librería como **`99designs/keyring`** sería el equivalente directo y más seguro a `ProtectedData`.
