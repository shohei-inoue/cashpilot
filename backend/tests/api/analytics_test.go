package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransactionSummary(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)
	req := httptest.NewRequest(http.MethodGet, "/api/transactions/summary", nil)
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}

	var res struct {
		Summary struct {
			TotalIncome     int `json:"total_income"`
			TotalExpense    int `json:"total_expense"`
			NetCashflow     int `json:"net_cashflow"`
			TransactionCount int `json:"transaction_count"`
		} `json:"summary"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode: %v", err)
	}
}

func TestTransactionSummaryUnauthorized(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/transactions/summary", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

func TestTransactionSummaryWithTransactions(t *testing.T) {
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

	// Create income
	txBody1 := map[string]interface{}{
		"account_id":  accRes.ID,
		"amount":      50000,
		"occurred_at": "2025-01-10T12:00:00Z",
	}
	tb1, _ := json.Marshal(txBody1)
	txReq1 := httptest.NewRequest(http.MethodPost, "/api/transactions", bytes.NewReader(tb1))
	txReq1.Header.Set("Content-Type", "application/json")
	for _, c := range cookies {
		txReq1.AddCookie(c)
	}
	txRec1 := httptest.NewRecorder()
	r.ServeHTTP(txRec1, txReq1)
	if txRec1.Code != http.StatusCreated {
		t.Fatalf("income create failed: status = %d", txRec1.Code)
	}

	// Create expense
	txBody2 := map[string]interface{}{
		"account_id":  accRes.ID,
		"amount":      -20000,
		"occurred_at": "2025-01-15T12:00:00Z",
	}
	tb2, _ := json.Marshal(txBody2)
	txReq2 := httptest.NewRequest(http.MethodPost, "/api/transactions", bytes.NewReader(tb2))
	txReq2.Header.Set("Content-Type", "application/json")
	for _, c := range cookies {
		txReq2.AddCookie(c)
	}
	txRec2 := httptest.NewRecorder()
	r.ServeHTTP(txRec2, txReq2)
	if txRec2.Code != http.StatusCreated {
		t.Fatalf("expense create failed: status = %d", txRec2.Code)
	}

	// Get summary
	summaryReq := httptest.NewRequest(http.MethodGet, "/api/transactions/summary", nil)
	for _, c := range cookies {
		summaryReq.AddCookie(c)
	}
	summaryRec := httptest.NewRecorder()
	r.ServeHTTP(summaryRec, summaryReq)

	if summaryRec.Code != http.StatusOK {
		t.Errorf("summary status = %d, want %d, body = %s", summaryRec.Code, http.StatusOK, summaryRec.Body.String())
	}

	var res struct {
		Summary struct {
			TotalIncome      int `json:"total_income"`
			TotalExpense     int `json:"total_expense"`
			NetCashflow      int `json:"net_cashflow"`
			TransactionCount int `json:"transaction_count"`
		} `json:"summary"`
	}
	if err := json.NewDecoder(summaryRec.Body).Decode(&res); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if res.Summary.TotalIncome != 50000 {
		t.Errorf("total_income = %d, want 50000", res.Summary.TotalIncome)
	}
	if res.Summary.TotalExpense != -20000 {
		t.Errorf("total_expense = %d, want -20000", res.Summary.TotalExpense)
	}
	if res.Summary.NetCashflow != 30000 {
		t.Errorf("net_cashflow = %d, want 30000", res.Summary.NetCashflow)
	}
	if res.Summary.TransactionCount != 2 {
		t.Errorf("transaction_count = %d, want 2", res.Summary.TransactionCount)
	}
}

func TestAnalyticsCashflow(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)
	req := httptest.NewRequest(http.MethodGet, "/api/analytics/cashflow", nil)
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}

	var res struct {
		Items   []struct {
			Period          string `json:"period"`
			TotalIncome     int    `json:"total_income"`
			TotalExpense    int    `json:"total_expense"`
			NetCashflow     int    `json:"net_cashflow"`
			TransactionCount int   `json:"transaction_count"`
		} `json:"items"`
		GroupBy string `json:"group_by"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if res.GroupBy != "monthly" {
		t.Errorf("group_by = %q, want monthly", res.GroupBy)
	}
}

func TestAnalyticsCashflowUnauthorized(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/analytics/cashflow", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
}

func TestAnalyticsCashflowGroupBy(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)

	for _, groupBy := range []string{"all", "daily", "weekly", "monthly", "yearly"} {
		req := httptest.NewRequest(http.MethodGet, "/api/analytics/cashflow?group_by="+groupBy, nil)
		for _, c := range cookies {
			req.AddCookie(c)
		}
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("group_by=%s: status = %d, want %d, body = %s", groupBy, rec.Code, http.StatusOK, rec.Body.String())
		}

		var res struct {
			Items   []struct {
				Period string `json:"period"`
			} `json:"items"`
			GroupBy string `json:"group_by"`
		}
		if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
			t.Fatalf("group_by=%s decode: %v", groupBy, err)
		}
		if res.GroupBy != groupBy {
			t.Errorf("group_by=%s: response group_by = %q", groupBy, res.GroupBy)
		}
	}
}

func TestAnalyticsCashflowGroupByAll(t *testing.T) {
	r, err := SetupRouterWithDB()
	if err != nil {
		t.Skipf("skip: DB not available: %v", err)
	}

	cookies := signupAndGetCookies(t, r)
	req := httptest.NewRequest(http.MethodGet, "/api/analytics/cashflow?group_by=all", nil)
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}

	var res struct {
		Items   []struct {
			Period string `json:"period"`
		} `json:"items"`
		GroupBy string `json:"group_by"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if res.GroupBy != "all" {
		t.Errorf("group_by = %q, want all", res.GroupBy)
	}
	if len(res.Items) != 1 {
		t.Errorf("len(items) = %d, want 1", len(res.Items))
	}
	if len(res.Items) > 0 && res.Items[0].Period != "all" {
		t.Errorf("items[0].period = %q, want all", res.Items[0].Period)
	}
}
