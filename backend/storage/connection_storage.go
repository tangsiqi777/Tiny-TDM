package storage

import (
	"encoding/json"
	"golang.org/x/net/context"
	"sync"
)

type ConnectionStorage struct {
	ctx     context.Context
	storage *LocalStorage
	//mutex   sync.Mutex
}

type ConnectionConfig struct {
	Id       int    `json:"id" yaml:"name"`
	Name     string `json:"name" yaml:"name"`
	Addr     string `json:"addr,omitempty" yaml:"addr,omitempty"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

/*var connectionList []ConnectionConfig
var onceConnection sync.Once

func ConnectionList() []ConnectionConfig {
	if connectionList == nil {
		onceConnection.Do(func() {
			var tmp []ConnectionConfig
			connectionList = tmp
		})
	}
	return connectionList
}*/

var connectionStorage *ConnectionStorage
var onceConnection sync.Once

func GetConnectionStorage() *ConnectionStorage {
	if connectionStorage == nil {
		onceConnection.Do(func() {
			connectionStorage = &ConnectionStorage{
				storage: NewLocalStore("tiny-tdm-connections"),
			}

		})
	}
	return connectionStorage
}

func (c *ConnectionStorage) Start(ctx context.Context) {
	c.ctx = ctx
}

func (c *ConnectionStorage) listConnection() []ConnectionConfig {
	b, err := c.storage.Load()
	if err != nil {
		return make([]ConnectionConfig, 0)
	}
	var configs []ConnectionConfig
	err1 := json.Unmarshal(b, &configs)
	if err1 != nil {
		return make([]ConnectionConfig, 0)
	}
	return configs
}

func (c *ConnectionStorage) saveConnection(connectionList []ConnectionConfig) error {
	b, err := json.Marshal(connectionList)
	if err != nil {
		return err
	}
	if err = c.storage.Store(b); err != nil {
		return err
	}
	return nil
}
