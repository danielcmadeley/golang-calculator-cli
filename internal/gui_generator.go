package internal

import (
	"fmt"
	"strings"
	"text/template"
)

// GUIGenerator handles the generation of GUI-based Python calculator applications
type GUIGenerator struct {
	config CalculatorConfig
}

// NewGUIGenerator creates a new GUI calculator generator instance
func NewGUIGenerator(config CalculatorConfig) *GUIGenerator {
	return &GUIGenerator{config: config}
}

// GenerateGUICalculator creates a Tkinter-based desktop calculator
func (g *GUIGenerator) GenerateGUICalculator() (string, error) {
	data := g.prepareGUITemplateData()
	return g.renderGUITemplate(data)
}

// prepareGUITemplateData prepares data for GUI template rendering
func (g *GUIGenerator) prepareGUITemplateData() TemplateData {
	imports := g.generateGUIImports()
	functions := g.generateGUIFunctions()
	mainContent := g.generateGUIMainContent()

	return TemplateData{
		Config:      g.config,
		Imports:     imports,
		Functions:   functions,
		MainContent: mainContent,
		Version:     "1.0.0",
	}
}

// generateGUIImports creates the list of Python imports for GUI calculator
func (g *GUIGenerator) generateGUIImports() []string {
	var imports []string

	// Standard library imports for GUI
	imports = append(imports, "import tkinter as tk")
	imports = append(imports, "from tkinter import ttk, messagebox, simpledialog")
	imports = append(imports, "import math")
	imports = append(imports, "import sys")
	imports = append(imports, "import os")

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

// generateGUIFunctions creates calculator function implementations for GUI
func (g *GUIGenerator) generateGUIFunctions() []string {
	var functions []string

	// Add all the mathematical functions from the original generator
	if g.config.Features.BasicArithmetic {
		functions = append(functions, g.generateBasicArithmeticFunctions()...)
	}

	if g.config.Features.Trigonometric {
		functions = append(functions, g.generateTrigonometricFunctions()...)
	}

	if g.config.Features.Logarithmic {
		functions = append(functions, g.generateLogarithmicFunctions()...)
	}

	if g.config.Features.Statistical && g.config.Libraries.UseNumpy {
		functions = append(functions, g.generateStatisticalFunctions()...)
	}

	if g.config.Features.Memory {
		functions = append(functions, g.generateMemoryClass())
	}

	if g.config.Features.History {
		functions = append(functions, g.generateHistoryClass())
	}

	return functions
}

// generateBasicArithmeticFunctions creates basic arithmetic functions for GUI
func (g *GUIGenerator) generateBasicArithmeticFunctions() []string {
	return []string{
		`def safe_eval(expression):
    """Safely evaluate mathematical expressions"""
    try:
        # Replace common symbols
        expression = expression.replace('^', '**')
        expression = expression.replace('×', '*')
        expression = expression.replace('÷', '/')

        # Create safe evaluation context
        safe_dict = {
            "__builtins__": {},
            "abs": abs, "round": round, "min": min, "max": max,
            "sqrt": math.sqrt, "pi": math.pi, "e": math.e,
            "sin": math.sin, "cos": math.cos, "tan": math.tan,
            "asin": math.asin, "acos": math.acos, "atan": math.atan,
            "log": math.log10, "ln": math.log, "log10": math.log10,
            "exp": math.exp, "pow": pow
        }

        return eval(expression, safe_dict)
    except Exception as e:
        raise ValueError(f"Invalid expression: {str(e)}")`,
	}
}

// generateTrigonometricFunctions creates trigonometric functions for GUI
func (g *GUIGenerator) generateTrigonometricFunctions() []string {
	return []string{
		`def deg_to_rad(degrees):
    """Convert degrees to radians"""
    return math.radians(degrees)

def rad_to_deg(radians):
    """Convert radians to degrees"""
    return math.degrees(radians)`,
	}
}

// generateLogarithmicFunctions creates logarithmic functions for GUI
func (g *GUIGenerator) generateLogarithmicFunctions() []string {
	return []string{
		`def safe_log(x, base=10):
    """Safe logarithm function"""
    if x <= 0:
        raise ValueError("Logarithm input must be positive")
    if base == math.e:
        return math.log(x)
    return math.log(x, base)`,
	}
}

// generateStatisticalFunctions creates statistical functions for GUI
func (g *GUIGenerator) generateStatisticalFunctions() []string {
	return []string{
		`def calculate_stats(data_str):
    """Calculate statistics from comma-separated data"""
    try:
        data = [float(x.strip()) for x in data_str.split(',')]
        return {
            'mean': np.mean(data),
            'median': np.median(data),
            'std': np.std(data),
            'var': np.var(data),
            'min': np.min(data),
            'max': np.max(data)
        }
    except Exception as e:
        raise ValueError(f"Invalid data format: {str(e)}")`,
	}
}

// generateMemoryClass creates memory functionality for GUI
func (g *GUIGenerator) generateMemoryClass() string {
	return `class MemoryManager:
    """Handles calculator memory operations"""
    def __init__(self):
        self.memory = 0

    def store(self, value):
        """Store value in memory"""
        self.memory = float(value)

    def recall(self):
        """Recall value from memory"""
        return self.memory

    def clear(self):
        """Clear memory"""
        self.memory = 0

    def add(self, value):
        """Add value to memory"""
        self.memory += float(value)

    def subtract(self, value):
        """Subtract value from memory"""
        self.memory -= float(value)`
}

// generateHistoryClass creates history functionality for GUI
func (g *GUIGenerator) generateHistoryClass() string {
	return `class HistoryManager:
    """Handles calculation history"""
    def __init__(self, max_entries=100):
        self.history = []
        self.max_entries = max_entries

    def add_entry(self, expression, result):
        """Add calculation to history"""
        entry = {
            'timestamp': datetime.now().strftime('%H:%M:%S'),
            'expression': expression,
            'result': str(result)
        }
        self.history.append(entry)

        # Keep only max_entries
        if len(self.history) > self.max_entries:
            self.history = self.history[-self.max_entries:]

    def get_history(self):
        """Get all history entries"""
        return self.history

    def clear(self):
        """Clear history"""
        self.history = []

    def save_to_file(self, filename):
        """Save history to file"""
        with open(filename, 'w') as f:
            json.dump(self.history, f, indent=2)`
}

// generateGUIMainContent creates the main GUI calculator class
func (g *GUIGenerator) generateGUIMainContent() string {
	var content strings.Builder

	// Start of Calculator class
	content.WriteString(`class CalculatorGUI:
    """Main GUI Calculator Application"""

    def __init__(self):
        self.root = tk.Tk()
        self.root.title("` + g.config.ProjectName + `")
        self.root.geometry("400x600")
        self.root.resizable(False, False)

        # Calculator state
        self.display_var = tk.StringVar()
        self.display_var.set("0")
        self.current_expression = ""
        self.result_shown = False
        self.angle_unit = "` + g.config.UI.AngleUnit + `"
        self.precision = ` + fmt.Sprintf("%d", g.config.UI.Precision) + `

`)

	// Add memory and history if enabled
	if g.config.Features.Memory {
		content.WriteString("        self.memory = MemoryManager()\n")
	}

	if g.config.Features.History {
		content.WriteString("        self.history = HistoryManager()\n")
	}

	// Continue with GUI setup
	content.WriteString(`
        # Setup GUI
        self.create_widgets()
        self.setup_layout()
        self.setup_bindings()

        # Apply theme
        self.apply_theme()

    def create_widgets(self):
        """Create all GUI widgets"""
        # Main frame
        self.main_frame = ttk.Frame(self.root, padding="10")

        # Display frame
        self.display_frame = ttk.Frame(self.main_frame)
        self.display = ttk.Entry(
            self.display_frame,
            textvariable=self.display_var,
            justify='right',
            state='readonly'
        )

        # Expression display
        self.expr_var = tk.StringVar()
        self.expr_display = ttk.Label(
            self.display_frame,
            textvariable=self.expr_var
        )

        # Button frame
        self.button_frame = ttk.Frame(self.main_frame)

        # Create calculator buttons
        self.create_buttons()
`)

	// Add menu if advanced features are enabled
	if g.config.Features.Memory || g.config.Features.History || g.config.Features.Statistical {
		content.WriteString(`
        # Create menu
        self.create_menu()
`)
	}

	// Add button creation method
	content.WriteString(`
    def create_buttons(self):
        """Create calculator buttons"""
        # Button configuration
        button_config = {
            'width': 5
        }

        # Number buttons (0-9)
        self.num_buttons = {}
        for i in range(10):
            self.num_buttons[i] = ttk.Button(
                self.button_frame,
                text=str(i),
                command=lambda n=i: self.append_number(str(n)),
                **button_config
            )

        # Operator buttons
        self.op_buttons = {
            '+': ttk.Button(self.button_frame, text='+', command=lambda: self.append_operator('+'), **button_config),
            '-': ttk.Button(self.button_frame, text='-', command=lambda: self.append_operator('-'), **button_config),
            '*': ttk.Button(self.button_frame, text='×', command=lambda: self.append_operator('*'), **button_config),
            '/': ttk.Button(self.button_frame, text='÷', command=lambda: self.append_operator('/'), **button_config),
            '=': ttk.Button(self.button_frame, text='=', command=self.calculate, **button_config),
            '.': ttk.Button(self.button_frame, text='.', command=lambda: self.append_number('.'), **button_config),
            'C': ttk.Button(self.button_frame, text='C', command=self.clear, **button_config),
            'CE': ttk.Button(self.button_frame, text='CE', command=self.clear_entry, **button_config),
        }

        # Advanced function buttons`)

	// Add scientific function buttons if enabled
	if g.config.Features.Trigonometric {
		content.WriteString(`
        if True:  # Trigonometric functions
            self.trig_buttons = {
                'sin': ttk.Button(self.button_frame, text='sin', command=lambda: self.append_function('sin'), **button_config),
                'cos': ttk.Button(self.button_frame, text='cos', command=lambda: self.append_function('cos'), **button_config),
                'tan': ttk.Button(self.button_frame, text='tan', command=lambda: self.append_function('tan'), **button_config),
            }`)
	}

	if g.config.Features.Logarithmic {
		content.WriteString(`
        if True:  # Logarithmic functions
            self.log_buttons = {
                'log': ttk.Button(self.button_frame, text='log', command=lambda: self.append_function('log'), **button_config),
                'ln': ttk.Button(self.button_frame, text='ln', command=lambda: self.append_function('ln'), **button_config),
            }`)
	}

	// Add memory buttons if enabled
	if g.config.Features.Memory {
		content.WriteString(`
        if True:  # Memory functions
            self.mem_buttons = {
                'MS': ttk.Button(self.button_frame, text='MS', command=self.memory_store, **button_config),
                'MR': ttk.Button(self.button_frame, text='MR', command=self.memory_recall, **button_config),
                'MC': ttk.Button(self.button_frame, text='MC', command=self.memory_clear, **button_config),
            }`)
	}

	// Add layout setup
	content.WriteString(`

    def setup_layout(self):
        """Setup widget layout"""
        self.main_frame.pack(fill='both', expand=True)

        # Display layout
        self.display_frame.pack(fill='x', pady=(0, 10))
        self.expr_display.pack(fill='x')
        self.display.pack(fill='x', ipady=10)

        # Button layout
        self.button_frame.pack(fill='both', expand=True)

        # Layout buttons in grid`)

	// Different layouts for basic vs scientific
	if g.config.Type == ScientificCalculator {
		content.WriteString(g.generateScientificLayout())
	} else {
		content.WriteString(g.generateBasicLayout())
	}

	// Add event handling methods
	content.WriteString(`

    def setup_bindings(self):
        """Setup keyboard bindings"""
        self.root.bind('<Key>', self.on_key_press)
        self.root.focus_set()

    def on_key_press(self, event):
        """Handle keyboard input"""
        key = event.char
        if key.isdigit():
            self.append_number(key)
        elif key in '+-*/':
            self.append_operator(key)
        elif key == '.':
            self.append_number('.')
        elif key == '\r' or key == '=':
            self.calculate()
        elif key.lower() == 'c':
            self.clear()
        elif event.keysym == 'BackSpace':
            self.backspace()

    def append_number(self, number):
        """Add number to current expression"""
        if self.result_shown:
            self.current_expression = ""
            self.result_shown = False

        if number == '.' and '.' in self.current_expression.split()[-1]:
            return  # Don't allow multiple decimal points

        self.current_expression += number
        self.update_display()

    def append_operator(self, operator):
        """Add operator to current expression"""
        if self.result_shown:
            self.result_shown = False

        if self.current_expression and self.current_expression[-1] in '+-*/':
            self.current_expression = self.current_expression[:-1]

        self.current_expression += operator
        self.update_display()

    def append_function(self, function):
        """Add function to current expression"""
        if self.result_shown:
            self.current_expression = ""
            self.result_shown = False

        self.current_expression += function + "("
        self.update_display()

    def calculate(self):
        """Perform calculation"""
        try:
            if not self.current_expression:
                return

            # Handle angle units for trig functions
            expression = self.current_expression`)

	// Add angle unit handling for trigonometric functions
	if g.config.Features.Trigonometric {
		content.WriteString(`
            if self.angle_unit == "degrees":
                # Convert degrees to radians for trig functions
                import re
                trig_functions = ['sin', 'cos', 'tan']
                for func in trig_functions:
                    pattern = f'{func}\\(([^)]+)\\)'
                    def replace_trig(match):
                        angle = match.group(1)
                        return f'{func}(math.radians({angle}))'
                    expression = re.sub(pattern, replace_trig, expression)`)
	}

	content.WriteString(`

            result = safe_eval(expression)
            formatted_result = self.format_result(result)

            # Update display
            self.display_var.set(str(formatted_result))
            self.expr_var.set(f"{self.current_expression} =")

            # Add to history`)

	if g.config.Features.History {
		content.WriteString(`
            self.history.add_entry(self.current_expression, formatted_result)`)
	}

	content.WriteString(`

            # Set up for next calculation
            self.current_expression = str(formatted_result)
            self.result_shown = True

        except Exception as e:
            self.display_var.set("Error")
            self.expr_var.set(str(e))
            self.current_expression = ""

    def format_result(self, result):
        """Format calculation result"""
        if isinstance(result, (int, float)):
            if result == int(result):
                return int(result)
            else:
                return round(result, self.precision)
        return result

    def clear(self):
        """Clear everything"""
        self.current_expression = ""
        self.display_var.set("0")
        self.expr_var.set("")
        self.result_shown = False

    def clear_entry(self):
        """Clear current entry"""
        self.current_expression = ""
        self.display_var.set("0")
        self.update_display()

    def backspace(self):
        """Remove last character"""
        if self.current_expression:
            self.current_expression = self.current_expression[:-1]
            self.update_display()

    def update_display(self):
        """Update the display"""
        if self.current_expression:
            self.display_var.set(self.current_expression)
        else:
            self.display_var.set("0")`)

	// Add memory methods if enabled
	if g.config.Features.Memory {
		content.WriteString(`

    def memory_store(self):
        """Store current value in memory"""
        try:
            current_value = float(self.display_var.get())
            self.memory.store(current_value)
            messagebox.showinfo("Memory", f"Stored {current_value} in memory")
        except ValueError:
            messagebox.showerror("Error", "Invalid value to store")

    def memory_recall(self):
        """Recall value from memory"""
        value = self.memory.recall()
        self.current_expression = str(value)
        self.display_var.set(str(value))
        self.result_shown = True

    def memory_clear(self):
        """Clear memory"""
        self.memory.clear()
        messagebox.showinfo("Memory", "Memory cleared")`)
	}

	// Add menu creation if advanced features enabled
	if g.config.Features.Memory || g.config.Features.History || g.config.Features.Statistical {
		content.WriteString(`

    def create_menu(self):
        """Create application menu"""
        menubar = tk.Menu(self.root)
        self.root.config(menu=menubar)

        # Tools menu
        tools_menu = tk.Menu(menubar, tearoff=0)
        menubar.add_cascade(label="Tools", menu=tools_menu)`)

		if g.config.Features.Statistical {
			content.WriteString(`
        tools_menu.add_command(label="Statistics Calculator", command=self.show_stats_dialog)`)
		}

		if g.config.Features.History {
			content.WriteString(`
        tools_menu.add_command(label="Show History", command=self.show_history)
        tools_menu.add_command(label="Clear History", command=self.clear_history)`)
		}

		content.WriteString(`
        tools_menu.add_separator()
        tools_menu.add_command(label="About", command=self.show_about)`)
	}

	// Add statistical dialog if enabled
	if g.config.Features.Statistical {
		content.WriteString(`

    def show_stats_dialog(self):
        """Show statistics calculator dialog"""
        data_str = simpledialog.askstring(
            "Statistics",
            "Enter comma-separated numbers:"
        )
        if data_str:
            try:
                stats = calculate_stats(data_str)
                result = "\\n".join([f"{k.title()}: {v}" for k, v in stats.items()])
                messagebox.showinfo("Statistics Results", result)
            except Exception as e:
                messagebox.showerror("Error", str(e))`)
	}

	// Add history methods if enabled
	if g.config.Features.History {
		content.WriteString(`

    def show_history(self):
        """Show calculation history"""
        history = self.history.get_history()
        if not history:
            messagebox.showinfo("History", "No calculations in history")
            return

        # Create history window
        hist_window = tk.Toplevel(self.root)
        hist_window.title("Calculation History")
        hist_window.geometry("400x300")

        # Create text widget with scrollbar
        text_frame = ttk.Frame(hist_window)
        text_frame.pack(fill='both', expand=True, padding=10)

        text_widget = tk.Text(text_frame, wrap='word')
        scrollbar = ttk.Scrollbar(text_frame, orient='vertical', command=text_widget.yview)
        text_widget.configure(yscrollcommand=scrollbar.set)

        # Add history entries
        for entry in history:
            text_widget.insert('end', f"{entry['timestamp']}: {entry['expression']} = {entry['result']}\\n")

        text_widget.config(state='disabled')
        text_widget.pack(side='left', fill='both', expand=True)
        scrollbar.pack(side='right', fill='y')

    def clear_history(self):
        """Clear calculation history"""
        self.history.clear()
        messagebox.showinfo("History", "History cleared")`)
	}

	// Add theme and about methods
	content.WriteString(`

    def apply_theme(self):
        """Apply visual theme"""
        # Configure styles based on theme setting
        style = ttk.Style()`)

	if g.config.UI.Theme == "dark" {
		content.WriteString(`
        # Dark theme
        self.root.configure(bg='#2b2b2b')
        style.configure('TFrame', background='#2b2b2b')
        style.configure('TButton', background='#404040', foreground='white')
        style.configure('TLabel', background='#2b2b2b', foreground='white')`)
	} else {
		content.WriteString(`
        # Light theme (default)
        style.configure('TButton', padding=5)`)
	}

	content.WriteString(`

    def show_about(self):
        """Show about dialog"""
        about_text = f"""` + g.config.ProjectName + `
Version: 1.0.0
Author: ` + g.config.Author + `

` + g.config.Description + `

Generated by Calculator Generator"""
        messagebox.showinfo("About", about_text)

    def run(self):
        """Start the calculator application"""
        self.root.mainloop()

def main():
    """Main application entry point"""
    calculator = CalculatorGUI()
    calculator.run()

if __name__ == "__main__":
    main()`)

	return content.String()
}

// generateBasicLayout creates button layout for basic calculator
func (g *GUIGenerator) generateBasicLayout() string {
	layout := `

        # Basic calculator layout (4x5 grid)
        row = 0

        # Memory buttons (if enabled)`

	if g.config.Features.Memory {
		layout += `
        if hasattr(self, 'mem_buttons'):
            col = 0
            for text, button in self.mem_buttons.items():
                button.grid(row=row, column=col, padx=2, pady=2, sticky='nsew')
                col += 1
            row += 1`
	}

	layout += `

        # Clear buttons
        self.op_buttons['CE'].grid(row=row, column=0, padx=2, pady=2, sticky='nsew')
        self.op_buttons['C'].grid(row=row, column=1, padx=2, pady=2, sticky='nsew')
        row += 1

        # Number and operator layout
        # Row 1: 7, 8, 9, /
        self.num_buttons[7].grid(row=row, column=0, padx=2, pady=2, sticky='nsew')
        self.num_buttons[8].grid(row=row, column=1, padx=2, pady=2, sticky='nsew')
        self.num_buttons[9].grid(row=row, column=2, padx=2, pady=2, sticky='nsew')
        self.op_buttons['/'].grid(row=row, column=3, padx=2, pady=2, sticky='nsew')
        row += 1

        # Row 2: 4, 5, 6, *
        self.num_buttons[4].grid(row=row, column=0, padx=2, pady=2, sticky='nsew')
        self.num_buttons[5].grid(row=row, column=1, padx=2, pady=2, sticky='nsew')
        self.num_buttons[6].grid(row=row, column=2, padx=2, pady=2, sticky='nsew')
        self.op_buttons['*'].grid(row=row, column=3, padx=2, pady=2, sticky='nsew')
        row += 1

        # Row 3: 1, 2, 3, -
        self.num_buttons[1].grid(row=row, column=0, padx=2, pady=2, sticky='nsew')
        self.num_buttons[2].grid(row=row, column=1, padx=2, pady=2, sticky='nsew')
        self.num_buttons[3].grid(row=row, column=2, padx=2, pady=2, sticky='nsew')
        self.op_buttons['-'].grid(row=row, column=3, padx=2, pady=2, sticky='nsew')
        row += 1

        # Row 4: 0, ., =, +
        self.num_buttons[0].grid(row=row, column=0, columnspan=2, padx=2, pady=2, sticky='nsew')
        self.op_buttons['.'].grid(row=row, column=2, padx=2, pady=2, sticky='nsew')
        self.op_buttons['+'].grid(row=row, column=3, padx=2, pady=2, sticky='nsew')
        row += 1

        # Equals button
        self.op_buttons['='].grid(row=row, column=0, columnspan=4, padx=2, pady=2, sticky='nsew')

        # Configure grid weights
        for i in range(4):
            self.button_frame.columnconfigure(i, weight=1)
        for i in range(row + 1):
            self.button_frame.rowconfigure(i, weight=1)`

	return layout
}

// generateScientificLayout creates button layout for scientific calculator
func (g *GUIGenerator) generateScientificLayout() string {
	layout := `

        # Scientific calculator layout (6x8 grid)
        row = 0

        # Memory buttons (if enabled)`

	if g.config.Features.Memory {
		layout += `
        if hasattr(self, 'mem_buttons'):
            col = 0
            for text, button in self.mem_buttons.items():
                button.grid(row=row, column=col, padx=1, pady=1, sticky='nsew')
                col += 1
            row += 1`
	}

	// Add scientific function buttons
	if g.config.Features.Trigonometric {
		layout += `

        # Trigonometric functions
        if hasattr(self, 'trig_buttons'):
            col = 0
            for text, button in self.trig_buttons.items():
                button.grid(row=row, column=col, padx=1, pady=1, sticky='nsew')
                col += 1
            row += 1`
	}

	if g.config.Features.Logarithmic {
		layout += `

        # Logarithmic functions
        if hasattr(self, 'log_buttons'):
            col = 3
            for text, button in self.log_buttons.items():
                button.grid(row=row-1, column=col, padx=1, pady=1, sticky='nsew')
                col += 1`
	}

	layout += `

        # Clear buttons
        self.op_buttons['CE'].grid(row=row, column=0, padx=1, pady=1, sticky='nsew')
        self.op_buttons['C'].grid(row=row, column=1, padx=1, pady=1, sticky='nsew')
        row += 1

        # Standard calculator layout
        # Row: 7, 8, 9, /
        self.num_buttons[7].grid(row=row, column=0, padx=1, pady=1, sticky='nsew')
        self.num_buttons[8].grid(row=row, column=1, padx=1, pady=1, sticky='nsew')
        self.num_buttons[9].grid(row=row, column=2, padx=1, pady=1, sticky='nsew')
        self.op_buttons['/'].grid(row=row, column=3, padx=1, pady=1, sticky='nsew')
        row += 1

        # Row: 4, 5, 6, *
        self.num_buttons[4].grid(row=row, column=0, padx=1, pady=1, sticky='nsew')
        self.num_buttons[5].grid(row=row, column=1, padx=1, pady=1, sticky='nsew')
        self.num_buttons[6].grid(row=row, column=2, padx=1, pady=1, sticky='nsew')
        self.op_buttons['*'].grid(row=row, column=3, padx=1, pady=1, sticky='nsew')
        row += 1

        # Row: 1, 2, 3, -
        self.num_buttons[1].grid(row=row, column=0, padx=1, pady=1, sticky='nsew')
        self.num_buttons[2].grid(row=row, column=1, padx=1, pady=1, sticky='nsew')
        self.num_buttons[3].grid(row=row, column=2, padx=1, pady=1, sticky='nsew')
        self.op_buttons['-'].grid(row=row, column=3, padx=1, pady=1, sticky='nsew')
        row += 1

        # Row: 0, ., =, +
        self.num_buttons[0].grid(row=row, column=0, columnspan=2, padx=1, pady=1, sticky='nsew')
        self.op_buttons['.'].grid(row=row, column=2, padx=1, pady=1, sticky='nsew')
        self.op_buttons['+'].grid(row=row, column=3, padx=1, pady=1, sticky='nsew')
        row += 1

        # Equals button
        self.op_buttons['='].grid(row=row, column=0, columnspan=4, padx=1, pady=1, sticky='nsew')

        # Configure grid weights
        for i in range(6):
            self.button_frame.columnconfigure(i, weight=1)
        for i in range(row + 1):
            self.button_frame.rowconfigure(i, weight=1)`

	return layout
}

// renderGUITemplate renders the GUI calculator template with the given data
func (g *GUIGenerator) renderGUITemplate(data TemplateData) (string, error) {
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
`

	t, err := template.New("gui_calculator").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var result strings.Builder
	if err := t.Execute(&result, data); err != nil {
		return "", err
	}

	return result.String(), nil
}
