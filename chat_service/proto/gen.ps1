$ErrorActionPreference = "Stop"
$GoBin = (go env GOPATH) + "\bin"
$ProtoGenGo     = Join-Path $GoBin "protoc-gen-go.exe"
$ProtoGenGrpc   = Join-Path $GoBin "protoc-gen-go-grpc.exe"
if (!(Test-Path $ProtoGenGo)) { Write-Error "protoc-gen-go missing"; exit 1 }
if (!(Test-Path $ProtoGenGrpc)) { Write-Error "protoc-gen-go-grpc missing"; exit 1 }
$Dir = Split-Path -Parent $MyInvocation.MyCommand.Path
$Out = Join-Path $Dir "..\grpc"
$null = New-Item -ItemType Directory -Force -Path $Out
protoc --plugin=protoc-gen-go="$ProtoGenGo" --plugin=protoc-gen-go-grpc="$ProtoGenGrpc" `
  --proto_path="$Dir" `
  --go_out="$Out" --go_opt=paths=source_relative `
  --go-grpc_out="$Out" --go-grpc_opt=paths=source_relative `
  "$Dir\chat.proto"
Write-Host "Generated into: $Out"