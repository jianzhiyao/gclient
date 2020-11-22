package request

import (
	"bytes"
	"encoding"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"github.com/jianzhiyao/gclient/consts"
	"github.com/jianzhiyao/gclient/consts/content_type"
	"github.com/jianzhiyao/gclient/request/form"
	"github.com/jianzhiyao/gclient/request/multipart_form"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodConnect,
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPut,
		http.MethodTrace,
	}
	url := "https://cn.bing.com"
	for _, method := range methods {
		req, err := New(method, url)
		if err != nil {
			t.Error(err)
			return
		}

		if req.method != method {
			t.Error(err)
			return
		}

		if req.url != url {
			t.Error(err)
			return
		}
	}

}

func TestNew2(t *testing.T) {
	method, url := `1`, "https://cn.bing.com"
	_, err := New(method, url)
	if err == nil {
		t.Error(err)
		return
	}
}

func TestRequest_Method(t *testing.T) {
	method, url := http.MethodPatch, "https://cn.bing.com"
	req, _ := New(method, url)

	if req.GetMethod() != method {
		t.Error()
		return
	}

	if req.GetMethod() == http.MethodGet {
		t.Error()
		return
	}
}

func TestRequest_Headers(t *testing.T) {
	method, url := http.MethodPatch, "https://cn.bing.com"
	req, _ := New(method, url)

	req.SetHeader("a", "1")
	req.SetHeader("a", "1")
	req.SetHeader("b", "1")

	if len(req.GetHeaders()) != 2 {
		t.Error()
		return
	}
}

func TestRequest_SetHeader(t *testing.T) {
	method, url := http.MethodPatch, "https://cn.bing.com"
	req, _ := New(method, url)

	req.SetHeader("a", "1")
	req.SetHeader("a", "1")
	req.SetHeader("b", "1")

	if _, ok := req.headers["a"]; !ok {
		t.Error()
		return
	}
	if _, ok := req.headers["b"]; !ok {
		t.Error()
		return
	}
	if _, ok := req.headers["c"]; ok {
		t.Error()
		return
	}
}

//todo build request

func TestRequest_Body(t *testing.T) {
	method, url := http.MethodPatch, "https://cn.bing.com"
	req, _ := New(method, url)

	content := `TestRequest_Body`
	if err := req.Body(content); err != nil {
		t.Error(err)
		return
	}

	if body, err := ioutil.ReadAll(req.GetBody()); err != nil {
		t.Error(err)
		return
	} else {
		if string(body) != content {
			t.Error()
			return
		}
	}
}

func TestRequest_Body2(t *testing.T) {
	method, url := http.MethodPatch, "https://cn.bing.com"
	req, _ := New(method, url)

	content := `TestRequest_Body2`
	if err := req.Body([]byte(content)); err != nil {
		t.Error(err)
		return
	}

	if body, err := ioutil.ReadAll(req.GetBody()); err != nil {
		t.Error(err)
		return
	} else {
		if string(body) != content {
			t.Error()
			return
		}
	}
}

func TestRequest_Body3(t *testing.T) {
	method, url := http.MethodPatch, "https://cn.bing.com"
	req, _ := New(method, url)

	content := `TestRequest_Body3`
	if err := req.Body(bytes.NewBufferString(content)); err != nil {
		t.Error(err)
		return
	}

	if body, err := ioutil.ReadAll(req.GetBody()); err != nil {
		t.Error(err)
		return
	} else {
		if string(body) != content {
			t.Error()
			return
		}
	}
}

func TestRequest_Body4(t *testing.T) {
	method, url := http.MethodPatch, "https://cn.bing.com"
	req, _ := New(method, url)

	content := selfContent(`TestRequest_Body4`)
	if err := req.Body(&content); err != nil {
		t.Error(err)
		return
	}

	if body, err := ioutil.ReadAll(req.GetBody()); err != nil {
		t.Error(err)
		return
	} else {
		if string(body) != string(content) {
			t.Error()
			return
		}
	}
}

type selfContent string

func (c *selfContent) MarshalBinary() (data []byte, err error) {
	return []byte(*c), nil
}

func TestRequest_Form(t *testing.T) {
	method, url := http.MethodPost, "https://cn.bing.com"
	req, _ := New(method, url)

	err := req.Form(
		form.Value(`field1`, `value1`),
		form.Values(`field2`, `value2_1`, `value2_2`),
	)

	if err != nil {
		t.Error()
		return
	}

	if body, err := ioutil.ReadAll(req.GetBody()); err != nil {
		t.Error(err)
		return
	} else {
		strBody := string(body)

		t.Log("form", strBody)
		if !strings.Contains(strBody, `field1=value1`) {
			t.Error(err)
			return
		}
		if !strings.Contains(strBody, `field2=value2_1`) {
			t.Error(err)
			return
		}
		if !strings.Contains(strBody, `field2=value2_2`) {
			t.Error(err)
			return
		}
	}

	if value, ok := req.GetHeader(consts.HeaderContentType); ok {
		if value != content_type.ApplicationXWwwFormUrlencoded {
			t.Error(value)
			return
		}
	} else {
		t.Error()
		return
	}
}

type jsonBody struct {
	Value1 int    `json:"value1"`
	Value2 int    `json:"value2"`
	Value3 string `json:"value3"`
}

var _ encoding.BinaryMarshaler = new(jsonBody)

func (c *jsonBody) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

func TestRequest_Json(t *testing.T) {
	method, url := http.MethodPost, "https://cn.bing.com"
	req, _ := New(method, url)

	v1, v2, v3 := 1223455, 232123123, "asdiu1o2i3jlk"
	err := req.Json(&jsonBody{
		Value1: v1,
		Value2: v2,
		Value3: v3,
	})
	if err != nil {
		t.Error()
		return
	}

	if body, err := ioutil.ReadAll(req.GetBody()); err != nil {
		t.Error(err)
		return
	} else {
		strBody := string(body)

		t.Log("json", strBody)
		if !strings.Contains(strBody, fmt.Sprint(v1)) {
			t.Error(strBody)
			return
		}
		if !strings.Contains(strBody, fmt.Sprint(v2)) {
			t.Error(strBody)
			return
		}
		if !strings.Contains(strBody, fmt.Sprint(v3)) {
			t.Error(strBody)
			return
		}
	}

	if value, ok := req.GetHeader(consts.HeaderContentType); ok {
		if value != content_type.ApplicationJson {
			t.Error(value)
			return
		}
	} else {
		t.Error()
		return
	}
}

type xmlBody struct {
	Value1 int    `json:"value1" xml:"value_1"`
	Value2 int    `json:"value2" xml:"value_2"`
	Value3 string `json:"value3" xml:"value_3"`
}

var _ encoding.BinaryMarshaler = new(xmlBody)

func (c *xmlBody) MarshalBinary() (data []byte, err error) {
	return xml.Marshal(c)
}

func TestRequest_Xml(t *testing.T) {
	method, url := http.MethodPost, "https://cn.bing.com"
	req, _ := New(method, url)

	v1, v2, v3 := 2323124344, 123412, "ssdas412322"
	err := req.Xml(&xmlBody{
		Value1: v1,
		Value2: v2,
		Value3: v3,
	})
	if err != nil {
		t.Error()
		return
	}

	if body, err := ioutil.ReadAll(req.GetBody()); err != nil {
		t.Error(err)
		return
	} else {
		strBody := string(body)

		t.Log("xml", strBody)
		if !strings.Contains(strBody, fmt.Sprint(v1)) {
			t.Error(strBody)
			return
		}
		if !strings.Contains(strBody, fmt.Sprint(v2)) {
			t.Error(strBody)
			return
		}
		if !strings.Contains(strBody, fmt.Sprint(v3)) {
			t.Error(strBody)
			return
		}
	}

	if value, ok := req.GetHeader(consts.HeaderContentType); ok {
		if value != content_type.ApplicationXml {
			t.Error(value)
			return
		}
	} else {
		t.Error()
		return
	}
}

func TestRequest_MultiForm(t *testing.T) {
	method, url := http.MethodPost, "https://cn.bing.com"
	req, _ := New(method, url)

	bd := uuid.New().String()
	uid := uuid.New().String()
	filePath := `multipart_form/test_files/test_file1.txt`
	err := req.MultiForm(
		multipart_form.Boundary(bd),
		multipart_form.Field("uuid", uid),
		multipart_form.File("file", filePath),
	)
	if err != nil {
		t.Error(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var body []byte
		tmpBytes := make([]byte, 128)
		for {
			if count, err := req.GetBody().Read(tmpBytes); err != io.EOF {
				for i := 0; i < count; i++ {
					body = append(body, tmpBytes[i])
				}
			} else {
				break
			}
		}

		strBody := string(body)
		if strings.Count(strBody, fmt.Sprint(bd)) != 3 {
			t.Error(strBody)
			return
		}
		if !strings.Contains(strBody, fmt.Sprint(uid)) {
			t.Error(strBody)
			return
		}
		if !strings.Contains(strBody, fmt.Sprint(filePath)) {
			t.Error(strBody)
			return
		}

		wg.Done()
	}()

	wg.Wait()
}
