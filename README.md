# hello-go-wasm

Go WASM の素振りです

## Usage

```console
cp .env.sample .env
mise install -y
mise run setup
mise run build hello
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" bin/
mise run server
open http://localhost:8080/work/
```

## Project Making

```console
mkdir work
touch work/index.html
```
