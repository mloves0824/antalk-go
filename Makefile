.PHONY:all apigw data client auth
 
all: apigw data client auth

apigw: 
	go build -i -o bin/apigw ./apigw

data: 
	go build -i -o bin/data ./data

auth: 
	go build -i -o bin/auth ./auth

client: 
	go build -i -o bin/client ./client

clean: 
	@rm -rf bin

docker:
	docker build --force-rm -t antalk-image .
