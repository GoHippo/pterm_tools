package pterm_menu

import (
	"github.com/pterm/pterm"

	"log/slog"
)

type MenuOption struct {
	Log       *slog.Logger
	Cfg       any
	MenuStart func(opt MenuOption)
}

func (MenuOption) ClearTerminal() {
	pterm.DefaultBasicText.Println("\033[H\033[J")
}

func (mo MenuOption) BackMainMenu() {
	mo.MenuStart(mo)
}
