package main

import (
	"fmt"
	"strings"
)

// AssignmentsMap is a `flag.Value` for `KEY=VALUE` arguments.
// The value of the `Separator` field is used instead  of `"="` when set.
type AssignmentsMap struct {
	Separator string

	Values map[string]string
	Texts  []string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *AssignmentsMap) Help() string {
	separator := "="
	if fv.Separator != "" {
		separator = fv.Separator
	}
	return fmt.Sprintf("a key/value pair KEY%sVALUE", separator)
}

// Set is flag.Value.Set
func (fv *AssignmentsMap) Set(v string) error {
	separator := "="
	if fv.Separator != "" {
		separator = fv.Separator
	}
	i := strings.Index(v, separator)
	if i < 0 {
		return fmt.Errorf(`"%s" must have the form KEY%sVALUE`, v, separator)
	}
	fv.Texts = append(fv.Texts, v)
	if fv.Values == nil {
		fv.Values = make(map[string]string)
	}
	fv.Values[v[:i]] = v[i+len(separator):]
	return nil
}

func (fv *AssignmentsMap) String() string {
	return strings.Join(fv.Texts, ", ")
}
