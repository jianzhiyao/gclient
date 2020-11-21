package form

import (
	"net/url"
	"testing"
)

func TestValue(t *testing.T) {
	m := url.Values{}

	f1, v1 := `f1`, `q`
	f2, v2 := `f2`, `w`
	f3, v3 := `f3`, `e`

	options := []Option{
		Value(f1, v1),
		Value(f2, v2),
		Value(f3, v3),
	}

	for _, option := range options {
		_ = option(m)
	}

	if v, ok := m[f1]; ok {
		if v[0] != v1 {
			t.Error(v, v1)
			return
		}
	} else {
		t.Error()
		return
	}

	if v, ok := m[f2]; ok {
		if v[0] != v2 {
			t.Error(v, v2)
			return
		}
	} else {
		t.Error()
		return
	}

	if v, ok := m[f3]; ok {
		if v[0] != v3 {
			t.Error(v, v3)
			return
		}
	} else {
		t.Error()
		return
	}
}

func TestValues(t *testing.T) {
	m := url.Values{}

	f1, v11, v12, v13 := `f1`, `q`, `w`, `e`
	f2, v21, v22, v23, v24 := `f2`, `a`, `s`, `d`, `f`
	f3, v31, v32 := `f3`, `z`, `x`

	options := []Option{
		Values(f1, v11, v12, v13),
		Values(f2, v21, v22, v23, v24),
		Values(f3, v31, v32),
	}

	for _, option := range options {
		_ = option(m)
	}

	if v, ok := m[f1]; ok {
		i := uint(0)

		if v[i] != v11 {
			t.Error(v, v11)
			return
		}
		i++

		if v[i] != v12 {
			t.Error(v, v12)
			return
		}
		i++

		if v[i] != v13 {
			t.Error(v, v13)
			return
		}
		i++

	} else {
		t.Error()
		return
	}

	if v, ok := m[f2]; ok {
		i := uint(0)

		if v[i] != v21 {
			t.Error(v, v21)
			return
		}
		i++

		if v[i] != v22 {
			t.Error(v, v22)
			return
		}
		i++

		if v[i] != v23 {
			t.Error(v, v23)
			return
		}
		i++

		if v[i] != v24 {
			t.Error(v, v24)
			return
		}
		i++
	} else {
		t.Error()
		return
	}

	if v, ok := m[f3]; ok {
		i := uint(0)

		if v[i] != v31 {
			t.Error(v, v31)
			return
		}
		i++

		if v[i] != v32 {
			t.Error(v, v32)
			return
		}
		i++
	} else {
		t.Error()
		return
	}
}
