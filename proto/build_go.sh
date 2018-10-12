#!/bin/bash

protoc ./common/*.proto --go_out=plugins=grpc:./
protoc ./apigw/*.proto --go_out=plugins=grpc:./
protoc ./data/*.proto --go_out=plugins=grpc:./
protoc ./auth/*.proto --go_out=plugins=grpc:./
protoc ./msg/*.proto --go_out=plugins=grpc:./
protoc ./seq/*.proto --go_out=plugins=grpc:./

sed -i 's/common \"common\"/common \"github.com\/mloves0824\/antalk-go\/proto\/common\"/' ./apigw/apigw.pb.go ./data/data.pb.go ./auth/auth.pb.go ./msg/msg.pb.go ./seq/seq.pb.go
