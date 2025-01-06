A web interpreter for [Nuru](https://github.com/NuruProgramming/Nuru) -- A Swahili Programming Language built from the ground up -- powered by WebAssembly.  

### Getting started

**Prerequisites**
+ Go (^1.19.0)
+ Node.js (^18.13)

**Working with the Wasm Interpreter**
To build the wasm binary from the go interpreter: 

1. Change directories to `/wasm`

```shell
cd wasm
```

2. Install the required go dependencies
   
``` shell
   go mod tidy
```

3. To build the wasm binary

```shell
wasm && GOOS=js GOARCH=wasm go build -o main.wasm
```

or if you are on **Windows**:

```shell
$env:GOOS="js"; $env:GOARCH="wasm"; go build -o main.wasm
```

**Web app**
Powered by [Svelte](https://svelte.dev/). To work with it:

1. Change directories to `/app`
   
```shell
   cd app
```

2. Install dependencies

```shell
npm i
```

3. To start a development server:

```bash
npm run dev

# or start the server and open the app in a new browser tab

npm run dev -- --open
```

**Workspace scripts**
[/package.json](https://github.com/Heracraft/nuru-playground/blob/0af828dc5a8fff8eb92ecb8ccd188eecd1cf1dc6/package.json#L6) redefines the above commands as scripts you can run from the root folder.

```shell
npm run dev #runs the dev server
npm run build:wasm #Builds the wasm binary and copies it over to the web app
```

> If you are on windows, edit the [build:wasm](https://github.com/Heracraft/nuru-playground/blob/0af828dc5a8fff8eb92ecb8ccd188eecd1cf1dc6/package.json#L11C18-L11C101) script and replace `cp` with `copy` or `Copy-Item`
