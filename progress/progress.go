package progress

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type ProgressBar struct {
	title     string
	progress  float64
	completed bool
	width     int
}

func NewProgressBar(title string) *ProgressBar {
	if title == "" {
		title = "Untitled"
	}

	width := 40
	out, err := exec.Command("tput", "cols").Output()
	if err == nil {
		width, err = strconv.Atoi(string(out))
		if err != nil {
			width = 40
		}
	}

	width -= 16

	return &ProgressBar{
		title:     title,
		progress:  0,
		completed: false,
		width:     width,
	}
}

func (p *ProgressBar) SetTitle(title string) {
	p.title = title
}

func (p *ProgressBar) AdvanceProgress(amount float64) {
	p.SetProgress(p.progress + amount)
}

func (p *ProgressBar) SetProgress(progress float64) {
	if p.completed {
		return
	}

	p.progress = progress
	if p.progress < 0 {
		p.progress = 0
	}

	if p.progress > 1 {
		p.progress = 1
	}

	len := int(float64(p.width) * p.progress)
	bar := strings.Repeat("▉", len)
	space := strings.Repeat(" ", p.width-len)
	fmt.Printf("\r%s [%s]", p.title, bar+space)

	if p.progress == 1 {
		p.completed = true
	}
}

func (p *ProgressBar) CompleteProgress() {
	p.SetProgress(1)
	println()
}
