build:
	go build -o api cmd/api/main.go

clean:
	rm api

.PHONY: clean
