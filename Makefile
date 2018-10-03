build:
	find . -name '*.proto' -prune -type f | xargs -IX protoc -I. -I$$GOPATH/src/ X --go_out=plugins=grpc:$$GOPATH/src/		
	find . -name '*.pb.go' -prune -type f | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'