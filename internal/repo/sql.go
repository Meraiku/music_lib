package repo

import (
	"fmt"
	"strings"
)

type SQLUpdate struct {
	assignments []string
	values      []any
}

func (s *SQLUpdate) add(key string, value any) {
	if s.assignments == nil {
		s.assignments = make([]string, 0, 1)
	}
	if s.values == nil {
		s.values = make([]any, 0, 1)
	}

	s.assignments = append(s.assignments, fmt.Sprintf("%s = ?", key))
	s.values = append(s.values, value)
}

func (s *SQLUpdate) Assignments() string {
	return strings.Join(s.assignments, ", ")
}

func (s *SQLUpdate) Values() []any {
	return s.values
}
