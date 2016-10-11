package depl

import (
	"testing"
)

func TestIsDisposableDomain(t *testing.T) {
	checks := map[string]bool{
		"gmail.com":        false,
		"123-m.com":        true,
		"yahoo.com":        false,
		"20minutemail.com": true,
		"21cn.com":         true,
		"emz.net":          true,
		"zomg.info":        true,
		"":                 false,
	}

	for dom, val := range checks {
		c, err := IsDomainDisposable(dom)
		if dom == "" && err != ErrEmptyString {
			t.Error("Empty string should return an error")
		}

		if c != val {
			t.Errorf("%s should not be %v", dom, val)
		}
	}
}
