package types

type Client struct {
	Nickname string
	MsgChan  chan *Msg
}
