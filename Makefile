IMAGE_NAME = test-hls:latest

run:
	go run ./cmd/main.go

build-docker:
	docker build -t $(IMAGE_NAME) -f ./build/Dockerfile .

run-docker:
	docker run -p 8090:8090 --rm $(IMAGE_NAME)