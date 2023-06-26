package requests

type SummonerById struct {
	Id            string
	AccountId     string
	Puuid         string
	Name          string
	ProfileIconId int64
	// RevisionDatetime uint64
	SummonerLevel int32
}

type Event struct {
	EventId int32
	EventName string
	EventTime float64
	Assisters []string
	KillerName string
	VictimName string
}

type Events struct {
	Events []Event
}