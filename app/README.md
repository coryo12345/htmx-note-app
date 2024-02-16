# Project noteapp

A simple note app using HTMX, Go, and TEMPL. Using tailwindcss for styling. Using a local SQLite database to store data.

## Getting Started

First, you must create create a `.env` file. Copy the `.env_example` and rename to get started. (`cp .env_example .env`).  

You can run the code using the Makefile provided, see the section below for more info.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```