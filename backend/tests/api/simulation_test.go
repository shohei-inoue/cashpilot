package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSimulationRun(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)
	body := map[string]interface{}{
		"period_months":  12,
		"monthly_income": 300000,
		"monthly_expense": 200000,
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/simulation/run", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}

	var res struct {
		StartBalance    int `json:"start_balance"`
		MinBalance      int `json:"min_balance"`
		EndBalance      int `json:"end_balance"`
		MonthlyBalances []struct {
			Month   string `json:"month"`
			Balance int    `json:"balance"`
		} `json:"monthly_balances"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(res.MonthlyBalances) != 12 {
		t.Errorf("len(monthly_balances) = %d, want 12", len(res.MonthlyBalances))
	}
	// start=0, monthly_net=100000, end after 12 months = 1,200,000
	expectedEnd := res.StartBalance + 12*100000
	if res.EndBalance != expectedEnd {
		t.Errorf("end_balance = %d, want %d", res.EndBalance, expectedEnd)
	}
}

func TestSimulationRunUnauthorized(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	body := map[string]interface{}{"period_months": 12}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/simulation/run", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

func TestSimulationRunInvalidPeriod(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)

	for _, invalid := range []int{0, -1, 121} {
		body := map[string]interface{}{"period_months": invalid}
		b, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/api/simulation/run", bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
		for _, c := range cookies {
			req.AddCookie(c)
		}
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("period_months=%d: status = %d, want 400, body = %s", invalid, rec.Code, rec.Body.String())
		}
	}
}

func TestSimulationRunWithTransactions(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)

	// Create account
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

	// Create transaction (current balance = 50000)
	txBody := map[string]interface{}{
		"account_id":  accRes.ID,
		"amount":      50000,
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
		t.Fatalf("transaction create failed: status = %d", txRec.Code)
	}

	// Run simulation
	simBody := map[string]interface{}{
		"period_months":  3,
		"monthly_income": 100000,
		"monthly_expense": 30000,
	}
	sb, _ := json.Marshal(simBody)
	simReq := httptest.NewRequest(http.MethodPost, "/api/simulation/run", bytes.NewReader(sb))
	simReq.Header.Set("Content-Type", "application/json")
	for _, c := range cookies {
		simReq.AddCookie(c)
	}
	simRec := httptest.NewRecorder()
	r.ServeHTTP(simRec, simReq)

	if simRec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d, body = %s", simRec.Code, http.StatusOK, simRec.Body.String())
	}

	var res struct {
		StartBalance int `json:"start_balance"`
		EndBalance   int `json:"end_balance"`
	}
	if err := json.NewDecoder(simRec.Body).Decode(&res); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if res.StartBalance != 50000 {
		t.Errorf("start_balance = %d, want 50000", res.StartBalance)
	}
	// 50000 + 3 * (100000 - 30000) = 50000 + 210000 = 260000
	expectedEnd := 50000 + 3*70000
	if res.EndBalance != expectedEnd {
		t.Errorf("end_balance = %d, want %d", res.EndBalance, expectedEnd)
	}
}
