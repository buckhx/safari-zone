# Safari Zone

A set of services demonstrating GRPC integration. There are no guarantees on any of this working as this is meant to be a learning exercise. I'd say it's 70% of the way there.

# Services

## Registry

The registry service is in charge of authorization/authentication.
It manages accounts and issues tokens (JWT) for other services to use.
The tokens are signed asymmetrically with ECDSA private key and the public key is available at an endpoint.
Tokens are scoped to roles.

## Pokedex

The pokedex service is meant to mimic a content API.
A user can only see info on pokemon they have caught, while a service has full access.

## Warden

The warden service is the meat of the gameplay and is meant to demonstrate a more complex service.
This service issues to tickets w/ expiries for battles.
The battles are delivered via bi-directional streams.

# Clients

## Bot

There is a bot client that manages the flow and user interaction via a state machine. Uses a pretty slick state machine pattern from [this](https://talks.golang.org/2011/lex.slide#1) talk.

## GUI

The gui has an embedded bot and exposes a UI using github.com/gizak/termui

## Other Neat things

* HTTP Gateway via grpc-gateway
* JWT Authorization Interceptor
* KVC Util
* Auto-gen service stubs

## Other Neat things

## Not Implemented

* Logging, Recovery Interceptors
* Api Proxy
* Meaningful Tests or Documentation

## Notes

* json openapi docs under service routes
* different packages for pbf & gw, to not muddy up vars
* UID is reproduces the same values every run
* openssl ecparam -out reg.pem -name secp256k1 -genkey (x509 lib can't read curve)
* gateway should translate grpc codes -> HTTP codes (16 -> 404)
* @authorize annotations in the proto would be awesome
* Slim docker containers
* k8s is not finished

```
OAK_KEY=$(curl -s "http://localhost:8080/registry/v0/trainer/52fdfc07/auth?scope=PROFESSOR" -H "Authorization: Basic NTJmZGZjMDc6c2FtK2RlbGlhNEVWRVI=" | jq --raw-output '.access')
ASH_KEY=$(curl -s "http://localhost:8080/registry/v0/trainer/037c4d7b/auth" -H "Authorization: Basic MDM3YzRkN2I6VEhFdmVyeWJlc3Q=" | jq --raw-output '.access')
curl "http://localhost:8080/registry/v0/trainer/52fdfc07" -H "Authorization: Bearer ${OAK_KEY}"; echo
curl "http://localhost:8080/registry/v0/trainer/037c4d7b" -H "Authorization: Bearer ${ASH_KEY}"; echo
```

Deploy
------

* go (gvm)
* protoc
* docker
* minikube
* xhyve
