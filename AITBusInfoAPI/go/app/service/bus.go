package service

import (
	"fmt"
	"strconv"
)

type TimeTable struct {
	Schedule map[string]map[string][]int
}

type BusTime struct {
	Hour   int
	Minute int
}

func (t TimeTable) GetNextTime(s string, hour, minute, offset int) BusTime {
	if offset == 0 {
		// 指定されたhourのminuteに近い値を探して返す
		for _, v := range t.Schedule[s][fmt.Sprintf("%d", hour)] {
			if v >= minute {
				return BusTime{hour, v}
			}
		}

		// もし指定されたhourのminuteに近い値がなかった場合はhour+1のminuteに近い値を探して返す
		for _, v := range t.Schedule[s][strconv.Itoa(hour+1)] {
			if v >= 0 {
				return BusTime{hour + 1, v}
			}
		}
	} else {
		result := t.GetNextTime(s, hour, minute, 0)
		for i := 0; i < offset; i++ {
			result = t.GetNextTime(s, result.Hour, result.Minute+1, 0)
		}
		return result
	}

	return BusTime{-1, -1}
}

func (t TimeTable) GetTimeTable(mode string) map[string][]int {
	return t.Schedule[mode]
}
