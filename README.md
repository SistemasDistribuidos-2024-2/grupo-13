# Grupo 13

### Integrantes:
- **Martín Pino** | Rol: 202073528-K
- **Luciano Yevenes** | Rol: 202173514-3

### Descripción:
Este proyecto implementa un sistema distribuido utilizando Docker y Docker Compose para su ejecución. Debido a problemas con las máquinas virtuales al momento de probar, se optó por realizar todas las pruebas en local. A pesar de esto, se sugiere utilizar Docker Compose para facilitar la ejecución en una máquina virtual, ya que este sistema configura automáticamente todo el entorno necesario para correr el proyecto.

### Requisitos:
- **GoLand** versión 1.23.1 (o superior).
- **Docker** instalado.
- **Docker Compose** para la configuración automática del entorno.

### Estructura del Proyecto:
Todos los códigos fuente están organizados en la carpeta `distribuidos`. 

### Instrucciones de ejecucion:
Ejecutar `docker-compose up --build` dentro de la carpeta `distribuidos`. 

### Consideraciones:
- En clientes se despliega una terminal para poder preguntar por el estado de un paquete, deben solo seguir la interfaz de esta, la cual es auto descriptiva.
- Los paquetes se envian automaticamente en segundo plano.