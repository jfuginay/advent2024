# Advent of Code 2024

Welcome to my Advent of Code 2024 repository! This project contains my solutions to the Advent of Code 2024 challenges. I will start solving the problems using Python and then branch out to other languages such as TypeScript and Go.

## Table of Contents

- [Introduction](#introduction)
- [Languages Used](#languages-used)
- [Setup](#setup)
- [Project Structure](#project-structure)
- [How to Run](#how-to-run)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Advent of Code is an annual event that provides daily programming challenges from December 1st to December 25th. This repository contains my personal solutions to these challenges.

## Languages Used

- Python
- TypeScript
- Go

## Setup

### TypeScript Setup
1. Initialize a new npm project in the typescript directory:
```sh
cd typescript
npm init -y
```

2. Install required dependencies:
```sh
npm install typescript ts-node @types/node
```

3. Create a tsconfig.json (if not already present):
```sh
npx tsc --init
```

## Project Structure

```
/advent2024
│
├── python/
│   ├── day1.py
│   ├── day2.py
│   └── ...
│
├── typescript/
│   ├── day1.ts
│   ├── day2.ts
│   └── ...
│
├── go/
│   ├── day1.go
│   ├── day2.go
│   └── ...
│
└── README.md
```

## How to Run

### Python

Navigate to the `python` directory and run the desired day's solution:

```sh
cd python
python day1.py
```

### TypeScript

Navigate to the `typescript` directory and ensure you have the necessary dependencies:

```sh
cd typescript
npm install typescript ts-node @types/node
```

Then you can run the TypeScript files directly using ts-node:

```sh
npx ts-node day1.ts
```

Alternatively, you can still use the traditional compile-and-run method:

```sh
tsc day1.ts
node day1.js
```

### Go

Navigate to the `go` directory and run the desired day's solution:

```sh
cd go
go run day1.go
```

## Contributing

Feel free to fork this repository and submit pull requests. Contributions are welcome!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

