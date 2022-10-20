package utils_test

import (
	"testing"

	"github.com/GerardoHP/qmk_layout_helper_service/src/utils"
)

func TestGetKeyMissingReading(t *testing.T) {
	keyMapper := utils.NewFileKeyMapper([]utils.MappingType{})
	_, e := keyMapper.GetKey("")
	if e == nil {
		t.Error("expected error, but found nil")
	}
}
