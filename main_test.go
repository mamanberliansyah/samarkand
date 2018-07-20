//main_test.go

package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "io/ioutil"
)

//function for testing handler
func TestHandler(t *testing.T){
  req, err := http.NewRequest("GET", "", nil)
  if err != nil {
    t.Fatal(err)
  }

  recorder := httptest.NewRecorder()

  hf := http.HandlerFunc(handler)
  hf.ServeHTTP(recorder, req)

  if status := recorder.Code; status != http.StatusOK{
    t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }

  expected := `SAMARKAND:8080`
  actual := recorder.Body.String()
  if actual != expected {
    t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
  }
}

func TestRouter(t *testing.T){
  router := newRouter()
  mockServer := httptest.NewServer(router)

  resp, err := http.Get(mockServer.URL + "/")
  if err != nil {
    t.Fatal(err)
  }

  if resp.StatusCode != http.StatusOK {
    t.Errorf("Status should be ok, got %d", resp.StatusCode)
  }

  defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
  respString := string(b)
  expected := "SAMARKAND:8080"

  if respString != expected {
    t.Errorf("Response should be %s, got %s", expected, respString)
  }
}

func TestRouterForNonExistentRoute(t *testing.T) {
	router := newRouter()
	mockServer := httptest.NewServer(router)
	resp, err := http.Post(mockServer.URL+"/", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	// The code to test the body is also mostly the same, except this time, we
	// expect an empty body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}
