/*
  Emanuel Calvo - 2016
*/

package main

import (
    "github.com/gorilla/mux"
    "encoding/json"
    "log"
    "fmt"
    "net/http"
    "strconv"
)

type ResFiboSerie struct {
    ReqNumber   uint64
    Numbers   []uint64
}


func fibonacci() func() uint64 {
	first, second := uint64(0), uint64(1)
	return func() uint64 {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte(fmt.Sprintf("This is Home. Please use /fibo/<number>")))
}



func FiboHandler(w http.ResponseWriter, r *http.Request) {
          fx := fibonacci()
          var sl []uint64
          vars := mux.Vars(r)
          value := vars["value"]
          valueUint, err := strconv.ParseUint(value,10,64)

          if err != nil {
            w.Write([]byte(fmt.Sprintf("There was an error converting the value into an unisgned 64 integer ")))
          }

          if valueUint < 0 {
            w.Write([]byte(fmt.Sprintf("Negative numbers not allowed.")))
          }

	// Return the fist N numbers of the serie OR return the serie UNTIL N?
          for i := uint64(0); i < valueUint; i++ {
	  //for {
	     ret := fx()
		if ret > valueUint { break ; } 
              sl = append(sl, ret)
	     
          }  
  	

          serieFibo := ResFiboSerie{
            ReqNumber: valueUint,
            Numbers:   sl,
          }

          responseArrayJson, _ := json.Marshal(serieFibo)

          w.Write([]byte(fmt.Sprintf("%s", responseArrayJson )))
  }

func main() {

   r := mux.NewRouter()
   r.HandleFunc("/", HomeHandler)
   r.HandleFunc("/fibo/{value:[0-9]+}", FiboHandler)

   http.Handle("/",r)
   log.Fatal(http.ListenAndServe(":3333", nil))

}
