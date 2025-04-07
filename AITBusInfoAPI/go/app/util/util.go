package util

import (
	"log"
	"time"
)

type Util struct{}

func (Util) SliceUniqueString(target []string) (unique []string) {
	m := map[string]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

func (Util) SliceUniqueNumber(target []int64) (unique []int64) {
	m := map[int64]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

// 配列の中に特定の文字列が含まれているか
func (Util) ArrayStringContains(target []string, str string) bool {
	for _, v := range target {
		if v == str {
			return true
		}
	}
	return false
}

// 引数datetime文字列とタイムゾーン文字列を受け取りTime型に変換する関数
func (Util) ConvertDatetimeToLocationTime(datetime string, timezone string) (time.Time, error) {
	timeZone, _ := time.LoadLocation(timezone)
	locationTime, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, timeZone)
	if err != nil {
		log.Fatal(err.Error())
		return time.Time{}, err
	}
	return locationTime, nil
}

func (Util) TimeToUnixMilli(t time.Time) int64 {
	return t.UnixNano() / 1000000
}
