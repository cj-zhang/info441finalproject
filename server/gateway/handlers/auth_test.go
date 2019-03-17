package handlers

import (
	"assignments-cj-zhang/servers/gateway/models/users"
	"assignments-cj-zhang/servers/gateway/sessions"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestUsersHandler(t *testing.T) {
	cases := []string{
		`{ 
			"email": "test@test.com",
			"password": "passHash",
			"passwordConf": "passHash",
			"userName": "usernameTest",
			"firstName": "firstName",
			"lastName": "lastName"
		}`,
		`
			"": "test@test.com",
			"password": "passHash",
			"passwordConf": "passHash",
			"userName": "usernameTest",
			"firstName": "firstName",
			"lastName": "lastName"
		}`,
		`{ 
			"password": "passHash",
			"passwordConf": "passHash",
			"userName": "usernameTest",
			"firstName": "firstName",
			"lastName": "lastName"
		}`,
	}

	ctx := makeNewContext()
	req, err := http.NewRequest("POST", "/v1/users", strings.NewReader(cases[0]))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctx.UsersHandler)

	// trigger with no json header
	handler.ServeHTTP(rr, req)
	// trigger correctly
	req.Header.Set(contentType, applicationJSON)
	handler.ServeHTTP(rr, req)
	for _, c := range cases {
		req, err = http.NewRequest("POST", "/v1/users", strings.NewReader(c))
		req.Header.Set(contentType, applicationJSON)
		handler.ServeHTTP(rr, req)
	}

	_, err = ctx.UserStore.GetByEmail("test@test.com")
	if err != nil {
		t.Fatal("couldn't find user: " + err.Error())
	}

	// test bad method
	req, err = http.NewRequest("GET", "/v1/users", nil)
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Fatal("failed getting the users:" + err.Error())
	}
}

func TestSpecificUserHandler(t *testing.T) {
	// reused variables
	ctx := makeNewContext()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ctx.SpecificUserHandler)

	// test the "me" request for GET
	req, err := http.NewRequest("GET", "/v1/users/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)

	// test an ID request for GET
	req, err = http.NewRequest("GET", "/v1/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)

	// test PATCH
	body, _ := json.Marshal(&users.Updates{
		FirstName: "firstUpdates",
		LastName:  "lastUpdates",
	})

	// invalid user in URL
	req, err = http.NewRequest("PATCH", "/v1/users/1293", strings.NewReader(string(body)))
	req.Header.Set(contentType, applicationJSON)
	handler.ServeHTTP(rr, req)

	req, err = http.NewRequest("PATCH", "/v1/users/1", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	// before setting header to json
	handler.ServeHTTP(rr, req)
	// setting the header
	req.Header.Set(contentType, applicationJSON)
	handler.ServeHTTP(rr, req)

	// incorrect method type
	req, err = http.NewRequest("POST", "/v1/users/1", strings.NewReader(string(body)))
	handler.ServeHTTP(rr, req)
}

func TestSessionsHandler(t *testing.T) {
	correctCred := strings.NewReader(`{
		"email": "test@test.com",
		"password": "passHash"
	}`)
	incorrect := strings.NewReader(`{
		"email": "test@test.com",
		"password": "passwordHash\n"
	}`)

	ctx := makeNewContext()
	handler := http.HandlerFunc(ctx.SessionsHandler)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/v1/sessions", correctCred)
	if nil != err {
		t.Fatal(err)
	}
	req.Header.Set(contentType, applicationJSON)

	handler.ServeHTTP(rr, req)

	// incorrect json
	req, err = http.NewRequest("POST", "/v1/sessions", incorrect)
	handler.ServeHTTP(rr, req)

	// incorrect method type
	req, err = http.NewRequest("PATCH", "/v1/sessions", incorrect)
	handler.ServeHTTP(rr, req)
}

func TestSpecificSessionsHandler(t *testing.T) {
	ctx := makeNewContext()
	handler := http.HandlerFunc(ctx.SpecificSessionHandler)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/v1/sessions/mine", nil)
	if nil != err {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)

	// incorrect URL
	req, err = http.NewRequest("DELETE", "/v1/sessions/", nil)
	handler.ServeHTTP(rr, req)

	// incorrect method type
	req, err = http.NewRequest("PATCH", "/v1/sessions/", nil)
	handler.ServeHTTP(rr, req)

}

func makeNewContext() (ctx *HandlerContext) {
	return &HandlerContext{
		SigningKey: "signing Key",
		UserStore:  &MyMockStore{},
		SessStore:  sessions.NewMemStore(time.Minute, 0),
	}
}
