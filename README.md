# LeetCode Clone

A simple LeetCode clone with a React frontend and Express.js backend.

## Features

- List of coding problems with difficulty indicators
- Problem detail page with description and constraints
- Advanced code editor with:
  - Syntax highlighting for JavaScript, Python, and C++
  - Monaco Editor (same as VS Code)
  - Language selector
  - Code submission capability
  - Smart code templates based on problem type
- Code execution with test cases:
  - Built-in executors for JavaScript, Python, and C++ 
- Advanced Learning Capabilities
  - Spaced Repetition for better path to learning
  - AI hints and doubt resolution for custom hints during study
  - AI generated custom learning paths based on proficiency

## Project Structure

- `/client` - React frontend
- `/server` - Golang backend

## Getting Started

### Prerequisites

- Node.js
- npm
- Docker and Docker Compose
- Golang

### Installation

1. Clone the repository
2. Install dependencies:

```bash
make prepare
```

### Running the Application

```bash
# From the root directory
npm start
```

This will concurrently start:
- Frontend on http://localhost:3000
- Backend on http://localhost:5000

## Technologies Used

- Frontend: React, React Router, Axios, Monaco Editor
- Backend: Golang, Gin
- Code Execution: [Piston](https://github.com/engineer-man/piston)