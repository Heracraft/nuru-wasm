/wasm/modified/ Contains modified parts of the Nuru intepreter.

These modifications are made to make the interpreter work inside the browser's WebAssembly environment. 

## Changes
- [builtins.go] Modified the nuru 'jaza' function to make a system call to the browser's prompt function to get user input.