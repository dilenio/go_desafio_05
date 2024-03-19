package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTemperatureByCEP(t *testing.T) {
	// Cria um request de teste com um CEP válido
	req, err := http.NewRequest("GET", "/45208643", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Cria um response recorder para gravar a resposta do handler
	rr := httptest.NewRecorder()

	// Cria um handler falso usando o roteador que você definiu
	handler := http.HandlerFunc(handleGetTemperatureByCEP)

	// Simula uma solicitação para a rota com o request de teste e o response recorder
	handler.ServeHTTP(rr, req)

	// Verifica o código de status da resposta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verifica se o corpo da resposta contém os dados esperados
	expected := `{"temp_C":28.5,"temp_F":83.3,"temp_K":301.5}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
