package types

type NewClient struct {
	ClientNick string
	MsgChan    chan *Msg
}
