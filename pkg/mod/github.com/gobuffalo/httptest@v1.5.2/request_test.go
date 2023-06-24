package httptest

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HTML_Headers_Dont_Overwrite_App_Headers(t *testing.T) {
	r := require.New(t)
	w := New(App())
	w.Headers["foo"] = "bar"

	req := w.HTML("/")
	req.Headers["foo"] = "baz"
	r.Equal("baz", req.Headers["foo"])
	r.Equal("bar", w.Headers["foo"])
}

func Test_Get(t *testing.T) {
	r := require.New(t)
	w := New(App())

	req := w.HTML("/get")
	r.Equal("/get", req.URL)

	res := req.Get()
	r.Equal(201, res.Code)
	r.Contains(res.Body.String(), "METHOD:GET")
	r.Contains(res.Body.String(), "Hello from Get!")
}

func Test_Delete(t *testing.T) {
	r := require.New(t)
	w := New(App())

	req := w.HTML("/delete")
	r.Equal("/delete", req.URL)

	res := req.Delete()
	r.Contains(res.Body.String(), "METHOD:DELETE")
	r.Contains(res.Body.String(), "Goodbye")
}

func Test_Post_Struct(t *testing.T) {
	r := require.New(t)
	w := New(App())

	req := w.HTML("/post")
	res := req.Post(User{Name: "Mark"})
	r.Contains(res.Body.String(), "METHOD:POST")
	r.Contains(res.Body.String(), "NAME:Mark")
}

func Test_Post_Struct_Pointer(t *testing.T) {
	r := require.New(t)
	w := New(App())

	req := w.HTML("/post")
	res := req.Post(&User{Name: "Mark"})
	r.Contains(res.Body.String(), "METHOD:POST")
	r.Contains(res.Body.String(), "NAME:Mark")
}

func Test_Post_Values(t *testing.T) {
	r := require.New(t)
	w := New(App())

	req := w.HTML("/post")
	vals := url.Values{}
	vals.Add("name", "Mark")
	res := req.Post(vals)
	r.Contains(res.Body.String(), "METHOD:POST")
	r.Contains(res.Body.String(), "NAME:Mark")
}

func Test_Put(t *testing.T) {
	r := require.New(t)
	w := New(App())

	req := w.HTML("/put")
	res := req.Put(User{Name: "Mark"})
	r.Contains(res.Body.String(), "METHOD:PUT")
	r.Contains(res.Body.String(), "NAME:Mark")
}

func Test_Put_Struct_Pointer(t *testing.T) {
	r := require.New(t)
	w := New(App())

	req := w.HTML("/put")
	res := req.Put(&User{Name: "Mark"})
	r.Contains(res.Body.String(), "METHOD:PUT")
	r.Contains(res.Body.String(), "NAME:Mark")
}

func Test_Put_Values(t *testing.T) {
	r := require.New(t)
	w := New(App())

	req := w.HTML("/put")
	vals := url.Values{}
	vals.Add("name", "Mark")
	res := req.Put(vals)
	r.Contains(res.Body.String(), "METHOD:PUT")
	r.Contains(res.Body.String(), "NAME:Mark")
}
