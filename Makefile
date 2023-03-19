
default: build run.local

build: build.apidocs build.server

build.apidocs:
	apidoc -i internal -o apidoc

build.server:
	go build -o server

run.local: deps.up
	./server

deps.up:
	docker-compose up -d

deps.down:
	docker-compose down -v
