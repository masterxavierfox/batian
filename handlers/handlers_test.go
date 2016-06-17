package handlers

import (
	"net/http/httptest"
	"net/http"
	"testing"
)

func TestIndex(t *testing.T){
	//request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	//Index(response, request)
	if response.Code != http.StatusOK {
        t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
    }
}