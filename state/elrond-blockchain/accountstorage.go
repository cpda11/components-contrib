package elrond_blockchain

import (
	elrond "github.com/cpda11/dapr-elrond/pkg/elrond-blockchain"
	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"
)

type StateStore struct {
	state.DefaultBulkStore
	accountStorage *elrond.AccountStorage
	pemFilePath    string // TODO
	features       []state.Feature
	logger         logger.Logger
}

func (s *StateStore) Init(metadata state.Metadata) error {
	var (
		err error
	)
	s.logger.Infof("Metadata: %v", metadata)

	s.pemFilePath = "../../../../dev-wallet-0.pem" // TODO Config
	s.accountStorage, err = elrond.InitAccountStorage(s.pemFilePath)
	if err != nil {
		return err
	}
	return nil
}

func (s *StateStore) Get(req *state.GetRequest) (*state.GetResponse, error) {
	var (
		res   = &state.GetResponse{}
		value string
		err   error
	)

	value, err = elrond.GetValue(s.accountStorage, req.Key, true) // TODO Check req options
	if err != nil {
		return nil, err
	}
	res.Data = []byte(value)
	return res, nil
}
