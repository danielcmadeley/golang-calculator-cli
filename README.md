# Calculator Generator

A powerful CLI tool written in Go that generates customizable Python calculators with advanced mathematical capabilities.

## 🚀 Features

- **Two Calculator Types**: Basic and Scientific calculators
- **Extensive Library Support**: NumPy, Pandas, SciPy, SymPy, Plotly integration
- **Rich Feature Set**: Memory, history, plotting, equation solving, and more
- **Interactive Wizard**: Guided calculator creation process
- **Highly Customizable**: Configure UI, precision, libraries, and features
- **Professional Output**: Generated Python code with proper documentation

## 📦 Installation

### Prerequisites

- Go 1.21 or higher
- Python 3.8+ (for running generated calculators)

### Build from Source

```bash
git clone <repository-url>
cd golang-cli
go mod tidy
go build -o calculator-generator
```

### Install Dependencies (for generated calculators)

The tool automatically generates a `requirements.txt` file. Install Python dependencies:

```bash
pip install -r requirements.txt
```

## 🎯 Quick Start

### Generate a Basic Calculator

```bash
./calculator-generator generate --type basic --output my_calculator.py
python my_calculator.py
```

### Generate a Scientific Calculator

```bash
./calculator-generator generate --type scientific --output scientific_calc.py
python scientific_calc.py
```

### Interactive Mode

```bash
./calculator-generator interactive
```

## 📋 Command Reference

### Generate Command

Create a calculator with specific configuration:

```bash
calculator-generator generate [flags]
```

#### Flags

**Calculator Type:**
- `--type, -t`: Calculator type (`basic`, `scientific`)

**Project Information:**
- `--name, -n`: Project name
- `--author, -a`: Author name
- `--description, -d`: Project description
- `--output, -o`: Output file path

**Libraries:**
- `--libraries`: Comma-separated list (`numpy,pandas,scipy,sympy,plotly`)
- `--math`: Include math library (default: true)

**Features:**
- `--features`: Comma-separated feature list
- `--memory`: Include memory functionality
- `--history`: Include calculation history
- `--interactive`: Create interactive calculator (default: true)

**UI Configuration:**
- `--style`: UI style (`cli`, `gui`, `web`)
- `--theme`: UI theme (`light`, `dark`, `colorful`)
- `--precision`: Decimal precision (1-20, default: 10)
- `--angle-unit`: Angle unit (`degrees`, `radians`)
- `--show-help`: Show help information (default: true)
- `--show-banner`: Show application banner (default: true)

### Interactive Command

Launch the interactive wizard:

```bash
calculator-generator interactive
```

### List Commands

Explore available options:

```bash
calculator-generator list features    # List all features
calculator-generator list libraries   # List all libraries
calculator-generator list types       # List calculator types
calculator-generator list examples    # Show example commands
```

## 🔧 Available Features

### Basic Features
- `basic-arithmetic` - Addition, subtraction, multiplication, division
- `memory` - Store and recall values (M+, MR, MC)
- `history` - Keep track of calculation history

### Scientific Features
- `trigonometric` - sin, cos, tan, asin, acos, atan
- `logarithmic` - log, ln, log10, log2
- `exponential` - exp, power functions
- `complex-numbers` - Complex number arithmetic

### Statistical Features
- `statistical` - mean, median, std dev, variance
- `data-analysis` - Advanced data manipulation

### Advanced Mathematical Features
- `linear-algebra` - Matrix operations, eigenvalues
- `calculus` - Derivatives, integrals
- `equation-solver` - Solve algebraic equations
- `matrix-operations` - Advanced matrix functions

### Visualization Features
- `plotting` - Create 2D plots and charts
- `graphing` - Interactive graphing capabilities

### Utility Features
- `unit-conversion` - Convert between different units
- `programming` - Programmable calculator functions

## 📚 Supported Libraries

### Standard Library
- **math** - Basic mathematical functions (always recommended)

### Third-Party Libraries
- **numpy** - Numerical computing (required for statistical, linear-algebra features)
- **pandas** - Data manipulation (required for data-analysis)
- **scipy** - Scientific computing (advanced statistical functions)
- **sympy** - Symbolic mathematics (required for equation-solver, calculus)
- **plotly** - Interactive plotting (required for plotting, graphing)

## 💡 Examples

### Basic Calculator with Memory

```bash
calculator-generator generate \
  --type basic \
  --features "memory,history" \
  --name "Basic Calculator with Memory" \
  --output basic_calc.py
```

### Scientific Calculator with Plotting

```bash
calculator-generator generate \
  --type scientific \
  --libraries "numpy,plotly,sympy" \
  --features "trigonometric,plotting,equation-solver" \
  --precision 15 \
  --output sci_calc.py
```

### Data Analysis Calculator

```bash
calculator-generator generate \
  --type basic \
  --libraries "numpy,pandas" \
  --features "statistical,data-analysis" \
  --name "Data Analysis Calculator" \
  --output data_calc.py
```

### High-Precision Engineering Calculator

```bash
calculator-generator generate \
  --type scientific \
  --precision 20 \
  --angle-unit radians \
  --features "trigonometric,logarithmic,linear-algebra,complex-numbers" \
  --name "Engineering Calculator" \
  --author "Your Name" \
  --output engineering_calc.py
```

## 🎮 Using Generated Calculators

### Basic Usage

```python
# Run the calculator
python calculator.py

# Interactive mode
calc> 2 + 3 * 4
Result: 14

calc> sin(30)
Result: 0.5

calc> help
# Shows available commands

calc> quit
```

### Memory Commands

```python
calc> mem store 42
Stored 42 in memory

calc> mem recall
Memory: 42

calc> mem clear
Memory cleared
```

### History Commands

```python
calc> hist show
# Shows recent calculations

calc> hist show 20
# Shows last 20 calculations

calc> hist save my_history.json
History saved to my_history.json

calc> hist clear
History cleared
```

### Advanced Features (Scientific Calculator)

```python
# Equation solving
calc> solve_equation("x**2 - 4", "x")
Result: [-2, 2]

# Plotting
calc> plot_function("x**2", (-5, 5))
# Opens interactive plot

# Matrix operations
calc> matrix_multiply([[1,2],[3,4]], [[5,6],[7,8]])
Result: [[19, 22], [43, 50]]
```

## 🛠️ Development

### Project Structure

```
golang-cli/
├── main.go              # Entry point
├── go.mod               # Go module file
├── cmd/                 # CLI commands
│   ├── root.go         # Root command
│   ├── generate.go     # Generate command
│   ├── interactive.go  # Interactive wizard
│   └── list.go         # List commands
├── internal/           # Internal packages
│   ├── types.go        # Type definitions
│   └── generator.go    # Calculator generator
└── templates/          # Template files (future use)
```

### Building

```bash
# Build for current platform
go build -o calculator-generator

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o calculator-generator-linux
GOOS=windows GOARCH=amd64 go build -o calculator-generator-windows.exe
GOOS=darwin GOARCH=amd64 go build -o calculator-generator-macos
```

### Testing

```bash
# Run tests
go test ./...

# Test generated calculator
./calculator-generator generate --type basic --output test_calc.py
python test_calc.py
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Adding New Features

To add a new calculator feature:

1. Update the `Features` struct in `internal/types.go`
2. Add feature generation logic in `internal/generator.go`
3. Update the feature list in `cmd/list.go`
4. Add flag handling in `cmd/generate.go` and `cmd/interactive.go`

### Adding New Libraries

To add support for a new Python library:

1. Update the `Libraries` struct in `internal/types.go`
2. Add import generation logic in `internal/generator.go`
3. Update the library list in `cmd/list.go`
4. Add dependency handling in the requirements generation

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Viper](https://github.com/spf13/viper) - Configuration management
- NumPy, SciPy, SymPy, Pandas, Plotly - Python scientific libraries

## 📞 Support

- Create an issue for bug reports or feature requests
- Check the [examples](#-examples) section for common use cases
- Use `calculator-generator list examples` for command examples
- Use `calculator-generator interactive` for guided setup

---

**Happy Calculating!** 🧮✨# golang-calculator-cli
