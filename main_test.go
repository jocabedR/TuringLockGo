package main

import "testing"

func TestJumpTo(t *testing.T) {
	result := jumpTo(0, "+19", 49)

	if result != 19 {
		t.Errorf("jumpTo was incorrect, got: %d, want: %d", result, 19)
	}
}

func TestGetRegisterValue(t *testing.T) {
	testCases := []struct {
		Name     string
		Register string
		a        int
		b        int
		Expect   int
	}{
		{
			Name:     "getRegisterValue returns a value.",
			Register: "a",
			a:        1,
			b:        5,
			Expect:   1,
		},
		{
			Name:     "getRegisterValue returns b value.",
			Register: "b",
			a:        1,
			b:        5,
			Expect:   5,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			result := getRegisterValue(tc.Register, tc.a, tc.b)
			if result != tc.Expect {
				t.Errorf("getRegisterValue was incorrect, got: %d, want: %d", result, tc.Expect)
			}
		})
	}
}

func TestSetRegisterValue(t *testing.T) {
	testCases := []struct {
		Name     string
		Operator string
		Register int
		Expect   int
	}{
		{
			Name:     "setRegisterValue has doene a hlf operation.",
			Operator: "hlf",
			Register: 6,
			Expect:   3,
		},
		{
			Name:     "setRegisterValue has done a tpl operation.",
			Operator: "tpl",
			Register: 6,
			Expect:   18,
		},
		{
			Name:     "setRegisterValue has done a inc operation.",
			Operator: "inc",
			Register: 6,
			Expect:   7,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			result := setRegisterValue(tc.Operator, tc.Register)
			if result != tc.Expect {
				t.Errorf("setRegisterValue was incorrect, got: %d, want: %d", result, tc.Expect)
			}
		})
	}
}
