package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// signupAndGetCookies はサインアップして認証済み Cookie を返す
func signupAndGetCookies(t *testing.T, r http.Handler) []*http.Cookie {
	t.Helper()
	body := map[string]string{"email": uniqueEmail(), "password": "password123"}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/signup", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("signup failed: status = %d", rec.Code)
	}
	return rec.Result().Cookies()
}

func TestAccountList(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)
	req := httptest.NewRequest(http.MethodGet, "/api/accounts", nil)
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}
}

func TestAccountListUnauthorized(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/accounts", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

func TestAccountCreateAndList(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)

	// POST create
	createBody := map[string]string{"type": "bank", "name": "三井住友銀行"}
	cb, _ := json.Marshal(createBody)
	createReq := httptest.NewRequest(http.MethodPost, "/api/accounts", bytes.NewReader(cb))
	createReq.Header.Set("Content-Type", "application/json")
	for _, c := range cookies {
		createReq.AddCookie(c)
	}
	createRec := httptest.NewRecorder()
	r.ServeHTTP(createRec, createReq)

	if createRec.Code != http.StatusCreated {
		t.Errorf("create status = %d, want %d, body = %s", createRec.Code, http.StatusCreated, createRec.Body.String())
	}

	// GET list
	listReq := httptest.NewRequest(http.MethodGet, "/api/accounts", nil)
	for _, c := range cookies {
		listReq.AddCookie(c)
	}
	listRec := httptest.NewRecorder()
	r.ServeHTTP(listRec, listReq)

	if listRec.Code != http.StatusOK {
		t.Errorf("list status = %d, want %d", listRec.Code, http.StatusOK)
	}
}

func TestAccountCreateInvalidType(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)
	body := map[string]string{"type": "invalid", "name": "テスト口座"}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/accounts", bytes.NewReader(b))
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
