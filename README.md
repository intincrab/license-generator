# License Generator

License Generator is a CLI tool that helps you quickly generate licenses for your projects. Written in [Go](https://go.dev/).

## Installation

Choose the appropriate installation method for your operating system:

### Linux (64-bit)

1. Download the `license-generator-linux-amd64.tar.gz` file from the latest release.
2. Open a terminal and navigate to the download directory.
3. Extract the file:
   ```
   tar -xzf license-generator-linux-amd64.tar.gz
   ```
4. Move the extracted `license-generator` file to a directory in your PATH, e.g.:
   ```
   sudo mv license-generator /usr/local/bin/
   ```
5. Make it executable:
   ```
   sudo chmod +x /usr/local/bin/license-generator
   ```

### Windows (64-bit)

1. Download the `license-generator-windows-amd64.zip` file from the latest release.
2. Extract the ZIP file to a location of your choice.
3. Add the extracted directory to your system's PATH environment variable.
4. Open a new Command Prompt and run `license-generator.exe` to use the tool.

### macOS (64-bit)

1. Download the `license-generator-darwin-amd64.tar.gz` file from the latest release.
2. Open Terminal and navigate to the download directory.
3. Extract the file:
   ```
   tar -xzf license-generator-darwin-amd64.tar.gz
   ```
4. Move the extracted `license-generator` file to a directory in your PATH, e.g.:
   ```
   sudo mv license-generator /usr/local/bin/
   ```
5. Make it executable:
   ```
   sudo chmod +x /usr/local/bin/license-generator
   ```

## Usage

After installation, you can use the License Generator tool from the command line:

```
license-generator [options]
```

### Options:

- `-license`: Specify the license name (optional)
- `-skip-prompt`: Skip prompts and use defaults (optional)

### Examples:

1. Generate a license with interactive prompts:
   ```
   license-generator
   ```

2. Generate a specific license (e.g., MIT):
   ```
   license-generator -license MIT
   ```

3. Generate a license using defaults without prompts:
   ```
   license-generator -skip-prompt
   ```

4. Generate a specific license using defaults:
   ```
   license-generator -license Apache-2.0 -skip-prompt
   ```

The generated license will be saved as `LICENSE` in your current directory.

## Available Licenses

The License Generator supports various open-source licenses. Run the tool without arguments to see the list of available licenses.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
