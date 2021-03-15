package utils_go

import (
	"crypto/rand"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"regexp"
	"strconv"
	"time"
)

const (
	Day                = time.Hour * 24
	DaySecond          = Day / time.Second
	Month              = Day * 30
	MonthSecond        = DaySecond * 30
	Year               = Month * 365
	YearSecond         = MonthSecond * 365
	RegexpPatternEmail = "[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?"
	YYYYMMDDHHMMSS     = "2006-01-02 15:04:05"
	YYYYMMDDHHMM       = "2006-01-02 15:04"
	YYYYMMDDHH         = "2006-01-02 15"
	YYYYMMDD           = "2006-01-02"
	HHMMSS             = "15:04:05"
	HHMM               = "15:04"
	MMSS               = "04:05"
	YYYY               = "2006"
	MM1                = "01"
	DD                 = "02"
	HH                 = "15"
	MM2                = "04"
	SS                 = "05"
)

func RandInt(max int64, min int64) int64 {
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	d := r.Int64()
	if d < min {
		return d + min
	} else {
		return d
	}
}

func RandString(n int, num bool, lowcase bool, upcase bool) string {
	var code string
	var letterBytes = ""
	if num {
		letterBytes += "0123456789"
	}
	if lowcase {
		letterBytes += "abcdefghijklmnopqrstuvwxyz"
	}
	if upcase {
		letterBytes += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	length := len(letterBytes)
	if length == 0 {
		return ""
	}
	for i := 0; i < n; i++ {
		c := letterBytes[RandInt(int64(length), 0)]
		code += string(c)
	}
	return code
}

func RandMix(n int) string {
	return RandString(n, true, true, true)
}

func RandLowcase(n int) string {
	return RandString(n, false, true, false)
}

func RandUpcase(n int) string {
	return RandString(n, false, false, true)
}

func RandIntString(n int) string {
	return RandString(n, true, false, false)
}

func StringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func StringToFloatDecimal(s string, exp int32) float64 {
	f1, _ := decimal.RequireFromString(s).Float64()
	f2, _ := decimal.NewFromFloatWithExponent(f1, exp).Float64()
	return f2
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func FloatToInt(f float64) int64 {
	return int64(f)
}

func IntToFloat(i int64) float64 {
	return float64(i)
}

func IntToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func InsertByte(source []byte, i uint8, idx int64) []byte {
	result := make([]byte, 0)
	result = append(result, source[:idx]...)
	result = append(result, i)
	result = append(result, source[idx:]...)
	return result
}

func Regexp(pattern string, s string) bool {
	r, _ := regexp.MatchString(pattern, s)
	return r
}

func UnixString() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func IntDateFmt(date int64, inFmt string, outFmt string) string {
	var t time.Time
	var err error
	if t, err = time.Parse(inFmt, IntToString(date)); err != nil {
		return ""
	}
	return t.Format(outFmt)
}

func StringDateFmt(date, inFmt, outFmt string) string {
	var t time.Time
	var err error
	if t, err = time.Parse(inFmt, date); err != nil {
		return ""
	}
	return t.Format(outFmt)
}

func IsDate(fmt, str string) bool {
	_, err := time.Parse(fmt, str)
	if err == nil {
		return true
	} else {
		return false
	}
}
