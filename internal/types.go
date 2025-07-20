package internal

// CalculatorType represents the type of calculator to generate
type CalculatorType string

const (
	BasicCalculator      CalculatorType = "basic"
	ScientificCalculator CalculatorType = "scientific"
)

// CalculatorConfig holds all configuration options for the calculator generator
type CalculatorConfig struct {
	Type        CalculatorType `json:"type"`
	OutputFile  string         `json:"output_file"`
	ProjectName string         `json:"project_name"`
	Author      string         `json:"author"`
	Description string         `json:"description"`
	Interactive bool           `json:"interactive"`
	Libraries   Libraries      `json:"libraries"`
	Features    Features       `json:"features"`
	UI          UIConfig       `json:"ui"`
}

// Libraries configuration for Python dependencies
type Libraries struct {
	UseNumpy  bool `json:"use_numpy"`
	UsePandas bool `json:"use_pandas"`
	UseScipy  bool `json:"use_scipy"`
	UseMath   bool `json:"use_math"`
	UseSympy  bool `json:"use_sympy"`
	UsePlotly bool `json:"use_plotly"`
}

// Features configuration for calculator capabilities
type Features struct {
	// Basic features
	BasicArithmetic bool `json:"basic_arithmetic"`
	History         bool `json:"history"`
	Memory          bool `json:"memory"`

	// Scientific features
	Trigonometric  bool `json:"trigonometric"`
	Logarithmic    bool `json:"logarithmic"`
	Exponential    bool `json:"exponential"`
	Statistical    bool `json:"statistical"`
	LinearAlgebra  bool `json:"linear_algebra"`
	Calculus       bool `json:"calculus"`
	Plotting       bool `json:"plotting"`
	UnitConversion bool `json:"unit_conversion"`
	ComplexNumbers bool `json:"complex_numbers"`

	// Advanced features
	EquationSolver   bool `json:"equation_solver"`
	MatrixOperations bool `json:"matrix_operations"`
	DataAnalysis     bool `json:"data_analysis"`
	Graphing         bool `json:"graphing"`
	Programming      bool `json:"programming"`
}

// UIConfig configuration for user interface options
type UIConfig struct {
	Style      string `json:"style"` // "cli", "gui", "web"
	Theme      string `json:"theme"` // "light", "dark", "colorful"
	ShowHelp   bool   `json:"show_help"`
	ShowBanner bool   `json:"show_banner"`
	Precision  int    `json:"precision"`  // decimal places
	AngleUnit  string `json:"angle_unit"` // "degrees", "radians"
}

// TemplateData holds data for template rendering
type TemplateData struct {
	Config      CalculatorConfig
	Imports     []string
	Functions   []string
	MainContent string
	Version     string
	Timestamp   string
}

// ValidationError represents a configuration validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

// GetDefaultConfig returns a default calculator configuration
func GetDefaultConfig() CalculatorConfig {
	return CalculatorConfig{
		Type:        BasicCalculator,
		OutputFile:  "calculator.py",
		ProjectName: "Python Calculator",
		Author:      "Calculator Generator",
		Description: "A customizable calculator application",
		Interactive: true,
		Libraries: Libraries{
			UseNumpy:  false,
			UsePandas: false,
			UseScipy:  false,
			UseMath:   true,
			UseSympy:  false,
			UsePlotly: false,
		},
		Features: Features{
			BasicArithmetic:  true,
			History:          false,
			Memory:           false,
			Trigonometric:    false,
			Logarithmic:      false,
			Exponential:      false,
			Statistical:      false,
			LinearAlgebra:    false,
			Calculus:         false,
			Plotting:         false,
			UnitConversion:   false,
			ComplexNumbers:   false,
			EquationSolver:   false,
			MatrixOperations: false,
			DataAnalysis:     false,
			Graphing:         false,
			Programming:      false,
		},
		UI: UIConfig{
			Style:      "cli",
			Theme:      "light",
			ShowHelp:   true,
			ShowBanner: true,
			Precision:  10,
			AngleUnit:  "degrees",
		},
	}
}

// GetScientificConfig returns a scientific calculator configuration
func GetScientificConfig() CalculatorConfig {
	config := GetDefaultConfig()
	config.Type = ScientificCalculator
	config.Description = "A scientific calculator with advanced mathematical functions"

	// Enable scientific libraries
	config.Libraries.UseNumpy = true
	config.Libraries.UseScipy = true
	config.Libraries.UseSympy = true

	// Enable scientific features
	config.Features.Trigonometric = true
	config.Features.Logarithmic = true
	config.Features.Exponential = true
	config.Features.Statistical = true
	config.Features.LinearAlgebra = true
	config.Features.ComplexNumbers = true
	config.Features.History = true
	config.Features.Memory = true

	return config
}
