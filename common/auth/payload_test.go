package auth

import (
	"testing"
)

func TestCan(t *testing.T) {
	payload := &Payload{
		Authorities: []string{"userCreate", "userUpdate", "userDeleteRestore"},
	}

	tests := []struct {
		authoritiesToCheck []string
		expectedResults    []bool
	}{
		{[]string{"userCreate", "orderDeleteRestore"}, []bool{true, false}},
		{[]string{"userUpdate", "userDeleteRestore"}, []bool{true, true}},
		{[]string{"nonExistentAuthority"}, []bool{false}},
		{[]string{}, []bool{}}, // Empty authorities to check
		{[]string{"userCreate", "userUpdate", "userDeleteRestore"}, []bool{true, true, true}}, // All authorities exist
		{[]string{"userCreate", "userUpdate", "userDeleteRestore", "nonExistentAuthority"}, []bool{true, true, true, false}},
	}

	for _, test := range tests {
		results := payload.Can(test.authoritiesToCheck)

		if len(results) != len(test.expectedResults) {
			t.Errorf("For authorities %v, expected %v, but got %v", test.authoritiesToCheck, test.expectedResults, results)
			continue
		}

		for i, expected := range test.expectedResults {
			if results[i] != expected {
				t.Errorf("For authorities %v, expected %v at index %d, but got %v", test.authoritiesToCheck, expected, i, results[i])
			}
		}
	}
}
