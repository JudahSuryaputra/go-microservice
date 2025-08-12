APP_NAME=go-microservice
ENV_FILE=docker.env
FILEBEAT_CONFIG=filebeat.yml

setup: env-config filebeat-perms docker-up

env-config:
	@echo ">> Ensuring $(ENV_FILE) exists"
	@if [ ! -f $(ENV_FILE) ]; then \
		echo 'APP_MODE=docker' > $(ENV_FILE); \
		echo 'DB_HOST=db' >> $(ENV_FILE); \
		echo 'DB_PORT=5432' >> $(ENV_FILE); \
		echo 'DB_USER=postgres' >> $(ENV_FILE); \
		echo 'DB_PASSWORD=postgres' >> $(ENV_FILE); \
		echo 'DB_NAME=postgres' >> $(ENV_FILE); \
		echo 'REDIS_HOST=redis' >> $(ENV_FILE); \
		echo 'REDIS_PORT=6379' >> $(ENV_FILE); \
		echo "Created default $(ENV_FILE)"; \
	fi

filebeat-perms:
	@echo ">> Setting correct permissions for $(FILEBEAT_CONFIG)"
	@if [ -f $(FILEBEAT_CONFIG) ]; then \
		chmod go-w $(FILEBEAT_CONFIG); \
	else \
		echo "WARNING: $(FILEBEAT_CONFIG) not found"; \
	fi

docker-up:
	@echo ">> Building and starting containers"
	docker compose up --build -d

stop:
	@echo ">> Stopping containers"
	docker compose down

logs:
	@docker compose logs -f

rebuild:
	@docker compose build --no-cache

