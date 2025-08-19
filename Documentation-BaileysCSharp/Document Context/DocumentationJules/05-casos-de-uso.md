# 5. Casos de Uso de Negocio

Esta sección traduce los flujos técnicos a narrativas desde la perspectiva del sistema o del negocio. Sirven para construir un backlog, definir criterios de aceptación y alinear el desarrollo técnico con los objetivos funcionales.

---

### **Caso de Uso 1: Establecer una nueva sesión de cliente**

-   **Narrativa**: Como sistema, cuando un nuevo cliente sin credenciales previas intenta conectarse, debo guiarlo a través del proceso de escaneo de código QR para establecer una sesión segura y persistente.
-   **Precondiciones**:
    -   El cliente no tiene un archivo de credenciales (`creds.json`) válido.
    -   La aplicación tiene acceso a la red para conectar con los servidores de WhatsApp.
-   **Flujo Principal**:
    1.  El sistema intenta conectarse y detecta la ausencia de credenciales.
    2.  El sistema solicita y recibe un código QR de los servidores de WhatsApp.
    3.  El sistema expone el código QR para que el usuario final pueda escanearlo.
    4.  Una vez escaneado, el sistema recibe las credenciales del servidor.
    5.  El sistema guarda estas credenciales de forma segura en el almacenamiento persistente.
    6.  El sistema confirma que la conexión está abierta y operativa.
-   **Postcondiciones**:
    -   Se ha creado un archivo de credenciales (`creds.json`) cifrado.
    -   La conexión está en estado `Open`.
    -   El sistema está listo para enviar y recibir mensajes.
-   **KPIs**:
    -   Tiempo desde la solicitud hasta la visualización del QR < 2 segundos.
    -   Tasa de éxito de establecimiento de sesión > 99%.

---

### **Caso de Uso 2: Reconectar una sesión existente**

-   **Narrativa**: Como sistema, cuando un cliente con credenciales válidas se conecta, debo usar esas credenciales para reanudar la sesión rápidamente sin requerir intervención del usuario.
-   **Precondiciones**:
    -   Existe un archivo de credenciales válido y descifrable.
-   **Flujo Principal**:
    1.  El sistema carga las credenciales del almacenamiento persistente.
    2.  El sistema se conecta al servidor de WhatsApp y presenta las credenciales durante el handshake.
    3.  El servidor valida las credenciales y reanuda la sesión.
    4.  El sistema confirma que la conexión está en estado `Open`.
-   **Postcondiciones**:
    -   La conexión está en estado `Open` y la sesión anterior ha sido reanudada.
-   **KPIs**:
    -   Tiempo de reconexión < 3 segundos.

---

### **Caso de Uso 3: Enviar un mensaje de texto**

-   **Narrativa**: Como sistema, cuando un cliente solicita enviar un mensaje a un destinatario, debo asegurar que el mensaje se cifre correctamente para ese destinatario y se envíe de forma fiable.
-   **Precondiciones**:
    -   La conexión está en estado `Open`.
    -   Existen claves de sesión válidas para el destinatario.
-   **Flujo Principal**:
    1.  El sistema recibe la solicitud de envío con el JID del destinatario y el contenido.
    2.  Recupera las claves de sesión del destinatario del almacén de Signal.
    3.  Cifra el contenido del mensaje.
    4.  Serializa el mensaje cifrado en el formato de `WebMessageInfo`.
    5.  Envía el mensaje binario a través del WebSocket.
-   **Postcondiciones**:
    -   El mensaje ha sido enviado al servidor.
    -   (Opcional) El estado del mensaje se actualiza localmente a "enviado".
-   **KPIs**:
    -   Latencia de envío (desde la llamada hasta el envío por el socket) < 100ms.

---

### **Caso de Uso 4: Recibir y procesar un mensaje entrante**

-   **Narrativa**: Como sistema, debo escuchar continuamente los mensajes entrantes, descifrarlos de forma segura y notificarlos al cliente de la aplicación en tiempo real.
-   **Precondiciones**:
    -   La conexión está en estado `Open`.
-   **Flujo Principal**:
    1.  El sistema recibe un frame binario del WebSocket.
    2.  Parsea el frame y extrae el mensaje cifrado.
    3.  Identifica al remitente y recupera las claves de sesión correspondientes.
    4.  Descifra el mensaje.
    5.  Emite un evento (`Message.Upsert`) con el mensaje descifrado.
    6.  Persiste el mensaje en la base de datos local.
-   **Postcondiciones**:
    -   El cliente de la aplicación ha sido notificado del nuevo mensaje.
    -   El mensaje se ha guardado en el historial local.
-   **KPIs**:
    -   Latencia de recepción (desde la llegada al socket hasta la emisión del evento) < 150ms.

---

### **Caso de Uso 5: Gestionar una desconexión inesperada**

-   **Narrativa**: Como sistema, si la conexión se pierde por un error de red, debo intentar restablecerla automáticamente sin intervención del usuario, utilizando una estrategia de reintentos inteligente.
-   **Precondiciones**:
    -   La conexión estaba previamente en estado `Open`.
-   **Flujo Principal**:
    1.  Se detecta una interrupción en el WebSocket.
    2.  El sistema identifica que la causa no es un cierre de sesión voluntario.
    3.  El sistema inicia una política de reintentos con backoff exponencial (espera 1s, luego 2s, 4s, etc.).
    4.  En cada intento, ejecuta el flujo del **Caso de Uso 2: Reconectar una sesión existente**.
    5.  Si tiene éxito, la conexión se restaura. Si falla después de N intentos, se notifica un error fatal.
-   **Postcondiciones**:
    -   La conexión se ha restablecido o se ha notificado un estado de fallo permanente.
-   **KPIs**:
    -   Tasa de reconexión automática exitosa > 95%.
