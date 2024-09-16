# Variables para los nombres de las imágenes y contenedores
CARAVANAS_IMAGE=caravanas_service
CLIENTES_IMAGE=clientes_service
FINANZAS_IMAGE=finanzas_service
LOGISTICA_IMAGE=logistica_service
RABBITMQ_IMAGE=rabbitmq:3-management

CARAVANAS_CONTAINER=caravanas_container
CLIENTES_CONTAINER=clientes_container
FINANZAS_CONTAINER=finanzas_container
LOGISTICA_CONTAINER=logistica_container
RABBITMQ_CONTAINER=rabbitmq_container

# Construir imágenes Docker para cada servicio

.PHONY: build-caravanas
build-caravanas:
	@echo "Construyendo la imagen de Caravanas..."
	docker build -t $(CARAVANAS_IMAGE) ./caravanas/
	@echo "Imagen de Caravanas construida."

.PHONY: build-clientes
build-clientes:
	@echo "Construyendo la imagen de Clientes..."
	docker build -t $(CLIENTES_IMAGE) ./clientes/
	@echo "Imagen de Clientes construida."

.PHONY: build-finanzas
build-finanzas:
	@echo "Construyendo la imagen de Finanzas..."
	docker build -t $(FINANZAS_IMAGE) ./finanzas/
	@echo "Imagen de Finanzas construida."

.PHONY: build-logistica
build-logistica:
	@echo "Construyendo la imagen de Logística..."
	docker build -t $(LOGISTICA_IMAGE) ./logistica/
	@echo "Imagen de Logística construida."

.PHONY: build-rabbitmq
build-rabbitmq:
	@echo "Usando la imagen predeterminada de RabbitMQ..."
	docker pull $(RABBITMQ_IMAGE)
	@echo "Imagen de RabbitMQ obtenida."

# Ejecutar los contenedores individualmente

.PHONY: run-caravanas
run-caravanas:
	@echo "Iniciando contenedor de Caravanas..."
	docker run -d --name $(CARAVANAS_CONTAINER) \
		$(CARAVANAS_IMAGE)
	@echo "Contenedor de Caravanas en ejecución."

.PHONY: run-clientes
run-clientes:
	@echo "Iniciando contenedor de Clientes..."
	docker run -d --name $(CLIENTES_CONTAINER) \
		-it \
		$(CLIENTES_IMAGE)
	@echo "Contenedor de Clientes en ejecución."

.PHONY: run-finanzas
run-finanzas:
	@echo "Iniciando contenedor de Finanzas..."
	docker run -d --name $(FINANZAS_CONTAINER) \
		$(FINANZAS_IMAGE)
	@echo "Contenedor de Finanzas en ejecución."

.PHONY: run-logistica
run-logistica:
	@echo "Iniciando contenedor de Logística..."
	docker run -d --name $(LOGISTICA_CONTAINER) \
		$(LOGISTICA_IMAGE)
	@echo "Contenedor de Logística en ejecución."

.PHONY: run-rabbitmq
run-rabbitmq:
	@echo "Iniciando contenedor de RabbitMQ..."
	docker run -d --name $(RABBITMQ_CONTAINER) \
		-p 5672:5672 -p 15672:15672 \
		$(RABBITMQ_IMAGE)
	@echo "Contenedor de RabbitMQ en ejecución."

# Detener y eliminar los contenedores

.PHONY: stop-caravanas
stop-caravanas:
	@echo "Deteniendo contenedor de Caravanas..."
	docker stop $(CARAVANAS_CONTAINER) && docker rm $(CARAVANAS_CONTAINER)
	@echo "Contenedor de Caravanas detenido y eliminado."

.PHONY: stop-clientes
stop-clientes:
	@echo "Deteniendo contenedor de Clientes..."
	docker stop $(CLIENTES_CONTAINER) && docker rm $(CLIENTES_CONTAINER)
	@echo "Contenedor de Clientes detenido y eliminado."

.PHONY: stop-finanzas
stop-finanzas:
	@echo "Deteniendo contenedor de Finanzas..."
	docker stop $(FINANZAS_CONTAINER) && docker rm $(FINANZAS_CONTAINER)
	@echo "Contenedor de Finanzas detenido y eliminado."

.PHONY: stop-logistica
stop-logistica:
	@echo "Deteniendo contenedor de Logística..."
	docker stop $(LOGISTICA_CONTAINER) && docker rm $(LOGISTICA_CONTAINER)
	@echo "Contenedor de Logística detenido y eliminado."

.PHONY: stop-rabbitmq
stop-rabbitmq:
	@echo "Deteniendo contenedor de RabbitMQ..."
	docker stop $(RABBITMQ_CONTAINER) && docker rm $(RABBITMQ_CONTAINER)
	@echo "Contenedor de RabbitMQ detenido y eliminado."

# Limpiar todas las imágenes y contenedores

.PHONY: clean
clean:
	@echo "Eliminando todas las imágenes y contenedores..."
	docker stop $(CARAVANAS_CONTAINER) $(CLIENTES_CONTAINER) $(FINANZAS_CONTAINER) $(LOGISTICA_CONTAINER) $(RABBITMQ_CONTAINER) || true
	docker rm $(CARAVANAS_CONTAINER) $(CLIENTES_CONTAINER) $(FINANZAS_CONTAINER) $(LOGISTICA_CONTAINER) $(RABBITMQ_CONTAINER) || true
	docker rmi $(CARAVANAS_IMAGE) $(CLIENTES_IMAGE) $(FINANZAS_IMAGE) $(LOGISTICA_IMAGE) $(RABBITMQ_IMAGE) || true
	@echo "Limpieza completada."
