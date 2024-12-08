package pterm_menu

import (
	"fmt"
	"github.com/pterm/pterm"
	"golang.org/x/term"
	"os"
	"strconv"
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

func PrintMenuNubmer(menuItems []MenuItem, mode_without_enter bool) int {

	for i, text := range MenuItemBuild(menuItems) {
		pterm.DefaultBasicText.Println(fmt.Sprintf("[%v] ", i) + text)
	}
	max_menu_item := len(menuItems)
	title := "Menu number: "

	if !mode_without_enter {
		for {

			var input string
			fmt.Scan(&input) // Считывает до первого пробела

			s := input[len(input)-1]

			num, err := strconv.Atoi(string(s))
			if err != nil {
				continue
			}

			if num >= max_menu_item || num < 0 {
				continue
			}
			return num
		}
	}

	fmt.Print(title)

	for {
		// Переводим терминал в сырой режим
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			continue
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState) // Возврат терминала в обычный режим

		// Считываем один символ
		var buf = make([]byte, 1)
		_, err = os.Stdin.Read(buf)
		if err != nil {
			continue
		}

		// Проверяем, что это цифра
		if buf[0] >= '0' && buf[0] <= '9' {
			s := string(buf[0])
			num, _ := strconv.Atoi(s)

			if num >= max_menu_item || num < 0 {
				continue
			}

			fmt.Println("")
			return num
		} else {
			continue
		}
	}

}
