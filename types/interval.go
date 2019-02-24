package types

import "strconv"

type Interval struct {
	Lower uint8
	Upper uint8
}

func newInterval(lw uint8, upp uint8) *Interval{
	result := Interval{}
	if(lw <= upp) {
		result.Lower = lw
		result.Upper = upp
	} else
	{
		result.Lower = upp
		result.Upper = lw
	}
	return &result
}

func ToString(intv Interval) string{
	var result string = "["
	result += strconv.Itoa(int(intv.Lower))
	result += ", "
	result += strconv.Itoa(int(intv.Upper))
	result += "]"
	return result
}