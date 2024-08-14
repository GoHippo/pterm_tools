package pterm_pb

import (
	"fmt"
	"github.com/paulbellamy/ratecounter"
	"github.com/pterm/pterm"
	"time"
)

type ConsoleBar struct {
	spinner  *pterm.SpinnerPrinter
	counter  *ratecounter.RateCounter
	signal   chan int
	count    int
	isClosed bool
}

func NewSpinnerBar(text string) *ConsoleBar {
	counter := ratecounter.NewRateCounter(time.Second)

	fmt.Println("")
	spinner, _ := pterm.DefaultSpinner.Start(text)
	consoleBar := &ConsoleBar{
		spinner: spinner,
		counter: counter,
		signal:  make(chan int),
	}
	consoleBar.goLisner(text)
	return consoleBar
}

func (cb ConsoleBar) Add(i int) {
	cb.signal <- i
}

func (cb *ConsoleBar) Close(text_success string) {
	cb.signal <- -1
	cb.isClosed = true
	close(cb.signal)
	cb.spinner.Success(fmt.Sprintf("%v [%v]", text_success, cb.count))
}

func (cb *ConsoleBar) goLisner(text string) {
	go func() {
		for {
			i := <-cb.signal
			if i == -1 {
				return
			}
			cb.count++
			cb.counter.Incr(1)
		}
	}()

	go func() {
		for {
			if cb.isClosed {
				return
			}
			cb.spinner.UpdateText(fmt.Sprintf("%s [%v] [%v n/s]", text, cb.count, cb.counter.Rate())) // Обновляем текст спиннера для отображения активности
			time.Sleep(time.Second)
		}
	}()
}
