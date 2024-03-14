package components

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
	purple    = lipgloss.Color("99")
	gray      = lipgloss.Color("245")
	lightGray = lipgloss.Color("241")
	red       = lipgloss.Color("9")
	green     = lipgloss.Color("10")
)

func Table(headers ...string) *table.Table {
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		// HeaderStyle is the lipgloss style used for the table headers.
		HeaderStyle = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		// CellStyle is the base lipgloss style used for the table rows.
		CellStyle = re.NewStyle().Padding(0, 2)
		// OddRowStyle is the lipgloss style used for odd-numbered table rows.
		OddRowStyle = CellStyle.Copy().Foreground(gray)
		// EvenRowStyle is the lipgloss style used for even-numbered table rows.
		EvenRowStyle = CellStyle.Copy().Foreground(lightGray)
		// BorderStyle is the lipgloss style used for the table border.
		BorderStyle = lipgloss.NewStyle().Foreground(purple)
	)
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, _ int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		}).
		Headers(headers...).
		Rows([][]string{}...)

	return t
}

func ErrorBox(msg string) string {
	errorStyle := lipgloss.NewStyle().Foreground(red).Bold(true).Padding(1).Border(lipgloss.RoundedBorder()).BorderForeground(red)
	return errorStyle.Render("🚨 " + msg + " 🚨")
}

func SuccessBox(msg string) string {
	successStyle := lipgloss.NewStyle().Foreground(green).Bold(true).Padding(1).Border(lipgloss.RoundedBorder()).BorderForeground(green)
	return successStyle.Render("✅ " + msg + " ✅")
}

func Box(msg string) string {
	boxStyle := lipgloss.NewStyle().Padding(1).Border(lipgloss.RoundedBorder()).BorderForeground(purple)
	return boxStyle.Render(msg)
}
