# Virtual Environment Destroyer

`Virtual Environment Destroyer` is a Golang written CLI tool that helps you destroy all virtual environments below the current working directory.

<!-- ## Installation

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install foobar.

```bash
pip install foobar
``` -->

## Development Setup

1. Clone the repo

```bash
git clone https://github.com/bautsi/venv-destroyer.git
cd venv-destroyer
```

2. Download and install dependencies

```bash
go mod tidy
```

3. Build the executable file

```bash
go build -o .
```

4. Start the program and see the introduction message

```bash
./venv-destroyer.exe
```

## Usage

```bash
./venv-destroyer.exe

# Introduction
./venv-destroyer.exe -h

# Check existing virtual environments below the current working directory.
./venv-destroyer.exe -check

# Destroy all virtual environments below the current working directory.
./venv-destroyer.exe -destroy
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

<!-- Please make sure to update tests as appropriate. -->

## License

[MIT](https://choosealicense.com/licenses/mit/)
