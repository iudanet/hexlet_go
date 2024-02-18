package solution

import "fmt"

// BEGIN (write your solution here)

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		return fmt.Sprint("en."+ domain)
	} else {
		return fmt.Sprint(locale + "." + domain)
	}
}
