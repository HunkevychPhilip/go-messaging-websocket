package types

type ChatChannels struct {
	ClientRequests    chan *Client
	ClientDisconnects chan string
	Messages          chan *Msg
}
