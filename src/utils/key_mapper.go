package utils

import "github.com/GerardoHP/qmk_layout_helper_service/src/models"

type KeyMapper interface {
	Read() error
	GetKey(k string) (*models.KeyMap, error)
	GetAllKeys() map[string]models.KeyMap
}
