package respect

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/hellcattc/respectng/internal/requests"
)

type RespectEngine struct {
	requestHandler    requests.RequestHandler
	respectController RespectController
}

func NewRespectEngine(handler requests.RequestHandler) RespectEngine {
	return RespectEngine{
		requestHandler:    handler,
		respectController: RespectController{},
	}
}

func (ng *RespectEngine) getAllClientEvents() (*requests.Events, error) {
	body, err := ng.requestHandler.ExecClientLiveRequest("eventdata")
	if err != nil {
		return nil, err
	} else {
		var events requests.Events
		json.Unmarshal(body, &events)
		return &events, nil
	}
}

func (ng *RespectEngine) getLiveClientEvents(ctx context.Context) (<-chan *requests.Event, error) {
	eventsChan := make(chan *requests.Event)
	var lastEventId int32 = 0
	// eventsSet := make(set.Set) скорее всего не нужно
	go func() {
		defer close(eventsChan)
		prevEvents, err := ng.getAllClientEvents()
		if err != nil {
			log.Println("Game hasn't started yet")
		} else {
			for _, event := range prevEvents.Events {
				eventCopy := event
				// eventsSet.Add(event.EventTime)
				eventsChan <- &eventCopy
			}
			lastEventId = prevEvents.Events[len(prevEvents.Events)-1].EventId
		}
		for {
			select {
			case <-ctx.Done():
				return
			default:
				events, err := ng.getAllClientEvents()
				if err != nil {
					log.Println("Couldn't execute long-polling request: ", err)
				} else if (len(events.Events) > 0) && ((events.Events[len(events.Events)-1]).EventId != lastEventId) {
					if err != nil {
						log.Println("Couldn't unmarshal events: ", err)
					}
					for _, event := range(events.Events[lastEventId+1:]) {
						eventCopy := event
						eventsChan <- &eventCopy
					}
					lastEventId = events.Events[len(events.Events)-1].EventId
					// if !(eventsSet.Has(lastevent.EventTime)) {
					// 	eventsSet.Add(lastevent.EventTime)
					// 	eventsChan <- &lastevent
					// }
				}
				time.Sleep(time.Second * 1)
			}
		}
	}()
	return eventsChan, nil
}

func (ng *RespectEngine) ListenForEvents() {
	ctx := context.TODO()
	subChan, err := ng.getLiveClientEvents(ctx)
	if err != nil {
		log.Println("Couldn't subscribe for events: ", err)
	}
	for val := range subChan {
		ng.respectController.HandleGameEvent(val)
	}
}
