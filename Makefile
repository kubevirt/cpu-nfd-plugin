all: docker clean
VERSION=`git describe --tags`
push-image:
	docker push quay.io/kubevirt/cpu-nfd-plugin:${VERSION}

image:
	docker build -t quay.io/kubevirt/cpu-nfd-plugin:${VERSION} .

binary: test
	go build cmd/cpu-nfd-plugin/cpu-nfd-plugin.go

test:
	go test ./...

clean:
	rm -f cpu-nfd-plugin

.PHONY: all push docker binary test clean
