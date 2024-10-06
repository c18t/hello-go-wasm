# hello-go-wasm

Go WASMの素振りです

## Usage

```console
mise install -y
mise run setup
mise run build
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" bin/
mise run server
open http://localhost:8080/work/
```

## Project Making

```console
mkdir work
touch work/index.html
```
