.PHONY: help run stop restart build rebuild logs clean ps shell fresh

# Default target
help:
	@echo "Available commands:"
	@echo "  make run         - Start the container"
	@echo "  make stop        - Stop the container"
	@echo "  make restart     - Restart the container"
	@echo "  make build       - Build the Docker image"
	@echo "  make rebuild     - Rebuild the image without cache"
	@echo "  make logs        - Show container logs"
	@echo "  make ps          - Show running containers"
	@echo "  make clean       - Remove container, image, and prune system"
	@echo "  make shell       - Enter container shell"
	@echo "  make fresh       - Rebuild and restart container"

# Variables
IMAGE_NAME=docker-go-fiber
CONTAINER_NAME=docker-go-container

# Start container
run:
	docker run -d --name $(CONTAINER_NAME) -p 5000:5000 $(IMAGE_NAME)

# Stop container
stop:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

# Restart container
restart: stop run

# Build image

build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME) .

# Rebuild image without cache
rebuild:
	docker build --no-cache -t $(IMAGE_NAME) .

# Show logs
logs:
	docker logs -f $(CONTAINER_NAME)

# Show running containers
ps:
	docker ps

# Clean up everything
clean:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true
	docker rmi $(IMAGE_NAME) || true
	docker system prune -f

# Enter container shell
shell:
	docker exec -it $(CONTAINER_NAME) sh

# Rebuild and restart
fresh: stop rebuild run
