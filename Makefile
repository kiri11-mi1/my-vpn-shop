up:
	docker-compose up -d --build

down:
	docker-compose down

restart:
	docker-compose down && docker-compose up -d --build

test:
	docker-compose exec bot go test ./tests/...
