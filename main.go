package main
import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Eu sou Full Cycle -  Desafio 1 - Go Go Go")
	

}


func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}