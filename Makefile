gen-go:
	protoc --go_out=. --go-grpc_out=. proto/main.proto
gen-python:
	protoc --python_out=python proto/main.proto


gen: gen-go gen-python
