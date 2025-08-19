# 01. Resumen ejecutivo unificado

BaileysCSharp es un port en C# de la librería **Baileys** (Node.js) que permite interactuar con WhatsApp Web mediante WebSocket.  El proyecto sirve como núcleo de otros productos y expone un API para conectar, enviar y recibir mensajes, así como gestionar eventos y almacenar sesiones.

## Estado actual

- **Port parcial:** el código se adaptó desde la versión TypeScript y mantiene estructuras y patrones propios de Node.js.  Esto provoca acoplamiento y complejidad innecesaria en C#.
- **Persistencia simplista:** utiliza SQLite para guardar credenciales y sesiones.  No hay control fino del esquema ni de la concurrencia.
- **Gestión de conexión:** implementa un cliente WebSocket que simula a WhatsApp Web.  La reconexión y la detección de caídas son básicas.
- **Cifrado:** emplea la librería LibSignal para el intercambio de claves y el cifrado de extremo a extremo, pero la gestión de claves y el ciclo de vida carece de documentación.
- **Pruebas limitadas:** sólo existen test unitarios para credenciales; carece de pruebas de integración, contractuales o de carga.
- **Arquitectura monolítica:** el ensamblado central agrupa lógica de transporte, persistencia, dominio y presentación.  Falta separación clara de responsabilidades y uso de inyección de dependencias.
- **Riesgos de licencia:** al ser un fork de un proyecto público, hay obligaciones de licencia (MIT), y se reutilizan partes de Baileys.  Debe revisarse el cumplimiento.

## Problemas críticos

1. **Seguridad y sesión:** el almacenamiento de credenciales en archivos JSON y SQLite sin cifrado ataca la privacidad.  No hay rotación de claves ni protección frente a fugas.
2. **Robustez y resiliencia:** la conexión WebSocket no implementa backoff ni retry adecuados.  Cualquier error de red provoca desconexiones que exigen reescanear el código QR.
3. **Pruebas y calidad:** la falta de test limita la confianza en refactorizaciones y cambios.  La cobertura actual es insignificante.
4. **Arquitectura y mantenimiento:** el código mezcla conceptos (core, eventos, provider de plataforma) dificultando la extensión y el reemplazo de componentes.
5. **Rendimiento:** algunas operaciones (procesamiento de proto, manejo de archivos multimedia) bloquean hilos y usan memoria excesiva.

## Recomendación inicial

- **Semáforo de decisión:** se propone **amarillo** para una adaptación incremental.  Hay valor reutilizable en el código actual (modelos proto, bindings a LibSignal), pero la arquitectura necesita reestructuración.
- **Alternativa:** iniciar un proyecto nuevo con diseño hexagonal en **.NET** o **Go** (según análisis en el capítulo 07) y migrar funcionalidades gradualmente.  Esta opción permite corregir de raíz la arquitectura pero requiere más esfuerzo inicial.

Proveniencia: Codex, Jules, Copilot y análisis propio.
