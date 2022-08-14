package elrond_blockchain

import (
	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

const defaultPemFilePath = "../../../../dev-wallet-0.pem"

var metadata = state.Metadata{
	Properties: map[string]string{
		"pemFilePath": defaultPemFilePath,
	},
}

func TestStateStore_Init(t *testing.T) {
	t.Run("Init", func(t *testing.T) {
		as := NewAccountStorage(logger.NewLogger("testing"))
		err := as.Init(metadata)
		assert.Nil(t, err)
	})
}

// TODO This is an integration test
func TestStateStore_Get(t *testing.T) {
	t.Run("Get value", func(t *testing.T) {
		as := NewAccountStorage(logger.NewLogger("testing"))
		err := as.Init(metadata)
		assert.Nil(t, err)
		err = nil

		key := "foo2" // TODO
		resp, err := as.Get(&state.GetRequest{
			Key:      key,
			Metadata: nil,
			Options:  state.GetStateOption{},
		})
		assert.EqualValues(t, "bar", string(resp.Data))
		as.logger.Infof("Got value: %v for key: %v", string(resp.Data), key)
		assert.Nil(t, err)
	})
}
