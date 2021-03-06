package matcher

import (
	"regexp"
)

// Matcher is a regular expression matcher
type Matcher interface {
	Match(value string) bool
	MatchGroups(value string) map[string]string
}

type matcher struct {
	matcher *regexp.Regexp
}

// New creates a new regular expression matcher or returns error
func New(expression string) (Matcher, error) {
	m, err := regexp.Compile(expression)
	if err != nil {
		return nil, err
	}
	return &matcher{m}, nil
}

// NewMust creates a new regular expression matcher or panics
func NewMust(expression string) Matcher {
	m := regexp.MustCompile(expression)
	return &matcher{m}
}

func (m *matcher) Match(value string) bool {
	return m.matcher.MatchString(value)
}

// MatchGroups for a given regular expression and a string,
// it matches and returns the group values defined in the expression.
func (m *matcher) MatchGroups(value string) map[string]string {
	groups := map[string]string{}

	match := m.matcher.FindStringSubmatch(value)
	if match == nil {
		return groups // no match
	}

	for i, name := range m.matcher.SubexpNames() {
		if i == 0 {
			continue // skip the zero group
		}
		groups[name] = match[i]
	}
	return groups
}
