package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hellcattc/respectng/internal/requests"
	"github.com/hellcattc/respectng/internal/respect"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
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
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	// defer cancel()
	engine := respect.NewRespectEngine(*handler)
	engine.ListenForEvents()
}
