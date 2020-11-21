package form

type Option func(m map[string]interface{}) (err error)

func Field(field string, value interface{}) Option {
	return func(m map[string]interface{}) (err error) {
		m[field] = value
		return
	}
}
