# Serve-Ready CLI

Serve-Ready is a CLI tool designed to check if a server meets the necessary requirements for running selected frameworks, databases, caches, and web servers. It dynamically checks for required versions and configurations based on predefined `.yml` files.

## Project Structure

```bash
/project-root
│
├── /builds                 # Folder for build outputs
│   └── /cli                # Contains the CLI binary
├── /src                    # Source files for the project
│   ├── /caches             # Cache configuration files
│   ├── /cmd                # Main entry point for the CLI
│   ├── /databases          # Database configuration files
│   ├── /frameworks         # Framework configuration files
│   └── /internal           # Internal logic for CLI and requirements
│       ├── /cli            # CLI prompt logic
│       └── /requirements   # Requirement loader and checker
├── go.mod                  # Go module definition
└── go.sum                  # Go dependencies checksum file
```

### Frameworks

The supported frameworks for the tool are defined in `.yml` files located in the `src/frameworks/` folder.

- **Supported Frameworks**:
  - Laravel
  - Next.js
  - Django

### Databases

The supported databases for the tool are defined in `.yml` files located in the `src/databases/` folder.

- **Supported Databases**:
  - MySQL
  - PostgreSQL

### Caches

The supported caches for the tool are defined in `.yml` files located in the `src/caches/` folder.

- **Supported Caches**:
  - Redis
  - Firebase

## Installation

1. Clone the repository.
2. Make sure you have [Go](https://golang.org/dl/) installed (Go 1.18 or higher).

```bash
git clone https://your-repo-url
cd serve-ready
```

3. Initialize the Go modules if necessary:

```bash
go mod tidy
```

## Building the CLI

To build the Serve-Ready CLI, run the following command:

```bash
go build -o ./builds/cli/serve-ready ./src/cmd/main.go
```

This will create a binary called `serve-ready` inside the `builds/cli/` directory.

## Running the CLI

Once you have built the CLI, you can run it using the following command:

```bash
./builds/cli/serve-ready
```

The tool will guide you through selecting a framework, database, cache, and web server, and will check if the server meets the requirements for running the selected options.

## How the Tool Works

1. **Framework Selection**: The tool will prompt you to select a framework from the available options.
2. **Database Selection**: Optionally, you can select a database from the available options.
3. **Cache Selection**: Optionally, you can select a cache system from the available options.
4. **Web Server Selection**: Optionally, you can select a web server from the available options.
5. **Requirement Checking**: The tool will check for the required versions of PHP, Node.js, Python, and any additional packages or extensions required by the selected framework.

## Project Setup

The project source files are located inside the `src` directory. Here is the breakdown:

- **`src/cmd/main.go`**: The entry point for the CLI tool.
- **`src/internal/requirements`**: Contains the logic for loading `.yml` files and checking server requirements.
- **`src/internal/cli`**: Contains the logic for prompting the user for input.

## Development

To modify the source code, update the files under the `src/` directory. For example, to add more frameworks or databases, you can add `.yml` files in the respective folders (`src/frameworks/`, `src/databases/`, `src/caches/`).

After making changes, rebuild the CLI:

```bash
go build -o ./builds/cli/serve-ready ./src/cmd/main.go
```

## Known Issues

- Ensure that the `src/frameworks/`, `src/databases/`, and `src/caches/` folders contain valid `.yml` configuration files.
- The tool must be run from the root directory where the `src` folder is located.

## Contributing

If you would like to contribute to the project:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request with a detailed description of the changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
