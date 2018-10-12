.PHONY:all apigw data client auth msg
 
all: apigw data client auth msg

apigw: 
	go build -i -o bin/apigw ./apigw

data: 
	go build -i -o bin/data ./data

auth: 
	go build -i -o bin/auth ./auth

client: 
	go build -i -o bin/client ./client

msg: 
	go build -i -o bin/msg ./msg

clean: 
	@rm -rf bin

docker:
	docker build --force-rm -t antalk-image .
