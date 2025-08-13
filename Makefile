APP_NAME=go-microservice
ENV_FILE=docker.env

setup: env-config docker-up fix-filebeat-perms restart-filebeat
	@echo "✅ Environment setup complete. All services are running."

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
	else \
		echo "$(ENV_FILE) already exists"; \
	fi

docker-up:
	@echo ">> Building and starting containers"
	docker compose up --build -d --no-recreate

fix-filebeat-perms:
	@echo ">> Fixing filebeat.yml permissions before starting Filebeat..."
	@docker compose run --rm --entrypoint "" filebeat \
		sh -c "chmod go-w filebeat.yml && chown root:root filebeat.yml"
	@echo "Permissions fixed ✅"

restart-filebeat:
	@echo ">> Restarting Filebeat container..."
	@docker compose restart filebeat
	@echo "Filebeat restarted ✅"

stop:
	@echo ">> Stopping containers"
	docker compose down

logs:
	@docker compose logs -f

rebuild:
	@docker compose build --no-cache
