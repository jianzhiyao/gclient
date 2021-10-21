package response

import (
	"github.com/jianzhiyao/gclient/tests"
	"net/http"
	"testing"
)

func TestResponse_Bytes(t *testing.T) {
	url := tests.GetServerUrl() + `ok`
	c := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	if resp, err := c.Do(req); err != nil {
		t.Error(err)
	} else {
		r := Response{resp: resp}

		if b, e := r.Bytes(); e != nil {
			t.Error(e)
		} else if string(b) != `ok` {
			t.Error(string(b))
		}
	}
}

func TestResponse_String(t *testing.T) {
	url := tests.GetServerUrl() + `ok`
	c := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	if resp, err := c.Do(req); err != nil {
		t.Error(err)
	} else {
		r := Response{resp: resp}

		if b, e := r.String(); e != nil {
			t.Error(e)
		} else if b != `ok` {
			t.Error(b)
		}
	}
}

func TestResponse_JsonUnmarshal(t *testing.T) {
	type J struct {
		Code    int      `json:"code"`
		Message string   `json:"message"`
		Users   []string `json:"users"`
	}

	var j J

	url := tests.GetServerUrl() + `json`
	c := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	if resp, err := c.Do(req); err != nil {
		t.Error(err)
	} else {
		r := Response{resp: resp}

		if e := r.JsonUnmarshal(&j); e != nil {
			t.Error(e)
		}

		if j.Code != 1 {
			t.Error()
		}

		if j.Message != `ok` {
			t.Error()
		}

		if len(j.Users) != 2 || j.Users[0] != `aaron` || j.Users[1] != `john` {
			t.Error()
		}
	}
}

func TestResponse_XmlUnmarshal(t *testing.T) {
	type X struct {
		Message string `xml:"message"`
	}

	var x X

	url := tests.GetServerUrl() + `xml`
	c := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	if resp, err := c.Do(req); err != nil {
		t.Error(err)
	} else {
		r := Response{resp: resp}
		if e := r.XmlUnmarshal(&x); e != nil {
			t.Error(e)
		}
		if x.Message != `ok` {
			t.Error()
		}
	}
}

func TestResponse_YamlUnmarshal(t *testing.T) {
	type Y struct {
		Code    int      `json:"code"`
		Message string   `json:"message"`
		Users   []string `json:"users"`
	}

	var y Y

	url := tests.GetServerUrl() + `yaml`
	c := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	if resp, err := c.Do(req); err != nil {
		t.Error(err)
	} else {
		r := Response{resp: resp}
		if e := r.YamlUnmarshal(&y); e != nil {
			t.Error(e)
		}
		if y.Code != 1 {
			t.Error()
		}

		if y.Message != `ok` {
			t.Error()
		}

		if len(y.Users) != 2 || y.Users[0] != `aaron` || y.Users[1] != `john` {
			t.Error()
		}
	}
}
