package proto_rename

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/protobuf/proto"
	newproto "github.com/prateek/proto_rename_test/newproto"
	oldproto "github.com/prateek/proto_rename_test/oldproto"
)

func TestRenameCompatibility(t *testing.T) {
	old := oldproto.Something{
		FieldX: 123,
		FieldY: 456,
		FieldZ: 789,
	}

	oldBytes, err := proto.Marshal(&old)
	require.NoError(t, err)

	var new newproto.Something
	require.NoError(t, proto.Unmarshal(oldBytes, &new))

	require.Equal(t, old.FieldX, new.FieldX)
	require.Equal(t, old.FieldY, new.FieldYY)
	require.Equal(t, old.FieldZ, new.FieldZ)
}

func TestBytesStringCompatiblility(t *testing.T) {
	/*
		This test generates a message of with String fields, encodes it and then decodes it
		into a message which uses []byte instead.

		message StringBackedMessage {
			string fieldX  = 1;
			string fieldYY = 2;
			string fieldZ  = 3;
		}

		message BytesBackedMessage {
			bytes fieldX  = 1;
			bytes fieldYY = 2;
			bytes fieldZ  = 3;
		}
	*/
	stringMessage := newproto.StringBackedMessage{
		FieldX:  string("abc"),
		FieldYY: string("def"),
		FieldZ:  string("some unicode chars 日本語"),
	}

	stringMessageBytes, err := proto.Marshal(&stringMessage)
	require.NoError(t, err)

	var bytesMessage newproto.BytesBackedMessage
	require.NoError(t, proto.Unmarshal(stringMessageBytes, &bytesMessage))

	require.Equal(t, stringMessage.FieldX, string(bytesMessage.FieldX))
	require.Equal(t, stringMessage.FieldYY, string(bytesMessage.FieldYY))
	require.Equal(t, stringMessage.FieldZ, string(bytesMessage.FieldZ))

	require.Equal(t, []byte(stringMessage.FieldX), bytesMessage.FieldX)
	require.Equal(t, []byte(stringMessage.FieldYY), bytesMessage.FieldYY)
	require.Equal(t, []byte(stringMessage.FieldZ), bytesMessage.FieldZ)
}
