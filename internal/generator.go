package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// Generator handles the generation of Python calculator scripts
type Generator struct {
	config CalculatorConfig
}

// NewGenerator creates a new calculator generator instance
func NewGenerator(config CalculatorConfig) *Generator {
	return &Generator{config: config}
}

// Generate creates the Python calculator script based on the configuration
func (g *Generator) Generate() error {
	// Validate configuration
	if err := g.validateConfig(); err != nil {
		return fmt.Errorf("configuration validation failed: %w", err)
	}

	// Prepare template data
	data := g.prepareTemplateData()

	// Generate the Python script
	var content string
	var err error

	if g.config.UI.Style == "gui" {
		// Generate GUI calculator
		guiGen := NewGUIGenerator(g.config)
		content, err = guiGen.GenerateGUICalculator()
	} else {
		// Generate CLI calculator
		content, err = g.renderTemplate(data)
	}

	if err != nil {
		return fmt.Errorf("template rendering failed: %w", err)
	}

	// Write to file
	if err := g.writeToFile(content); err != nil {
		return fmt.Errorf("file writing failed: %w", err)
	}

	// Generate requirements.txt
	if err := g.generateRequirements(); err != nil {
		return fmt.Errorf("requirements generation failed: %w", err)
	}

	return nil
}

// validateConfig validates the calculator configuration
func (g *Generator) validateConfig() error {
	if g.config.OutputFile == "" {
		return ValidationError{Field: "output_file", Message: "output file cannot be empty"}
	}

	if g.config.ProjectName == "" {
		return ValidationError{Field: "project_name", Message: "project name cannot be empty"}
	}

	if g.config.UI.Precision < 1 || g.config.UI.Precision > 20 {
		return ValidationError{Field: "precision", Message: "precision must be between 1 and 20"}
	}

	return nil
}

// prepareTemplateData prepares data for template rendering
func (g *Generator) prepareTemplateData() TemplateData {
	imports := g.generateImports()
	functions := g.generateFunctions()
	mainContent := g.generateMainContent()

	return TemplateData{
		Config:      g.config,
		Imports:     imports,
		Functions:   functions,
		MainContent: mainContent,
		Version:     "1.0.0",
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
	}
}

// generateImports creates the list of Python imports based on configuration
func (g *Generator) generateImports() []string {
	var imports []string

	// Standard library imports
	imports = append(imports, "import sys")
	imports = append(imports, "import os")

	if g.config.Libraries.UseMath {
		imports = append(imports, "import math")
	}

	if g.config.Features.History {
		imports = append(imports, "import json")
		imports = append(imports, "from datetime import datetime")
	}

	// Third-party library imports
	if g.config.Libraries.UseNumpy {
		imports = append(imports, "import numpy as np")
	}

	if g.config.Libraries.UsePandas {
		imports = append(imports, "import pandas as pd")
	}

	if g.config.Libraries.UseScipy {
		imports = append(imports, "import scipy as sp")
		imports = append(imports, "from scipy import stats, optimize, integrate")
	}

	if g.config.Libraries.UseSympy {
		imports = append(imports, "import sympy as sym")
		imports = append(imports, "from sympy import symbols, solve, diff, integrate as sym_integrate")
	}

	if g.config.Libraries.UsePlotly || g.config.Features.Plotting {
		imports = append(imports, "import plotly.graph_objects as go")
		imports = append(imports, "import plotly.express as px")
	}

	if g.config.Features.ComplexNumbers {
		imports = append(imports, "import cmath")
	}

	return imports
}

// generateFunctions creates calculator function implementations
func (g *Generator) generateFunctions() []string {
	var functions []string

	// Basic arithmetic functions
	if g.config.Features.BasicArithmetic {
		functions = append(functions, g.generateBasicArithmetic()...)
	}

	// Memory functions
	if g.config.Features.Memory {
		functions = append(functions, g.generateMemoryFunctions()...)
	}

	// History functions
	if g.config.Features.History {
		functions = append(functions, g.generateHistoryFunctions()...)
	}

	// Scientific functions
	if g.config.Features.Trigonometric {
		functions = append(functions, g.generateTrigonometricFunctions()...)
	}

	if g.config.Features.Logarithmic {
		functions = append(functions, g.generateLogarithmicFunctions()...)
	}

	if g.config.Features.Statistical && g.config.Libraries.UseNumpy {
		functions = append(functions, g.generateStatisticalFunctions()...)
	}

	if g.config.Features.LinearAlgebra && g.config.Libraries.UseNumpy {
		functions = append(functions, g.generateLinearAlgebraFunctions()...)
	}

	if g.config.Features.Plotting && g.config.Libraries.UsePlotly {
		functions = append(functions, g.generatePlottingFunctions()...)
	}

	if g.config.Features.EquationSolver && g.config.Libraries.UseSympy {
		functions = append(functions, g.generateEquationSolverFunctions()...)
	}

	return functions
}

// generateBasicArithmetic creates basic arithmetic functions
func (g *Generator) generateBasicArithmetic() []string {
	return []string{
		`def add(a, b):
    """Addition operation"""
    return a + b`,

		`def subtract(a, b):
    """Subtraction operation"""
    return a - b`,

		`def multiply(a, b):
    """Multiplication operation"""
    return a * b`,

		`def divide(a, b):
    """Division operation"""
    if b == 0:
        raise ValueError("Cannot divide by zero")
    return a / b`,

		`def power(a, b):
    """Power operation"""
    return a ** b`,

		`def modulo(a, b):
    """Modulo operation"""
    if b == 0:
        raise ValueError("Cannot calculate modulo with zero")
    return a % b`,
	}
}

// generateMemoryFunctions creates memory-related functions
func (g *Generator) generateMemoryFunctions() []string {
	return []string{
		`class Memory:
    """Calculator memory functionality"""
    def __init__(self):
        self.value = 0

    def store(self, value):
        """Store value in memory"""
        self.value = value
        return f"Stored {value} in memory"

    def recall(self):
        """Recall value from memory"""
        return self.value

    def clear(self):
        """Clear memory"""
        self.value = 0
        return "Memory cleared"

    def add_to_memory(self, value):
        """Add value to memory"""
        self.value += value
        return f"Added {value} to memory, new value: {self.value}"`,
	}
}

// generateHistoryFunctions creates history-related functions
func (g *Generator) generateHistoryFunctions() []string {
	return []string{
		`class History:
    """Calculator history functionality"""
    def __init__(self, max_entries=100):
        self.entries = []
        self.max_entries = max_entries

    def add_entry(self, operation, result):
        """Add calculation to history"""
        entry = {
            "timestamp": datetime.now().isoformat(),
            "operation": operation,
            "result": result
        }
        self.entries.append(entry)

        # Keep only max_entries
        if len(self.entries) > self.max_entries:
            self.entries = self.entries[-self.max_entries:]

    def get_history(self, count=10):
        """Get recent history entries"""
        return self.entries[-count:]

    def clear_history(self):
        """Clear calculation history"""
        self.entries = []
        return "History cleared"

    def save_to_file(self, filename="calculator_history.json"):
        """Save history to file"""
        with open(filename, 'w') as f:
            json.dump(self.entries, f, indent=2)
        return f"History saved to {filename}"`,
	}
}

// generateTrigonometricFunctions creates trigonometric functions
func (g *Generator) generateTrigonometricFunctions() []string {
	functions := []string{
		`def sin(x, angle_unit="degrees"):
    """Sine function"""
    if angle_unit == "degrees":
        x = math.radians(x)
    return math.sin(x)`,

		`def cos(x, angle_unit="degrees"):
    """Cosine function"""
    if angle_unit == "degrees":
        x = math.radians(x)
    return math.cos(x)`,

		`def tan(x, angle_unit="degrees"):
    """Tangent function"""
    if angle_unit == "degrees":
        x = math.radians(x)
    return math.tan(x)`,

		`def asin(x, angle_unit="degrees"):
    """Arcsine function"""
    result = math.asin(x)
    if angle_unit == "degrees":
        result = math.degrees(result)
    return result`,

		`def acos(x, angle_unit="degrees"):
    """Arccosine function"""
    result = math.acos(x)
    if angle_unit == "degrees":
        result = math.degrees(result)
    return result`,

		`def atan(x, angle_unit="degrees"):
    """Arctangent function"""
    result = math.atan(x)
    if angle_unit == "degrees":
        result = math.degrees(result)
    return result`,
	}

	return functions
}

// generateLogarithmicFunctions creates logarithmic functions
func (g *Generator) generateLogarithmicFunctions() []string {
	return []string{
		`def log(x, base=10):
    """Logarithm function"""
    if x <= 0:
        raise ValueError("Logarithm input must be positive")
    if base == math.e:
        return math.log(x)
    return math.log(x, base)`,

		`def ln(x):
    """Natural logarithm"""
    if x <= 0:
        raise ValueError("Natural logarithm input must be positive")
    return math.log(x)`,

		`def log10(x):
    """Base-10 logarithm"""
    if x <= 0:
        raise ValueError("Logarithm input must be positive")
    return math.log10(x)`,

		`def log2(x):
    """Base-2 logarithm"""
    if x <= 0:
        raise ValueError("Logarithm input must be positive")
    return math.log2(x)`,
	}
}

// generateStatisticalFunctions creates statistical functions
func (g *Generator) generateStatisticalFunctions() []string {
	return []string{
		`def mean(data):
    """Calculate mean of data"""
    return np.mean(data)`,

		`def median(data):
    """Calculate median of data"""
    return np.median(data)`,

		`def std(data):
    """Calculate standard deviation"""
    return np.std(data)`,

		`def variance(data):
    """Calculate variance"""
    return np.var(data)`,

		`def correlation(x, y):
    """Calculate correlation coefficient"""
    return np.corrcoef(x, y)[0, 1]`,
	}
}

// generateLinearAlgebraFunctions creates linear algebra functions
func (g *Generator) generateLinearAlgebraFunctions() []string {
	return []string{
		`def matrix_multiply(a, b):
    """Matrix multiplication"""
    return np.dot(a, b)`,

		`def matrix_inverse(matrix):
    """Matrix inverse"""
    return np.linalg.inv(matrix)`,

		`def matrix_determinant(matrix):
    """Matrix determinant"""
    return np.linalg.det(matrix)`,

		`def eigenvalues(matrix):
    """Calculate eigenvalues"""
    return np.linalg.eigvals(matrix)`,
	}
}

// generatePlottingFunctions creates plotting functions
func (g *Generator) generatePlottingFunctions() []string {
	return []string{
		`def plot_function(func_str, x_range=(-10, 10), num_points=100):
    """Plot a mathematical function"""
    x = np.linspace(x_range[0], x_range[1], num_points)
    y = eval(func_str.replace('x', 'x'))

    fig = go.Figure()
    fig.add_trace(go.Scatter(x=x, y=y, mode='lines', name=func_str))
    fig.update_layout(title=f"Plot of {func_str}", xaxis_title="x", yaxis_title="y")
    fig.show()`,

		`def plot_data(x_data, y_data, title="Data Plot"):
    """Plot data points"""
    fig = go.Figure()
    fig.add_trace(go.Scatter(x=x_data, y=y_data, mode='markers+lines'))
    fig.update_layout(title=title, xaxis_title="x", yaxis_title="y")
    fig.show()`,
	}
}

// generateEquationSolverFunctions creates equation solver functions
func (g *Generator) generateEquationSolverFunctions() []string {
	return []string{
		`def solve_equation(equation_str, variable='x'):
    """Solve algebraic equation"""
    x = symbols(variable)
    equation = sym.sympify(equation_str)
    solutions = solve(equation, x)
    return solutions`,

		`def differentiate(expr_str, variable='x'):
    """Calculate derivative"""
    x = symbols(variable)
    expr = sym.sympify(expr_str)
    return diff(expr, x)`,

		`def integrate_symbolic(expr_str, variable='x'):
    """Calculate symbolic integral"""
    x = symbols(variable)
    expr = sym.sympify(expr_str)
    return sym_integrate(expr, x)`,
	}
}

// generateMainContent creates the main calculator interface
func (g *Generator) generateMainContent() string {
	var content strings.Builder

	content.WriteString(`class Calculator:
    """Main calculator class"""

    def __init__(self):
        self.precision = ` + fmt.Sprintf("%d", g.config.UI.Precision) + `
        self.angle_unit = "` + g.config.UI.AngleUnit + `"
`)

	if g.config.Features.Memory {
		content.WriteString("        self.memory = Memory()\n")
	}

	if g.config.Features.History {
		content.WriteString("        self.history = History()\n")
	}

	content.WriteString(`
    def format_result(self, result):
        """Format calculation result"""
        if isinstance(result, (int, float)):
            return round(result, self.precision)
        return result

    def run(self):
        """Run the calculator interface"""
`)

	if g.config.UI.ShowBanner {
		content.WriteString(`        self.show_banner()
`)
	}

	if g.config.Interactive {
		content.WriteString(`        self.interactive_mode()

    def interactive_mode(self):
        """Interactive calculator mode"""
        print("Calculator started. Type 'help' for commands, 'quit' to exit.")

        while True:
            try:
                user_input = input("calc> ").strip()

                if user_input.lower() in ['quit', 'exit', 'q']:
                    break
                elif user_input.lower() == 'help':
                    self.show_help()
                elif user_input.lower() == 'clear':
                    os.system('cls' if os.name == 'nt' else 'clear')
`)

		if g.config.Features.Memory {
			content.WriteString(`                elif user_input.lower().startswith('mem'):
                    self.handle_memory_commands(user_input)
`)
		}

		if g.config.Features.History {
			content.WriteString(`                elif user_input.lower().startswith('hist'):
                    self.handle_history_commands(user_input)
`)
		}

		content.WriteString(`                else:
                    result = self.evaluate_expression(user_input)
                    formatted_result = self.format_result(result)
                    print(f"Result: {formatted_result}")
`)

		if g.config.Features.History {
			content.WriteString(`                    self.history.add_entry(user_input, formatted_result)
`)
		}

		content.WriteString(`
            except KeyboardInterrupt:
                print("\nGoodbye!")
                break
            except Exception as e:
                print(f"Error: {e}")
`)
	}

	if g.config.UI.ShowBanner {
		content.WriteString(`
    def show_banner(self):
        """Display calculator banner"""
        print("="*50)
        print(f"  ` + g.config.ProjectName + `")
        print(f"  ` + g.config.Description + `")
        print("="*50)
`)
	}

	if g.config.UI.ShowHelp {
		content.WriteString(`
    def show_help(self):
        """Display help information"""
        help_text = """
Available commands:
  Basic operations: +, -, *, /, **, %
  Functions: sin(), cos(), tan(), log(), ln(), sqrt()
`)

		if g.config.Features.Memory {
			content.WriteString(`  Memory: mem store <value>, mem recall, mem clear
`)
		}

		if g.config.Features.History {
			content.WriteString(`  History: hist show, hist clear, hist save
`)
		}

		content.WriteString(`  Other: help, clear, quit
        """
        print(help_text)
`)
	}

	content.WriteString(`
    def evaluate_expression(self, expression):
        """Evaluate mathematical expression"""
        # Basic expression evaluation
        try:
            # Replace common functions
            expression = expression.replace('^', '**')
            `)

	if g.config.Libraries.UseMath {
		content.WriteString(`
            # Add math functions to evaluation context
            safe_dict = {
                "__builtins__": {},
                "abs": abs, "round": round, "min": min, "max": max,
                "sqrt": math.sqrt, "pi": math.pi, "e": math.e
            }`)

		// Add trigonometric functions if enabled
		if g.config.Features.Trigonometric {
			content.WriteString(`
            safe_dict.update({
                "sin": lambda x: sin(x, self.angle_unit),
                "cos": lambda x: cos(x, self.angle_unit),
                "tan": lambda x: tan(x, self.angle_unit),
                "asin": lambda x: asin(x, self.angle_unit),
                "acos": lambda x: acos(x, self.angle_unit),
                "atan": lambda x: atan(x, self.angle_unit)
            })`)
		}

		// Add logarithmic functions if enabled
		if g.config.Features.Logarithmic {
			content.WriteString(`
            safe_dict.update({
                "log": log, "ln": ln, "log10": log10, "log2": log2
            })`)
		}
	} else {
		content.WriteString(`
            safe_dict = {
                "__builtins__": {},
                "abs": abs, "round": round, "min": min, "max": max
            }`)
	}

	if g.config.Libraries.UseNumpy {
		content.WriteString(`
            safe_dict.update({
                "np": np, "array": np.array, "mean": mean,
                "median": median, "std": std
            })`)
	}

	content.WriteString(`

            return eval(expression, safe_dict)
        except Exception as e:
            raise ValueError(f"Invalid expression: {e}")
`)

	if g.config.Features.Memory {
		content.WriteString(`
    def handle_memory_commands(self, command):
        """Handle memory-related commands"""
        parts = command.split()
        if len(parts) < 2:
            print("Memory commands: mem store <value>, mem recall, mem clear")
            return

        action = parts[1].lower()
        if action == "store" and len(parts) > 2:
            try:
                value = float(parts[2])
                print(self.memory.store(value))
            except ValueError:
                print("Invalid value for memory storage")
        elif action == "recall":
            print(f"Memory: {self.memory.recall()}")
        elif action == "clear":
            print(self.memory.clear())
        else:
            print("Unknown memory command")
`)
	}

	if g.config.Features.History {
		content.WriteString(`
    def handle_history_commands(self, command):
        """Handle history-related commands"""
        parts = command.split()
        if len(parts) < 2:
            print("History commands: hist show, hist clear, hist save")
            return

        action = parts[1].lower()
        if action == "show":
            count = 10
            if len(parts) > 2:
                try:
                    count = int(parts[2])
                except ValueError:
                    pass
            history = self.history.get_history(count)
            for entry in history:
                print(f"{entry['timestamp']}: {entry['operation']} = {entry['result']}")
        elif action == "clear":
            print(self.history.clear_history())
        elif action == "save":
            filename = "calculator_history.json"
            if len(parts) > 2:
                filename = parts[2]
            print(self.history.save_to_file(filename))
        else:
            print("Unknown history command")
`)
	}

	return content.String()
}

// renderTemplate renders the calculator template with the given data
func (g *Generator) renderTemplate(data TemplateData) (string, error) {
	tmpl := `#!/usr/bin/env python3
"""
{{.Config.ProjectName}}
{{.Config.Description}}

Generated by Calculator Generator
Author: {{.Config.Author}}
Version: {{.Version}}
Generated: {{.Timestamp}}
"""

{{range .Imports}}{{.}}
{{end}}

{{range .Functions}}
{{.}}

{{end}}

{{.MainContent}}

if __name__ == "__main__":
    calculator = Calculator()
    calculator.run()
`

	t, err := template.New("calculator").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	if err := t.Execute(&result, data); err != nil {
		return "", err
	}

	return result.String(), nil
}

// writeToFile writes the generated content to the specified file
func (g *Generator) writeToFile(content string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(g.config.OutputFile)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Write the file
	return os.WriteFile(g.config.OutputFile, []byte(content), 0755)
}

// generateRequirements creates a requirements.txt file based on enabled libraries
func (g *Generator) generateRequirements() error {
	var requirements []string

	if g.config.Libraries.UseNumpy {
		requirements = append(requirements, "numpy>=1.21.0")
	}

	if g.config.Libraries.UsePandas {
		requirements = append(requirements, "pandas>=1.3.0")
	}

	if g.config.Libraries.UseScipy {
		requirements = append(requirements, "scipy>=1.7.0")
	}

	if g.config.Libraries.UseSympy {
		requirements = append(requirements, "sympy>=1.9.0")
	}

	if g.config.Libraries.UsePlotly || g.config.Features.Plotting {
		requirements = append(requirements, "plotly>=5.0.0")
	}

	// Add tkinter note for GUI calculators
	if g.config.UI.Style == "gui" {
		requirements = append(requirements, "# tkinter (included with Python)")
	}

	if len(requirements) == 0 {
		return nil // No requirements file needed
	}

	// Create requirements.txt in the same directory as the output file
	dir := filepath.Dir(g.config.OutputFile)
	requirementsPath := filepath.Join(dir, "requirements.txt")

	content := strings.Join(requirements, "\n") + "\n"
	return os.WriteFile(requirementsPath, []byte(content), 0644)
}
