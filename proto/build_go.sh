#!/bin/bash

protoc ./common/*.proto --go_out=plugins=grpc:./
protoc ./apigw/*.proto --go_out=plugins=grpc:./
protoc ./data/*.proto --go_out=plugins=grpc:./
protoc ./auth/*.proto --go_out=plugins=grpc:./
protoc ./msg/*.proto --go_out=plugins=grpc:./
