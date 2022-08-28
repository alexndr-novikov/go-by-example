package main

import (
	"go-by-example/internal/array"
	atomic_counters "go-by-example/internal/atomic-counters"
	channelBuffering "go-by-example/internal/channel-buffering"
	channelDirections "go-by-example/internal/channel-directions"
	channelRange "go-by-example/internal/channel-range"
	channelSync "go-by-example/internal/channel-sync"
	"go-by-example/internal/channels"
	closingChannels "go-by-example/internal/closing-channels"
	"go-by-example/internal/closures"
	"go-by-example/internal/constant"
	"go-by-example/internal/errors"
	_for "go-by-example/internal/for"
	_func "go-by-example/internal/func"
	multiReturn "go-by-example/internal/func-multiple-return"
	"go-by-example/internal/generics"
	"go-by-example/internal/goroutines"
	"go-by-example/internal/hello"
	"go-by-example/internal/interfaces"
	"go-by-example/internal/maps"
	"go-by-example/internal/methods"
	"go-by-example/internal/mutex"
	nonBlockingChannels "go-by-example/internal/non-blocking-channels"
	"go-by-example/internal/pointers"
	_range "go-by-example/internal/range"
	"go-by-example/internal/ratelimiting"
	"go-by-example/internal/recursion"
	"go-by-example/internal/runes"
	selectExamples "go-by-example/internal/select"
	"go-by-example/internal/slices"
	stateful_goroutines "go-by-example/internal/stateful-goroutines"
	structEmbedding "go-by-example/internal/struct-embedding"
	"go-by-example/internal/structs"
	_switch "go-by-example/internal/switch"
	"go-by-example/internal/tickers"
	"go-by-example/internal/timeouts"
	"go-by-example/internal/timers"
	"go-by-example/internal/values"
	"go-by-example/internal/variables"
	"go-by-example/internal/variadic"
	"go-by-example/internal/waitgroups"
	workerPool "go-by-example/internal/worker-pool"
)

func main() {
	stateful_goroutines.RunSample()
	mutex.RunSample()
	atomic_counters.RunSample()
	ratelimiting.RunSample()
	waitgroups.RunSample()
	workerPool.RunSample()
	tickers.RunSample()
	timers.RunSample()
	channelRange.RunSample()
	hello.RunSample()
	values.RunSample()
	variables.RunSample()
	constant.RunSample()
	_for.RunSample()
	_switch.RunSample()
	array.RunSample()
	slices.RunSample()
	maps.RunSample()
	_range.RunSample()
	_func.RunSample()
	multiReturn.RunSample()
	variadic.RunSample()
	closures.RunSample()
	recursion.RunSample()
	pointers.RunSample()
	runes.RunSample()
	structs.RunSample()
	methods.RunSample()
	interfaces.RunSample()
	structEmbedding.RunSample()
	generics.RunSample()
	errors.RunSample()
	goroutines.RunSample()
	channels.RunSample()
	channelBuffering.RunSample()
	channelSync.RunSample()
	channelDirections.RunSample()
	selectExamples.RunSample()
	timeouts.RunSample()
	nonBlockingChannels.RunSample()
	closingChannels.RunSample()
}
