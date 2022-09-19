# WSR: WASM standalone runtime based on [wazero](https://github.com/tetratelabs/wazero/)

## Motivation

The main goal is to be able to run WASM modules on old arm devices such as Pi Zero.

Pi Zero (2015) uses BCM2835 which contains [ARM1176JZF-S](https://developer.arm.com/documentation/ddi0301/h/introduction/about-the-processor) processor which implements the ARM11 ARM architecture v6. armv6 (armv6l precisely) is outdated and not a single modern WASM runtime is able to compile to it ([wasmtime](https://github.com/bytecodealliance/wasmtime/issues/1173), [wasmer](https://github.com/wasmerio/wasmer/issues/1652), wasmedge [works in theory](https://github.com/WasmEdge/WasmEdge/releases/tag/0.8.1) but crushes).

wazero can help with that since it's a runtime written completely in go (so zero deps) and go compiler is great at cross-compilation and has decent [arm support](https://github.com/golang/go/wiki/GoArm).

## Usage

To use it compile it and run. Provide wasm file path as first argument to the program. If necessary you can specify entry function using flag `-entrypoint=functionName` (specify before file path).
