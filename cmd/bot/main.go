package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hellcattc/respectengine/internal/requests"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Couldn't load API Key")
	}
	apiKey := os.Getenv("RIOT_API_KEY")
	handler := requests.NewRequestHandler(apiKey)
	resp, err := handler.GetSummonerByName("DarkReaper13")
	if err != nil {
		log.Fatalf("Couldn't execute request: %v", err)
	}
	form, _ := json.MarshalIndent(resp, "", "   ")
	log.Println(string(form))
	resp1, err1 := handler.GetLiveClientEvents()
	form, _ = json.MarshalIndent(resp1, "", "    ")
	if err1 != nil {
		log.Fatalf("Couldn't retrieve events: %v", err)
	}
	log.Println(json.MarshalIndent(form, "", "   "))
}