all: push clean

push: docker
	docker push quay.io/ksimon/nfd-host-supported-cpus:latest

docker: binary
	docker build -t quay.io/ksimon/nfd-host-supported-cpus:latest .

binary: test
	go build cmd/nfd-host-cpus/nfd-host-cpus.go

test:
	go test ./...

clean:
	rm -f nfd-host-cpus

.PHONY: all push docker binary test clean
