package handler

import (
	"encoding/json"
	"hoge/src/go/usecase/sample"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Add POSTのサンプル
func (h *ProdHandlingFunc) Add(w http.ResponseWriter, r *http.Request) {
}

//Update PUTのサンプル
func (h *ProdHandlingFunc) Update(w http.ResponseWriter, r *http.Request) {

}

//Delete DELETEのサンプル
func (h *ProdHandlingFunc) Delete(w http.ResponseWriter, r *http.Request) {

}

//Get GETのサンプル /get/firstName/lastName でfirstName+lastNameをjsonで返すだけ
func (h *ProdHandlingFunc) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := sample.GetFullName(vars)
	log.Println(name)
	json.NewEncoder(w).Encode(name)
}

func outputError(w *http.ResponseWriter, e error, message string) {
	io.WriteString(*w, e.Error())
	log.Println(message, " エラーが発生しました:", e)
}
