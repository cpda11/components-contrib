package elrond_blockchain

import (
	"errors"
	elrond "github.com/cpda11/dapr-elrond/pkg/elrond-blockchain"
	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"
)

type StateStore struct {
	state.DefaultBulkStore
	accountStorage *elrond.AccountStorage
	pemFilePath    string
	features       []state.Feature
	logger         logger.Logger
}

func (s *StateStore) Init(metadata state.Metadata) error {
	var (
		err error
	)
	if metadata.Properties["pemFilePath"] == "" {
		err = errors.New("metadata property pemFilePath required for initialization")
		s.logger.Errorf("Metadata: %v", metadata)
		return err
	}

	// Store for later use (store item)
	s.pemFilePath = metadata.Properties["pemFilePath"] // TODO Re-check if avoiding pem storage is a good solution
	s.accountStorage, err = elrond.InitAccountStorageWithPemFile(s.pemFilePath)
	if err != nil {
		return err
	}
	return nil
}

func (s *StateStore) Features() []state.Feature {
	return s.features
}

func (s *StateStore) Get(req *state.GetRequest) (*state.GetResponse, error) {
	var (
		res   = &state.GetResponse{}
		value string
		err   error
	)

	// https://docs.dapr.io/reference/api/state_api/#key-scheme
	// Request: &{Key:myapp||foo2 Metadata:map[] Options:{Consistency:}}
	s.logger.Infof("Request: %+v", req.Key)
	value, err = elrond.GetValue(s.accountStorage, req.Key, true)
	if err != nil {
		return nil, err
	}
	s.logger.Infof("Received value: %+v", value)
	res.Data = []byte(value)
	return res, nil
}

func (s *StateStore) Set(req *state.SetRequest) error {
	value, ok := req.Value.(string)
	if !ok {
		err := errors.New("currently only a string value is valid for elrond storage")
		s.logger.Debugf("Received request: %+v", req)
		return err
	}
	return elrond.StoreItemWithPemFile(s.accountStorage, s.pemFilePath, elrond.NewAccountStorageItem(req.GetKey(), value))
}

func (s *StateStore) Delete(req *state.DeleteRequest) error {
	return elrond.DeleteItemWithPemFile(s.accountStorage, s.pemFilePath, req.GetKey())
}

func NewAccountStorage(logger logger.Logger) *StateStore {
	s := &StateStore{
		features: []state.Feature{state.Eventual},
		logger:   logger,
	}
	s.DefaultBulkStore = state.NewDefaultBulkStore(s)
	return s
}
