gobuild:
	go build ./cmd/server/main.go

gorun:
	go run ./cmd/server/main.go

build:
	docker build -t my-mysql-container .

run:
	docker run -d --name my-mysql-container -p 3306:3306 my-mysql-container

clean:
	docker stop my-mysql-container
	docker rm my-mysql-container
