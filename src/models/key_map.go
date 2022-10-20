package models

import "strings"

type KeyMap struct {
	Key     string
	Value   string
	Aliases []string
}

func (k KeyMap) HasAliases() (bool, string) {
	l := len(k.Aliases)
	switch l {
	case 0:
		return false, ""
	case 1:
		return true, k.Aliases[0]
	default:
		return true, greatestString(k.Aliases)
	}
}

func NewKeyMapFromFileLine(str string) *KeyMap {
	attributes := strings.Split(str, ",")
	aliases := strings.Split(attributes[1], "|")
	for i := range aliases {
		aliases[i] = strings.Trim(aliases[i], " ")
	}

	return &KeyMap{Key: strings.Trim(attributes[0], " "), Aliases: aliases, Value: strings.Trim(attributes[2], " ")}
}

func greatestString(aliases []string) string {
	m := ""
	for _, v := range aliases {
		if len(v) > len(m) {
			m = v
		}
	}

	return m
}
