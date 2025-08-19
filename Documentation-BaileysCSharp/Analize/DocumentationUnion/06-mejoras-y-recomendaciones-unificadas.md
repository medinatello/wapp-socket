# 06. Mejoras y recomendaciones unificadas

Integrando las observaciones de las auditorías y el análisis del código se proponen las siguientes mejoras.

## Manejo de errores

- **Clasificar errores:** diferenciar entre errores de red transitorios, errores de autenticación y errores fatales del protocolo.  
- **Reintentos y backoff:** utilizar bibliotecas como `Polly` para implementar reintentos exponenciales con límites; en Go usar `cenkalti/backoff`.  
- **Idempotencia:** asegurar que las operaciones de envío y recepción sean idempotentes para evitar duplicados en caso de reintento.

## Logging y telemetría

- **Logging estructurado:** integrar `Serilog` o `Microsoft.Extensions.Logging` con formato JSON; incluir correlación entre eventos (IDs).  
- **Trazas distribuidas:** agregar soporte a OpenTelemetry para seguir el recorrido de mensajes y medir latencias.  
- **Niveles de log:** definir niveles claros (Debug, Info, Warning, Error) y permitir configurarlos por entorno.

## Refactor de arquitectura

- **Arquitectura hexagonal:** definir puertos (interfaces) para transporte, persistencia y cifrado; implementar adaptadores.  
- **Inyección de dependencias:** usar `Microsoft.Extensions.DependencyInjection` para componer módulos y facilitar pruebas unitarias.  
- **Módulos independientes:** separar el código generado de Proto, la lógica de Signal y el cliente WebSocket en proyectos o bibliotecas diferentes.

## Eficiencia y rendimiento

- **Control de memoria:** evitar cargar archivos grandes (e.g., WAProto) en memoria al inicio; cargar clases generadas bajo demanda.  
- **Procesamiento en streaming:** manejar medios (imágenes, vídeos) mediante streams y buffers, evitando copiar en memoria.  
- **Backpressure:** aplicar límites a la cantidad de mensajes en vuelo y usar colas asíncronas para regular la producción y el consumo.

## Seguridad

- **Cifrado en reposo:** cifrar la base de datos SQLite con `sqlcipher` o utilizar DPAPI; no almacenar claves en texto plano.  
- **Rotación de claves:** implementar rotación periódica de las claves de sesión y eliminación segura de claves antiguas.  
- **Gestión de secretos:** externalizar secretos (APIs, tokens) en un gestor seguro como Azure Key Vault o vault.

## Compatibilidad y portabilidad

- **Interfaz estable:** documentar y mantener estable la API pública para facilitar la migración de los usuarios.  
- **Soporte multiplataforma:** comprobar la librería en Windows, Linux y macOS; evitar llamadas específicas de Windows.

Estas mejoras apuntan a aumentar la robustez y escalabilidad del proyecto, independientemente de si se continúa en .NET o se migra a Go.

Proveniencia: Codex, Jules, Copilot y análisis propio.
