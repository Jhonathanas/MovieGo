package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"math/rand"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var Movies []Movie

func main(){

}