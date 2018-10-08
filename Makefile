.PHONY:all apigw client

all: apigw client

apigw: 
	go build -i -o bin/apigw ./apigw

client: 
	go build -i -o bin/client ./client

clean: 
	@rm -rf bin

docker:
	docker build --force-rm -t antalk-image .
