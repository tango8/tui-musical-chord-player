package main

import (
	"fmt"
	"strings"
	"time"

	//"github.com/tango8/good-person-checker/message"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type unpressed int

type pressed int

type KeybLayout int

const (
	QWERTY KeybLayout = iota
	DVORAK
)

type choice struct {
	name    string
	pressed bool
}

type model struct {
	choices [7]choice // possibly make this up to 9/10 because there is numbers up to 9 and 0
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

var sustainTime time.Duration = 280 * time.Millisecond

var stateTimers [7]*time.Timer

func initialModel(mode Mode) model {
	var diatonic [7]string

	switch mode {
	default:
		diatonic = ionianDiatonic
	case Ionian:
		diatonic = ionianDiatonic
	case Aeolian:
		diatonic = aeolianDiatonic
	}

	var choices [7]choice
	for i, name := range diatonic {
		choices[i] = choice{
			name:    name,
			pressed: false,
		}
	}

	return model{
		choices: choices,
		layout:  QWERTY,
	}
}

func (m model) Init() tea.Cmd {
	loadSounds(filesToLoad)
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case unpressed:
		m.choices[msg].pressed = false
	case pressed:
		m.choices[msg].pressed = true

	case tea.KeyPressMsg:
		switch msg.String() {
		/*case "m":
		activeMod	*/
		case "M":
		case "ctrl+c":
			return m, tea.Quit
		case "enter", "space":
			//playChord(getChord(Major, []Extension{MajorSeventh})
		case "1":
			return m, func() tea.Msg {
				playSound(1)
				playSound(5)
				playSound(8)
				playSound(12)
				return pressed(0)
			}
		case "2":
			return m, func() tea.Msg {
				playSound(3)
				playSound(6)
				playSound(10)
				//playSound(1)
				playSound(13)
				return pressed(1)
			}
		case "3":
			return m, func() tea.Msg {
				playSound(5)
				playSound(8)
				playSound(12)
				//playSound(3)
				playSound(15)
				return pressed(2)
			}
		case "4":
			return m, func() tea.Msg {
				playSound(6)
				playSound(10)
				//playSound(1)
				//playSound(5)
				playSound(13)
				playSound(17)
				return pressed(3)
			}
		case "5":
			return m, func() tea.Msg {
				playSound(8)
				playSound(12)
				//playSound(3)
				//playSound(7)
				playSound(15)
				playSound(19)
				return pressed(4)
			}
		case "6":
			return m, func() tea.Msg {
				playSound(10)
				//playSound(1)
				//playSound(5)
				//playSound(8)
				playSound(13)
				playSound(17)
				playSound(20)
				return pressed(5)
			}
		case "7":
			return m, func() tea.Msg {
				playSound(12)
				//playSound(3)
				//playSound(6)
				//playSound(9)
				playSound(15)
				playSound(18)
				playSound(22)
				return pressed(6)
			}
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
	s.WriteString(getCurrentKey() + "\n\n")

	for i, choice := range m.choices {

		keyboardKey := i + 1

		if choice.pressed {
			s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, pressedModelStyle.Render(choice.name)))
		} else {
			s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, modelStyle.Render(choice.name)))
		}

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
