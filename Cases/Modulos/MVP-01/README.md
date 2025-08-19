# MVP_13 — Testing y Deployment

Este módulo corresponde al Sprint 13 y tiene como objetivo fortalecer las pruebas del proyecto y automatizar el deployment.

## Objetivos

- Completar la suite de pruebas unitarias e integración, alcanzando al menos un 80 % de cobertura global.
- Crear pruebas de API utilizando frameworks como Testify o GoConvey.
- Desarrollar mocks para dependencias externas.
- Preparar Dockerfiles optimizados y archivos de docker-compose para entornos de desarrollo.
- Configurar una pipeline de CI/CD básica para build, tests, generación de imágenes Docker y despliegue.

## Versiones y Dependencias Clave

- Go ≥ 1.20
- Librerías de testing (`github.com/stretchr/testify`), frameworks de integración.
- Docker y docker-compose.

## Criterios de Aceptación

- La cobertura de pruebas supera el 80 % a nivel global.
- Las pruebas de integración validan los endpoints principales.
- Los Dockerfiles generan imágenes reproducibles y ligeras.
- La pipeline de CI/CD construye, prueba y despliega la aplicación de forma automatizada.
- La documentación y `resultado.md` están actualizadas.

## Riesgos y Consideraciones

- Complejidad de configurar CI/CD en entornos limitados.
- Garantizar que los contenedores se comporten de forma similar al entorno de producción.