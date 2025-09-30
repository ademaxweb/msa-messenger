@echo off
setlocal enabledelayedexpansion

set VENDOR_PROTO_PATH=%~dp0vendor.protobuf

echo [1/6] Cleaning vendor directory...
if exist "%VENDOR_PROTO_PATH%" rmdir /s /q "%VENDOR_PROTO_PATH%"
mkdir "%VENDOR_PROTO_PATH%"

echo [2/6] Cloning google/protobuf...
git clone -b main --single-branch -n --depth=1 --filter=tree:0 ^
    https://github.com/protocolbuffers/protobuf "%VENDOR_PROTO_PATH%/protobuf"
cd "%VENDOR_PROTO_PATH%/protobuf"
git sparse-checkout set --no-cone src/google/protobuf
git checkout
mkdir "%VENDOR_PROTO_PATH%/google/protobuf"
for /r "%VENDOR_PROTO_PATH%/protobuf/src/google/protobuf" %%f in (*.proto) do (
    move "%%f" "%VENDOR_PROTO_PATH%/google/protobuf/"
)
rmdir /s /q "%VENDOR_PROTO_PATH%/protobuf"

echo [3/6] Cloning googleapis...
git clone -b master --single-branch -n --depth=1 --filter=tree:0 ^
    https://github.com/googleapis/googleapis "%VENDOR_PROTO_PATH%/googleapis"
cd "%VENDOR_PROTO_PATH%/googleapis"
git sparse-checkout set --no-cone google/api
git checkout
mkdir "%VENDOR_PROTO_PATH%/google/api"
for /r "%VENDOR_PROTO_PATH%/googleapis/google/api" %%f in (*.proto) do (
    move "%%f" "%VENDOR_PROTO_PATH%/google/api/"
)
rmdir /s /q "%VENDOR_PROTO_PATH%/googleapis"

echo [4/6] Cloning grpc-gateway...
git clone -b main --single-branch -n --depth=1 --filter=tree:0 ^
    https://github.com/grpc-ecosystem/grpc-gateway "%VENDOR_PROTO_PATH%/grpc-gateway"
cd "%VENDOR_PROTO_PATH%/grpc-gateway"
git sparse-checkout set --no-cone protoc-gen-openapiv2/options
git checkout
mkdir "%VENDOR_PROTO_PATH%/protoc-gen-openapiv2"
if exist "%VENDOR_PROTO_PATH%/grpc-gateway/protoc-gen-openapiv2/options" (
    move "%VENDOR_PROTO_PATH%/grpc-gateway/protoc-gen-openapiv2/options" "%VENDOR_PROTO_PATH%/protoc-gen-openapiv2/"
)
rmdir /s /q "%VENDOR_PROTO_PATH%/grpc-gateway"

echo [5/6] Cloning protovalidate...
git clone -b main --single-branch --depth=1 --filter=tree:0 ^
    https://github.com/bufbuild/protovalidate "%VENDOR_PROTO_PATH%/protovalidate"
cd "%VENDOR_PROTO_PATH%/protovalidate"
git checkout

:: Используем xcopy для надежного копирования папки buf
if exist "%VENDOR_PROTO_PATH%/protovalidate/proto/protovalidate/buf" (
    mkdir "%VENDOR_PROTO_PATH%/buf"
    xcopy /s /e /i /y "%VENDOR_PROTO_PATH%/protovalidate/proto/protovalidate/buf" "%VENDOR_PROTO_PATH%/buf\"
    rmdir /s /q "%VENDOR_PROTO_PATH%/protovalidate/proto/protovalidate/buf"
)
rmdir /s /q "%VENDOR_PROTO_PATH%/protovalidate"

echo [6/6] Cleaning up non-proto files...
powershell -Command "Get-ChildItem -Path '%VENDOR_PROTO_PATH%' -Recurse -File | Where-Object { $_.Extension -ne '.proto' } | Remove-Item -Force"
powershell -Command "Get-ChildItem -Path '%VENDOR_PROTO_PATH%' -Recurse -File | Where-Object { $_.Name -like '*unittest*' -or $_.Name -like '*test*' -or $_.Name -like '*sample*' } | Remove-Item -Force"
REM powershell -Command "Get-ChildItem -Path '%VENDOR_PROTO_PATH%' -Recurse -Directory | Where-Object { (Get-ChildItem -Path $_.FullName).Count -eq 0 } | Remove-Item -Force"

echo.
echo ✅ Vendor setup completed successfully!
echo.