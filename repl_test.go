package main

import "testing"

func TestParseInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Test input",
			expected: []string{"test", "input"},
		},
	}

	for _, cs := range cases {
		_input := parseInput(cs.input)

		if len(_input) != len(cs.expected) {
			t.Errorf("Expected %v, got %v", len(cs.expected), len(_input))
			continue
		}

		for i := range _input {
			_act := _input[i]
			_exp := cs.expected[i]

			if _act != _exp {
				t.Errorf("Expected %v, got %v", _exp, _act)
			}
		}
	}
}
