gen-go:
	protoc --go_out=. --go-grpc_out=. proto/main.proto
gen-python:
	protoc --python_out=python proto/main.proto


gen-normalize: gen-go gen-python

gen-auth:
	protoc --go_out=. --go-grpc_out=. proto/auth.proto