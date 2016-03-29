package megamatchers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kr/pretty"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

// DiffEqual is an extension to the standard Gomega Equal matcher that also provides a diff of actual and expected.
func DiffEqual(expected interface{}) types.GomegaMatcher {
	return &DiffEqualMatcher{
		Expected: expected,
	}
}

// DiffEqualMatcher is a custom Gomega matcher that checks equality and provides a diff of any changes.
type DiffEqualMatcher struct {
	Expected interface{}
}

// Match satisfies gomega.GomegaMatcher by checking equality of actual and matcher.Expected with reflect.DeepEqual.
func (m *DiffEqualMatcher) Match(actual interface{}) (bool, error) {
	if actual == nil && m.Expected == nil {
		return false, fmt.Errorf("Refusing to compare <nil> to <nil>.\nBe explicit and use BeNil() instead.  This is to avoid mistakes where both sides of an assertion are erroneously uninitialized.")
	}
	return reflect.DeepEqual(actual, m.Expected), nil
}

// FailureMessage satisfies gomega.GomegaMatcher by returning a failure message that includes actual, expected and the
// diff of actual and expected.
func (m *DiffEqualMatcher) FailureMessage(actual interface{}) string {
	return m.message(actual, "to equal")
}

// NegatedFailureMessage satisfies gomega.GomegaMatcher by returning a failure message that includes actual, expected
// and the diff of actual and expected.
func (m *DiffEqualMatcher) NegatedFailureMessage(actual interface{}) string {
	return m.message(actual, "not to equal")
}

func (m *DiffEqualMatcher) message(actual interface{}, equalityStr string) string {
	diffs := pretty.Diff(actual, m.Expected)
	return fmt.Sprintf(
		"Expected\n%s\n%s\n%s\n%s\n%s",
		format.Object(actual, 1),
		equalityStr,
		format.Object(m.Expected, 1),
		"with diff",
		format.Object(strings.Join(diffs, "\n"), 1),
	)
}
