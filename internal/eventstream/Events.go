package eventstream;

func ConsumeChannel[K any](channel chan K, handler func(K)) {
	go func() {
		for val := range channel {
			handler(val)
		}
	}()
}