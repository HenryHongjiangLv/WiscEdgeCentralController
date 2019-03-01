#!/bin/bash

protoc -I heartbeat/ heartbeat/heartbeat.proto --go_out=plugins=grpc:heartbeat
