package repository

import "github.com/KumKeeHyun/PDK/logic-core/domain/model"

type MetaRepo interface {
	GetNode(key string) (*model.Node, error)
	GetSensor(key string) (*model.Sensor, error)
	NewNode(key string, n *model.Node) error
	NewSensor(key string, s *model.Sensor) error
	DelNode(key string) error
	DelSensor(key string) error
}
