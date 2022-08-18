package main

import (
	"go-by-example/internal/array"
	channel_buffering "go-by-example/internal/channel-buffering"
	channel_directions "go-by-example/internal/channel-directions"
	channel_sync "go-by-example/internal/channel-sync"
	"go-by-example/internal/channels"
	"go-by-example/internal/closures"
	"go-by-example/internal/constant"
	"go-by-example/internal/errors"
	"go-by-example/internal/for"
	_func "go-by-example/internal/func"
	multReturn "go-by-example/internal/func-multiple-return"
	"go-by-example/internal/generics"
	"go-by-example/internal/goroutines"
	"go-by-example/internal/hello"
	"go-by-example/internal/interfaces"
	"go-by-example/internal/maps"
	"go-by-example/internal/methods"
	non_blocking_channels "go-by-example/internal/non-blocking-channels"
	"go-by-example/internal/pointers"
	"go-by-example/internal/range"
	"go-by-example/internal/recursion"
	"go-by-example/internal/runes"
	select_examples "go-by-example/internal/select"
	"go-by-example/internal/slices"
	struct_embedding "go-by-example/internal/struct-embedding"
	"go-by-example/internal/structs"
	"go-by-example/internal/switch"
	"go-by-example/internal/timeouts"
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
	generics.Generics()
	errors.Errors()
	goroutines.Goroutines()
	channels.Channels()
	channel_buffering.ChannelBuffering()
	channel_sync.ChannelSync()
	channel_directions.ChannelDirections()
	select_examples.Select()
	timeouts.Timeouts()
	non_blocking_channels.NonBlockingChannels()
}
