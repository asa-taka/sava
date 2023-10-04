IMG=astk03/sava

dev:
	air

build:
	go build -o sava

docker-build:
	docker build -t $(IMG) .

docker-run:
	docker run -p 3000:3000 -v $(shell pwd)/data:/data $(IMG) --host 0.0.0.0 --port 3000

docker-push:
	docker push $(IMG)
