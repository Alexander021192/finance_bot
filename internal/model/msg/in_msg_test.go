package model

import (
	"testing"

	mocks "github.com/Alexander021192/finance_bot/internal/mocks/msg"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestOnStartIncomingMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := NewModel(sender)

	sender.EXPECT().SendMessage("hello", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: 123,
	})

	assert.NoError(t, err)

}

type mockSender struct {
	text   string
	userID int64
}

func (m *mockSender) SendMessage(text string, userID int64) error {
	m.text = text
	m.userID = userID
	return nil
}
