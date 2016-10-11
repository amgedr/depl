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

func TestIsEmailDisposable(t *testing.T) {
	checks := map[string]bool{
		"qwerty@gmail.com":        false,
		"qwerty@123-m.com":        true,
		"qwerty@yahoo.com":        false,
		"qwerty@20minutemail.com": true,
		"@21cn.com":               false,
		"zomg.info":               false,
		"":                        false,
	}

	for email, val := range checks {
		c, err := IsEmailDisposable(email)
		if email == "" && err != ErrEmptyString {
			t.Error("Empty string should return an error")
			continue
		}

		if (email == "@21cn.com" || email == "zomg.info") && err != ErrInvalidEmail {
			t.Errorf("%s should return: %s", email, ErrInvalidEmail.Error())
		}

		if c != val {
			t.Errorf("%s should not be %v", email, val)
		}
	}
}
