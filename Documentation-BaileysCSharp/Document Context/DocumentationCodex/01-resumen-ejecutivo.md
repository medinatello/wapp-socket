# 01. Resumen ejecutivo

**BaileysCSharp** es un fork en .NET 9 de la librería Baileys (NodeJS) que permite interactuar con WhatsApp Web mediante WebSocket. Hoy provee un cliente capaz de:
- Iniciar sesión mediante código QR y mantener una sesión local.
- Enviar y recibir mensajes de texto y multimedia.
- Gestionar eventos de presencia, historial y algunas funciones de newsletter/grupos.

## Estado actual
- Código parcialmente portado desde TypeScript sin una arquitectura clara.
- Uso de LiteDB para almacenamiento local de estado.
- Pruebas mínimas: solo se incluyen credenciales de ejemplo sin tests automatizados reales.
- Alto acoplamiento entre transporte, cifrado y lógica de dominio.

## Problemas críticos
1. **Seguridad**: credenciales (`creds.json`) en el repositorio y manejo de claves en texto plano.
2. **Robustez limitada**: reconexión, manejo de errores y backoff apenas implementados.
3. **Cobertura de tests nula**: dificulta refactor y evolución segura.

## Caminos posibles
| Opción | Descripción | Ventajas | Desventajas |
|-------|-------------|----------|-------------|
| Mejorar/adaptar | Refactor gradual sobre el código actual en .NET | Aprovecha código existente, menor curva de aprendizaje | Hereda deuda técnica y acoplamientos; refactor complejo |
| Reescribir | Rediseño total (en .NET o Go) aplicando arquitectura hexagonal | Código limpio, control de dependencias, mejoras de seguridad | Mayor costo inicial, requiere definir compatibilidad de protocolo |

## Recomendación inicial
**Semáforo: Amarillo.** Se sugiere iniciar con un refactor profundo en .NET para evaluar viabilidad. Si en 2-3 sprints la deuda permanece elevada, considerar reescritura controlada (preferentemente en Go para servicios de alto rendimiento) por estas razones:
1. Go ofrece manejo de concurrencia más simple y eficiente para sockets.
2. El código actual requiere rediseño total de almacenamiento y seguridad.
3. Ecosistema .NET para WebSocket/crypto es sólido pero la implementación está inmadura.
