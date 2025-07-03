package formatter

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/georgedolee/amindi/internal/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func PrintForecast(forecast *model.WeatherForecast, unit string, days int) error {
	printCurrentForecast(forecast, unit)

	if len(forecast.Forecast.ForecastDay) > 0 {
		fmt.Println()
		printDailyForecast(forecast.Forecast.ForecastDay, unit, days)
	}

	return nil
}

func printCurrentForecast(forecast *model.WeatherForecast, unit string) {
	t := createTable()
	setBorderOptions(t, false, false, false)

	city := color.New(color.FgWhite, color.Bold).Sprint(forecast.Location.Name)
	country := color.New(color.FgHiWhite).Sprint(forecast.Location.Country)
	condition := color.New(color.FgHiWhite).Sprint(forecast.Current.Condition.Text)

	temp, unitSymbol := formatTemp(forecast.Current.TempC, forecast.Current.TempF, unit)
	coloredTemp := colorizeTemp(temp, unitSymbol)

	icon := symbolFor(forecast.Current.Condition.Code)

	t.AppendRows([]table.Row{
		{city, "", coloredTemp},
		{"", icon, ""},
		{country, "", condition},
	})

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignRight},
		{Number: 2, Align: text.AlignCenter},
		{Number: 3, Align: text.AlignLeft},
	})

	t.Render()
}

func printDailyForecast(days []model.ForecastDay, unit string, daysN int) {
	t := createTable()
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true

	title := color.New(color.FgWhite, color.Italic).Sprintf("Next %d Day Forecast", daysN)
	fmt.Printf("%s\n", title)

	for _, day := range days {
		date := time.Unix(int64(day.DateEpoch), 0).Format("Mon, Jan 2")

		temp, unitSymbol := formatTemp(day.Day.AvgTempC, day.Day.AvgTempF, unit)
		coloredTemp := colorizeTemp(temp, unitSymbol)

		icon := symbolFor(day.Day.Condition.Code)
		condition := color.New(color.FgWhite).Sprint(day.Day.Condition.Text)

		t.AppendRow([]interface{}{date, coloredTemp, icon, condition})
	}

	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignLeft},
		{Number: 2, Align: text.AlignCenter},
		{Number: 3, Align: text.AlignCenter},
		{Number: 4, Align: text.AlignLeft},
	})

	t.Render()
}

func formatTemp(tempC, tempF float64, unit string) (float64, string) {
	if unit == "fahrenheit" {
		return tempF, "°F"
	}
	return tempC, "°C"
}
