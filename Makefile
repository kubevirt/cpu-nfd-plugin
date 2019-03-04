all: docker clean
VERSION="0.0.2"
push-image:
	docker push quay.io/ksimon/kubevirt-cpu-model-nfd-plugin:${VERSION}

image: binary
	docker build -t quay.io/ksimon/kubevirt-cpu-model-nfd-plugin:${VERSION} .

binary: test
	go build cmd/cpu-model-nfd-plugin/cpu-model-nfd-plugin.go

test:
	go test ./...

clean:
	rm -f cpu-model-nfd-plugin

.PHONY: all push docker binary test clean
