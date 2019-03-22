all: docker clean
VERSION="v0.0.2"
push-image:
	docker push quay.io/ksimon/kubevirt-cpu-nfd-plugin:${VERSION}

image: binary
	docker build -t quay.io/ksimon/kubevirt-cpu-nfd-plugin:${VERSION} .

binary: test
	go build cmd/cpu-model-nfd-plugin/kubevirt-cpu-nfd-plugin.go

test:
	go test ./...

clean:
	rm -f kubevirt-cpu-nfd-plugin

.PHONY: all push docker binary test clean
