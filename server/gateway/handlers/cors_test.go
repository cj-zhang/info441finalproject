package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	wrappedHandler := &CorsMiddleWare{Handler: handler}

	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	rec := httptest.NewRecorder()

	wrappedHandler.ServeHTTP(rec, req)

	resp := rec.Result()

	allowOriginHeader := resp.Header.Get(headerAccessControlAllowOrigin)
	if allowOriginHeader != allowOriginValue {
		t.Errorf("allow origin header is incorrectly set")
	}

	allowMethodsHeader := resp.Header.Get(headerAccessControlAllowMethods)
	if allowMethodsHeader != allowMethodsValue {
		t.Errorf("allow methods header is incorrectly set")
	}

	allowHeadersHeader := resp.Header.Get(headerAccessControlAllowHeaders)
	if allowHeadersHeader != allowHeadersValue {
		t.Errorf("allow methods header is incorrectly set")
	}

	exposeHeadersHeader := resp.Header.Get(headerAccessControlExposeHeaders)
	if exposeHeadersHeader != exposeHeadersValue {
		t.Errorf("allow methods header is incorrectly set")
	}

	maxAgeHeader := resp.Header.Get(headerAccessControlMaxAge)
	if maxAgeHeader != controlMaxAgeValue {
		t.Errorf("allow methods header is incorrectly set")
	}

}
