package main 

import (
	"runtime"
	"github.com/xie3245/merge_intervals_go/merge"
	"github.com/xie3245/merge_intervals_go/ui"
	"github.com/xie3245/merge_intervals_go/types"
)

func main() {
	runtime.GOMAXPROCS(8)
	data_channel := make(chan types.Interval)
	result_channel := make(chan types.Interval)

	go ui.GetIntervals(data_channel)
	go merge.MergeIntervals(data_channel, result_channel)

	ui.PrintResult(result_channel)
}