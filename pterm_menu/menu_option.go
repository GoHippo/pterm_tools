package pterm_menu

import (
	"github.com/pterm/pterm"
	"gmail_token_recover/pkg/config"
	"log/slog"
)

type MenuOption struct {
	Log       *slog.Logger
	Cfg       *config.Config
	MenuStart func(opt MenuOption)
}

func (MenuOption) ClearTerminal() {
	pterm.DefaultBasicText.Println("\033[H\033[J")
}

func (mo MenuOption) BackMainMenu() {
	mo.MenuStart(mo)
}
