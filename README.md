# Safari Zone

A set of services demonstrating GRPC integration


## TODO

* Busines logic
* Containerize services

## Not Implemented

* Logging Interceptor
* Panic Revcovery Interceptor

## Notes

* auto-gen service stubs
* json openapi docs under service routes
* different packages for pbf & gw, to not muddy up vars
* UID is reproduces the same values every run
* openssl ecparam -out reg.pem -name secp256k1 -genkey (x509 lib can't read curve)
* gateway should translate grpc codes -> HTTP codes (16 -> 404)
* @authorize annotations in the proto would be awesome

```
OAK_KEY=$(curl -s "http://localhost:8080/registrysrv/v0/trainer/52fdfc07/auth?scope=PROFESSOR" -H "Authorization: Basic NTJmZGZjMDc6c2FtK2RlbGlhNEVWRVI=" | jq --raw-output '.access')
ASH_KEY=$(curl -s "http://localhost:8080/registrysrv/v0/trainer/037c4d7b/auth" -H "Authorization: Basic MDM3YzRkN2I6VEhFdmVyeWJlc3Q=" | jq --raw-output '.access')
curl "http://localhost:8080/registrysrv/v0/trainer/52fdfc07" -H "Authorization: Bearer ${OAK_KEY}"; echo
curl "http://localhost:8080/registrysrv/v0/trainer/037c4d7b" -H "Authorization: Bearer ${ASH_KEY}"; echo
```
