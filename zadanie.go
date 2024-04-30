package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeResponse struct {
	Date string `json:"Date"`
	Time string `json:"Time"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		answer := TimeResponse{
			Date: currentTime.Format("02.01.2006"),
			Time: currentTime.Format("15:04"),
		}

		jsonanswer, _ := json.Marshal(answer)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonanswer)
	})
	http.ListenAndServe(":10533", nil)
}
