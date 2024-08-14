package pterm_menu

import (
	"fmt"
	"github.com/pterm/pterm"
	"strings"
)

type MenuItem struct {
	Description string
	Value       string
}

func (m MenuItem) SetVal(s string) MenuItem {
	m.Value = s
	return m
}

// Функция для построения элементов меню с 5 пробелами после самой длинной строки описания
func MenuItemBuild(menuItems []MenuItem) []string {
	// Определение максимальной длины описания без учета управляющих последовательностей для цвета
	maxDescLength := 0
	for _, item := range menuItems {
		cleanDesc := pterm.RemoveColorFromString(item.Description)
		if len(cleanDesc) > maxDescLength {
			maxDescLength = len(cleanDesc)
		}
	}

	// Добавление 5 пробелов к максимальной длине описания
	space := 5
	paddedLength := maxDescLength + space

	options := make([]string, len(menuItems))
	for i, item := range menuItems {
		cleanDesc := pterm.RemoveColorFromString(item.Description)
		padding := strings.Repeat(" ", paddedLength-len(cleanDesc))
		options[i] = fmt.Sprintf("%s%s%s", item.Description, padding, item.Value)
	}
	return options
}
