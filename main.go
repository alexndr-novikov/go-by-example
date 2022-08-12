package main

import (
	"go-by-example/internal/array"
	"go-by-example/internal/closures"
	"go-by-example/internal/constant"
	"go-by-example/internal/for"
	_func "go-by-example/internal/func"
	multReturn "go-by-example/internal/func-multiple-return"
	"go-by-example/internal/hello"
	"go-by-example/internal/interfaces"
	"go-by-example/internal/maps"
	"go-by-example/internal/methods"
	"go-by-example/internal/pointers"
	"go-by-example/internal/range"
	"go-by-example/internal/recursion"
	"go-by-example/internal/runes"
	"go-by-example/internal/slices"
	struct_embedding "go-by-example/internal/struct-embedding"
	"go-by-example/internal/structs"
	"go-by-example/internal/switch"
	"go-by-example/internal/values"
	"go-by-example/internal/variables"
	"go-by-example/internal/variadic"
)

func main() {
	hello.Hello()
	values.Values()
	variables.Variables()
	constant.Const()
	_for.For()
	_switch.Switch()
	array.Array()
	slices.Slices()
	maps.Maps()
	_range.Range()
	_func.Func()
	multReturn.FuncMultipleReturnValues()
	variadic.Variadic()
	closures.Closures()
	recursion.Recursion()
	pointers.Pointers()
	runes.Runes()
	structs.Structs()
	methods.Methods()
	interfaces.Interfaces()
	struct_embedding.StructEmbedding()
}
