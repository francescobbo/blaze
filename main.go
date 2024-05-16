package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"golang.design/x/clipboard"

	"github.com/francescobbo/blaze/calculator"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error

	calculatorResult calculator.Value
	calculatorError  error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Calculate anything..."
	ti.Focus()
	ti.CharLimit = 156

	return model{
		textInput:        ti,
		err:              nil,
		calculatorResult: calculator.Value{},
		calculatorError:  errors.New(""),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			clipboard.Write(clipboard.FmtText, []byte(formatResult(m.calculatorResult)))
			time.Sleep(100 * time.Millisecond)
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)

	res, err := calculator.Evaluate(m.textInput.Value())
	m.calculatorResult = res
	m.calculatorError = err

	if math.IsNaN(res.Number) {
		m.calculatorError = errors.New("not a number")
	}

	return m, cmd
}

func formatResult(value calculator.Value) string {
	// Fix negative zero
	if value.Number == 1./math.Inf(-1) {
		value.Number = 0
	}

	var numberStr string

	if value.Number == math.Inf(1) {
		numberStr = "∞"
	} else if value.Number == math.Inf(-1) {
		numberStr = "-∞"
	} else {
		// Convert the number to a string with up to 6 decimal places, without trailing zeros
		numberStr = strconv.FormatFloat(value.Number, 'f', 6, 64)
		numberStr = strings.TrimRight(numberStr, "0")
		numberStr = strings.TrimRight(numberStr, ".")
	}

	if value.Unit != "" {
		return fmt.Sprintf("%s %s", numberStr, value.Unit)
	}

	return numberStr
}

func (m model) View() string {
	v := m.textInput.View()

	if m.calculatorError == nil {
		return v + "\n" + formatResult(m.calculatorResult) + "\n\nPress Enter to copy the result to the clipboard."
	}

	return v
}
