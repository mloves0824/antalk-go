.PHONY:all apigw data client auth msg seq
 
all: apigw data client auth msg seq

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

seq: 
	go build -i -o bin/seq ./cmd/seq

clean: 
	@rm -rf bin

docker:
	docker build --force-rm -t antalk-image .
