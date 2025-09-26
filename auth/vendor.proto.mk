# Путь до завендореных protobuf файлов
VENDOR_PROTO_PATH := $(CURDIR)/vendor.protobuf

SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command

# vendor
vendor: .vendor-reset .vendor-googleapis .vendor-google-protobuf .vendor-protovalidate .vendor-protoc-gen-openapiv2 .vendor-tidy

# delete VENDOR_PROTO_PATH
.vendor-reset:
	powershell -Command "if (Test-Path '$(VENDOR_PROTO_PATH)') { Remove-Item -Recurse -Force '$(VENDOR_PROTO_PATH)' }"
	powershell -Command "New-Item -ItemType Directory -Path '$(VENDOR_PROTO_PATH)' -Force | Out-Null"

# Устанавливаем proto описания google/protobuf
.vendor-google-protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 ^
		https://github.com/protocolbuffers/protobuf "$(VENDOR_PROTO_PATH)/protobuf" && ^
	cd "$(VENDOR_PROTO_PATH)/protobuf" && ^
	git sparse-checkout set --no-cone src/google/protobuf && ^
	git checkout
	powershell -Command "New-Item -ItemType Directory -Path '$(VENDOR_PROTO_PATH)/google/protobuf' -Force | Out-Null"
	powershell -Command "Get-ChildItem -Path '$(VENDOR_PROTO_PATH)/protobuf/src/google/protobuf' -Filter '*.proto' -File | ForEach-Object { Move-Item $$_.FullName '$(VENDOR_PROTO_PATH)/google/protobuf' -Force }"
	powershell -Command "Remove-Item -Recurse -Force '$(VENDOR_PROTO_PATH)/protobuf'"

# Устанавливаем proto описания google/api
.vendor-googleapis:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 ^
		https://github.com/googleapis/googleapis "$(VENDOR_PROTO_PATH)/googleapis" && ^
	cd "$(VENDOR_PROTO_PATH)/googleapis" && ^
	git sparse-checkout set --no-cone google/api && ^
	git checkout
	powershell -Command "New-Item -ItemType Directory -Path '$(VENDOR_PROTO_PATH)/google/api' -Force | Out-Null"
	powershell -Command "Get-ChildItem -Path '$(VENDOR_PROTO_PATH)/googleapis/google/api' -Filter '*.proto' -File | ForEach-Object { Move-Item $$_.FullName '$(VENDOR_PROTO_PATH)/google/api' -Force }"
	powershell -Command "Remove-Item -Recurse -Force '$(VENDOR_PROTO_PATH)/googleapis'"

# Устанавливаем proto описания protoc-gen-openapiv2/options
.vendor-protoc-gen-openapiv2:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 ^
		https://github.com/grpc-ecosystem/grpc-gateway "$(VENDOR_PROTO_PATH)/grpc-gateway" && ^
	cd "$(VENDOR_PROTO_PATH)/grpc-gateway" && ^
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && ^
	git checkout
	powershell -Command "New-Item -ItemType Directory -Path '$(VENDOR_PROTO_PATH)/protoc-gen-openapiv2' -Force | Out-Null"
	powershell -Command "if (Test-Path '$(VENDOR_PROTO_PATH)/grpc-gateway/protoc-gen-openapiv2/options') { Move-Item '$(VENDOR_PROTO_PATH)/grpc-gateway/protoc-gen-openapiv2/options' '$(VENDOR_PROTO_PATH)/protoc-gen-openapiv2/' -Force }"
	powershell -Command "Remove-Item -Recurse -Force '$(VENDOR_PROTO_PATH)/grpc-gateway'"

# Устанавливаем proto описания buf/validate/validate.proto
.vendor-protovalidate:
	git clone -b main --single-branch --depth=1 --filter=tree:0 ^
		https://github.com/bufbuild/protovalidate "$(VENDOR_PROTO_PATH)/protovalidate" && ^
	cd "$(VENDOR_PROTO_PATH)/protovalidate" && ^
	git checkout
	powershell -Command "if (Test-Path '$(VENDOR_PROTO_PATH)/protovalidate/proto/protovalidate/buf') { Move-Item '$(VENDOR_PROTO_PATH)/protovalidate/proto/protovalidate/buf' '$(VENDOR_PROTO_PATH)/' -Force }"
	powershell -Command "Remove-Item -Recurse -Force '$(VENDOR_PROTO_PATH)/protovalidate'"

# delete all non .proto files
.vendor-tidy:
	powershell -Command "Get-ChildItem -Path '$(VENDOR_PROTO_PATH)' -Recurse -File | Where-Object { $$_.Extension -ne '.proto' } | Remove-Item -Force"
	powershell -Command "Get-ChildItem -Path '$(VENDOR_PROTO_PATH)' -Recurse -File | Where-Object { $$_.Name -like '*unittest*' -or $$_.Name -like '*test*' -or $$_.Name -like '*sample*' } | Remove-Item -Force"
	powershell -Command "Get-ChildItem -Path '$(VENDOR_PROTO_PATH)' -Recurse -Directory | Where-Object { (Get-ChildItem -Path $$_.FullName).Count -eq 0 } | Remove-Item -Force"

# Объявляем, что текущие команды не являются файлами и
# интсрументируем Makefile не искать изменения в файловой системе
.PHONY: ^
	.vendor-reset ^
	.vendor-google-protobuf ^
	.vendor-googleapis ^
	.vendor-protoc-gen-openapiv2 ^
	.vendor-protovalidate ^
	.vendor-tidy ^
	vendor