# 3a. Análisis de Persistencia con LiteDB

## 3a.1. Esquema y Uso Actual

La librería `BaileysCSharp` emplea una estrategia de persistencia dual:

1.  **`FileKeyStore`**: Para datos de sesión y credenciales críticas. Serializa objetos individuales en archivos `.json` en el disco. No utiliza LiteDB.
2.  **`MemoryStore`**: Para datos de la aplicación como chats, mensajes y contactos. Utiliza una única base de datos **LiteDB** llamada `store.db`.

`MemoryStore` instancia y gestiona la base de datos `LiteDB`, creando las siguientes colecciones (equivalentes a tablas en SQL):

-   **`chats`**: Almacena metadatos de las conversaciones (`ChatModel`).
-   **`messages`**: Almacena los mensajes individuales (`MessageModel`).
-   **`contacts`**: Guarda la información de los contactos (`ContactModel`).
-   **`groupMetaData`**: Mantiene los metadatos de los grupos (`GroupMetadataModel`).

El `MemoryStore` se acopla fuertemente al `EventEmitter` de la librería. Se suscribe a eventos (`Message.Upsert`, `Chats.Update`, etc.) y actualiza las colecciones de LiteDB en respuesta. Esto significa que la persistencia está intrínsecamente ligada a la lógica de eventos.

**Rutas de Archivos**:
-   La base de datos se encuentra en: `<config.CacheRoot>\\store.db`.
-   Las credenciales de `FileKeyStore` se guardan en: `<config.CacheRoot>\\<FolderPrefix>\\<id>.json`.

**Locking y Concurrencia**:
-   Se utilizan `lock(locker)` en varios métodos de `MemoryStore` y `FileKeyStore` para evitar condiciones de carrera al acceder tanto a la base de datos como a los archivos.
-   LiteDB, por defecto, opera en modo `Exclusive` (un único proceso escritor). Esto puede ser un cuello de botella si se intentara escalar a múltiples hilos escribiendo de forma intensiva. El uso de `Checkpoint()` cada 30 segundos ayuda a liberar el log y mantener el archivo de datos.

## 3a.2. Alternativas y Mejoras en .NET

La principal debilidad del enfoque actual es el alto acoplamiento. La lógica de negocio está mezclada con la de persistencia.

| Alternativa             | Ventajas                                                              | Desventajas                                                                 |
| ----------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| **SQLite con Dapper/EF Core** | Motor SQL robusto, ampliamente soportado, buen rendimiento con WAL (Write-Ahead Logging). | Requiere definir un esquema relacional (migraciones). Más verboso que LiteDB. |
| **LiteDB (con Repositorio)** | Se mantiene la simplicidad de LiteDB, pero se desacopla la lógica de negocio. | Requiere un refactor para introducir interfaces de repositorio.             |
| **Realm .NET**          | Base de datos mobile-first, reactiva, puede ser más rápida para ciertos casos de uso. | Ecosistema más pequeño, puede ser excesivo para una librería de backend.     |
| **Archivos Cifrados (JSON/Protobuf)** | Control total sobre el formato, sin dependencias de DB. Se puede usar `System.Security.Cryptography`. | Lento, propenso a errores, se debe implementar toda la lógica de consulta y transacción. |

**Recomendación para .NET**:
La mejora más impactante y con menor esfuerzo sería **introducir un patrón de Repositorio**.
1.  Definir interfaces como `IMessageRepository`, `IChatRepository`, etc.
2.  La implementación actual (`LiteDBMessageRepository`) usaría LiteDB.
3.  La lógica de negocio (`WASocket`, `MemoryStore`) dependería de las interfaces, no de LiteDB directamente.
Esto desacopla el código y facilita enormemente las pruebas unitarias (usando repositorios en memoria).

## 3a.3. Alternativas y Equivalentes en Go

Si se considera una reescritura en Go, estas son las alternativas de persistencia:

| Alternativa             | Equivalencia .NET       | Ventajas en Go                                                                   | Desventajas                                                              |
| ----------------------- | ----------------------- | -------------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
| **`mattn/go-sqlite3`**  | SQLite (Dapper)         | El driver más popular y estable para SQLite en Go. CGO-based.                   | Requiere compilador C (CGO), puede complicar el cross-compiling.         |
| **`modernc.org/sqlite`**| SQLite (EF Core)        | Implementación de SQLite en Go puro. Sin CGO, facilita la compilación cruzada.   | Relativamente más nuevo que `mattn/go-sqlite3`.                          |
| **`bbolt` (antes BoltDB)** | LiteDB                  | Base de datos de clave-valor embebida, muy rápida para lecturas. Popular (usada en `etcd`). | API de bajo nivel. No es una base de datos de documentos como LiteDB.      |
| **`badger`**            | LiteDB / Realm          | KV store embebido de alto rendimiento de Dgraph. Optimizado para SSDs.          | Más complejo que bbolt.                                                  |

**Recomendación para Go**:
Para un caso de uso como este, **`modernc.org/sqlite`** es probablemente la mejor opción. Proporciona el poder de SQL para consultas complejas (ej. "buscar mensajes con media de un usuario específico") sin la sobrecarga de CGO. Si la simplicidad de clave-valor es suficiente, **`bbolt`** es una excelente alternativa, robusta y probada en batalla.

## 3a.4. Plan de Migración (si aplica)

Si se decide **mejorar el proyecto .NET actual**, el plan sería:
1.  **Sprint 1**: Introducir interfaces de repositorio (`IRepository<T>`) y un `UnitOfWork` para la gestión de transacciones.
2.  **Sprint 2**: Refactorizar `MemoryStore` para que use las nuevas interfaces en lugar de acceder a `LiteDB` directamente.
3.  **Sprint 3**: Escribir implementaciones de prueba (en memoria) de los repositorios y añadir pruebas unitarias a la lógica de negocio que antes estaba acoplada.
4.  **Sprint 4 (Opcional)**: Crear una implementación de los repositorios con `SQLite` para evaluar el rendimiento y la viabilidad de una migración.
