#!/bin/bash

# generate proto file golang
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.