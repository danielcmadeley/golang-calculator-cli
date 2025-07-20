# Calculator Generator - Usage Guide

## Table of Contents
1. [Getting Started](#getting-started)
2. [Command Line Interface](#command-line-interface)
3. [Calculator Types](#calculator-types)
4. [Features and Capabilities](#features-and-capabilities)
5. [Library Dependencies](#library-dependencies)
6. [Configuration Options](#configuration-options)
7. [Examples and Use Cases](#examples-and-use-cases)
8. [Using Generated Calculators](#using-generated-calculators)
9. [Troubleshooting](#troubleshooting)
10. [Advanced Usage](#advanced-usage)

## Getting Started

### Installation and Setup

1. **Build the CLI tool:**
   ```bash
   git clone <repository-url>
   cd golang-cli
   go mod tidy
   go build -o calculator-generator
   ```

2. **Verify installation:**
   ```bash
   ./calculator-generator --help
   ```

3. **Quick start - Generate your first calculator:**
   ```bash
   ./calculator-generator generate --type basic --output my_calculator.py
   python my_calculator.py
   ```

### Basic Workflow

1. **Choose your approach:**
   - **Quick generation:** Use `generate` command with flags
   - **Guided setup:** Use `interactive` command for step-by-step wizard
   - **Explore options:** Use `list` commands to see available features

2. **Generate calculator:**
   ```bash
   ./calculator-generator generate [options]
   ```

3. **Install dependencies (if needed):**
   ```bash
   pip install -r requirements.txt
   ```

4. **Run your calculator:**
   ```bash
   python your_calculator.py
   ```

## Command Line Interface

### Main Commands

#### `generate` - Create a calculator directly
```bash
calculator-generator generate [flags]
```

**Key Flags:**
- `--type`: Calculator type (`basic`, `scientific`)
- `--output`: Output file path
- `--features`: Comma-separated feature list
- `--libraries`: Comma-separated library list
- `--name`: Project name
- `--author`: Author name
- `--precision`: Decimal precision (1-20)

#### `interactive` - Guided wizard
```bash
calculator-generator interactive
```
Walk through a step-by-step wizard to configure your calculator.

#### `list` - Explore options
```bash
calculator-generator list features    # Show all features
calculator-generator list libraries   # Show all libraries
calculator-generator list types       # Show calculator types
calculator-generator list examples    # Show example commands
```

### Global Flags
- `-o, --output`: Output file path (default: "calculator.py")
- `-a, --author`: Author name (default: "Calculator Generator")
- `--config`: Config file path
- `--verbose`: Enable verbose output

## Calculator Types

### Basic Calculator
- **Purpose:** Simple arithmetic operations
- **Best for:** Learning, basic calculations, minimal dependencies
- **Default features:** Basic arithmetic (+, -, *, /, %, **)
- **Dependencies:** Python standard library only

```bash
./calculator-generator generate --type basic
```

### Scientific Calculator
- **Purpose:** Advanced mathematical computations
- **Best for:** Engineering, science, research, education
- **Default features:** All basic features plus trigonometric, logarithmic, statistical functions
- **Dependencies:** NumPy, SciPy, SymPy

```bash
./calculator-generator generate --type scientific
```

## Features and Capabilities

### Basic Features
| Feature | Description | Command Flag | Dependencies |
|---------|-------------|--------------|--------------|
| `basic-arithmetic` | +, -, *, /, %, ** | `arithmetic` | None |
| `memory` | Store/recall values | `memory` | None |
| `history` | Calculation history | `history` | None |

### Scientific Features
| Feature | Description | Command Flag | Dependencies |
|---------|-------------|--------------|--------------|
| `trigonometric` | sin, cos, tan, etc. | `trigonometric` | math |
| `logarithmic` | log, ln, log10, log2 | `logarithmic` | math |
| `exponential` | exp, power functions | `exponential` | math |
| `complex-numbers` | Complex arithmetic | `complex-numbers` | None |

### Statistical Features
| Feature | Description | Command Flag | Dependencies |
|---------|-------------|--------------|--------------|
| `statistical` | mean, median, std | `statistical` | numpy |
| `data-analysis` | Advanced data ops | `data-analysis` | pandas |

### Advanced Features
| Feature | Description | Command Flag | Dependencies |
|---------|-------------|--------------|--------------|
| `linear-algebra` | Matrix operations | `linear-algebra` | numpy |
| `equation-solver` | Solve equations | `equation-solver` | sympy |
| `plotting` | Create graphs | `plotting` | plotly |
| `calculus` | Derivatives, integrals | `calculus` | sympy |

### Usage Examples:
```bash
# Single feature
./calculator-generator generate --features "memory"

# Multiple features
./calculator-generator generate --features "memory,history,trigonometric"

# Scientific with specific features
./calculator-generator generate --type scientific --features "plotting,equation-solver"
```

## Library Dependencies

### Standard Library
- **math**: Basic mathematical functions (always recommended)

### Third-Party Libraries
- **numpy**: Numerical computing, arrays, mathematical functions
- **pandas**: Data manipulation and analysis
- **scipy**: Scientific computing, optimization, integration
- **sympy**: Symbolic mathematics, equation solving
- **plotly**: Interactive plotting and visualization

### Automatic Dependency Management
The tool automatically:
1. Generates `requirements.txt` with needed libraries
2. Includes proper import statements
3. Provides installation instructions

```bash
# Example: Install generated dependencies
pip install -r requirements.txt
```

## Configuration Options

### Project Information
```bash
--name "My Calculator"              # Project name
--author "John Doe"                 # Author name
--description "Custom calculator"   # Description
--output "custom_calc.py"          # Output file
```

### UI Configuration
```bash
--precision 15                 # Decimal places (1-20)
--angle-unit radians          # degrees or radians
--style cli                   # UI style (cli, gui, web)
--theme dark                  # light, dark, colorful
--show-banner true            # Show application banner
--show-help true              # Show help information
```

### Behavior Options
```bash
--interactive true            # Interactive mode
--memory true                 # Enable memory functions
--history true                # Enable history tracking
```

## Examples and Use Cases

### 1. Student Calculator
Perfect for students learning mathematics:
```bash
./calculator-generator generate \
  --type basic \
  --features "memory,history,trigonometric" \
  --name "Student Calculator" \
  --precision 10 \
  --output student_calc.py
```

### 2. Engineering Calculator
High-precision calculator for engineering work:
```bash
./calculator-generator generate \
  --type scientific \
  --features "trigonometric,logarithmic,linear-algebra,complex-numbers" \
  --precision 15 \
  --angle-unit radians \
  --name "Engineering Calculator" \
  --output engineering_calc.py
```

### 3. Data Analysis Calculator
Optimized for statistical and data analysis tasks:
```bash
./calculator-generator generate \
  --type basic \
  --libraries "numpy,pandas" \
  --features "statistical,data-analysis" \
  --name "Data Analysis Calculator" \
  --output data_calc.py
```

### 4. Graphing Calculator
For visualization and equation solving:
```bash
./calculator-generator generate \
  --type scientific \
  --libraries "numpy,plotly,sympy" \
  --features "plotting,equation-solver,graphing" \
  --name "Graphing Calculator" \
  --output graphing_calc.py
```

### 5. Research Calculator
Full-featured calculator for research work:
```bash
./calculator-generator generate \
  --type scientific \
  --libraries "numpy,scipy,sympy,plotly" \
  --features "trigonometric,statistical,calculus,plotting,equation-solver" \
  --precision 20 \
  --name "Research Calculator" \
  --output research_calc.py
```

## Using Generated Calculators

### Interactive Mode
```python
# Start the calculator
python my_calculator.py

# Basic calculations
calc> 2 + 3 * 4
Result: 14

calc> sqrt(16)
Result: 4.0

# Scientific functions (if enabled)
calc> sin(30)        # degrees by default
Result: 0.5

calc> log(100)
Result: 2.0
```

### Memory Operations (if enabled)
```python
calc> 42
Result: 42

calc> mem store 42
Stored 42 in memory

calc> 5 + 3
Result: 8

calc> mem recall
Memory: 42

calc> mem clear
Memory cleared
```

### History Operations (if enabled)
```python
calc> hist show
2024-07-20T15:30:15: 2 + 3 = 5
2024-07-20T15:30:20: sqrt(16) = 4.0

calc> hist show 5        # Show last 5 entries
calc> hist clear         # Clear history
calc> hist save my_history.json  # Save to file
```

### Advanced Functions (Scientific calculators)
```python
# Equation solving (SymPy required)
calc> solve_equation("x**2 - 4", "x")
Result: [-2, 2]

# Matrix operations (NumPy required)
calc> matrix_multiply([[1,2],[3,4]], [[5,6],[7,8]])
Result: [[19, 22], [43, 50]]

# Statistical functions (NumPy required)
calc> mean([1, 2, 3, 4, 5])
Result: 3.0

calc> std([1, 2, 3, 4, 5])
Result: 1.4142135623730951
```

### Built-in Commands
- `help`: Show available commands
- `clear`: Clear screen
- `quit` or `exit`: Exit calculator
- `mem store <value>`: Store value in memory
- `mem recall`: Recall memory value
- `mem clear`: Clear memory
- `hist show [count]`: Show calculation history
- `hist clear`: Clear history
- `hist save [filename]`: Save history to file

## Troubleshooting

### Common Issues and Solutions

#### 1. Calculator fails to start
**Error:** `ModuleNotFoundError: No module named 'numpy'`
**Solution:** Install required dependencies:
```bash
pip install -r requirements.txt
```

#### 2. Function not recognized
**Error:** `NameError: name 'sin' is not defined`
**Cause:** Feature not enabled during generation
**Solution:** Regenerate with the required feature:
```bash
./calculator-generator generate --features "trigonometric" --output new_calc.py
```

#### 3. Import errors in generated code
**Cause:** Library not properly included
**Solution:** Check and regenerate with correct libraries:
```bash
./calculator-generator generate --libraries "numpy,scipy" --features "statistical"
```

#### 4. Precision issues
**Cause:** Default precision too low
**Solution:** Generate with higher precision:
```bash
./calculator-generator generate --precision 15
```

#### 5. Angle unit confusion
**Issue:** Getting unexpected trigonometric results
**Solution:** Check angle unit setting:
```bash
# For degrees (default)
./calculator-generator generate --angle-unit degrees

# For radians
./calculator-generator generate --angle-unit radians
```

### Debug Mode
Use verbose output to diagnose issues:
```bash
./calculator-generator generate --verbose --type scientific
```

### Validation Errors
The tool validates configuration and shows helpful error messages:
```bash
# Example validation error
Error: configuration validation failed: precision: precision must be between 1 and 20
```

## Advanced Usage

### Configuration Files
Create a configuration file for repeated use:

**`.calculator-generator.yaml`:**
```yaml
type: scientific
author: "Your Name"
precision: 15
angle_unit: radians
libraries:
  use_numpy: true
  use_scipy: true
  use_sympy: true
features:
  trigonometric: true
  logarithmic: true
  statistical: true
  equation_solver: true
ui:
  show_banner: true
  show_help: true
  theme: dark
```

Use with:
```bash
./calculator-generator generate --config .calculator-generator.yaml
```

### Batch Generation
Generate multiple calculators for different purposes:

**`generate_all.sh`:**
```bash
#!/bin/bash

# Basic calculator
./calculator-generator generate --type basic --output calculators/basic.py

# Student calculator
./calculator-generator generate \
  --type basic \
  --features "memory,trigonometric" \
  --output calculators/student.py

# Engineering calculator
./calculator-generator generate \
  --type scientific \
  --precision 15 \
  --angle-unit radians \
  --output calculators/engineering.py

# Data analysis calculator
./calculator-generator generate \
  --libraries "numpy,pandas" \
  --features "statistical,data-analysis" \
  --output calculators/data.py
```

### Custom Templates (Future Feature)
The tool is designed to support custom templates for specialized calculator types.

### Integration with IDEs
Generated calculators work well with Python IDEs and can be extended with additional functionality.

### Performance Considerations
- Basic calculators start instantly
- Scientific calculators may take 1-2 seconds to load dependencies
- Memory usage depends on enabled libraries (NumPy, SciPy, etc.)

### Security Notes
- Generated calculators use `eval()` with restricted context
- Only safe mathematical operations are allowed
- No file system or network access in evaluation context

## Best Practices

1. **Start Simple:** Begin with basic calculator, add features as needed
2. **Match Purpose:** Choose features that match your use case
3. **Consider Dependencies:** More libraries = more capabilities but larger dependencies
4. **Test Generation:** Always test generated calculators before deployment
5. **Version Control:** Keep track of generation commands for reproducibility
6. **Documentation:** Document any customizations made to generated code

## Getting Help

- Use `--help` with any command for detailed information
- Use `list` commands to explore available options
- Use `interactive` mode for guided setup
- Check generated code for inline documentation
- Review this guide for comprehensive usage information

---

**Happy Calculating!** 🧮✨