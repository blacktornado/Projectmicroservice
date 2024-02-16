package trackmicroservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type TopTracksResponse struct {
	Message TopTracksMessage `json:"message"`
}

type Header_detail struct {
	Status_Code  int
	Execute_Time float64
}

type TopTracksMessage struct {
	Body TopTracksBody `json:"body"`
}

type TopTracksBody struct {
	Track_List []Track_ `json:"track_list"`
}

type Track_ struct {
	Track Track
}

type Track struct {
	Track_ID    int
	Track_Name  string
	Album_Id    int
	Album_Name  string
	Artist_Id   int
	Artist_Name string
}

func GetTopTracks(location string) ([]Track_, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	musixmatchAPIBaseURL := os.Getenv("MUSIXMATCHAPIBASEURL")
	apiKey := os.Getenv("MUMATCH_API_KEY")
	fmt.Println(location)
	url := musixmatchAPIBaseURL + "chart.tracks.get?country=" + location + "&has_lyrics=1&chart_name=top&page=1&page_size=2&apikey=" + apiKey
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var trackresponse TopTracksResponse

	err = decoder.Decode(&trackresponse)
	if err != nil {
		return nil, err
	}
	fmt.Println(trackresponse)
	return trackresponse.Message.Body.Track_List, nil
}
