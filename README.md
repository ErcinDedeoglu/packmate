# PackMate

PackMate is a lightweight and efficient Docker container utility designed to create ZIP archives from Docker volumes and directories. It provides a simple way to backup and archive data with both automatic and custom naming options.

## Features

- Create ZIP archives from any mounted directory
- Support for custom archive names
- Read-only source protection
- JSON and text output formats
- No compression option for faster archiving
- Docker Hub and GitHub Container Registry support

## Installation

Pull the image from Docker Hub:
```bash
docker pull dublok/packmate:latest
```

Or from GitHub Container Registry:
```bash
docker pull ghcr.io/ercindedeoglu/packmate:latest
```

## Usage

### Basic Usage

```bash
docker run --rm \
  -v /path/to/source:/source:ro \
  -v /path/to/output:/output \
  dublok/packmate:latest \
  --path /source \
  --output /output \
  --format json
```

### With Custom Archive Name

```bash
docker run --rm \
  -v /path/to/source:/source:ro \
  -v /path/to/output:/output \
  dublok/packmate:latest \
  --path /source \
  --output /output \
  --name "my-backup" \
  --format json
```

### Parameters

- `--path`: Source directory to archive (required)
- `--output`: Output directory for the archive (required)
- `--name`: Custom name for the archive file (optional)
- `--format`: Output format - 'json' or 'text' (default: json)

### Example Output

JSON format:
```json
{
  "path": "/source",
  "archivePath": "/output/my-backup.zip",
  "status": "Success"
}
```

Text format:
```
Archive Creation Result:
=======================

✅ Path: /source
   Archive: /output/my-backup.zip
   Status: Success
```

## Docker Volume Backup Example

```bash
docker run --rm \
  -v your_volume_name:/source:ro \
  -v /path/to/backups:/output \
  dublok/packmate:latest \
  --path /source \
  --output /output \
  --name "backup-$(date +%Y%m%d)" \
  --format json
```

## Building from Source

```bash
git clone https://github.com/ercindedeoglu/packmate.git
cd packmate
docker build -t packmate:latest -f src/Dockerfile src
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

If you encounter any issues or have questions, please file an issue on the GitHub repository.
