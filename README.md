# Serve-Ready CLI

Serve-Ready is a CLI tool designed to check if a server meets the necessary requirements for running selected frameworks, databases, caches, and web servers. It dynamically checks for required versions and configurations based on predefined `.yml` files.

## Project Structure

```bash
/project-root
│
├── /builds                 # Folder for build outputs
│   └── /cli                # Contains the CLI binary
├── /src                    # Source files for the project
│   ├── /cmd                # Main entry point for the CLI
│   ├── /internal           # Internal logic for CLI and requirements
│   │   ├── /cli            # CLI prompt logic
│   │   ├── /frameworks     # Framework configuration files
│   │   ├── /requirements   # Requirement loader and checker
│   │   ├── /services       # Services for caches, databases, runtimes, webservers
│   │   │   ├── /caches     # Cache services (e.g., Redis, Memcached)
│   │   │   ├── /databases  # Database services (e.g., MySQL, PostgreSQL)
│   │   │   ├── /runtimes   # Runtime services (e.g., PHP, Node.js)
│   │   │   └── /webservers # Web server services (e.g., Nginx, Apache)
│   └── /pkg                # Package utilities such as configuration and colors
├── go.mod                  # Go module definition
└── go.sum                  # Go dependencies checksum file
```

### Frameworks

The supported frameworks for the tool are defined in `.yml` files located in the `src/internal/frameworks/` folder.

- **Supported Frameworks**:
  - Laravel
  - Next.js
  - Django
  - FastAPI
  - Flask
  - Remix.js

### Databases

The supported databases for the tool are defined in `.yml` files located in the `src/internal/services/databases/` folder.

- **Supported Databases**:
  - MySQL
  - PostgreSQL
  - MariaDB
  - MongoDB

### Caches

The supported caches for the tool are defined in `.yml` files located in the `src/internal/services/caches/` folder.

- **Supported Caches**:
  - Redis
  - Memcached

### Web Servers

The supported web servers for the tool are defined in `.yml` files located in the `src/internal/services/webservers/` folder.

- **Supported Web Servers**:
  - Nginx
  - Apache
  - Caddy
  - Lighttpd
  - Tomcat

### Runtimes

The supported runtimes for the tool are located in the `src/internal/services/runtimes/` folder.

- **Supported Runtimes**:
  - PHP
  - Node.js
  - Python
  - Java
  - Go
  - Rust
  - C#

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
2. **Runtime & Package Manager Check**: Depending on the selected framework, the tool will check the runtime (e.g., PHP, Node.js) and relevant package managers (e.g., Composer, NPM).
3. **Database Selection**: Optionally, you can select a database from the available options.
4. **Cache Selection**: Optionally, you can select a cache system from the available options.
5. **Web Server Selection**: Optionally, you can select a web server from the available options.
6. **Requirement Checking**: The tool will check for the required versions of PHP, Node.js, Python, and any additional packages or extensions required by the selected framework.

## Project Setup

The project source files are located inside the `src` directory. Here is the breakdown:

- **`src/cmd/main.go`**: The entry point for the CLI tool.
- **`src/internal/requirements`**: Contains the logic for loading `.yml` files and checking server requirements.
- **`src/internal/cli`**: Contains the logic for prompting the user for input.
- **`src/internal/services`**: Contains the logic for checking services such as caches, databases, runtimes, and web servers.

## Development

To modify the source code, update the files under the `src/` directory. For example, to add more frameworks or databases, you can add `.yml` files in the respective folders (`src/internal/frameworks/`, `src/internal/services/databases/`, `src/internal/services/caches/`).

After making changes, rebuild the CLI:

```bash
go build -o ./builds/cli/serve-ready ./src/cmd/main.go
```

## Known Issues

- Ensure that the `src/internal/frameworks/`, `src/internal/services/databases/`, and `src/internal/services/caches/` folders contain valid `.yml` configuration files.
- The tool must be run from the root directory where the `src` folder is located.

## Contributing

If you would like to contribute to the project:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request with a detailed description of the changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
