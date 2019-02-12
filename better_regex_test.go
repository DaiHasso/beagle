package beagle

import (
    "testing"
    "fmt"

	gm "github.com/onsi/gomega"
)

func TestSimpleRegex(t *testing.T) {
	g := gm.NewGomegaWithT(t)

    regex := MustRegex(`I heard .*`)

    testString := "I heard that go regexps are hard! But maybe they are " +
        "easy..."
    result := regex.Match(testString)

    g.Expect(result.NamedGroup("thing")).To(gm.BeNil())
    g.Expect(result.NamedGroup("difficulty")).To(gm.BeNil())
    g.Expect(result.UnamedGroups()).To(gm.BeNil())
    g.Expect(result.MatchedSubstrings()).To(gm.ConsistOf(testString))
}

func TestRegexWithNamedGroups(t *testing.T) {
	g := gm.NewGomegaWithT(t)

    regex := MustRegex(`(?P<thing>\S+) (are|is) (?P<difficulty>\S+)`)

    result := regex.Match(
        "I heard that go regexps are hard! But maybe it is easy...",
    )

    g.Expect(result.NamedGroup("thing")[0]).To(gm.Equal("regexps"))
    g.Expect(result.NamedGroup("difficulty")[0]).To(gm.Equal("hard!"))
    g.Expect(result.NamedGroup("thing")[1]).To(gm.Equal("it"))
    g.Expect(result.NamedGroup("difficulty")[1]).To(gm.Equal("easy..."))
    g.Expect(result.UnamedGroups()).To(gm.Equal([]string{"are", "is"}))
    g.Expect(result.MatchedSubstrings()[0]).To(gm.Equal("regexps are hard!"))
    g.Expect(result.MatchedSubstrings()[1]).To(gm.Equal("it is easy..."))
}

func TestReadmeExample(t *testing.T) {
	g := gm.NewGomegaWithT(t)

    regex := MustRegex(`(?P<thing>\S+) (are|is) (?P<difficulty>\S+)`)

    result := regex.Match(
        "I heard that go regexps are hard! But maybe it is easy.",
    )

    expected := "Actually regexps are easy."

    g.Expect(fmt.Sprintf(
        "Actually %s %s %s",
        result.NamedGroup("thing")[0],
        result.UnamedGroups()[0],
        result.NamedGroup("difficulty")[1],
    )).To(gm.Equal(expected))
}
