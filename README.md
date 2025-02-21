# PackMate

PackMate is a high-performance, Docker-ready archiving tool designed specifically for backing up Docker volumes and directories. It offers flexible compression options and supports both JSON and human-readable outputs.

## 🚀 Features

- **Fast & Efficient**: Optimized for handling large datasets with minimal resource usage
- **Docker-Ready**: Purpose-built for Docker volume backups
- **Configurable Compression**: Choose between no compression for speed or maximum compression for space savings
- **Multiple Output Formats**: Support for both JSON and human-readable outputs
- **Custom Naming**: Flexible archive naming options
- **Safe Operations**: Read-only source handling and comprehensive error reporting

## 📋 Usage

### Docker Command

```bash
docker run --rm \
  -v /your/source/path:/source:ro \
  -v /your/backup/path:/output \
  dublok/packmate:latest \
  --path /source \
  --output /output \
  --name "backup-$(date +%Y%m%d)" \
  --compression -2 \
  --format json
```

### Parameters

| Parameter | Description | Required | Default |
|-----------|-------------|----------|---------|
| `--path` | Source path to archive | Yes | - |
| `--output` | Output directory for archive | Yes | - |
| `--name` | Custom name for archive file | No | Base64 encoded path |
| `--compression` | Compression level | No | -2 |
| `--format` | Output format (json/text) | No | json |

### Compression Levels

- `-2`: No compression (default, fastest)
- `1`: Best speed
- `9`: Best compression

## 💡 Examples

### Basic Backup with No Compression
```bash
docker run --rm \
  -v /var/lib/docker/volumes/myapp_data/_data:/source:ro \
  -v /backup:/output \
  dublok/packmate:latest \
  --path /source \
  --output /output \
  --name "backup-$(date +%Y%m%d)"
```

### Maximum Compression Backup
```bash
docker run --rm \
  -v /var/lib/docker/volumes/myapp_data/_data:/source:ro \
  -v /backup:/output \
  dublok/packmate:latest \
  --path /source \
  --output /output \
  --name "backup-$(date +%Y%m%d)" \
  --compression 9
```

### Human-Readable Output
```bash
docker run --rm \
  -v /var/lib/docker/volumes/myapp_data/_data:/source:ro \
  -v /backup:/output \
  dublok/packmate:latest \
  --path /source \
  --output /output \
  --format text
```

## 📤 Output Examples

### JSON Format
```json
{
  "path": "/source",
  "archivePath": "/output/backup-20250221.zip",
  "status": "Success",
  "compressionLevel": -2
}
```

### Text Format
```
Archive Creation Result:
=======================
✅ Path: /source
   Archive: /output/backup-20250221.zip
   Compression Level: -2
   Status: Success
```

## 🔒 Security

- Source volumes are mounted read-only (`ro`)
- No root privileges required
- Minimal container footprint

## 🏗️ Building

```bash
# Build the Docker image
docker build -t packmate:latest -f src/Dockerfile src

# Optional: Tag and push to your registry
docker tag packmate:latest your-registry/packmate:latest
docker push your-registry/packmate:latest
```

## 📦 Dependencies

- Go 1.23+
- Alpine Linux (base container)
- No external runtime dependencies

## ⚙️ Technical Details

- Written in Go for maximum performance
- Uses efficient buffering for large files
- Minimal memory footprint
- Docker multi-stage builds for smaller image size

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## ✨ Acknowledgments

- Inspired by the need for efficient Docker volume backups
- Built with Go's standard library for maximum compatibility
