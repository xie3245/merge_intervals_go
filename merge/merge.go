package merge

import(
	"github.com/xie3245/merge_intervals_go/types"
)

var whole_range_mask uint64

func MergeIntervals(data_ch chan types.Interval, result_ch chan types.Interval) {
	for interval := range data_ch {
		set_interval(interval)
	}

	traverse_result(result_ch)
	close(result_ch)
}

func set_interval(interval types.Interval) {
	if interval.Lower == interval.Upper {
		set_bit(interval.Lower)
	} else {
		set_range(interval.Lower, interval.Upper)
	}
}

func set_bit(pos uint8) {
	whole_range_mask |= (1 << pos)
}

func is_set(pos uint8) bool {
	return (whole_range_mask & (1 << pos)) > 0
}

func set_range(lower uint8, upper uint8) {
	const all_set uint64= 0xFFFFFFFFFFFFFFFF
	mask_start := all_set << lower
	mask_end := all_set >> (63 - upper)
	whole_range_mask |= (mask_start & mask_end)
}

var in_a_interval bool = false
var current_lower uint8
var current_upper uint8

func traverse_result(result_ch chan types.Interval) {
	var current_bit uint8 = 0
	for current_bit < 64 {
		handle_current_bit(current_bit, result_ch)
		current_bit++
	}

	if in_a_interval {
		result_ch <- types.Interval{current_lower, 63}
	}
}

func handle_current_bit(pos uint8, ch chan types.Interval) {
	var status bool = is_set(pos)
	if status {
		handle_bit_set(pos)
	} else {
		handle_bit_not_set(pos, ch)
	}
}

func handle_bit_set(pos uint8) {
	if !in_a_interval {
		current_lower = pos
		in_a_interval = true;
	}
}

func handle_bit_not_set(pos uint8, ch chan types.Interval) {
	if in_a_interval {
		current_upper = pos - 1
		in_a_interval = false;
		ch <- types.Interval{current_lower, current_upper}		
	}
}