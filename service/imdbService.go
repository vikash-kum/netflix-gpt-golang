package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ImdbService struct {
	TMDBToken string
}

func NewImdbService(token string) *ImdbService {
	return &ImdbService{TMDBToken: token}
}

func (s *ImdbService) NowPlayingTMDBApi(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIyNzkwN2ZkZDczOWM4OGVlYTVkOTMyM2U4MDJlYzIzNSIsInN1YiI6IjY1NjZkOGUyMTU2Y2M3MDE0ZTY2NjQzOSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.xqSzX87E-qH58F9bjogjfq-CNh4GBd23z8Vqv3ZCSFw")

	fmt.Print(req.Header)
	resp, err := client.Do(req)

	fmt.Println(err)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch data from TMDB API: " + string(body))
	}

	return string(body), nil
}
