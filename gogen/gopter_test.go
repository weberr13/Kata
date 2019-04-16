package gogen

import (
	"math/rand"
	"regexp"
	"testing"
	"time"
	"unicode"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
	"github.com/leanovate/gopter/gen"
)

func TestTimeParser(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	Convey("test rfc regexp date parser", t, func() {
		parameters := gopter.DefaultTestParameters()
		parameters.MinSuccessfulTests = 100
		parameters.MinSize = 1
		parameters.MaxSize = 10
		properties := gopter.NewProperties(parameters)

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
		properties.Property("radom data after date", a.ForAll(badParseRFC3339Time))
		properties.TestingRun(t, gopter.ConsoleReporter(false))
	})
}