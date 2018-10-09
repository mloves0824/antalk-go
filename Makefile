.PHONY:all apigw data client

all: apigw data client

apigw: 
	go build -i -o bin/apigw ./apigw

data: 
	go build -i -o bin/data ./data

client: 
	go build -i -o bin/client ./client

clean: 
	@rm -rf bin

docker:
	docker build --force-rm -t antalk-image .
