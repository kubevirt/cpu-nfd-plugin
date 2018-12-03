all: push clean

push: docker
	docker push quay.io/ksimon/cpu-model-nfd-plugin:latest

docker: binary
	docker build -t quay.io/ksimon/cpu-model-nfd-plugin:latest .

binary: test
	go build cmd/cpu-model-nfd-plugin/cpu-model-nfd-plugin.go

test:
	go test ./...

clean:
	rm -f cpu-model-nfd-plugin

.PHONY: all push docker binary test clean
