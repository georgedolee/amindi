package cmd

import (
	"fmt"
	"os"

	"github.com/georgedolee/amindi/internal/apiclient"
	"github.com/georgedolee/amindi/internal/formatter"
	"github.com/georgedolee/amindi/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// Embedded API Key (Free Tier - Rate Limited)
// IMPORTANT: This is a shared key for initial user experience
// For production use, users should obtain their own key
const embeddedAPIKey = "b663718a012840be882181846250207" // Replace with actual key

var rootCmd = &cobra.Command{
	Use:   "amindi [location]",
	Short: "Check the weather forecast",
	Long: `amindi lets you check the weather forecast quickly and easily from the command line.

You can specify the temperature unit (Celsius or Fahrenheit) and the number of days for the forecast.  
Supports flags for choosing units, forecast length, and location input.  
Perfect for getting current weather and multi-day forecasts with simple commands.`,
	Args: cobra.ExactArgs(1),
	RunE: ShowWeather,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	_ = godotenv.Load()

	rootCmd.Flags().StringP(
		"unit", "u", "celsius",
		"Temperature unit: fahrenheit / celsius")
	rootCmd.Flags().IntP(
		"days", "d", 3,
		"Forecast days: 1 - 7")
}

func validateUnitFlag(cmd *cobra.Command) (string, error) {
	unit, err := cmd.Flags().GetString("unit")
	if err != nil {
		return "", fmt.Errorf("error parsing --unit flag: %w", err)
	}
	if unit != "celsius" && unit != "fahrenheit" {
		return "", fmt.Errorf("invalid unit. Use \"celsius\" or \"fahrenheit\"")
	}
	return unit, nil
}

func validateDaysFlag(cmd *cobra.Command) (int, error) {
	days, err := cmd.Flags().GetInt("days")
	if err != nil {
		return 0, fmt.Errorf("error parsing --days flag: %w", err)
	}
	if days < 0 || days > 7 {
		return 0, fmt.Errorf("invalid value for --days. Must be between 0 and 7")
	}
	return days, nil
}

func ShowWeather(cmd *cobra.Command, args []string) error {
	location := args[0]
	unit, err := validateUnitFlag(cmd)
	if err != nil {
		return err
	}

	days, err := validateDaysFlag(cmd)
	if err != nil {
		return err
	}

	apiKey := os.Getenv("AMINDI_API_KEY")
	if apiKey == "" {
		apiKey = embeddedAPIKey
	}

	apiClient := apiclient.NewClient(apiKey)
	forecastService := service.NewForecastService(apiClient)

	forecast, err := forecastService.Get(location, days)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	formatter.PrintForecast(forecast, unit, days)

	return nil
}
