# 03e. Procesamiento de media

## Dependencias
- **FFMpegCore**: wrapper para ffmpeg utilizado en transcodificación de audio y video.
- **SkiaSharp**: librería de gráficos para manipulación de imágenes y stickers.

## Uso actual
- Conversión de archivos multimedia al formato exigido por WhatsApp (por ejemplo, compresión de imágenes o audio).
- Descarga de archivos y escritura directa en disco sin verificación de tamaño.

## Problemas
1. Posible fuga de procesos `ffmpeg` si se cancelan operaciones abruptamente.
2. Sin control de límites de tamaño o memoria antes de procesar.
3. Falta de abstracción para reemplazar implementaciones.

## Mejoras sugeridas
- Implementar un servicio `IMediaProcessor` con colas y *pooling* de procesos.
- Validar metadatos (mime, tamaño) antes de invocar a `ffmpeg` o `SkiaSharp`.
- En Go, evaluar `ffmpeg-go` o bindings a `libvips` para imágenes.
