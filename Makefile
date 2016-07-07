all: pbf docs

docs:
	mkdir -p ./pbf/docs
	protoc -I /usr/local/include -I ./pbf/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:./pbf/ \
		./pbf/*.proto
	bootprint openapi ./pbf/pokedex.swagger.json ./pbf/docs/
	html-inline ./pbf/docs/index.html > ./pbf/docs/static.html

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


pbf:
	@mkdir -p ./pbf
	protoc -I /usr/local/include -I ./pbf/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--gogoslick_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:./pbf/ \
		./pbf/*.proto
	protoc -I /usr/local/include -I ./pbf/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:./pbf/ \
		./pbf/*.proto

	@ls -htlr ./pbf/*

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
	pbf \
	run \
	vendor \
