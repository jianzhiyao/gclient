package form

import (
	"fmt"
	"net/url"
)

//Option form option
type Option func(url.Values) (err error)

//Value set form key-value
func Value(field string, value interface{}) Option {
	return Values(field, value)
}

//Values set form key-values
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
