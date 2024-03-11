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
)

func Table(headers ...string) *table.Table {
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		// HeaderStyle is the lipgloss style used for the table headers.
		HeaderStyle = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		// CellStyle is the base lipgloss style used for the table rows.
		CellStyle = re.NewStyle().Padding(0, 1)
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
	errorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true).Padding(1).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("9"))
	return errorStyle.Render("! " + msg + " !")
}
