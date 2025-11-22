#!/usr/bin/env bash
set -euo pipefail
if ! command -v protoc >/dev/null 2>&1; then echo "protoc not found" >&2; exit 1; fi
DIR="$(cd "$(dirname "$0")" && pwd)"
PROTO_DIR="$DIR"
OUT_DIR="$DIR/../grpc"
mkdir -p "$OUT_DIR"
GOPATH="$(go env GOPATH || true)"
if [ -z "$GOPATH" ]; then echo "Go not found" >&2; exit 1; fi
GO_PLUGIN="$GOPATH/bin/protoc-gen-go"
GRPC_PLUGIN="$GOPATH/bin/protoc-gen-go-grpc"
[ -f "$GO_PLUGIN" ] || { echo "protoc-gen-go not found at $GO_PLUGIN" >&2; exit 1; }
[ -f "$GRPC_PLUGIN" ] || { echo "protoc-gen-go-grpc not found at $GRPC_PLUGIN" >&2; exit 1; }
protoc --plugin=protoc-gen-go="$GO_PLUGIN" --plugin=protoc-gen-go-grpc="$GRPC_PLUGIN" --proto_path="$PROTO_DIR" --go_out="$OUT_DIR" --go_opt=paths=source_relative --go-grpc_out="$OUT_DIR" --go-grpc_opt=paths=source_relative "$PROTO_DIR/chat.proto"