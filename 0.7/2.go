package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(rw, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Printf("%s", dump)
		//fmt.Fprintf(rw, "srsrsrsr")
		//fmt.Fprintf(rw, "%q", dump)

	}))
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL, strings.NewReader("this is a demo"))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}
