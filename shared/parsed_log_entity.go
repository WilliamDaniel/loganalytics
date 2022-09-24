package shared

type ParsedLog struct {
	AuthenticatedEntity RequestAuthenticatedEntity `json:"authenticated_entity"`
	Service             RequestService             `json:"service"`
	Latencies           RequestLatencies           `json:"latencies"`
}

type AuthenticatedEntityConsumerID struct {
	UUID string `json:"uuid"`
}
type RequestAuthenticatedEntity struct {
	ConsumerID AuthenticatedEntityConsumerID `json:"consumer_id"`
}

type RequestService struct {
	ID string `json:"id"`
}

type RequestLatencies struct {
	Proxy   int64 `json:"proxy"`
	Gateway int64 `json:"gateway"`
	Request int64 `json:"request"`
}
