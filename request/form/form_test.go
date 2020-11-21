package form

import "testing"

func TestField(t *testing.T) {
	m := make(map[string]interface{})

	f1, v1 := `f1`, 1
	f2, v2 := `f2`, `2`
	f3, v3 := `f3`, 1e5

	options := []Option{
		Field(f1, v1),
		Field(f2, v2),
		Field(f3, v3),
	}

	for _, option := range options {
		_ = option(m)
	}

	if v, ok := m[f1]; ok {
		if v != v1 {
			t.Error(v, v1)
			return
		}
	} else {
		t.Error()
		return
	}

	if v, ok := m[f2]; ok {
		if v != v2 {
			t.Error(v, v2)
			return
		}
	} else {
		t.Error()
		return
	}

	if v, ok := m[f3]; ok {
		if v != v3 {
			t.Error(v, v3)
			return
		}
	} else {
		t.Error()
		return
	}
}
