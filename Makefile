postgres:
	docker run --name postgres1 --network api-network -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:alpine
user-api: postgres
	docker run --name rest-api1 --network api-network -p 8080:8080 -e HOST=postgres1 matejelenc/rest-api:latest
user-api: postgres
	docker run --name rest-api1 --network api-network -p 8080:8080 -e HOST=postgres1 matejelenc/rest-api:latest
check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models