package hltb

import "testing"

func TestSplitStrTerms(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		sep      string
		expected []string
	}{
		{name: "OneTerm", value: "Halo", sep: ",", expected: []string{"Halo"}},
		{name: "WhiteSpaceSep", value: "  God Of   War ", sep: " ", expected: []string{"God", "Of", "War"}},
		{name: "CommaSep", value: "Separated, Values", sep: ",", expected: []string{"Separated", "Values"}},
	}

	for _, test := range tests {
		split := splitStrTerms(test.value, test.sep)

		if len(split) != len(test.expected) {
			t.Errorf("%s: Unexpected term count. Expected: %d, Actual: %d", test.name, len(test.expected), len(split))
		}

		for i := 0; i < len(split); i++ {
			if split[i] != test.expected[i] {
				t.Errorf("%s: Unexpected term. Expected: \"%s\", Actual: \"%s\"", test.name, test.expected[i], split[i])
			}
		}
	}
}
