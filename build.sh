#!/bin/bash
echo "[start to build hrms]"
go build -o output main.go
chmod a+x output/main
./output/main
echo "[start hrms success]"