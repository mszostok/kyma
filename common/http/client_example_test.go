package http_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	bHTTP "github.com/kyma-project/kyma/common/http"
)

// ExampleDoerWithBasicAuthDo demonstrate the usage of the client with basic auth
func ExampleDoerWithBasicAuthDo() {
	stubServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, _ := r.BasicAuth()
		msg := fmt.Sprintf("Request basic auth username: %s and password %s", u, p)
		w.Write([]byte(msg))
		w.WriteHeader(http.StatusOK)
	}))
	defer stubServer.Close()

	doer := bHTTP.NewClientWithBasicAuth("b-user", "b-pass")

	req, err := http.NewRequest(http.MethodGet, stubServer.URL, nil)
	panicOnErr(err)

	resp, err := doer.Do(req)
	panicOnErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	panicOnErr(err)

	fmt.Printf("%s", body)

	// Output: Request basic auth username: b-user and password b-pass
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
