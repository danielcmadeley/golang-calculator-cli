package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available features and libraries",
	Long: `List all available features, libraries, and options that can be used
when generating a calculator.

This command helps you discover what's available and understand the
dependencies between features and libraries.

Examples:
  calculator-generator list features
  calculator-generator list libraries
  calculator-generator list types`,
}

// listFeaturesCmd lists all available features
var listFeaturesCmd = &cobra.Command{
	Use:   "features",
	Short: "List all available calculator features",
	Long:  `Display a comprehensive list of all features that can be included in your calculator.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📋 Available Calculator Features")
		fmt.Println("================================")
		fmt.Println()

		fmt.Println("🔢 Basic Features:")
		fmt.Println("  • basic-arithmetic    - Addition, subtraction, multiplication, division")
		fmt.Println("  • memory             - Store and recall values (M+, MR, MC)")
		fmt.Println("  • history            - Keep track of calculation history")
		fmt.Println()

		fmt.Println("🧮 Scientific Features:")
		fmt.Println("  • trigonometric      - sin, cos, tan, asin, acos, atan")
		fmt.Println("  • logarithmic        - log, ln, log10, log2")
		fmt.Println("  • exponential        - exp, power functions")
		fmt.Println("  • complex-numbers    - Complex number arithmetic")
		fmt.Println()

		fmt.Println("📊 Statistical Features:")
		fmt.Println("  • statistical        - mean, median, std dev, variance")
		fmt.Println("  • data-analysis      - Advanced data manipulation")
		fmt.Println()

		fmt.Println("🔬 Advanced Mathematical Features:")
		fmt.Println("  • linear-algebra     - Matrix operations, eigenvalues")
		fmt.Println("  • calculus           - Derivatives, integrals")
		fmt.Println("  • equation-solver    - Solve algebraic equations")
		fmt.Println("  • matrix-operations  - Advanced matrix functions")
		fmt.Println()

		fmt.Println("📈 Visualization Features:")
		fmt.Println("  • plotting           - Create 2D plots and charts")
		fmt.Println("  • graphing           - Interactive graphing capabilities")
		fmt.Println()

		fmt.Println("🔧 Utility Features:")
		fmt.Println("  • unit-conversion    - Convert between different units")
		fmt.Println("  • programming        - Programmable calculator functions")
		fmt.Println()

		fmt.Println("💡 Usage:")
		fmt.Println("  Use these feature names with the --features flag:")
		fmt.Println("  calculator-generator generate --features \"trigonometric,logarithmic,plotting\"")
	},
}

// listLibrariesCmd lists all available libraries
var listLibrariesCmd = &cobra.Command{
	Use:   "libraries",
	Short: "List all available Python libraries",
	Long:  `Display information about Python libraries that can be included in your calculator.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📚 Available Python Libraries")
		fmt.Println("=============================")
		fmt.Println()

		fmt.Println("🔢 Standard Library:")
		fmt.Println("  • math               - Basic mathematical functions")
		fmt.Println("                        Functions: sin, cos, tan, log, sqrt, pi, e")
		fmt.Println("                        Always recommended for scientific calculators")
		fmt.Println()

		fmt.Println("🧮 Numerical Computing:")
		fmt.Println("  • numpy              - Numerical computing with Python")
		fmt.Println("                        Features: Arrays, mathematical functions, linear algebra")
		fmt.Println("                        Required for: statistical, linear-algebra, matrix-operations")
		fmt.Println("                        Install: pip install numpy>=1.21.0")
		fmt.Println()

		fmt.Println("📊 Data Analysis:")
		fmt.Println("  • pandas             - Data manipulation and analysis")
		fmt.Println("                        Features: DataFrames, data import/export, statistics")
		fmt.Println("                        Required for: data-analysis")
		fmt.Println("                        Install: pip install pandas>=1.3.0")
		fmt.Println()

		fmt.Println("🔬 Scientific Computing:")
		fmt.Println("  • scipy              - Scientific computing library")
		fmt.Println("                        Features: Optimization, integration, interpolation")
		fmt.Println("                        Required for: advanced statistical functions")
		fmt.Println("                        Install: pip install scipy>=1.7.0")
		fmt.Println()

		fmt.Println("🔣 Symbolic Mathematics:")
		fmt.Println("  • sympy              - Symbolic mathematics")
		fmt.Println("                        Features: Algebraic manipulation, calculus, equation solving")
		fmt.Println("                        Required for: equation-solver, calculus")
		fmt.Println("                        Install: pip install sympy>=1.9.0")
		fmt.Println()

		fmt.Println("📈 Visualization:")
		fmt.Println("  • plotly             - Interactive plotting library")
		fmt.Println("                        Features: 2D/3D plots, interactive charts, web-based visualization")
		fmt.Println("                        Required for: plotting, graphing")
		fmt.Println("                        Install: pip install plotly>=5.0.0")
		fmt.Println()

		fmt.Println("💡 Usage:")
		fmt.Println("  Use these library names with the --libraries flag:")
		fmt.Println("  calculator-generator generate --libraries \"numpy,scipy,plotly\"")
	},
}

// listTypesCmd lists available calculator types
var listTypesCmd = &cobra.Command{
	Use:   "types",
	Short: "List available calculator types",
	Long:  `Display information about the different types of calculators you can generate.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔢 Available Calculator Types")
		fmt.Println("============================")
		fmt.Println()

		fmt.Println("📱 Basic Calculator:")
		fmt.Println("  • Type: basic")
		fmt.Println("  • Description: Simple arithmetic calculator")
		fmt.Println("  • Default Features:")
		fmt.Println("    - Basic arithmetic operations (+, -, *, /, %, **)")
		fmt.Println("    - Standard math library functions")
		fmt.Println("    - Interactive command-line interface")
		fmt.Println("  • Best For: Simple calculations, learning Python")
		fmt.Println("  • Dependencies: Python standard library only")
		fmt.Println()

		fmt.Println("🧮 Scientific Calculator:")
		fmt.Println("  • Type: scientific")
		fmt.Println("  • Description: Advanced mathematical calculator")
		fmt.Println("  • Default Features:")
		fmt.Println("    - All basic calculator features")
		fmt.Println("    - Trigonometric functions (sin, cos, tan)")
		fmt.Println("    - Logarithmic functions (log, ln)")
		fmt.Println("    - Statistical functions (mean, std dev)")
		fmt.Println("    - Linear algebra operations")
		fmt.Println("    - Complex number support")
		fmt.Println("    - Memory and history functionality")
		fmt.Println("  • Best For: Engineering, science, advanced mathematics")
		fmt.Println("  • Dependencies: numpy, scipy, sympy")
		fmt.Println()

		fmt.Println("🎨 Interface Options:")
		fmt.Println("  Both calculator types support multiple interfaces!")
		fmt.Println("  • CLI: Command-line interface for terminal use")
		fmt.Println("  • GUI: Desktop application with buttons and display")
		fmt.Println()

		fmt.Println("🎨 Customization:")
		fmt.Println("  Both calculator types are highly customizable!")
		fmt.Println("  • Add or remove features as needed")
		fmt.Println("  • Choose which libraries to include")
		fmt.Println("  • Customize UI appearance and behavior")
		fmt.Println("  • Set precision, angle units, and more")
		fmt.Println()

		fmt.Println("💡 Usage:")
		fmt.Println("  calculator-generator generate --type basic")
		fmt.Println("  calculator-generator generate --type scientific")
		fmt.Println("  calculator-generator generate --style gui  # For desktop GUI")
		fmt.Println("  calculator-generator interactive  # For guided selection")
	},
}

// listExamplesCmd shows example command combinations
var listExamplesCmd = &cobra.Command{
	Use:   "examples",
	Short: "Show example command combinations",
	Long:  `Display example commands showing different ways to use the calculator generator.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("💡 Example Commands")
		fmt.Println("==================")
		fmt.Println()

		fmt.Println("🔰 Basic Examples:")
		fmt.Println("  # Generate a simple basic calculator (CLI)")
		fmt.Println("  calculator-generator generate --type basic")
		fmt.Println()
		fmt.Println("  # Generate with custom output file")
		fmt.Println("  calculator-generator generate --type basic --output my_calc.py")
		fmt.Println()

		fmt.Println("🖥️ Desktop GUI Examples:")
		fmt.Println("  # Generate a basic GUI calculator")
		fmt.Println("  calculator-generator generate --type basic --style gui")
		fmt.Println()
		fmt.Println("  # Generate scientific GUI calculator")
		fmt.Println("  calculator-generator generate --type scientific --style gui")
		fmt.Println()
		fmt.Println("  # GUI calculator with memory and history")
		fmt.Println("  calculator-generator generate --style gui \\")
		fmt.Println("    --features \"memory,history,trigonometric\"")
		fmt.Println()

		fmt.Println("🧮 Scientific Examples:")
		fmt.Println("  # Generate a full scientific calculator")
		fmt.Println("  calculator-generator generate --type scientific")
		fmt.Println()
		fmt.Println("  # Scientific calculator with specific features")
		fmt.Println("  calculator-generator generate --type scientific \\")
		fmt.Println("    --features \"trigonometric,logarithmic,statistical\"")
		fmt.Println()

		fmt.Println("🎯 Custom Examples:")
		fmt.Println("  # Basic calculator with memory and history")
		fmt.Println("  calculator-generator generate --type basic \\")
		fmt.Println("    --features \"memory,history\" --libraries \"math\"")
		fmt.Println()
		fmt.Println("  # Data analysis calculator")
		fmt.Println("  calculator-generator generate --type basic \\")
		fmt.Println("    --libraries \"numpy,pandas\" \\")
		fmt.Println("    --features \"statistical,data-analysis\"")
		fmt.Println()
		fmt.Println("  # Graphing calculator")
		fmt.Println("  calculator-generator generate --type scientific \\")
		fmt.Println("    --libraries \"numpy,plotly\" \\")
		fmt.Println("    --features \"plotting,equation-solver\"")
		fmt.Println()

		fmt.Println("⚙️ Configuration Examples:")
		fmt.Println("  # High precision calculator with radians")
		fmt.Println("  calculator-generator generate --type scientific \\")
		fmt.Println("    --precision 15 --angle-unit radians")
		fmt.Println()
		fmt.Println("  # Custom GUI calculator with dark theme")
		fmt.Println("  calculator-generator generate --style gui --theme dark \\")
		fmt.Println("    --name \"My Calculator\" \\")
		fmt.Println("    --author \"John Doe\"")
		fmt.Println()

		fmt.Println("🎮 Interactive Mode:")
		fmt.Println("  # Launch interactive wizard")
		fmt.Println("  calculator-generator interactive")
		fmt.Println()

		fmt.Println("📋 Information Commands:")
		fmt.Println("  # List available features")
		fmt.Println("  calculator-generator list features")
		fmt.Println()
		fmt.Println("  # List available libraries")
		fmt.Println("  calculator-generator list libraries")
		fmt.Println()
		fmt.Println("  # Show calculator types")
		fmt.Println("  calculator-generator list types")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listFeaturesCmd)
	listCmd.AddCommand(listLibrariesCmd)
	listCmd.AddCommand(listTypesCmd)
	listCmd.AddCommand(listExamplesCmd)
}
