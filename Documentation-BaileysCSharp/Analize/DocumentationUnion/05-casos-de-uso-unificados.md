# 05. Casos de uso unificados

Los siguientes casos de uso reflejan las interacciones típicas con la librería.  Se redactan en un lenguaje orientado a negocio para facilitar su incorporación a un backlog de Sprints.

### CU‑01: Conexión mediante QR

*Como usuario desarrollador quiero iniciar una sesión de WhatsApp Web usando un código QR para que mi aplicación se autentique con una cuenta de WhatsApp.*

**Precondiciones:** El dispositivo del usuario tiene una cuenta de WhatsApp válida y acceso a internet.  
**Flujo básico:** La librería genera un QR, el usuario lo escanea con su móvil y la sesión queda autenticada.  
**Criterios de aceptación:** El QR se genera en menos de 5 segundos; la sesión se mantiene activa al menos 30 minutos sin interacción; se emiten eventos de éxito o error.

### CU‑02: Envío de mensaje de texto

*Como aplicación deseo enviar un mensaje de texto a un contacto específico para comunicar información a un usuario final.*

**Precondiciones:** Existe una sesión autenticada y se conoce el identificador del destinatario.  
**Flujo básico:** La aplicación invoca el método de envío, la librería cifra y envía el mensaje; se emite evento de confirmación.  
**Criterios de aceptación:** El mensaje llega al destinatario; los errores de red se reintentan automáticamente hasta 3 veces; se reportan fallos.

### CU‑03: Recepción de mensajes

*Como sistema quiero recibir y procesar mensajes entrantes para reaccionar a acciones del usuario.*

**Precondiciones:** Sesión activa con la API de WhatsApp.  
**Flujo básico:** La librería escucha el WebSocket y emite eventos de mensaje recibido; la aplicación los procesa según el tipo (texto, media, notificación).  
**Criterios de aceptación:** No se pierden mensajes; los eventos llegan con los metadatos completos; se puede reconocer mensajes de grupos y de usuarios individuales.

### CU‑04: Desconexión controlada

*Como aplicación quiero cerrar la sesión de forma limpia para liberar recursos y asegurar que no queden sesiones huérfanas.*

**Precondiciones:** Sesión activa.  
**Flujo básico:** La aplicación invoca una orden de cierre; la librería cierra el WebSocket, actualiza la base de datos y borra claves temporales.  
**Criterios de aceptación:** No queda conexión abierta; los recursos (memoria, ficheros) se liberan; se puede volver a conectar sin errores.

Proveniencia: Codex, Jules, Copilot y análisis propio.
