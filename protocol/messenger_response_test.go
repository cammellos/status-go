package protocol

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/status-im/status-go/protocol/common"
	"github.com/status-im/status-go/protocol/encryption/multidevice"
)

func TestMessengerResponseMergeChats(t *testing.T) {
	chat1 := &Chat{ID: "1"}
	modifiedChat1 := &Chat{ID: "1", Name: "name"}
	chat2 := &Chat{ID: "3"}
	response1 := &MessengerResponse{
		Chats: []*Chat{chat1},
	}

	response2 := &MessengerResponse{
		Chats: []*Chat{modifiedChat1, chat2},
	}

	require.NoError(t, response1.Merge(response2))

	require.Len(t, response1.Chats, 2)
	require.Equal(t, modifiedChat1, response1.Chats[0])
	require.Equal(t, chat2, response1.Chats[1])
}

func TestMessengerResponseMergeMessages(t *testing.T) {
	message1 := &common.Message{ID: "1"}
	modifiedMessage1 := &common.Message{ID: "1", From: "name"}
	message2 := &common.Message{ID: "3"}
	response1 := &MessengerResponse{
		Messages: []*common.Message{message1},
	}

	response2 := &MessengerResponse{
		Messages: []*common.Message{modifiedMessage1, message2},
	}

	require.NoError(t, response1.Merge(response2))

	require.Len(t, response1.Messages, 2)
	require.Equal(t, modifiedMessage1, response1.Messages[0])
	require.Equal(t, message2, response1.Messages[1])
}

func TestMessengerResponseMergeNotImplemented(t *testing.T) {
	response1 := &MessengerResponse{}

	response2 := &MessengerResponse{
		Contacts: []*Contact{&Contact{}},
	}
	require.Error(t, response1.Merge(response2))

	response2 = &MessengerResponse{
		Installations: []*multidevice.Installation{&multidevice.Installation{}},
	}
	require.Error(t, response1.Merge(response2))

	response2 = &MessengerResponse{
		EmojiReactions: []*EmojiReaction{&EmojiReaction{}},
	}
	require.Error(t, response1.Merge(response2))

	response2 = &MessengerResponse{
		Invitations: []*GroupChatInvitation{&GroupChatInvitation{}},
	}
	require.Error(t, response1.Merge(response2))

}
