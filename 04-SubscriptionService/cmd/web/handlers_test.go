package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"subscription-service/data"
	"testing"
)

var pageTests = []struct {
	name               string
	url                string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]any
	expectedHTML       string
}{
	{
		name:               "home",
		url:                "/",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.HomePage,
	},
	{
		name:               "login page",
		url:                "/login",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.LoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
	},
	{
		name:               "logout page",
		url:                "/logout",
		expectedStatusCode: http.StatusSeeOther,
		handler:            testApp.Logout,
		sessionData: map[string]any{
			"userID": 1,
			"user":   data.User{},
		},
	},
}

func TestConfig_Pages(t *testing.T) {
	pathToTemplates = "./templates"

	for _, p := range pageTests {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", p.url, nil)

		ctx := getCtx(req)
		req = req.WithContext(ctx)

		if len(p.sessionData) > 0 {
			for key, value := range p.sessionData {
				testApp.Sessions.Put(ctx, key, value)
			}
		}

		p.handler.ServeHTTP(rr, req)

		if rr.Code != p.expectedStatusCode {
			t.Errorf("%s failed: expected %d, but got %d", p.name, p.expectedStatusCode, rr.Code)
		}

		if len(p.expectedHTML) > 0 {
			html := rr.Body.String()
			if !strings.Contains(html, p.expectedHTML) {
				t.Errorf("%s, failed: expected to find %s, but did not ", p.name, p.expectedHTML)
			}
		}
	}
}

func TestConfig_PostLoginPage(t *testing.T) {
	pathToTemplates = "./templates"
	postedData := url.Values{
		"email":    {"admin@example.com"},
		"password": {"abc123abc123abc123abc123"},
	}
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testApp.PostLoginPage)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Error("wrong code returned")
	}

	if !testApp.Sessions.Exists(ctx, "userID") {
		t.Error("did not find userID in session")
	}
}
