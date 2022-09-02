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
	defet "go-by-example/internal/defer"
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
	"go-by-example/internal/panicpack"
	"go-by-example/internal/pointers"
	_range "go-by-example/internal/range"
	"go-by-example/internal/ratelimiting"
	"go-by-example/internal/recovery"
	"go-by-example/internal/recursion"
	"go-by-example/internal/runes"
	selectExamples "go-by-example/internal/select"
	"go-by-example/internal/sha256"
	"go-by-example/internal/slices"
	sort_functions "go-by-example/internal/sort-functions"
	"go-by-example/internal/sorting"
	stateful_goroutines "go-by-example/internal/stateful-goroutines"
	string_functions "go-by-example/internal/string-functions"
	structEmbedding "go-by-example/internal/struct-embedding"
	"go-by-example/internal/structs"
	_switch "go-by-example/internal/switch"
	"go-by-example/internal/tickers"
	time_package "go-by-example/internal/time-package"
	"go-by-example/internal/timeouts"
	"go-by-example/internal/timers"
	"go-by-example/internal/values"
	"go-by-example/internal/variables"
	"go-by-example/internal/variadic"
	"go-by-example/internal/waitgroups"
	workerPool "go-by-example/internal/worker-pool"
)

func main() {
	sha256.RunSample()
	time_package.RunSample()
	string_functions.RunSample()
	recovery.RunSample()
	defet.RunSample()
	sort_functions.RunSample()
	sorting.RunSample()
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
	panicpack.RunSample()
}
