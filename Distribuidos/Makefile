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


.PHONY: run-caravanas
run-caravanas:
	@echo "Iniciando contenedor de Caravanas con tiempo de operación..."
	docker run -d --name $(CARAVANAS_CONTAINER) \
		-e TIEMPO_OPERACION=$(TIEMPO_OPERACION) \
		$(CARAVANAS_IMAGE)
	@echo "Contenedor de Caravanas en ejecución con tiempo de operación $(TIEMPO_OPERACION)."

.PHONY: run-clientes
run-clientes:
	@echo "Iniciando contenedor de Clientes con tiempo de operación..."
	docker run -d --name $(CLIENTES_CONTAINER) \
		-it -e TIEMPO_OPERACION=$(TIEMPO_OPERACION) \
		$(CLIENTES_IMAGE)
	@echo "Contenedor de Clientes en ejecución con tiempo de operación $(TIEMPO_OPERACION)."

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
