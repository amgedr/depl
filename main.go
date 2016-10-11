package depl

import (
	"sort"
	"strings"
)

func init() {
	sort.Strings(Domains)
}

//IsDomainDisposable searches the list of the disposable email providers
//domains for dom and returns true if found.
func IsDomainDisposable(dom string) (bool, error) {
	if dom == "" {
		return false, ErrEmptyString
	}

	i := sort.Search(len(Domains),
		func(i int) bool {
			return Domains[i] >= dom
		})

	if i < len(Domains) && Domains[i] == dom {
		return true, nil
	} else {
		return false, nil
	}
}

//IsEmailDisposable searches the list of the disposable email providers
//domains for the domain part of email and returns true if found.
func IsEmailDisposable(email string) (bool, error) {
	if email == "" {
		return false, ErrEmptyString
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false, ErrInvalidEmail
	}

	if len(parts[0]) == 0 || len(parts[1]) == 0 {
		return false, ErrInvalidEmail
	}

	return IsDomainDisposable(parts[1])
}
