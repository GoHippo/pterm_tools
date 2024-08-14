package pterm_menu

import (
	"fmt"
	"github.com/GoHippo/slogpretty/sl"
	"github.com/pterm/pterm"
	"os"
)

type ReturnMenu struct {
	fMenu func(opt MenuOption)
	Opt   MenuOption
}

func NewReturnMenu(f func(opt MenuOption), opt MenuOption) ReturnMenu {
	return ReturnMenu{f, opt}
}

func (rm *ReturnMenu) Return(err error) {
	if err != nil {
		rm.Opt.Log.Error("", sl.Err(err))
		pterm.DefaultInteractiveTextInput.WithDefaultText(" - - - Press enter to back - - - ").WithDelimiter("").Show()
	}

	rm.Opt.ClearTerminal()
	rm.fMenu(rm.Opt)
}

// ====================== R ======================

func ExitEnter() {
	// Вывод сообщения и ожидание нажатия клавиши
	fmt.Println()
	pterm.DefaultInteractiveTextInput.WithDefaultText(" - - - Press enter to exit - - - ").WithDelimiter("").Show()

	// Завершение программы
	os.Exit(0)
}

func BackEnter() {
	fmt.Println("")
	pterm.DefaultInteractiveTextInput.WithDefaultText(" - - - Press enter to back - - - ").WithDelimiter("").Show()
}
