### Descripción:
Este proyecto implementa un sistema distribuido utilizando Docker y Docker Compose para su ejecución.

### Instrucciones de ejecucion:
Ejecutar `make docker-(programa a ejecutar)` dentro de la carpeta `Lab1` que esta cargada en el repositorio, las distribucion de las maquinas es:

- dist049 Konzu/logistica
- dist050 Clientes
- dist051 Caravanas
- dist052 Raquis/financiero

### El orden de ejecucion es (esperar que se monte completamente):

logistica -> Caravanas -> financiero -> Clientes

### Consideraciones:
- En clientes se despliega una terminal para poder preguntar por el estado de un paquete en su respectivo docker, deben solo seguir la interfaz de esta, la cual es auto descriptiva.
- Los paquetes se envian automaticamente en segundo plano.
- En raquis, los registros se guardan en un .txt que es llamado finanzas.txt en su respectivo docker.
- En caravanas se genera un .txt por cada caravana con sus logs respectivos en su respectivo docker.
- Los parametros de tiempo se ejecutan directamente en el `docker-compose.yml` con el nombre de `TIEMPO_OPERACION`, este es distinto tanto para `Carvanas` como para `Clientes`.

### Formato de entrada:

- El formato de entrada considerado es el mostrado en el correo que se envio exceptuando el ultimo campo:

  `int IdPaquete, string Tipo, String Contenido, int Precio, String Escolta, String Escolta, String Destino,` (todas las lineas deben ser terminadas con una `,`)

- Ejemplo del formato con nombre de `packages.txt` (igualmente hay un ejemplo cargado en `.\Distribuidos\clientes`):

````
0001,Ostronitas,Choripán,200,Escolta1,Destino1,
0002,Normal,Anticucho,450,Escolta2,Destino5,
0003,Prioritario,Pan,100,Escolta1,Destino2,
0004,Prioritario,Coca-Cola,75,Escolta2,Destino1,
0005,Prioritario,Coca-Cola,100,Carlos,Destino1,
0006,Normal,Coca-Cola,100,Carlos,Destino1,
0007,Ostronitas,Coca-Cola,100,Carlos,Destino1,
0008,Ostronitas,Coca-Cola,100,Carlos,Destino1,
````