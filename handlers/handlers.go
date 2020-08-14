package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func GetURLInput(r *http.Request, s string) []string {
	vars := mux.Vars(r)[s]
	return strings.Split(vars, ",")
}

func WelcomeHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome")

}

func AddHandler(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	inputs := GetURLInput(request, "numbers")
	first := inputs[0]
	second := inputs[1]
	firstNumber, firstErr := strconv.ParseFloat(first, 64)
	if firstErr != nil {
		http.Error(response, firstErr.Error(), http.StatusBadRequest)
	}

	secondNumber, secondErr := strconv.ParseFloat(second, 64)
	if secondErr != nil {
		http.Error(response, secondErr.Error(), http.StatusBadRequest)
	}
	sum := firstNumber + secondNumber
	fmt.Fprint(response, sum)

}

func SubstractHandler(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	inputs := GetURLInput(request, "numbers")
	first := inputs[0]
	second := inputs[1]
	firstNumber, firstErr := strconv.ParseFloat(first, 64)
	if firstErr != nil {
		http.Error(response, firstErr.Error(), http.StatusBadRequest)
	}
	secondNumber, secondErr := strconv.ParseFloat(second, 64)
	if secondErr != nil {
		http.Error(response, secondErr.Error(), http.StatusBadRequest)
	}
	subs := fmt.Sprintf("%.2f", (firstNumber - secondNumber))
	fmt.Fprint(response, subs)

}

func DivisionHandler(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	inputs := GetURLInput(request, "numbers")
	first := inputs[0]
	second := inputs[1]
	firstNumber, firstErr := strconv.ParseFloat(first, 64)
	if firstErr != nil {
		http.Error(response, firstErr.Error(), http.StatusBadRequest)
	}
	secondNumber, secondErr := strconv.ParseFloat(second, 64)
	if secondErr != nil {
		http.Error(response, secondErr.Error(), http.StatusBadRequest)
	}
	div := firstNumber / secondNumber
	fmt.Fprint(response, div)

}

func RandomHandler(w http.ResponseWriter, request *http.Request) {
	//start := time.Now()
	vars := mux.Vars(request)
	counter, _ := strconv.Atoi(vars["count"])
	rand.Seed(time.Now().Unix())
	if counter != 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		randm, _ := json.Marshal(rand.Perm(counter))
		w.Write([]byte(randm))
		// fmt.Fprint(w, rand.Perm(counter))
	} else {
		w.Header().Set("Content-Type", "application/json")
		randm, _ := json.Marshal(rand.Perm(10))
		w.Write([]byte(randm))
		// fmt.Fprint(w, rand.Perm(10))
	}
}

func LivenessCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `200`)

}

var StartTime time.Time

func ReadinessCheckHandler(w http.ResponseWriter, r *http.Request) {
	duration := time.Now().Sub(StartTime)
	if duration.Seconds() > 10 {
		w.WriteHeader(200)
		w.Write([]byte("200"))
	} else {
		w.WriteHeader(500)
		w.Write([]byte("Not Ready"))
	}
}
