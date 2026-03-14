package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransactionList(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)
	req := httptest.NewRequest(http.MethodGet, "/api/transactions", nil)
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}
}

func TestTransactionListUnauthorized(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/transactions", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

func TestTransactionCreateAndList(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)

	// Create account first
	accBody := map[string]string{"type": "bank", "name": "テスト銀行"}
	ab, _ := json.Marshal(accBody)
	accReq := httptest.NewRequest(http.MethodPost, "/api/accounts", bytes.NewReader(ab))
	accReq.Header.Set("Content-Type", "application/json")
	for _, c := range cookies {
		accReq.AddCookie(c)
	}
	accRec := httptest.NewRecorder()
	r.ServeHTTP(accRec, accReq)
	if accRec.Code != http.StatusCreated {
		t.Fatalf("account create failed: status = %d", accRec.Code)
	}
	var accRes struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(accRec.Body).Decode(&accRes); err != nil {
		t.Fatalf("decode account: %v", err)
	}

	// Create transaction
	txBody := map[string]interface{}{
		"account_id":  accRes.ID,
		"amount":      -1000,
		"memo":        "テスト支出",
		"occurred_at": "2025-01-15T12:00:00Z",
	}
	tb, _ := json.Marshal(txBody)
	txReq := httptest.NewRequest(http.MethodPost, "/api/transactions", bytes.NewReader(tb))
	txReq.Header.Set("Content-Type", "application/json")
	for _, c := range cookies {
		txReq.AddCookie(c)
	}
	txRec := httptest.NewRecorder()
	r.ServeHTTP(txRec, txReq)

	if txRec.Code != http.StatusCreated {
		t.Errorf("create status = %d, want %d, body = %s", txRec.Code, http.StatusCreated, txRec.Body.String())
	}

	// List
	listReq := httptest.NewRequest(http.MethodGet, "/api/transactions", nil)
	for _, c := range cookies {
		listReq.AddCookie(c)
	}
	listRec := httptest.NewRecorder()
	r.ServeHTTP(listRec, listReq)

	if listRec.Code != http.StatusOK {
		t.Errorf("list status = %d, want %d", listRec.Code, http.StatusOK)
	}
}
