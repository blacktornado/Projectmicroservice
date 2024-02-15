package main

import (
	"encoding/json"
	"log"
	lyricsmicroservice "lyrics/Api"
	"net/http"
)

func main() {
	http.HandleFunc("/lyrics", func(w http.ResponseWriter, r *http.Request) {
		trackID := r.URL.Query().Get("track_id")
		lyrics, err := lyricsmicroservice.GetLyrics(trackID)
		if err != nil {
			http.Error(w, "Failed to fetch lyrics", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lyrics)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
