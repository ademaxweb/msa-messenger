@echo off
setlocal enabledelayedexpansion

echo [I] Vendoring files
REM call vendor.proto.bat

REM ���� ���������� �������� ��������
timeout /t 2 /nobreak >nul

REM ���� �� protobuf ������
set PROTO_PATH=.\api
set PKG_PROTO_PATH=.\pkg
set VENDOR_PROTO_PATH=.\vendor.protobuf

echo [II] Installing plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

echo [III] Generating...
if not exist "%PKG_PROTO_PATH%" mkdir "%PKG_PROTO_PATH%"

protoc -I "%VENDOR_PROTO_PATH%" --proto_path=. ^
--go_out="%PKG_PROTO_PATH%" --go_opt paths=source_relative ^
--go-grpc_out="%PKG_PROTO_PATH%" --go-grpc_opt paths=source_relative ^
--grpc-gateway_out="%PKG_PROTO_PATH%" --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ^
"%PROTO_PATH%\chat\v1\messages.proto" "%PROTO_PATH%\chat\v1\service.proto"

go mod tidy

echo [Done!]