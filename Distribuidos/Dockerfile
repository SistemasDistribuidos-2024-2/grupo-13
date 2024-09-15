# Usa la imagen oficial de RabbitMQ con la interfaz de administración
FROM rabbitmq:3-management

# Comando para el healthcheck
HEALTHCHECK --interval=10s --timeout=5s --retries=5 \
  CMD rabbitmqctl status || exit 1

# Exponer los puertos necesarios
EXPOSE 15672 5672
