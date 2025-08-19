# 03c. Criptografía y claves

## Ubicación de claves
- `AuthenticationCreds` se serializa en `creds.json` en la carpeta de sesión.
- Claves de sesión y prekeys en disco sin cifrar.
- `LibSignal` implementa ECDH, HKDF y cifrado simétrico usando BouncyCastle.

## Generación y uso
- Durante el emparejamiento se generan pares de claves (`Noise`, `Signal`), almacenados en `FileKeyStore`.
- Los mensajes usan cifrado de extremo a extremo basado en el protocolo Signal.

## Problemas de seguridad
1. **Claves en repositorio**: varios `creds.json` versionados.
2. **Sin cifrado en reposo**: cualquier acceso al disco expone la sesión.
3. **Ausencia de rotación**: no se renuevan claves cuando se reestablece la conexión.

## Mejores prácticas .NET
- Usar `System.Security.Cryptography` para envolver claves y `ProtectedData` en Windows o `libsodium` en Linux.
- Guardar sesiones en archivos cifrados con contraseña o usar `Azure KeyVault`/`AWS KMS` si se despliega en la nube.
- Implementar rotación automática de prekeys.

## Equivalentes en Go
- Paquete estándar `crypto/*` para AES, HMAC y ECDH.
- Uso de `age` o `sops` para gestionar secretos.

## Recomendación
1. Mover archivos de sesión fuera del control de versiones.
2. Implementar `IKeyStore` con soporte de cifrado en reposo (por ejemplo, LiteDB con contraseña o SQLite cifrada).
3. Añadir validación de entrada y limpieza segura de buffers para evitar exposición en memoria.
