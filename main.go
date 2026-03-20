package main

import (
	"fmt"
	"strings"
	//"time"

	//"github.com/tango8/good-person-checker/message"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type pressedState uint

const (
	one pressedState = iota
	two
	three
	four
	five
	six
	seven
)

type KeybLayout int

const (
	QWERTY KeybLayout = iota
	DVORAK
)

type model struct {
	state   pressedState
	choices [7]string // possibly make this up to 9/10 because there is numbers up to 9 and 0
	cursor  int
	//selected map[int]struct{}
	layout KeybLayout
}

var filesToLoad [24]string = [24]string{
	"soundfiles/defaultpiano/C4.mp3",
	"soundfiles/defaultpiano/C#4.mp3",
	"soundfiles/defaultpiano/D4.mp3",
	"soundfiles/defaultpiano/D#4.mp3",
	"soundfiles/defaultpiano/E4.mp3",
	"soundfiles/defaultpiano/F4.mp3",
	"soundfiles/defaultpiano/F#4.mp3",
	"soundfiles/defaultpiano/G4.mp3",
	"soundfiles/defaultpiano/G#4.mp3",
	"soundfiles/defaultpiano/A4.mp3",
	"soundfiles/defaultpiano/A#4.mp3",
	"soundfiles/defaultpiano/B4.mp3",
	"soundfiles/defaultpiano/C5.mp3",
	"soundfiles/defaultpiano/C#5.mp3",
	"soundfiles/defaultpiano/D5.mp3",
	"soundfiles/defaultpiano/D#5.mp3",
	"soundfiles/defaultpiano/E5.mp3",
	"soundfiles/defaultpiano/F5.mp3",
	"soundfiles/defaultpiano/F#5.mp3",
	"soundfiles/defaultpiano/G5.mp3",
	"soundfiles/defaultpiano/G#5.mp3",
	"soundfiles/defaultpiano/A5.mp3",
	"soundfiles/defaultpiano/A#5.mp3",
	"soundfiles/defaultpiano/B5.mp3",
}

func initialModel(mode Mode) model {
	var diatonic [7]string

	switch mode {
	case Ionian:
		diatonic = ionianDiatonic
	case Aeolian:
		diatonic = aeolianDiatonic
	}

	return model{
		choices: diatonic,
	}
}

func (m model) Init() tea.Cmd {
	loadSounds(filesToLoad)
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter", "space":
			playSound(1)
			playSound(5)
			playSound(8)
			playSound(12)
		case "1":
			playSound(1)
			playSound(5)
			playSound(8)
			playSound(12)
			m.state = one
		case "2":
			playSound(3)
			playSound(6)
			playSound(10)
			//playSound(1)
			playSound(13)
			m.state = two
		case "3":
			playSound(5)
			playSound(8)
			playSound(12)
			//playSound(3)
			playSound(15)
			m.state = three
		case "4":
			playSound(6)
			playSound(10)
			playSound(1)
			//playSound(5)
			playSound(17)
			m.state = four
		case "5":
			playSound(8)
			playSound(12)
			//playSound(3)
			//playSound(7)
			playSound(15)
			playSound(19)
			m.state = five
		case "6":
			playSound(10)
			//playSound(1)
			//playSound(5)
			//playSound(8)
			playSound(13)
			playSound(17)
			playSound(20)
			m.state = six
		case "7":
			playSound(12)
			//playSound(3)
			//playSound(6)
			//playSound(9)
			playSound(15)
			playSound(18)
			playSound(21)
			m.state = seven
		}

		if m.layout == QWERTY {
			switch msg.String() {
			case "left", "h":
				if m.cursor > 0 {
					m.cursor--
				}

			case "right", "l":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}

			}
		}

		if m.layout == DVORAK {
			switch msg.String() {
			case "left", "d":
				if m.cursor > 0 {
					m.cursor--
				}

			case "right", "n":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	// Header
	var s strings.Builder
	//s = "go chord player\n\n"

	for i, choice := range m.choices {

		keyboardKey := i + 1

		switch m.state {
		case one:
			s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, pressedModelStyle.Render(fmt.Sprint(choice))))
			continue
		case two:
			s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, pressedModelStyle.Render(fmt.Sprint(choice))))
			continue
		case three:
			s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, pressedModelStyle.Render(fmt.Sprint(choice))))
			continue
		case four:
			s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, pressedModelStyle.Render(fmt.Sprint(choice))))
			continue
		case five:
			s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, pressedModelStyle.Render(fmt.Sprint(choice))))
			continue
		case six:
			s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, pressedModelStyle.Render(fmt.Sprint(choice))))
			continue
		case seven:
			s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, pressedModelStyle.Render(fmt.Sprint(choice))))
			continue
		}

		s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, modelStyle.Render(fmt.Sprint(choice))))

		s.WriteString(helpStyle.Render(fmt.Sprint(keyboardKey)))
	}

	s.WriteString("\nPress ctrl c to quit\n")

	view := tea.NewView(s.String())
	view.AltScreen = true

	return view
}

func main() {
	p := tea.NewProgram(initialModel(Ionian))
	if _, err := p.Run(); err != nil {
		fmt.Println("error " + err.Error())
	}
}
