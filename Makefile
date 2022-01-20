docker-run:
	docker rm -f redis-db
	docker rm -f url-shortener-api
	docker-compose -f  docker-compose.yml down --remove-orphans --volumes
	docker-compose up --force-recreate --build