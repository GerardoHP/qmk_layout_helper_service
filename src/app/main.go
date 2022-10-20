package main

import (
	"fmt"

	"github.com/GerardoHP/qmk_layout_helper_service/src/utils"
)

func main() {
	var km utils.KeyMapper = utils.NewFileKeyMapper([]utils.MappingType{})
	km.Read()
	fmt.Println(km.GetAllKeys())
}
