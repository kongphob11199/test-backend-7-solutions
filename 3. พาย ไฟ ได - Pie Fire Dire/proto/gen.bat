@echo off
set PROTO_FILE_BEEF=beef.proto

set OUT_DIR_SERVER=../internal

protoc %PROTO_FILE_BEEF% --go_out=%OUT_DIR_SERVER% --go-grpc_out=%OUT_DIR_SERVER%

echo Compilation complete!
pause