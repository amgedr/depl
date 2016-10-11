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
