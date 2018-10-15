all: push

TAG = 0.1
PREFIX = registry.hundsun.com/hcs/mem_alloc

build: mem_alloc.go
	CGO_ENABLED=0 go build -a ./mem_alloc.go

container: build
#	docker build -f Dockerfile-scratch -t $(PREFIX):$(TAG) .
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

clean:
	rm -f mem_alloc
