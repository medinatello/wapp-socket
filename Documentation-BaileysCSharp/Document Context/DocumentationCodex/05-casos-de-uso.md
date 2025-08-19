# 05. Casos de uso

## CU1. Enviar mensaje de texto
- **Como** bot o integración externa
- **Quiero** enviar un mensaje de texto a un número de WhatsApp
- **Para** notificar información a un usuario

**Precondiciones**
- Sesión autenticada y conexión activa.
- Número destino válido.

**Postcondiciones**
- Mensaje entregado o error registrado.

**Reglas de negocio**
- Respetar límites de WhatsApp (spam, tasa).
- Almacenar mensaje y estado en base local.

**KPIs**
- Tasa de entrega (<1% fallos).
- Tiempo medio de envío (<1s).

## CU2. Recibir mensajes y responder automáticamente
- **Como** servicio de atención
- **Quiero** procesar mensajes entrantes y responder
- **Para** automatizar soporte

**Precondiciones**: sesión y eventos registrados.
**Postcondiciones**: respuesta enviada, logs generados.
**Reglas**: parseo seguro, no ejecutar comandos arbitrarios.
**KPIs**: tiempo de respuesta, porcentaje de mensajes parseados.

## CU3. Gestión de grupos
- **Como** administrador
- **Quiero** crear grupos y añadir miembros
- **Para** organizar conversaciones

**Precondiciones**: credenciales con permisos.
**Postcondiciones**: grupo creado y miembros notificados.
**KPIs**: tiempo de creación, errores de membresía.

## CU4. Cierre seguro de sesión
- **Como** operador
- **Quiero** finalizar la conexión liberando recursos
- **Para** evitar fugas y mantener seguridad

**Precondiciones**: conexión establecida.
**Postcondiciones**: sockets cerrados, claves persistidas.
**KPIs**: memoria liberada, tiempo de cierre.
