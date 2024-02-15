package lyricsmicroservice

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type LyricsResponse struct {
	Message LyricsMessage `json:"message"`
}

type LyricsMessage struct {
	Body LyricsBody `json:"body"`
}

type LyricsBody struct {
	Lyrics Lyrics `json:"lyrics"`
}

type Lyrics struct {
	Lyrics_Body string `json:"lyrics_body"`
}

func GetLyrics(trackID string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return "", err
	}
	musixmatchAPIBaseURL := os.Getenv("MUSIXMATCHAPIBASEURL")
	apiKey := os.Getenv("MUMATCH_API_KEY")

	url := musixmatchAPIBaseURL + "track.lyrics.get?track_id=" + trackID + "&apikey=" + apiKey
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var lyricsResponse LyricsResponse
	err = json.NewDecoder(resp.Body).Decode(&lyricsResponse)
	if err != nil {
		return "", err
	}
	return lyricsResponse.Message.Body.Lyrics.Lyrics_Body, nil
}
