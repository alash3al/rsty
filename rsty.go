// rsty - just a http.Handler middleware for handling RESTful resources .
//
// rsty released under MIT License,
// Developed by Mohammed Al Ashaal <https://www.alash3al.xyz>
package rsty

import "encoding/json"
import "net/http"
import "net/url"

// A RESTful resource that supports HEAD, GET, POST, PUT, PATCH, DELETE .
type Resource interface {
	HEAD(input url.Values, headers http.Header) (int, http.Header, interface{})
	GET(input url.Values, headers http.Header) (int, http.Header, interface{})
	POST(input url.Values, headers http.Header) (int, http.Header, interface{})
	PUT(input url.Values, headers http.Header) (int, http.Header, interface{})
	PATCH(input url.Values, headers http.Header) (int, http.Header, interface{})
	DELETE(input url.Values, headers http.Header) (int, http.Header, interface{})
}

// This is the default resource type that could be used
// to fill any other resource .
type Defaults struct {}

// Called if the request method were HEAD .
func(Defaults) HEAD(input url.Values, headers http.Header) (int, http.Header, interface{}) {
	return 405, http.Header{"Content-Type": {"application/json; charset=UTF-8"}}, http.StatusText(405)
}

// Called if the request method were GET .
func(Defaults) GET(input url.Values, headers http.Header) (int, http.Header, interface{}) {
	return 405, http.Header{"Content-Type": {"application/json; charset=UTF-8"}}, http.StatusText(405)
}

// Called if the request method were POST .
func(Defaults) POST(input url.Values, headers http.Header) (int, http.Header, interface{}) {
	return 405, http.Header{"Content-Type": {"application/json; charset=UTF-8"}}, http.StatusText(405)
}

// Called if the request method were PUT .
func(Defaults) PUT(input url.Values, headers http.Header) (int, http.Header, interface{}) {
	return 405, http.Header{"Content-Type": {"application/json; charset=UTF-8"}}, http.StatusText(405)
}

// Called if the request method were PATCH .
func(Defaults) PATCH(input url.Values, headers http.Header) (int, http.Header, interface{}) {
	return 405, http.Header{"Content-Type": {"application/json; charset=UTF-8"}}, http.StatusText(405)
}

// Called if the request method were DELETE .
func(Defaults) DELETE(input url.Values, headers http.Header) (int, http.Header, interface{}) {
	return 405, http.Header{"Content-Type": {"application/json; charset=UTF-8"}}, http.StatusText(405)
}

// This will return a http.Handler that could be used anywhere .
func Handle(rsrc Resource) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.ParseForm() != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		var code int
		var body interface{}
		var header http.Header
		switch req.Method {
		case "HEAD":
			code, header, body = rsrc.HEAD(req.Form, req.Header)
		case "GET":
			code, header, body = rsrc.GET(req.Form, req.Header)
		case "POST":
			code, header, body = rsrc.POST(req.Form, req.Header)
		case "PUT":
			code, header, body = rsrc.PUT(req.Form, req.Header)
		case "PATCH":
			code, header, body = rsrc.PATCH(req.Form, req.Header)
		case "DELETE":
			code, header, body = rsrc.DELETE(req.Form, req.Header)
		default:
			code, header, body = 405, http.Header{}, ""
		}
		jbody, err := json.Marshal(body)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		for n, vals := range header {
			for _, val := range vals {
				res.Header().Add(n, val)
			}
		}
		res.WriteHeader(code)
		res.Write(jbody)
	})
}
