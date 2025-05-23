# LeetCode Clone Components

This directory contains the React components for the LeetCode clone application.

## Components

### CodeEditor

The `CodeEditor` component provides a full-featured code editor based on Monaco Editor (the same editor that powers VS Code). It includes:

- Syntax highlighting for JavaScript, Python, and C++
- Advanced editor features like auto-completion and code folding
- Custom dark theme that matches LeetCode's style
- Language selection
- Smart boilerplate code generation based on problem type

#### Props

- `language` (string, optional): The initial language for the editor (default: 'javascript')
- `problemTitle` (string, optional): The title of the problem, used to generate appropriate starter code

### Navbar

Simple navigation bar component for the application.

### ProblemDetail

Shows detailed information about a coding problem, including:
- Problem description
- Examples
- Constraints
- Code editor for solving the problem

### ProblemsList

Displays a list of available coding problems with difficulty levels.
