package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCategoryList(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)
	req := httptest.NewRequest(http.MethodGet, "/api/categories", nil)
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}
}

func TestCategoryListUnauthorized(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/categories", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

func TestCategoryCreateAndList(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)

	createBody := map[string]string{"type": "expense", "name": "食費"}
	cb, _ := json.Marshal(createBody)
	createReq := httptest.NewRequest(http.MethodPost, "/api/categories", bytes.NewReader(cb))
	createReq.Header.Set("Content-Type", "application/json")
	for _, c := range cookies {
		createReq.AddCookie(c)
	}
	createRec := httptest.NewRecorder()
	r.ServeHTTP(createRec, createReq)

	if createRec.Code != http.StatusCreated {
		t.Errorf("create status = %d, want %d, body = %s", createRec.Code, http.StatusCreated, createRec.Body.String())
	}

	listReq := httptest.NewRequest(http.MethodGet, "/api/categories", nil)
	for _, c := range cookies {
		listReq.AddCookie(c)
	}
	listRec := httptest.NewRecorder()
	r.ServeHTTP(listRec, listReq)

	if listRec.Code != http.StatusOK {
		t.Errorf("list status = %d, want %d", listRec.Code, http.StatusOK)
	}
}

func TestCategoryCreateInvalidType(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)
	body := map[string]string{"type": "invalid", "name": "テストカテゴリ"}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/categories", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d, body = %s", rec.Code, http.StatusBadRequest, rec.Body.String())
	}
}
