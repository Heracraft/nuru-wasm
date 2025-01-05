A wasm compiler/intepreter for [Nuru](https://github.com/NuruProgramming/Nuru)

Meant to power a web Nuru sandbox.

### Roadmap
- [x] Separate Nuru core from repl
- [ ] Get Wasm compiling!!!
- [ ] Import & use the wasm binary in a browser.
- [ ] Figure out weird array behaviour when acessing array length (`arr.idadi()`)
    refer to [sorting algorithm](https://nuruprogramming.org/en/kanuni-za-upangaji)

> Might need to rely on a build script for wasm compilation for windows support. 'cp' is not available on windows and it is used to copy the copiled wasm binary to the `static` folder