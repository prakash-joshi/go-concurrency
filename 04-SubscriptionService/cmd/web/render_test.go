package main

import (
	"net/http"
	"testing"
)

func TestConfig_AddDefault(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	ctx := getCtx(req)

	req = req.WithContext(ctx)

	testApp.Sessions.Put(ctx, "flash", "flash")
	testApp.Sessions.Put(ctx, "warning", "warning")
	testApp.Sessions.Put(ctx, "error", "error")

	td := testApp.AddDefaultData(&TemplateData{}, req)

	if td.Flash != "flash" {
		t.Error("failed to get flash data")
	}

	if td.Warning != "warning" {
		t.Error("failed to get warning data")
	}

	if td.Error != "error" {
		t.Error("failed to get error data")
	}

}

func TestConfig_IsAuthenticated(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	auth := testApp.isAuthenticated(req)

	if auth {
		t.Error("returns true for authenticated, when it should be false")
	}

	testApp.Sessions.Put(ctx, "userID", 1)

	auth = testApp.isAuthenticated(req)

	if !auth {
		t.Error("returns false for authenticated, when it should be true")
	}

}
