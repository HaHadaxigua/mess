package main

import (
	"fmt"
	"os"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	wasmBytes, err := os.ReadFile("./wasm_go/module.wasm")

	store := wasmer.NewStore(wasmer.NewEngine())
	module, _ := wasmer.NewModule(store, wasmBytes)

	wasiEnv, _ := wasmer.NewWasiStateBuilder("wasi-program").
		// Choose according to your actual situation
		// Argument("--foo").
		// Environment("ABC", "DEF").
		// MapDirectory("./", ".").
		Finalize()
	importObject, err := wasiEnv.GenerateImportObject(store, module)
	check(err)

	instance, err := wasmer.NewInstance(module, importObject)
	check(err)

	start, err := instance.Exports.GetWasiStartFunction()
	check(err)
	start()

	HelloWorld, err := instance.Exports.GetFunction("Hello")
	check(err)
	result, _ := HelloWorld()
	fmt.Println(result)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
