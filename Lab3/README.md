### Descripción:
Este proyecto implementa un sistema distribuido utilizando Docker y Docker Compose para su ejecución.

### Instrucciones de ejecucion:
Ejecutar `make docker-(programa a ejecutar)` dentro de la carpeta `Lab3` que esta cargada en el repositorio, las distribucion de las maquinas es:

- MV1:
  - Broker: `docker-broker`


- MV2:
  - Supervisor 1: `docker-s1`
  - Servidor hextech 1: `docker-h1`


- MV3:
  - Supervisor 2: `docker-s2`
  - Servidor hextech 2: `docker-h2`


- MV4:
  - Jayce: `docker-jayce`
  - Servidor hextech 3: `docker-h3`

### El orden de ejecucion es (esperar que se monte completamente):

Supervisores -> Jayce -> Broker -> Servidores hextech

### Consideraciones:
- Se asume que todos los comandos serán ejecutados manualmente desde la consola por el usuario. Tanto los Supervisores como Jayce cuentan con un menú propio que debe seguirse estrictamente (no se contemplan acciones fuera del menú).
- Los Supervisores utilizan memoria local, mientras que Jayce realiza escrituras en archivos .txt.
- En caso de una solicitud con errores (por ejemplo, consultar una región que no fue cargada), el sistema devolverá un reloj vectorial sin cambios respecto al último entregado.

