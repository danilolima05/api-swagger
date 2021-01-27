checkinstall:
	which swagger || go get -u https://github.com/go-swagger/go-swagger/cmd/swagger


swagger:
	swagger generate spec -o ./swagger.yaml --scan-models