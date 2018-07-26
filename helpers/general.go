package helpers

import (
	//"fmt"
	"strings"
    "unicode"
    "time"
    "bytes"
    "encoding/gob"
    "reflect"
    "strconv"
)

func CamelCaseToSnakeCase(camelCase string) (inputUnderScoreStr string) {
    //snake_case to camelCase
    for k, v := range camelCase {
        if(isUpperCase(string(v)) && k > 0 && unicode.IsLetter(v)){
            inputUnderScoreStr += "_"
        }
        inputUnderScoreStr += strings.ToLower(string(v))
    }
    return inputUnderScoreStr
}

func isUpperCase(str string) bool {
    if str == strings.ToUpper(str) {
        return true
    }
    return false
}

func CalProcessTime(start time.Time) string{
    var diffTime = time.Since(start)
    var second = diffTime.Seconds()
    return strconv.FormatFloat(second, 'f', 6, 64)
}

func CalNight(startDate string,endDate string) int{
    timeFormat := "2006-01-02"
    start, _ := time.Parse(timeFormat,startDate)
    end, _ := time.Parse(timeFormat,endDate)
    diff := end.Sub(start)

    diffDays := int(diff.Hours()/24)

    return diffDays
}

func GetBytes(key interface{}) ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

func InArray(val interface{}, array interface{}) (exists bool) {
    exists = false

    switch reflect.TypeOf(array).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                exists = true
                return
            }
        }
    }

    return
}

func Round(val float64) int {
    if val < 0 { return int(val-0.5) }
    return int(val+0.5)
}

func IsNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
 }

 func Explode(s string,del string) ([]string){
    //check delimiter
    dels := []string{",", "|","&","#","."}

    if r := InArray(del,dels); !r {
        return []string{}
    }

    c := Empty(s)

     if c != true  {
        stringSlice := strings.Split(s,del) 

        return stringSlice
    }

     return []string{}
 }

 func Empty(s string) bool {
    if len(s) > 0 || s != ""  {
        return false
    }

    return true
 }

 func GetDay(date string) int{
    timeFormat := "2006-01-02"
    dateParse, _ := time.Parse(timeFormat,date)
    day := dateParse.Day()

    return day
}

func GetDayOffWeek(date string) int{
    timeFormat := "2006-01-02"
    dateParse, _ := time.Parse(timeFormat,date)
    weekday := dateParse.Weekday()

    day := int(weekday)

    return day
}