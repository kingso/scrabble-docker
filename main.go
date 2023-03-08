package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unicode"
)

type Reply struct {
	Word   string  `json:"word"`
	Length int     `json:"length"`
	Score  int     `json:"score"`
	Ratio  float32 `json:"score-ratio"`
}

var alphabet = [26]int{1, 3, 3, 4, 1, 4, 4, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	if !IsLetter(word) {
		http.Error(w, "Your request was invalid", 400)
		return
	}
	score := GetScore(word)
	length := len(word)
	reply := Reply{
		Word:   word,
		Length: length,
		Score:  score,
		Ratio:  float32(length) / float32(score),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)
	if length != 0 {
		log.Print(reply.Word)
	}
}

func main() {
	http.HandleFunc("/", baseHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func GetScore(word string) (score int) {
	text := strings.ToUpper(word)
	for _, ltr := range text {
		ltr := string(ltr)
		ascii := []byte(ltr)[0] - 65
		score += alphabet[ascii]
	}
	return
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
