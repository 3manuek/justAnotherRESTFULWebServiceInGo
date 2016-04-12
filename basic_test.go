package main

import (
	"testing"
    "github.com/gorilla/mux"
//    "encoding/json"
//    "log"
 //   "fmt"
    "net/http"
//    "strconv"
)

func Test(t *testing.T) {
   r := mux.NewRouter()
   r.HandleFunc("/", HomeHandler)
   r.HandleFunc("/fibo/{value:[0-9]+}", FiboHandler)

   http.Handle("/",r)
   go http.ListenAndServe(":3333", nil)

   _, err := http.Get("http://localhost:3333/fibo/55") 
   if err != nil {
	t.Error(err)
   }
//	body, _ := ioutil.ReadAll(resp.Body)
//	if string(body) != "{\n  \"items\": [\n    \"item1\",\n    \"item2\"\n  ]\n}" {
//		t.Error("Not equal.")
//	}

        //go basic.go
	// commands := commands()
	// if len(commands) == 0 {
	// 	t.Fail()
	// }
}
