package gogen

import (
	"fmt"
	"math/rand"
	"regexp"
	"testing"
	"time"
	"unicode"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/leanovate/gopter/convey"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
	"github.com/leanovate/gopter/commands"
	"github.com/leanovate/gopter/gen"
)

func ShouldNotSucceedForAll(actual interface{}, expected ...interface{}) string {
	failstring := ShouldSucceedForAll(actual, expected...)
	if failstring == "" {
		return "Expected failure"
	}
	fmt.Println("got expected failure: \n", failstring)
	return ""
}
func TestTimeParser(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	Convey("test rfc regexp date parser", t, func() {
		parameters := gopter.DefaultTestParameters()
		parameters.MinSuccessfulTests = 100
		parameters.MinSize = 1
		parameters.MaxSize = 10

		rfc3339Regex := regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?[+-]\\d{2}:\\d{2}")
		parseRFC3339Time := func(s string) bool {
			_, err := time.Parse(time.RFC3339, s)
			if nil != err {
				return false
			}
			return true
		}
		badParseRFC3339Time := func(s string) bool {
			return rfc3339Regex.Match(([]byte)(s))
		}
		prependDatestr := func(s string) string {
			ts := rand.Int63n(253402300799) // end of 9999
			t := time.Unix(ts, 0)
			return t.Format("2006-01-02T15:04:05" + s)
		}

		So(parseRFC3339Time("2006-01-02T15:04:05Z"), ShouldBeTrue)
		a := arbitrary.DefaultArbitraries()
		chars := &unicode.RangeTable{R16: []unicode.Range16{{Lo: 0x002b, Hi: 0x005a, Stride: 1}}}
		a.RegisterGen(gen.UnicodeString(chars).Map(prependDatestr).SuchThat(parseRFC3339Time))
		So(badParseRFC3339Time, ShouldNotSucceedForAll, a, parameters)
	})
}
type WrapAPI2 struct{
	last int
}
func (w *WrapAPI2) Init() {
	if w.last == 0 {
		w.last = 1
	} 
}
func (w *WrapAPI2) Add() {
	if w.last == 1 {
		w.last = 2
	} else {
		w.last = 0
	}
}
func (w *WrapAPI2) Delete() {
	if w.last == 2 {
		w.last = 3
	} else {
		w.last = 0
	}
}
func (w *WrapAPI2) Reset() {
	if w.last == 3 {
		w.last = 4
	} else {
		w.last = 0
	}
}
func (w *WrapAPI2) IsBroken() bool {
	return w.last == 4
}
func TestFakeApi(t *testing.T) {
	Convey("generate a bad sequence of commands", t, func() {
		APIInit := &commands.ProtoCommand{
			Name: "INIT",
			RunFunc: func(systemUnderTest commands.SystemUnderTest) commands.Result {
				systemUnderTest.(*WrapAPI2).Init()
				return nil
			},
		}
		APIAdd := &commands.ProtoCommand{
			Name: "ADD",
			RunFunc: func(systemUnderTest commands.SystemUnderTest) commands.Result {
				systemUnderTest.(*WrapAPI2).Add()
				return nil
			},
		}
		APIDelete := &commands.ProtoCommand{
			Name: "DELETE",
			RunFunc: func(systemUnderTest commands.SystemUnderTest) commands.Result {
				systemUnderTest.(*WrapAPI2).Delete()
				return nil
			},
		}
		APIReset := &commands.ProtoCommand{
			Name: "RESET",
			RunFunc: func(systemUnderTest commands.SystemUnderTest) commands.Result {
				systemUnderTest.(*WrapAPI2).Reset()
				return nil
			},
		}
		APIIsBroken := &commands.ProtoCommand{
			Name: "IS_BROKEN",
			RunFunc: func(systemUnderTest commands.SystemUnderTest) commands.Result {
				return systemUnderTest.(*WrapAPI2).IsBroken()
			},
			PostConditionFunc: func(_ commands.State, result commands.Result) *gopter.PropResult {
				if result.(bool) {
					return &gopter.PropResult{Status: gopter.PropFalse}
				}
				return &gopter.PropResult{Status: gopter.PropTrue}
			},
		}
		APICommands := &commands.ProtoCommands{
			NewSystemUnderTestFunc: func(initialState commands.State) commands.SystemUnderTest {
				return &WrapAPI2{}
			},
			InitialStateGen: gen.Const(false),
			InitialPreConditionFunc: func(state commands.State) bool {
				return state.(bool) == false
			},
			GenCommandFunc: func(state commands.State) gopter.Gen {
				return gen.OneConstOf(APIInit, APIAdd, APIIsBroken, APIDelete, APIReset)
			},
		}
		properties := gopter.NewProperties(gopter.DefaultTestParameters())

		properties.Property("Wrapped API", commands.Prop(APICommands))
	
		properties.TestingRun(t, gopter.ConsoleReporter(false))
	})
}