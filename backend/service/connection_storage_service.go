package service

import (
	"context"
	"encoding/json"
	"sync"
	"tinytdm/backend/storage"
)

type ConnectionStorageService struct {
	LocalStorage *storage.LocalStorage
	ctx          context.Context
	//Mutex        sync.Mutex
}

type ConnectionConfig struct {
	Id       int    `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Addr     string `json:"addr,omitempty" yaml:"addr,omitempty"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

func (c *ConnectionStorageService) Startup(ctx context.Context) {
	c.ctx = ctx
}

var connectionStorageService *ConnectionStorageService
var onceConnectionStorageService sync.Once

func GetOneConnectionStorageService() *ConnectionStorageService {
	if connectionStorageService == nil {
		onceConnectionStorageService.Do(func() {
			connectionStorageService = &ConnectionStorageService{
				LocalStorage: storage.NewLocalStore("connection.txt"),
			}
		})
	}
	return connectionStorageService
}

// ListConnection 显示所有连接
func (c *ConnectionStorageService) ListConnection() []ConnectionConfig {
	b, err := c.LocalStorage.Load()
	if err != nil {
		return []ConnectionConfig{}
	}
	var connections []ConnectionConfig
	err1 := json.Unmarshal(b, &connections)
	if err1 != nil {
		return []ConnectionConfig{}
	}
	return connections
}

func (c *ConnectionStorageService) SaveConnection(newConnection ConnectionConfig) int {
	connections := c.ListConnection()
	maxId := 0
	for i := 0; i < len(connections); i++ {
		conn := connections[i]
		if conn.Id > maxId {
			maxId = conn.Id
		}
	}
	newConnection.Id = maxId
	connections = append(connections, newConnection)
	connectionsJson, err := json.Marshal(connections)
	if err != nil {
		return 0
	}
	err1 := c.LocalStorage.Store(connectionsJson)
	if err1 != nil {
		return 0
	}
	return 1
}

func (c *ConnectionStorageService) UpdateConnection(newConnection ConnectionConfig) int {
	connections := c.ListConnection()
	updateIndex := -1
	for i := 0; i < len(connections); i++ {
		conn := connections[i]
		if conn.Id == newConnection.Id {
			updateIndex = i
			break
		}
	}
	if updateIndex == -1 {
		return 0
	}
	connections[updateIndex] = newConnection
	connectionsJson, err := json.Marshal(connections)
	if err != nil {
		return 0
	}
	err1 := c.LocalStorage.Store(connectionsJson)
	if err1 != nil {
		return 0
	}
	return 1
}

func (c *ConnectionStorageService) DeleteConnection(id int) int {
	connections := c.ListConnection()
	newConnections := make([]ConnectionConfig, len(connections))
	deleteIndex := -1
	for i := 0; i < len(connections); i++ {
		conn := connections[i]
		if conn.Id == id {
			deleteIndex = i
		} else {
			newConnections = append(newConnections, conn)
		}
	}
	if deleteIndex == -1 {
		return 0
	}
	connectionsJson, err := json.Marshal(newConnections)
	if err != nil {
		return 0
	}
	err1 := c.LocalStorage.Store(connectionsJson)
	if err1 != nil {
		return 0
	}
	return 1
}
