#!/bin/bash

echo "[start to build hrms]"
rm -r hrms_app
echo "[start to build macos_app]"
go build -o hrms_app/macos_app main.go
chmod a+x hrms_app/macos_app

echo "[start to build windows_app]"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o hrms_app/windows_app main.go
chmod a+x hrms_app/windows_app

echo "[start to build linux_app]"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hrms_app/linux_app main.go
chmod a+x hrms_app/linux_app
echo "[start hrms success]"

echo "[start to copy folder]"
cp -R config hrms_app/config
cp -R static hrms_app/static
cp -R views hrms_app/views
echo "[copy folder success]"
