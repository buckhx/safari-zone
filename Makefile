all: proto docs

clean:
	rm -rf ./proto/pbf/*
	rm -rf ./proto/docs/*

docs:
	mkdir -p ./proto/docs
	protoc -I /usr/local/include -I ./proto/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
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
	go get -u github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/gengo/grpc-gateway/protoc-gen-swagger
	# docs
	# npm install -g bootprint bootprint-openapi html-inline


proto:
	@mkdir -p ./proto/pbf
	protoc -I ./proto/ -I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--gogoslick_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:./proto/pbf \
		./proto/*.proto
	protoc -I ./proto -I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:./proto/pbf \
		./proto/*.proto

run:
	go run cmd/srv/cli.go

vendor:
	govendor init
	govendor add -v +external
	govendor update -v +vendor


.PHONY: \
	all \
	docs \
	get \
	proto \
	run \
	vendor \
