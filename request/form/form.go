package form

import (
	"fmt"
	"net/url"
)

type Option func(url.Values) (err error)

func Value(field string, value interface{}) Option {
	return Values(field, value)
}

func Values(field string, values ...interface{}) Option {
	return func(v url.Values) (err error) {
		m := make([]string, 0)
		for _, value := range values {
			m = append(m, fmt.Sprint(value))
		}
		v[field] = m

		return
	}
}
