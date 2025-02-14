/wasm/modified/ Contains modified parts of the Nuru intepreter.

These modifications are made to make the interpreter work inside the browser's WebAssembly environment. 

## Changes
- [builtins.go] 
    - Modified the nuru 'jaza' function to make a system call to the browser's prompt function to get user input.
    - Replaced the default 'andika' function (print equivalent) to return output instead of printing to the console. This is so that we don't have to intercept console logs to obtain code output