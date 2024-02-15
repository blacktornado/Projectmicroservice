package main

import (
	"encoding/json"
	"log"
	"net/http"
	trackmicroservice "track/Api"
)

func main() {
	http.HandleFunc("/top-tracks", func(rw http.ResponseWriter, r *http.Request) {
		location := r.URL.Query().Get("location")
		tracks, err := trackmicroservice.GetTopTracks(location)
		if err != nil {
			http.Error(rw, "Failed to fetch top tracks", http.StatusInternalServerError)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(tracks)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
