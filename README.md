# Safari Zone

A set of services demonstrating GRPC integration


## TODO

* Busines logic
* Containerize services


## Notes

* auto-gen service stubs
* json openapi docs under service routes
* different packages for pbf & gw, to not muddy up vars
* UID is reproduces the same values every run
* openssl ecparam -out reg.pem -name secp256k1 -genkey
* gateway should translate grpc codes -> HTTP codes (16 -> 404)
