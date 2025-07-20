package cmd

import (
	"bufio"
	"calculator-generator/internal"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// interactiveCmd represents the interactive command
var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Interactive calculator generator wizard",
	Long: `Launch an interactive wizard to guide you through creating a customized Python calculator.

This wizard will ask you questions about what kind of calculator you want to create
and what features you'd like to include. It's perfect for beginners or when you
want to explore all available options.

The wizard will guide you through:
- Calculator type selection
- Feature selection
- Library dependencies
- UI customization
- Output configuration`,
	RunE: runInteractive,
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}

func runInteractive(cmd *cobra.Command, args []string) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🔢 Welcome to the Calculator Generator Interactive Wizard!")
	fmt.Println("This wizard will help you create a customized Python calculator.")
	fmt.Println()

	config := internal.GetDefaultConfig()

	// Project Information
	if err := askProjectInfo(reader, &config); err != nil {
		return err
	}

	// Calculator Type
	if err := askCalculatorType(reader, &config); err != nil {
		return err
	}

	// Features Selection
	if err := askFeatures(reader, &config); err != nil {
		return err
	}

	// Libraries Selection
	if err := askLibraries(reader, &config); err != nil {
		return err
	}

	// UI Configuration
	if err := askUIConfig(reader, &config); err != nil {
		return err
	}

	// Output Configuration
	if err := askOutputConfig(reader, &config); err != nil {
		return err
	}

	// Summary and Confirmation
	if err := showSummaryAndConfirm(reader, &config); err != nil {
		return err
	}

	// Generate the calculator
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

	if hasExternalLibrariesInteractive(config) {
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

func askProjectInfo(reader *bufio.Reader, config *internal.CalculatorConfig) error {
	fmt.Println("📋 Project Information")
	fmt.Println("======================")

	// Project name
	fmt.Printf("Project name [%s]: ", config.ProjectName)
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	name = strings.TrimSpace(name)
	if name != "" {
		config.ProjectName = name
	}

	// Author
	fmt.Printf("Author name [%s]: ", config.Author)
	author, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	author = strings.TrimSpace(author)
	if author != "" {
		config.Author = author
	}

	// Description
	fmt.Printf("Description [%s]: ", config.Description)
	description, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	description = strings.TrimSpace(description)
	if description != "" {
		config.Description = description
	}

	fmt.Println()
	return nil
}

func askCalculatorType(reader *bufio.Reader, config *internal.CalculatorConfig) error {
	fmt.Println("🔢 Calculator Type")
	fmt.Println("==================")
	fmt.Println("1. Basic Calculator - Simple arithmetic operations (+, -, *, /, etc.)")
	fmt.Println("2. Scientific Calculator - Advanced mathematical functions")
	fmt.Println()

	for {
		fmt.Print("Choose calculator type [1-2]: ")
		choice, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1", "basic", "b":
			config.Type = internal.BasicCalculator
			fmt.Println("✅ Basic calculator selected")
			fmt.Println()
			return nil
		case "2", "scientific", "s":
			config.Type = internal.ScientificCalculator
			// Apply scientific defaults
			scientificConfig := internal.GetScientificConfig()
			config.Libraries = scientificConfig.Libraries
			config.Features = scientificConfig.Features
			fmt.Println("✅ Scientific calculator selected")
			fmt.Println()
			return nil
		default:
			fmt.Println("❌ Invalid choice. Please enter 1 or 2.")
		}
	}
}

func askFeatures(reader *bufio.Reader, config *internal.CalculatorConfig) error {
	fmt.Println("🚀 Features Selection")
	fmt.Println("====================")

	features := []struct {
		name        string
		description string
		field       *bool
		dependency  string
	}{
		{"Memory", "Store and recall values", &config.Features.Memory, ""},
		{"History", "Keep track of calculations", &config.Features.History, ""},
		{"Trigonometric", "sin, cos, tan functions", &config.Features.Trigonometric, "math"},
		{"Logarithmic", "log, ln functions", &config.Features.Logarithmic, "math"},
		{"Statistical", "mean, median, std dev", &config.Features.Statistical, "numpy"},
		{"Linear Algebra", "Matrix operations", &config.Features.LinearAlgebra, "numpy"},
		{"Plotting", "Create graphs and charts", &config.Features.Plotting, "plotly"},
		{"Equation Solver", "Solve algebraic equations", &config.Features.EquationSolver, "sympy"},
		{"Complex Numbers", "Complex number arithmetic", &config.Features.ComplexNumbers, ""},
		{"Unit Conversion", "Convert between units", &config.Features.UnitConversion, ""},
	}

	for _, feature := range features {
		current := "no"
		if *feature.field {
			current = "yes"
		}

		dependency := ""
		if feature.dependency != "" {
			dependency = fmt.Sprintf(" (requires %s)", feature.dependency)
		}

		fmt.Printf("Include %s - %s%s? [y/N] (current: %s): ",
			feature.name, feature.description, dependency, current)

		choice, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		choice = strings.TrimSpace(strings.ToLower(choice))

		if choice == "y" || choice == "yes" {
			*feature.field = true
			// Enable required libraries
			switch feature.dependency {
			case "numpy":
				config.Libraries.UseNumpy = true
			case "plotly":
				config.Libraries.UsePlotly = true
			case "sympy":
				config.Libraries.UseSympy = true
			case "math":
				config.Libraries.UseMath = true
			}
		} else if choice == "n" || choice == "no" {
			*feature.field = false
		}
		// If empty, keep current value
	}

	fmt.Println()
	return nil
}

func askLibraries(reader *bufio.Reader, config *internal.CalculatorConfig) error {
	fmt.Println("📚 Library Dependencies")
	fmt.Println("=======================")
	fmt.Println("Additional libraries can provide more functionality:")

	libraries := []struct {
		name        string
		description string
		field       *bool
	}{
		{"NumPy", "Numerical computing library", &config.Libraries.UseNumpy},
		{"Pandas", "Data analysis and manipulation", &config.Libraries.UsePandas},
		{"SciPy", "Scientific computing functions", &config.Libraries.UseScipy},
		{"SymPy", "Symbolic mathematics", &config.Libraries.UseSympy},
		{"Plotly", "Interactive plotting library", &config.Libraries.UsePlotly},
	}

	for _, lib := range libraries {
		current := "no"
		if *lib.field {
			current = "yes"
		}

		fmt.Printf("Include %s - %s? [y/N] (current: %s): ",
			lib.name, lib.description, current)

		choice, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		choice = strings.TrimSpace(strings.ToLower(choice))

		if choice == "y" || choice == "yes" {
			*lib.field = true
		} else if choice == "n" || choice == "no" {
			*lib.field = false
		}
	}

	fmt.Println()
	return nil
}

func askUIConfig(reader *bufio.Reader, config *internal.CalculatorConfig) error {
	fmt.Println("🎨 User Interface Configuration")
	fmt.Println("==============================")

	// Interactive mode
	fmt.Printf("Create interactive calculator? [Y/n] (current: %v): ", config.Interactive)
	choice, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	choice = strings.TrimSpace(strings.ToLower(choice))
	if choice == "n" || choice == "no" {
		config.Interactive = false
	} else if choice == "y" || choice == "yes" || choice == "" {
		config.Interactive = true
	}

	// Precision
	fmt.Printf("Decimal precision (1-20) [%d]: ", config.UI.Precision)
	precisionStr, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	precisionStr = strings.TrimSpace(precisionStr)
	if precisionStr != "" {
		if precision, err := strconv.Atoi(precisionStr); err == nil && precision >= 1 && precision <= 20 {
			config.UI.Precision = precision
		}
	}

	// Angle unit
	// UI Style
	fmt.Println("Calculator interface:")
	fmt.Println("1. Command Line Interface (CLI)")
	fmt.Println("2. Desktop GUI Application")
	fmt.Printf("Choose [1-2] (current: %s): ", func() string {
		if config.UI.Style == "gui" {
			return "gui"
		}
		return "cli"
	}())
	choice, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	choice = strings.TrimSpace(choice)
	switch choice {
	case "1", "cli", "command", "c":
		config.UI.Style = "cli"
	case "2", "gui", "desktop", "g":
		config.UI.Style = "gui"
	}

	// Angle unit
	fmt.Println("Angle unit:")
	fmt.Println("1. Degrees")
	fmt.Println("2. Radians")
	fmt.Printf("Choose [1-2] (current: %s): ", config.UI.AngleUnit)
	choice, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	choice = strings.TrimSpace(choice)
	switch choice {
	case "1", "degrees", "deg":
		config.UI.AngleUnit = "degrees"
	case "2", "radians", "rad":
		config.UI.AngleUnit = "radians"
	}

	// Show banner
	fmt.Printf("Show application banner? [Y/n] (current: %v): ", config.UI.ShowBanner)
	choice, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	choice = strings.TrimSpace(strings.ToLower(choice))
	if choice == "n" || choice == "no" {
		config.UI.ShowBanner = false
	} else if choice == "y" || choice == "yes" || choice == "" {
		config.UI.ShowBanner = true
	}

	// Show help
	fmt.Printf("Show help information? [Y/n] (current: %v): ", config.UI.ShowHelp)
	choice, err = reader.ReadString('\n')
	if err != nil {
		return err
	}
	choice = strings.TrimSpace(strings.ToLower(choice))
	if choice == "n" || choice == "no" {
		config.UI.ShowHelp = false
	} else if choice == "y" || choice == "yes" || choice == "" {
		config.UI.ShowHelp = true
	}

	fmt.Println()
	return nil
}

func askOutputConfig(reader *bufio.Reader, config *internal.CalculatorConfig) error {
	fmt.Println("📁 Output Configuration")
	fmt.Println("======================")

	// Output file
	fmt.Printf("Output file path [%s]: ", config.OutputFile)
	outputFile, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	outputFile = strings.TrimSpace(outputFile)
	if outputFile != "" {
		config.OutputFile = outputFile
	}

	fmt.Println()
	return nil
}

func showSummaryAndConfirm(reader *bufio.Reader, config *internal.CalculatorConfig) error {
	fmt.Println("📋 Configuration Summary")
	fmt.Println("========================")
	fmt.Printf("Project Name: %s\n", config.ProjectName)
	fmt.Printf("Author: %s\n", config.Author)
	fmt.Printf("Description: %s\n", config.Description)
	fmt.Printf("Type: %s\n", config.Type)
	fmt.Printf("UI Style: %s\n", config.UI.Style)
	fmt.Printf("Output File: %s\n", config.OutputFile)
	fmt.Printf("Interactive: %v\n", config.Interactive)
	fmt.Printf("Precision: %d\n", config.UI.Precision)
	fmt.Printf("Angle Unit: %s\n", config.UI.AngleUnit)

	// Show enabled libraries
	var enabledLibs []string
	if config.Libraries.UseMath {
		enabledLibs = append(enabledLibs, "math")
	}
	if config.Libraries.UseNumpy {
		enabledLibs = append(enabledLibs, "numpy")
	}
	if config.Libraries.UsePandas {
		enabledLibs = append(enabledLibs, "pandas")
	}
	if config.Libraries.UseScipy {
		enabledLibs = append(enabledLibs, "scipy")
	}
	if config.Libraries.UseSympy {
		enabledLibs = append(enabledLibs, "sympy")
	}
	if config.Libraries.UsePlotly {
		enabledLibs = append(enabledLibs, "plotly")
	}
	fmt.Printf("Libraries: %s\n", strings.Join(enabledLibs, ", "))

	// Show enabled features
	var enabledFeatures []string
	if config.Features.BasicArithmetic {
		enabledFeatures = append(enabledFeatures, "basic-arithmetic")
	}
	if config.Features.Memory {
		enabledFeatures = append(enabledFeatures, "memory")
	}
	if config.Features.History {
		enabledFeatures = append(enabledFeatures, "history")
	}
	if config.Features.Trigonometric {
		enabledFeatures = append(enabledFeatures, "trigonometric")
	}
	if config.Features.Logarithmic {
		enabledFeatures = append(enabledFeatures, "logarithmic")
	}
	if config.Features.Statistical {
		enabledFeatures = append(enabledFeatures, "statistical")
	}
	if config.Features.LinearAlgebra {
		enabledFeatures = append(enabledFeatures, "linear-algebra")
	}
	if config.Features.Plotting {
		enabledFeatures = append(enabledFeatures, "plotting")
	}
	if config.Features.EquationSolver {
		enabledFeatures = append(enabledFeatures, "equation-solver")
	}
	if config.Features.ComplexNumbers {
		enabledFeatures = append(enabledFeatures, "complex-numbers")
	}
	if config.Features.UnitConversion {
		enabledFeatures = append(enabledFeatures, "unit-conversion")
	}
	fmt.Printf("Features: %s\n", strings.Join(enabledFeatures, ", "))

	fmt.Println()
	fmt.Print("Generate calculator with this configuration? [Y/n]: ")
	choice, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	choice = strings.TrimSpace(strings.ToLower(choice))

	if choice == "n" || choice == "no" {
		return fmt.Errorf("generation cancelled by user")
	}

	return nil
}

func hasExternalLibrariesInteractive(config internal.CalculatorConfig) bool {
	return config.Libraries.UseNumpy ||
		config.Libraries.UsePandas ||
		config.Libraries.UseScipy ||
		config.Libraries.UseSympy ||
		config.Libraries.UsePlotly
}
