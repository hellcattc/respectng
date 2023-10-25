package requests

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	// "log"
	"net/http"
)

type RequestHandler struct {
	RIOT_API_KEY  string
	networkClient http.Client
	localClient   http.Client
}

func NewRequestHandler(apiKey string) *RequestHandler {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return &RequestHandler{
		RIOT_API_KEY:  apiKey,
		networkClient: http.Client{},
		localClient:   *client,
	}
}

func (s *RequestHandler) execRequest(req *http.Request) ([]byte, error) {
	resp, err := s.networkClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s *RequestHandler) ExecClientLiveRequest(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("https://127.0.0.1:2999/liveclientdata/%s", endpoint)
	// log.Println(url)
	resp, err := s.localClient.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (s *RequestHandler) makeRequest(method string, endpoint string) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Riot-Token", s.RIOT_API_KEY)
	return req, nil
}

func (s *RequestHandler) GetSummonerByName(name string) (*SummonerById, error) {
	req, err := s.makeRequest("GET", fmt.Sprintf("https://ru.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", name))
	if err != nil {
		return nil, err
	}
	resp, err := s.execRequest(req)
	if err != nil {
		return nil, err
	}
	var summoner SummonerById
	json.Unmarshal(resp, &summoner)
	return &summoner, nil
}
