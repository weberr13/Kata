package gogen

import (
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"testing/quick"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMeaninglessThings(t *testing.T) {
	Convey("demonstrate quick.Check", t, func () {
		alwaysTrue := func (_ int) bool {
			return true
		}
		err := quick.Check(alwaysTrue, nil)
		So(err, ShouldBeNil)
	})
	Convey("demonstrate quick.CheckEqual", t, func() {
		alwaysInput := func (x int) int {
			return x
		}
		alsoAlwaysInput := func (x int) int {
			return x
		}
		err := quick.CheckEqual(alwaysInput, alsoAlwaysInput, nil)
		So(err, ShouldBeNil)
	})

	Convey("demonstrate quick.CheckError", t, func() {
		alwaysTrue := func (_ int) bool {
			return false
		}
		err := quick.Check(alwaysTrue, nil)
		So(err, ShouldNotBeNil)
		checkErr, ok := err.(*quick.CheckError)
		So(ok, ShouldBeTrue)
		So(checkErr.Count, ShouldBeGreaterThan, 0)
		t.Log(checkErr.In)
	})
	Convey("demonstrate quick.CheckEqualError", t, func() {
		alwaysInput := func (x int) int {
			return x
		}
		neverInput := func (x int) int {
			return x + 1
		}
		err := quick.CheckEqual(alwaysInput, neverInput, nil)
		So(err, ShouldNotBeNil)
		checkErr, ok := err.(*quick.CheckEqualError)
		So(ok, ShouldBeTrue)
		So(checkErr.Count, ShouldBeGreaterThan, 0)
		So(checkErr.In[0].(int), ShouldEqual, checkErr.Out1[0].(int))
		So(checkErr.Out2[0].(int), ShouldEqual, checkErr.Out1[0].(int)+1)
	})
}

func TestFirstOrderThings(t *testing.T) {
	offsetTimestamp := "2013-07-08T18:07:13.23-07:00"
	zuluTimestamp := "2013-07-08T18:07:13.23Z"
	rfc3339Regex := regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?[+-]\\d{2}:\\d{2}")
	parseRFC3339Time := func (s string) bool {
		_, err := time.Parse(time.RFC3339, s)
		if (nil != err) {
			return false
		}
		return true
	}
	badParseRFC3339Time := func (s string) bool {
		return rfc3339Regex.Match(([]byte)(s))
	}

	Convey("function that parses things unsuccessfuly with brute force", t, func() {

		So(badParseRFC3339Time(offsetTimestamp), ShouldBeTrue)
		So(parseRFC3339Time(offsetTimestamp), ShouldBeTrue)
		So(badParseRFC3339Time(zuluTimestamp), ShouldBeFalse)
		So(parseRFC3339Time(zuluTimestamp), ShouldBeTrue)

		// The following didn't work, obviously
		// err := quick.CheckEqual(parseRFC3339Time, badParseRFC3339Time, nil)
		// So(err, ShouldNotBeNil)
		// checkErr, ok := err.(*quick.CheckError)
		// So(ok, ShouldBeTrue)
		// So(checkErr.Count, ShouldBeGreaterThan, 0)
		// t.Log(checkErr.In)
	})

}
	
func RandString(r *rand.Rand, maxLen int32) string {
	len := r.Int31n(maxLen)
	var s strings.Builder
	for i := 0 ; i < (int)(len) ; i ++ {
		s.WriteRune((rune)(0x00ff & r.Int31n(256)))
	}
	return s.String()
}
//TimeValues are written to every element of "args"
func TimeValues(args []reflect.Value, r *rand.Rand) {
	for i := 0 ; i < len(args) ; i++ {
		ts := r.Int63n(253402300799) // end of 9999
		t := time.Unix(ts, 0)
		garbage := RandString(r, 10)
		v := reflect.New(reflect.TypeOf("string")).Elem()
		v.SetString(t.Format("2006-01-02T15:04:05" + garbage))
		args[i] = v
	}
}

func TestFirstOrderThingsGenerated(t *testing.T) {
	rfc3339Regex := regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?[+-]\\d{2}:\\d{2}")

	parseRFC3339Time := func (s string) bool {
		_, err := time.Parse(time.RFC3339, s)
		if (nil != err) {
			return false
		}
		return true
	}
	badParseRFC3339Time := func (s string) bool {
		return rfc3339Regex.Match(([]byte)(s))
	}
	Convey("function that parses things successfully by brute force", t, func() {
		err := quick.CheckEqual(parseRFC3339Time, badParseRFC3339Time, 
			&quick.Config{Values: TimeValues, MaxCount: 1000000})
		So(err, ShouldNotBeNil)
		checkErr, ok := err.(*quick.CheckEqualError)
		So(ok, ShouldBeTrue)
		So(checkErr.Count, ShouldBeGreaterThan, 0)
		t.Log(checkErr.In)
	})
}
