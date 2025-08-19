# 03a. SQLite y almacenamiento

## Esquema actual
- La librería utiliza **LiteDB** como base NoSQL embebida (`store.db`).
- Colecciones principales: chats, contactos, mensajes y claves (`Core/Types/*`).
- No existe cifrado ni índices definidos manualmente.
- Acceso mediante un singleton `MemoryStore`, lo que limita concurrencia.

## Problemas
- LiteDB no soporta alta concurrencia; riesgo de corrupción si se abren múltiples procesos.
- Los archivos `creds.json` se almacenan en texto plano sin encriptar.
- No hay estrategia de backup ni rotación de base.

## Alternativas en .NET
| Opción | Ventajas | Desventajas |
|-------|----------|-------------|
| SQLite con WAL | Mejor manejo de concurrencia, ecosistema maduro | Requiere mapeo relacional y migraciones |
| LiteDB con cifrado | Sencillo, API similar | No resuelve limitaciones de concurrencia |
| LiteDB + WAL-like (no soportado) | N/A | No disponible |
| Postgres embebido (p.ej. `npgsql` + servidor local) | Potente y escalable | Sobrecoste de instalación |

## Alternativas en Go
| Opción | Comentario |
|-------|------------|
| `mattn/go-sqlite3` | SQLite clásico, dependencias C |
| `modernc.org/sqlite` | SQLite puro Go, más fácil de compilar |
| `bbolt`/`etcd bbolt` | Key-value simple, ideal para sesiones |

## Recomendación
1. Definir interfaz `IStateStore` para desacoplar la persistencia.
2. Prototipar migración a SQLite con WAL + cifrado (p. ej. SQLCipher) para mayor seguridad.
3. En caso de reescritura en Go, evaluar `bbolt` para estado ligero y `postgres` externo para datos duraderos.
