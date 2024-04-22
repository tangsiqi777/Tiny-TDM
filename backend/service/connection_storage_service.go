package service

import (
	"context"
	"encoding/json"
	"github.com/labstack/gommon/log"
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

func (c *ConnectionStorageService) SaveConnection(newConnection ConnectionConfig) string {
	connections := c.ListConnection()
	maxId := 0
	name := newConnection.Name
	for i := 0; i < len(connections); i++ {
		conn := connections[i]
		if conn.Id > maxId {
			maxId = conn.Id
		}
		if conn.Name == name {
			log.Print("连接名称重复")
			return "连接名称重复"
		}
	}
	log.Print(maxId)
	newConnection.Id = maxId + 1
	connections = append(connections, newConnection)
	connectionsJson, err := json.Marshal(connections)
	if err != nil {
		return "发生未知错误"
	}
	err1 := c.LocalStorage.Store(connectionsJson)
	if err1 != nil {
		return "发生未知错误"
	}
	return "ok"
}

func (c *ConnectionStorageService) UpdateConnection(newConnection ConnectionConfig) string {
	connections := c.ListConnection()
	updateIndex := -1
	name := newConnection.Name
	for i := 0; i < len(connections); i++ {
		conn := connections[i]
		if conn.Id == newConnection.Id {
			updateIndex = i
		}
		if conn.Name == name && conn.Id != newConnection.Id {
			return "连接名称重复"
		}
	}
	if updateIndex == -1 {
		return "操作对象不存在请刷新"
	}
	connections[updateIndex] = newConnection
	connectionsJson, err := json.Marshal(connections)
	if err != nil {
		return "发生未知错误"
	}
	err1 := c.LocalStorage.Store(connectionsJson)
	if err1 != nil {
		return "发生未知错误"
	}
	return "ok"
}

func (c *ConnectionStorageService) DeleteConnection(id int) string {
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
		return "操作对象不存在请刷新"
	}
	connectionsJson, err := json.Marshal(newConnections)
	if err != nil {
		return "发生未知错误"
	}
	err1 := c.LocalStorage.Store(connectionsJson)
	if err1 != nil {
		return "发生未知错误"
	}
	return "ok"
}
