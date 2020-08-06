all : main.go
	go build -o cal -ldflags '-w -s -linkmode external -extldflags "-static-libgcc"' -a -v -trimpath
