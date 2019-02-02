#!/bin/bash

grpcc \
  -i \
  --address localhost:8080 \
  --proto gen/grpc/calc/pb/calc.proto \
  --eval 'client.add({a:1,b:2},printReply)'
