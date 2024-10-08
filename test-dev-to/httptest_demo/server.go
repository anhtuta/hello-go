package httptest_demo

import (
	"net/http"
)

// All Go HTTP handlers take in an implementation of the ResponseWriter interface, and a Request struct:
// The Request contains the data of the HTTP request, such as URL, headers, and body.
// The ResponseWriter is used to write the response back to the client.

func handleSlothfulMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// Not necessary, as 200/OK is the default code. Only need this when you want to set a different status code
	// w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"message": "Stay slothful!"}`))
}

func handleSlothfulMessagePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not allowed"}`))
		return
	}
	handleSlothfulMessage(w, r)
}

func handleErrorMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error_message": "description of the error"}`))
}

func AppRouter() http.Handler {
	rt := http.NewServeMux()
	rt.HandleFunc("/sloth", handleSlothfulMessage)
	rt.HandleFunc("/sloth-post", handleSlothfulMessagePost)
	rt.HandleFunc("/error", handleErrorMessage)
	return rt
}
