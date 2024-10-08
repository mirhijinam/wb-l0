_COMPOSE=docker compose -f docker-compose.yaml --env-file .env -p wb-l0

up:
	@echo "Starting Docker images..."
	${_COMPOSE} up

build:
	@echo "Building Docker images..."
	${_COMPOSE} build

down:
	${_COMPOSE} down -v

clean:
	${_COMPOSE} down --remove-orphans -v --rmi all

