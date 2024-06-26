package utils_go

import (
	"crypto/rand"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	Day                 = time.Hour * 24
	DaySecond           = Day / time.Second
	Month               = Day * 30
	MonthSecond         = DaySecond * 30
	Year                = Month * 365
	YearSecond          = MonthSecond * 365
	RegexpPatternEmail  = "[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?"
	RegexpPatternMobile = "^((13[0-9])|(14[5,7,9])|(15[^4])|(18[0-9])|(17[0,1,3,5,6,7,8])|(19)[0-9])\\d{8}$"
	TIMESTAMPZONE       = "2006-01-02 15:04:05.999999999 -0700"
	YYYYMMDDHHMMSS      = "2006-01-02 15:04:05"
	YYYYMMDDHHMM        = "2006-01-02 15:04"
	YYYYMMDDHH          = "2006-01-02 15"
	YYYYMMDD            = "2006-01-02"
	HHMMSS              = "15:04:05"
	HHMM                = "15:04"
	MMSS                = "04:05"
	YYYY                = "2006"
	MM1                 = "01"
	DD                  = "02"
	HH                  = "15"
	MM2                 = "04"
	SS                  = "05"
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

func StructToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		jsonName := strings.TrimSpace(strings.Split(tag, ",")[0])
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = StructToMap(field)
			} else {
				res[jsonName] = field
			}
		} else if v.Field(i).Type.Kind() == reflect.Struct {
			for kk, vv := range StructToMap(field) {
				res[kk] = vv
			}
		}
	}
	return res
}

func Max(first float64, rest ...float64) float64 {
	result := first
	for _, i := range rest {
		if i > result {
			result = i
		}
	}
	return result
}

func Min(first float64, rest ...float64) float64 {
	result := first
	for _, i := range rest {
		if i < result {
			result = i
		}
	}
	return result
}

func StringsContains(array []string, s string) bool {
	sort.Strings(array)
	idx := sort.SearchStrings(array, s)
	return idx < len(array) && strings.Compare(array[idx], s) == 0
}

func Yesterday() time.Time {
	return time.Now().Add(time.Hour * -24)
}

func LastWeek() (string, string) {
	n := time.Now()
	start := n.AddDate(0, 0, -7).AddDate(0, 0, int(-n.Weekday())+1)
	end := start.AddDate(0, 0, 6)
	return start.Format("2006-01-02"), end.Format("2006-01-02")
}

func LastMonth() (string, string) {
	n := time.Now()
	n.ISOWeek()
	end := n.AddDate(0, 0, -n.Day())
	start := time.Date(end.Year(), end.Month(), 1, 0, 0, 0, 0, time.Local)
	return start.Format("2006-01-02"), end.Format("2006-01-02")
}
