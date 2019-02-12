package beagle

import (
    "regexp"

    "github.com/pkg/errors"
)

type Regex struct {
    goRegexp *regexp.Regexp
}

type Result struct {
    fullMatches,
    unamedCaptureGroups []string
    namedCaptureGroups map[string][]string
    hasMatches bool
}

func makeResult(names []string, allMatches [][]string) *Result {
    hasMatches := false
    var (
        fullMatches []string
        unamedCaptureGroups []string
        namedCaptureGroups = make(map[string][]string)
    )
    for _, matches := range allMatches {
        hasMatches = true
        fullMatches = append(fullMatches, matches[0])

        if len(matches) > 1 {
            remMatches := matches[1:]
            remNames := names[1:]
            for i, captureKey := range remNames {
                match := remMatches[i]
                if captureKey == "" {
                    unamedCaptureGroups = append(unamedCaptureGroups, match)
                } else {
                    if _, ok := namedCaptureGroups[captureKey]; !ok {
                        namedCaptureGroups[captureKey] = []string{}
                    }
                    namedCaptureGroups[captureKey] = append(
                        namedCaptureGroups[captureKey], match,
                    )
                }
            }
        }
    }

    return &Result{
        fullMatches: fullMatches,
        unamedCaptureGroups: unamedCaptureGroups,
        namedCaptureGroups: namedCaptureGroups,
        hasMatches: hasMatches,
    }
}

func (self Regex) Match(target string) *Result {
    matches := self.goRegexp.FindAllStringSubmatch(target, -1)
    subexpNames := self.goRegexp.SubexpNames()

    return makeResult(subexpNames, matches)
}

func (self Result) NamedGroup(key string) []string {
    if vals, ok := self.namedCaptureGroups[key]; ok {
        return vals
    }
    return nil
}

func (self Result) UnamedGroups() []string {
    return self.unamedCaptureGroups
}

func (self Result) MatchedSubstrings() []string {
    return self.fullMatches
}


func MustRegex(regex string) *Regex {
    newRegex := regexp.MustCompile(regex)
    return &Regex{newRegex}
}

func MakeRegex(regex string) (*Regex, error) {
    newRegex, err := regexp.Compile(regex)
    if err != nil {
        return nil, errors.Wrap(err, "Error while compiling regexp")
    }
    return &Regex{
        goRegexp: newRegex,
    }, err
}

func MakeBetter(regex *regexp.Regexp) *Regex {
    return &Regex{
        goRegexp: regex,
    }
}
