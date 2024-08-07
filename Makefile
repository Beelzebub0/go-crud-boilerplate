run:
	go run src/cmd/main.go -config=./src/conf/conf.yaml

build:
	go build -o build/app src/cmd/main.go

run-docker-all:
	docker-compose -f docker-app.yaml up -d

stop-docker-all:
	docker-compose -f docker-app.yaml down

run-storage:
	docker-compose -f docker-storage.yaml up -d

init-storage:
	./files/sql/initdb.sh

stop-storage:
	docker-compose -f docker-storage.yaml down

image-build:
	go build -o build/app src/cmd/main.go && \
	docker build . -t ted/go-dummy-boilerplate
