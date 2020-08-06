all : main.go
	go build -o lcal -ldflags '-w -s -linkmode external -extldflags "-static-libgcc"' -a -v -trimpath
