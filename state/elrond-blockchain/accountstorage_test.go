package elrond_blockchain

import (
	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStateStore_Init(t *testing.T) {
	m := state.Metadata{}
	t.Run("Init", func(t *testing.T) {
		as := &StateStore{
			DefaultBulkStore: state.DefaultBulkStore{},
			accountStorage:   nil,
			features:         nil,
			logger:           logger.NewLogger("testing"),
		}
		err := as.Init(m)
		assert.Nil(t, err)
	})
}

func TestStateStore_Get(t *testing.T) {
	m := state.Metadata{}
	t.Run("Get value", func(t *testing.T) {
		as := &StateStore{
			DefaultBulkStore: state.DefaultBulkStore{},
			accountStorage:   nil,
			features:         nil,
			logger:           logger.NewLogger("testing"),
		}
		err := as.Init(m)
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
