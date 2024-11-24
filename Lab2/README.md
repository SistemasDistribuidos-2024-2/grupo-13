### Descripción:
Este proyecto implementa un sistema distribuido utilizando Docker y Docker Compose para su ejecución.

### Instrucciones de ejecucion:
Ejecutar `make docker-(programa a ejecutar)` dentro de la carpeta `distribuidos` que esta cargada en el repositorio, las distribucion de las maquinas es:

- MV1:
  - Data Node 1: `docker-dn1`
  - Isla File: `docker-if`


- MV2:
  - Continente Folder: `docker-cf`
  - Diaboromon: `docker-diaboromon`


- MV3:
  - Continente Server: `docker-cs`
  - Data Node 2: `docker-dn2`



- MV4:
  - Nodo Tai: `docker-tai`
  - Primary Node: `docker-pn`

### El orden de ejecucion es (esperar que se monte completamente):

Data Nodes -> Diaboromon -> Primary Node -> Continentes -> Nodo Tai

### Consideraciones:
- Se considera que en todas las carpetas estan los archivos de `INPUT.txt` con sus datos pertinentes y en los contienetes estaran los archivos `Digimons.txt` con sus datos pertinentes.
