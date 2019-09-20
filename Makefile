
build: 
	GOOS=linux GOARCH=amd64 go build -o terraform-provider-p2pub-linux-amd64
	GOOS=windows GOARCH=amd64 go build -o terraform-provider-p2pub-windows-amd64
	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-p2pub-darwin-amd64

test:
	go test ./...

clean:
	rm terraform-provider-p2pub-linux-amd64 terraform-provider-p2pub-windows-amd64 terraform-provider-p2pub-darwin-amd64
