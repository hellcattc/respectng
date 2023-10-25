package respect

import (
	"encoding/json"
	"log"

	"github.com/hellcattc/respectng/internal/requests"
)

type RespectController struct {
}

func NewRespectCalc() RespectController {
	return RespectController{}
}

func (rc *RespectController) HandleGameEvent(e *requests.Event) {
	event, _ := json.MarshalIndent(e, "", "   ")
	log.Println(string(event))
}
