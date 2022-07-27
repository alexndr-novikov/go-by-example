package main

import (
	"go-by-example/internal/array"
	"go-by-example/internal/constant"
	"go-by-example/internal/for"
	"go-by-example/internal/hello"
	"go-by-example/internal/maps"
	"go-by-example/internal/slices"
	"go-by-example/internal/switch"
	"go-by-example/internal/values"
	"go-by-example/internal/variables"
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
}
