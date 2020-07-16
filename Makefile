
.PHONY:  docker

BIN := go-bindata
# BIN := go-bindata-assetfs

all:
	$(BIN) -pkg data -o data/bindata.go static static/images

dev:
	$(BIN) -debug -pkg data -o data/bindata.go static static/images
	go install github.com/jmcarbo/graphspace/...

dep:
	go get github.com/jteeuwen/go-bindata/go-bindata
	go get github.com/elazarl/go-bindata-assetfs/...

	
docker:
	# build docker image
	docker build -t graphspace:1.0 .
docker-push:
	docker tag graphspace:1.0 jmcarbo/graphspace:1.0
	docker push jmcarbo/graphspace:1.0

