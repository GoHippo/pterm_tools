package pterm_input

import (
	"fmt"
	"github.com/pterm/pterm"
	"os"
	"strconv"
)

func ChangeIntBool(i int) int {
	if i == 0 {
		return 1
	} else {
		return 0
	}
}

func InputUserPath(title string) (string, error) {
	pterm.DefaultBasicText.Println("\033[H\033[J")

	path, err := pterm.DefaultInteractiveTextInput.Show(title)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	return path, nil
}

func InputUserString(title string) (string, error) {
	pterm.DefaultBasicText.Println("\033[H\033[J")

	path, err := pterm.DefaultInteractiveTextInput.Show(title)
	if err != nil {
		return "", err
	}

	return path, nil
}

func InpuntUserInt(title string) (int, error) {
	pterm.DefaultBasicText.Println("\033[H\033[J")

	result, _ := pterm.DefaultInteractiveTextInput.Show(title)
	i, err := strconv.Atoi(result)
	if err != nil {
		return 0, fmt.Errorf("failed to convert %s to int: %v", result, err)
	}

	if i < 1 {
		return 0, fmt.Errorf("The num cannot be less than 1")
	}
	return i, nil
}
