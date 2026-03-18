package main

import (
	"fmt"

	//"github.com/tango8/good-person-checker/message"

	tea "charm.land/bubbletea/v2"
)

type KeybLayout int

const (
	QWERTY KeybLayout = iota
	DVORAK
)

type Mode int

const (
	Ionian  = iota // major
	Aeolian        // natural minor
)

type model struct {
	choices [7]string // possibly make this up to 9/10 because there is numbers up to 9 and 0
	cursor  int
	//selected map[int]struct{}
	layout KeybLayout
}

func initialModel(mode Mode) model {
	var diatonic [7]string

	if mode == Ionian {
		diatonic = [7]string{"I", "ii", "iii", "IV", "V", "vi", "vii°"}
	}

	return model{
		choices: diatonic,
		//selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter", "space":
			/*_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}*/
			playSound("soundfiles/440hz-sin.mp3")
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
	s := "go chord player\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		keyboardKey := i + 1
		/*if _, err := m.selected[i]; err {
			checked = "x"
		}*/

		s += fmt.Sprintf("%s %d %s\n", cursor, keyboardKey, choice)
	}

	s += "\nPress ctrl c to quit\n"

	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel(Ionian))
	if _, err := p.Run(); err != nil {
		fmt.Println("error " + err.Error())
	}
}
