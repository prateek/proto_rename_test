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
