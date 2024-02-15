package apigateway

import (
	"encoding/json"
	"log"
	lyricsmicroservice "lyrics/Api"
	"net/http"
	"strconv"
)

type Track_ struct {
	Track Track
}

type Track struct {
	Track_ID    int    `json:"track_id"`
	Track_Name  string `json:"track_name"`
	Lyrics_Body string
	Album_Id    int
	Album_Name  string
	Artist_Id   int
	Artist_Name string
}

type TrackWithLyrics struct {
	Track_ID    int
	Track_Name  string
	Lyrics_Body string
	Album_Id    int
	Album_Name  string
	Artist_Id   int
	Artist_Name string
}

func GetTopTracksWithLyrics(location string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	// Step 1: Get top tracks
	url := "http://localhost:8080/top-tracks?location=" + location
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, " Something Went Wrong , Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	var tracks []Track_
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&tracks)
	if err != nil {
		return
	}

	// Step 2: Get lyrics for each track
	var tracksWithLyrics []TrackWithLyrics
	for _, track := range tracks {
		lyrics, err := lyricsmicroservice.GetLyrics(strconv.Itoa(track.Track.Track_ID))
		if err != nil {
			log.Printf("Failed to fetch lyrics for track %s: %v", track.Track.Track_Name, err)
			continue
		}
		tracksWithLyrics = append(tracksWithLyrics, TrackWithLyrics{Track_Name: track.Track.Track_Name,
			Lyrics_Body: lyrics, Track_ID: track.Track.Track_ID,
			Album_Id: track.Track.Album_Id, Album_Name: track.Track.Album_Name,
			Artist_Id: track.Track.Artist_Id, Artist_Name: track.Track.Artist_Name})
	}

	respjsonString, _ := json.Marshal(tracksWithLyrics)
	w.Write(respjsonString)
}

func GetTopTrackL(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method MISMATCH", http.StatusMethodNotAllowed)
		return
	}
	location := r.URL.Query().Get("location")
	if location != "" {
		GetTopTracksWithLyrics(location, w, r)
		return
	}

	http.Error(w, "Invalid Data , Error ", http.StatusUnprocessableEntity)
}
