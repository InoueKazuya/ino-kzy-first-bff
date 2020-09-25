package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler);
	_ = http.ListenAndServe(":8443", nil)
}

func handler (writer http.ResponseWriter, request *http.Request) {
	resp,err := http.Get("http://10.146.0.1:8080/hello");
	if err != nil {
		// error handling
	};
	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body);
	fmt.Fprint(writer, string(body));
}