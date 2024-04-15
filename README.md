# Blip
## Installation:

### Prerequisites
- Node.js (version 21.0.0 or higher)

### Installing and Running the Script

1. Open Command Prompt or PowerShell as an **administrator** (right-click -> Run as administrator).
2. Install the npm package globally:
```bash
npm i -g blip-cli
```
3. Setup the cli to install necessary dependencies:
```bash
blip --setup
```

## Usage:
   ```bash
   blip [command]
   ```

### Available Commands:
```bash
  help        Help about any command
  me          Me is a command that represents the user's profile
  ```

### Flags:
  ```bash
  -h, --help      help for blip
  -v, --version   View current version
  ```

Use `"blip [command] --help"` for more information about a command.

## Commands Usage:
  ```bash
  blip me [flags]
  blip me [command]
  ```
### Available Commands:      
```bash
login     Login to your Blip account
view      See the list of your available tool
```

[![Go Reference](https://pkg.go.dev/badge/github.com/scriptnsam/blip-v2.svg)](https://pkg.go.dev/github.com/scriptnsam/blip-v2)