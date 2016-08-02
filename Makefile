SRV_BIN="safari-srv"
VERSION=`git describe --always --tags`
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=--ldflags "-extldflags '-static' -X main.Version=${VERSION}"
PACKAGES=`go list ./... | grep -v /vendor/`

all: proto docs build

build: 
	mkdir -p ./dist
	CGO_ENABLED=0 go build -v ${LDFLAGS} -o ./dist/${SRV_BIN} cmd/srv/cli.go
	chmod +x ./dist/${SRV_BIN}

clean:
	rm -rf ./dist
	rm -rf ./proto/pbf/*
	rm -rf ./proto/docs/*
	docker rmi -f $(shell docker images -f dangling=true -q)

images: 
	docker build -f dev/srv.docker -t safari/srv:fat .
	docker export $(shell docker run -d safari/srv:fat /bin/true) | docker import - safari/srv
	docker build -f dev/registry.docker -t safari/registry .
	docker build -f dev/pokedex.docker -t safari/pokedex .
	docker build -f dev/warden.docker -t safari/warden .
	docker build -f dev/gateway.docker -t safari/gateway .

compose:
	docker-compose -f dev/docker-compose.yml up

eckey:
	#openssl ecparam -out dev/reg.pem -name secp256k1 -genkey -noout
	go run cmd/mint/cli.go > dev/reg.pem

docs:
	mkdir -p ./proto/docs
	protoc -I /usr/local/include -I ./proto/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:./proto/docs \
		./proto/*.proto
	bootprint openapi ./proto/*.swagger.json ./proto/docs/
	#html-inline ./pbf/docs/index.html > ./pbf/docs/static.html

get:
	# govendor
	go get -u -v github.com/kardianos/govendor
	# TODO proto
	go get -u -v github.com/golang/protobuf/protoc-gen-go
	# gogoproto
	go get -u -v github.com/gogo/protobuf/proto
	go get -u -v github.com/gogo/protobuf/protoc-gen-gogo
	go get -u -v github.com/gogo/protobuf/gogoproto
	go get -u -v github.com/gogo/protobuf/protoc-gen-gogoslick
	# grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	# docs
	# npm install -g bootprint bootprint-openapi html-inline

k8s-env:
	minikube start --vm-driver=xhyve
	$(shell dvm install 1.11.1)
	$(shell eval $(minikube docker-env))
	//shell make docker)

k8s:
	kubectl apply -f dev/cluster/
	


proto:
	@mkdir -p ./proto/pbf
	protoc -I ./proto/ -I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogoslick_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:./proto/pbf \
		./proto/*.proto
	protoc -I ./proto -I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:./proto/pbf \
		./proto/*.proto

run:
	go run cmd/srv/cli.go

test:
	go test -v $(PACKAGES)

vendor:
	govendor init
	govendor add -v +external
	govendor update -v +vendor


.PHONY: \
	all \
	images \
	docs \
	get \
	proto \
	run \
	vendor \
