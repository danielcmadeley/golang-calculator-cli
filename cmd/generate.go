package cmd

import (
	"calculator-generator/internal"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Python calculator script",
	Long: `Generate a customizable Python calculator script with the specified features.

You can create either a basic calculator with simple arithmetic operations
or a scientific calculator with advanced mathematical functions.

Examples:
  calculator-generator generate --type basic
  calculator-generator generate --type scientific --output scientific_calc.py
  calculator-generator generate --type scientific --features "trigonometric,logarithmic,statistical"
  calculator-generator generate --libraries "numpy,scipy,sympy" --features "plotting,linear-algebra"`,
	RunE: runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Calculator type
	generateCmd.Flags().StringP("type", "t", "basic", "calculator type (basic, scientific)")

	// Project information
	generateCmd.Flags().StringP("name", "n", "Python Calculator", "project name")
	generateCmd.Flags().StringP("description", "d", "", "project description")

	// Libraries
	generateCmd.Flags().String("libraries", "", "comma-separated list of libraries (numpy,pandas,scipy,sympy,plotly)")
	generateCmd.Flags().Bool("math", true, "include math library")

	// Features
	generateCmd.Flags().String("features", "", "comma-separated list of features")
	generateCmd.Flags().Bool("memory", false, "include memory functionality")
	generateCmd.Flags().Bool("history", false, "include calculation history")
	generateCmd.Flags().Bool("interactive", true, "create interactive calculator")

	// UI configuration
	generateCmd.Flags().String("style", "cli", "UI style (cli, gui)")
	generateCmd.Flags().String("theme", "light", "UI theme (light, dark, colorful)")
	generateCmd.Flags().Int("precision", 10, "decimal precision (1-20)")
	generateCmd.Flags().String("angle-unit", "degrees", "angle unit (degrees, radians)")
	generateCmd.Flags().Bool("show-help", true, "show help information")
	generateCmd.Flags().Bool("show-banner", true, "show application banner")

	// Bind flags to viper
	viper.BindPFlag("type", generateCmd.Flags().Lookup("type"))
	viper.BindPFlag("name", generateCmd.Flags().Lookup("name"))
	viper.BindPFlag("description", generateCmd.Flags().Lookup("description"))
	viper.BindPFlag("libraries", generateCmd.Flags().Lookup("libraries"))
	viper.BindPFlag("math", generateCmd.Flags().Lookup("math"))
	viper.BindPFlag("features", generateCmd.Flags().Lookup("features"))
	viper.BindPFlag("memory", generateCmd.Flags().Lookup("memory"))
	viper.BindPFlag("history", generateCmd.Flags().Lookup("history"))
	viper.BindPFlag("interactive", generateCmd.Flags().Lookup("interactive"))
	viper.BindPFlag("style", generateCmd.Flags().Lookup("style"))
	viper.BindPFlag("theme", generateCmd.Flags().Lookup("theme"))
	viper.BindPFlag("precision", generateCmd.Flags().Lookup("precision"))
	viper.BindPFlag("angle-unit", generateCmd.Flags().Lookup("angle-unit"))
	viper.BindPFlag("show-help", generateCmd.Flags().Lookup("show-help"))
	viper.BindPFlag("show-banner", generateCmd.Flags().Lookup("show-banner"))
}

func runGenerate(cmd *cobra.Command, args []string) error {
	// Get config from flags
	getConfigFromFlags()
	config := buildConfigFromFlags()

	// Validate calculator type
	calcType := viper.GetString("type")
	switch calcType {
	case "basic":
		config.Type = internal.BasicCalculator
	case "scientific":
		config = internal.GetScientificConfig()
		// Override with user flags
		config = buildConfigFromFlags()
		config.Type = internal.ScientificCalculator
	default:
		return fmt.Errorf("invalid calculator type: %s (must be 'basic' or 'scientific')", calcType)
	}

	// Apply libraries from flags
	if err := applyLibrariesFromFlags(&config); err != nil {
		return err
	}

	// Apply features from flags
	if err := applyFeaturesFromFlags(&config); err != nil {
		return err
	}

	// Generate calculator
	generator := internal.NewGenerator(config)
	if err := generator.Generate(); err != nil {
		return fmt.Errorf("generation failed: %w", err)
	}

	// Success message
	fmt.Printf("✅ Calculator generated successfully!\n")
	fmt.Printf("📁 Output file: %s\n", config.OutputFile)

	if config.UI.Style == "gui" {
		fmt.Printf("🖥️  Type: Desktop GUI Calculator\n")
	} else {
		fmt.Printf("💻 Type: Command Line Calculator\n")
	}

	if hasExternalLibraries(config) {
		fmt.Printf("📦 Requirements file: requirements.txt\n")
		fmt.Printf("💡 Install dependencies with: pip install -r requirements.txt\n")
	}

	if config.UI.Style == "gui" {
		fmt.Printf("🚀 Run GUI with: python %s\n", config.OutputFile)
		fmt.Printf("💡 Note: Tkinter is included with Python (no additional install needed)\n")
	} else {
		fmt.Printf("🚀 Run with: python %s\n", config.OutputFile)
	}

	return nil
}

func buildConfigFromFlags() internal.CalculatorConfig {
	config := internal.GetDefaultConfig()

	// Basic information
	config.ProjectName = viper.GetString("name")
	config.Author = viper.GetString("author")
	config.OutputFile = viper.GetString("output")

	description := viper.GetString("description")
	if description == "" {
		if viper.GetString("type") == "scientific" {
			description = "A scientific calculator with advanced mathematical functions"
		} else {
			description = "A basic calculator with essential arithmetic operations"
		}
	}
	config.Description = description

	// General settings
	config.Interactive = viper.GetBool("interactive")

	// Library settings
	config.Libraries.UseMath = viper.GetBool("math")

	// Feature settings
	config.Features.Memory = viper.GetBool("memory")
	config.Features.History = viper.GetBool("history")

	// UI settings
	config.UI.Style = viper.GetString("style")
	config.UI.Theme = viper.GetString("theme")
	config.UI.Precision = viper.GetInt("precision")
	config.UI.AngleUnit = viper.GetString("angle-unit")
	config.UI.ShowHelp = viper.GetBool("show-help")
	config.UI.ShowBanner = viper.GetBool("show-banner")

	return config
}

func applyLibrariesFromFlags(config *internal.CalculatorConfig) error {
	librariesStr := viper.GetString("libraries")
	if librariesStr == "" {
		return nil
	}

	libraries := strings.Split(librariesStr, ",")
	for _, lib := range libraries {
		lib = strings.TrimSpace(strings.ToLower(lib))
		switch lib {
		case "numpy":
			config.Libraries.UseNumpy = true
		case "pandas":
			config.Libraries.UsePandas = true
		case "scipy":
			config.Libraries.UseScipy = true
		case "sympy":
			config.Libraries.UseSympy = true
		case "plotly":
			config.Libraries.UsePlotly = true
		case "math":
			config.Libraries.UseMath = true
		default:
			return fmt.Errorf("unknown library: %s", lib)
		}
	}

	return nil
}

func applyFeaturesFromFlags(config *internal.CalculatorConfig) error {
	featuresStr := viper.GetString("features")
	if featuresStr == "" {
		return nil
	}

	features := strings.Split(featuresStr, ",")
	for _, feature := range features {
		feature = strings.TrimSpace(strings.ToLower(feature))
		switch feature {
		case "basic-arithmetic", "arithmetic":
			config.Features.BasicArithmetic = true
		case "trigonometric", "trig":
			config.Features.Trigonometric = true
			config.Libraries.UseMath = true
		case "logarithmic", "log":
			config.Features.Logarithmic = true
			config.Libraries.UseMath = true
		case "exponential", "exp":
			config.Features.Exponential = true
			config.Libraries.UseMath = true
		case "statistical", "stats":
			config.Features.Statistical = true
			config.Libraries.UseNumpy = true
		case "linear-algebra", "linalg":
			config.Features.LinearAlgebra = true
			config.Libraries.UseNumpy = true
		case "calculus":
			config.Features.Calculus = true
			config.Libraries.UseSympy = true
		case "plotting", "plot":
			config.Features.Plotting = true
			config.Libraries.UsePlotly = true
		case "unit-conversion", "units":
			config.Features.UnitConversion = true
		case "complex-numbers", "complex":
			config.Features.ComplexNumbers = true
		case "equation-solver", "solver":
			config.Features.EquationSolver = true
			config.Libraries.UseSympy = true
		case "matrix-operations", "matrix":
			config.Features.MatrixOperations = true
			config.Libraries.UseNumpy = true
		case "data-analysis", "data":
			config.Features.DataAnalysis = true
			config.Libraries.UsePandas = true
			config.Libraries.UseNumpy = true
		case "graphing", "graph":
			config.Features.Graphing = true
			config.Libraries.UsePlotly = true
		case "programming", "prog":
			config.Features.Programming = true
		case "memory", "mem":
			config.Features.Memory = true
		case "history", "hist":
			config.Features.History = true
		default:
			return fmt.Errorf("unknown feature: %s", feature)
		}
	}

	return nil
}

func hasExternalLibraries(config internal.CalculatorConfig) bool {
	return config.Libraries.UseNumpy ||
		config.Libraries.UsePandas ||
		config.Libraries.UseScipy ||
		config.Libraries.UseSympy ||
		config.Libraries.UsePlotly
}
