package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var worktimeLengthFlag = flag.Int("t", 25, "Work time length in minutes")
var breakLengthFlag = flag.Int("b", 5, "Break period length in minutes")

const (
	padding      = 2
	maxWidth     = 80
	messageWidth = 15
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type tickMsg time.Time
type finishedMsg time.Time

type model struct {
	workDuration       int
	breakDuration      int
	startTime          time.Time
	formattedRemaining string
	currentTimerType   string
	percent            float64
	progress           progress.Model
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			clearScreen()
			return m, tea.Quit
		default:
			return m, nil
		}
	case tea.WindowSizeMsg:
		clearScreen()
		m.progress.Width = msg.Width - padding*2 - messageWidth
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil
	case tickMsg:
		var duration int
		if m.currentTimerType == "work" {
			duration = m.workDuration
		} else {
			duration = m.breakDuration
		}
		remaining := time.Duration(duration)*time.Second - time.Since(m.startTime)
		m.formattedRemaining = getTimeRepresentation(remaining)
		m.percent = float64(time.Since(m.startTime).Seconds()) / float64(duration)
		if m.percent > 1.0 {
			m.percent = 1.0
			m.formattedRemaining = "00:00:00"
			return m, finishedCmd
		}
		return m, tickCmd()

	case finishedMsg:
		if m.currentTimerType == "work" {
			go playBellSound()
			m.currentTimerType = "break"
			m.startTime = time.Now()
			remaining := time.Duration(m.breakDuration)*time.Second - time.Since(m.startTime)
			m.formattedRemaining = getTimeRepresentation(remaining)
			return m, tickCmd()
		} else {
			m.formattedRemaining = "00:00:00"
			playBellSound()
			clearScreen()
			return m, tea.Quit
		}

	default:
		return m, nil
	}
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.ViewAs(m.percent) + " " + m.formattedRemaining +
		" " + m.currentTimerType + "\n\n" +
		pad + helpStyle("Press Ctrl-C to quit")
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func finishedCmd() tea.Msg {
	return finishedMsg(time.Now())
}

func main() {
	flag.Parse()
	prog := progress.New(progress.WithGradient("#ED2938", "#00FF7F"))
	clearScreen()

	workDuration := time.Duration(*worktimeLengthFlag) * time.Minute
	breakDuration := time.Duration(*breakLengthFlag) * time.Minute
	if _, err := tea.NewProgram(model{progress: prog, workDuration: int(workDuration.Seconds()), breakDuration: int(breakDuration.Seconds()), startTime: time.Now(), currentTimerType: "work", formattedRemaining: getTimeRepresentation(workDuration)}).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}
