# 03 - Decisión sobre Manejo de Eventos

Este documento analiza dos enfoques para manejar los eventos internos de la aplicación, principalmente los que provienen de la conexión WebSocket (mensajes entrantes, cambios de estado, etc.).

## Contexto

Los eventos de la conexión WebSocket (y otros eventos del dominio) necesitan ser consumidos por los casos de uso para orquestar la lógica de negocio. Por ejemplo, un evento `message_received` debe ser procesado por el caso de uso `ReceiveMessage`.

Se evalúan dos patrones principales para la distribución de estos eventos.

### Alternativa A: Bus de Eventos (Event Bus)

Un componente centralizado (`EventBus`) se encarga de recibir y distribuir eventos.

- **Diagrama**: `WebSocketAdapter -> EventBus.Publish(event) -> UseCase.Handle(event)`
- **Implementación**:
  - Una interfaz `port.EventBus` con métodos `Publish(topic, event)` y `Subscribe(topic) <-chan Event`.
  - Un adaptador en memoria que gestiona un mapa de `topic -> []chan Event`.
  - Los casos de uso se suscriben a los topics que les interesan al iniciar la aplicación.

#### Ventajas
- **Desacoplamiento**: El productor de eventos no conoce a los consumidores.
- **Flexibilidad**: Es fácil añadir nuevos consumidores sin modificar al productor.
- **Centralización**: El flujo de eventos es explícito y se gestiona en un solo lugar.

#### Desventajas
- **Complejidad**: Introduce un nuevo componente central que gestionar.
- **Backpressure**: Más difícil de manejar. Si un consumidor es lento, puede bloquear al bus o requerir buffers, aumentando el uso de memoria.
- **Descubrimiento**: Puede ser más difícil seguir el flujo de un evento a través del código.

### Alternativa B: Canales de Go Dedicados (Fan-out)

La conexión (`WebSocketConn`) expone un único canal de eventos salientes. Un componente "distribuidor" (o el propio caso de uso) lee de este canal y distribuye los eventos a los consumidores interesados (fan-out). Se puede usar un patrón "tee" para clonar el canal.

- **Diagrama**: `WebSocketAdapter.Events() <-chan Event -> Distributor -> UseCase1(<-chan Event), UseCase2(<-chan Event)`
- **Implementación**:
  - La interfaz `WebSocketConn` tiene un método `Events() <-chan Event`.
  - Un componente inicializador crea la conexión, obtiene el canal y lo pasa a los casos de uso que lo necesitan.

#### Ventajas
- **Simplicidad**: Menos abstracciones, usando primitivas nativas de Go.
- **Backpressure explícito**: La presión de un consumidor lento se propaga directamente al productor si los canales no tienen buffer.
- **Claridad**: El flujo de datos es más directo y fácil de seguir.

#### Desventajas
- **Acoplamiento**: El componente que inicializa el flujo necesita conocer a todos los consumidores.
- **Gestión de Goroutines**: Se debe tener cuidado de que todas las goroutines que leen del canal terminen correctamente para evitar fugas.

## Decisión para Sprint 1

Para el Sprint 1, se implementarán **ambas rutas** para poder experimentar y tomar una decisión más informada en el futuro.

- **Default**: La **Alternativa A (Event Bus)** será la implementación por defecto, habilitada.
- **Alternativa**: La **Alternativa B (Canales Dedicados)** se implementará detrás de un flag de configuración o un build tag (`alt_eventing`).

Ambos caminos conectarán a los mismos casos de uso, permitiendo un cambio rápido entre uno y otro para comparar la ergonomía y el rendimiento en un entorno simulado.
