#!/bin/bash
protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto
# you can run this file as a bash script on linux or you can just copy the line above and run it in a console like any other command
# P.S. make sure that you have protoc installed
