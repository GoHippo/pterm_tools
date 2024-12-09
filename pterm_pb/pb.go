package pterm_pb

import (
	"github.com/pterm/pterm"
	"sync"
)

type ProgressBar struct {
	pb   *pterm.ProgressbarPrinter
	lock *sync.Mutex
}

func NewPB(total int, title string) ProgressBar {
	pbLines, _ := pterm.DefaultProgressbar.WithTotal(total).Start(title)
	return ProgressBar{pb: pbLines, lock: &sync.Mutex{}}
}

func (p *ProgressBar) Add(i int) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.pb.Add(i)
}

func (p *ProgressBar) IsNil() bool {
	return p.pb == nil
}

func (p *ProgressBar) Close() {
	p.Add(p.pb.Total - p.pb.Current)
	p.pb.Stop()
}

// keka
