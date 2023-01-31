package msg

type MessageSender interface {
	SendMessage(text string, userID int64) error
}


type Model struct{
	tgClient MessageSender
}

func NewModel(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text string
	UserID int64
}


func (m *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {
		m.tgClient.SendMessage("hello", msg.UserID)
		return nil
	}
	m.tgClient.SendMessage("не знаю эту команду", msg.UserID)
	return nil
}
