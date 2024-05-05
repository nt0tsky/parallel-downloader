# Parallel Downloader

This application downloads files in parallel by chunks.

## Installation

To install the application, clone the repository:

```bash
git clone https://github.com/nt0tsky/parallel-downloader.git
```

## Usage

Build the application:

```bash
go build -o downloader ./cmd/main.go
```

Run the application with the desired command-line arguments:

```bash
./downloader -threads=3 -url=<URL> -destinationFolder=<DestinationFolder>
```

## Command-line Flags

- `-threads`: Limit the number of downloading goroutines (default is 3).
- `-url`: Specify the URL of the file to download.
- `-destinationFolder`: Specify the destination folder where the downloaded file will be saved.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
