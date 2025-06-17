package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name     string
		input    http.Header
		want     string
		errorMsg string
	}

	tests := []test{
		{name: "No auth", input: http.Header{}, want: "", errorMsg: auth.ErrNoAuthHeaderIncluded.Error()},
		{name: "Mallformed no keys", input: http.Header{"Authorization": []string{"ApiKey"}}, want: "", errorMsg: "malformed authorization header"},
		{name: "Mallformed No ApiKey", input: http.Header{"Authorization": []string{"NoApiKey"}}, want: "", errorMsg: "malformed authorization header"},
		{name: "Good test", input: http.Header{"Authorization": []string{"ApiKey MyKey"}}, want: "MyKey", errorMsg: ""},
		//{name: "BAD TEST", input: http.Header{}, want: "WRONG", errorMsg: ""},
	}

	var output string
	var err error
	for _, test := range tests {
		output, err = auth.GetAPIKey(test.input)
		if output != test.want {
			t.Errorf("%v: Expected (%v) got (%v)", test.name, test.want, output)
		}

		expectError := len(test.errorMsg) > 0
		if expectError {
			if err == nil {
				t.Errorf("%v: Expeted error (%v) but no error was returned", test.name, test.errorMsg)
			}

			errMsg := err.Error()
			if test.errorMsg != errMsg {
				t.Errorf("%v: Expected error (%v) didn't match received error (%v)", test.name, test.errorMsg, errMsg)
			}
		} else {
			if err != nil {
				t.Errorf("%v: Got error (%v) when no error was expected", test.name, err.Error())
			}
		}

	}
}

func unusedFunction() {
	//only exists to test ci.yml staticcheck
}
