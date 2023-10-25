package eventstream

import "log"

func ConsumeChannel[K any](channel <-chan K, handler func(K)) {
	log.Println("consuming")
	go func() {
		log.Println("consuming")
		for val := range channel {
			handler(val)
		}
	}()
}
