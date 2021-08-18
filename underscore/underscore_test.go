package underscore

import (
	"fmt"
	"testing"
)

// Using table driven tests
// func TestCamel(t *testing.T) {
// 	testCases := []struct {
// 		arg  string
// 		want string
// 	}{
// 		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
// 		{"with a space", "with a space"},
// 		{"endsWithA", "ends_with_a"},
// 	}
// 	for _, tc := range testCases {
// 		t.Logf("Testing: %q", tc.arg)
// 		got := Camel(tc.arg)
// 		if got != tc.want {
// 			t.Errorf("Camel(%q) = %q; want %q", tc.arg, got, tc.want)
// 		}
// 	}
// }

// func TestCamel(t *testing.T) {
// 	type args struct {
// 		str string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{"simple", args{"thisIsACamelCaseString"}, "this_is_a_camel_case_string"},
// 		{"spaces", args{"with a space"}, "with a space"},
// 		{"end with capital", args{"endsWithA"}, "ends_with_a"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Camel(tt.args.str); got != tt.want {
// 				t.Errorf("Camel() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// Use args as test's name
func TestCamel(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
		{"with a space", "with a space"},
		{"endsWithA", "ends_with_a"},
	}
	// so three tests will all run at the same time individually
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := Camel(tt.arg); got != tt.want {
				t.Fatalf("Camel() = %v, want %v", got, tt.want)
			}
			fmt.Println("this won't print if it fails...")
			// check2
			// check3
		})
	}
}
