<div align="center">

# 🗄️ PackMate

A high-performance, Docker-ready archiving tool for backing up Docker volumes and directories

[![Docker Pulls](https://img.shields.io/docker/pulls/dublok/packmate.svg)](https://hub.docker.com/r/dublok/packmate)
[![Docker Image Size](https://img.shields.io/docker/image-size/dublok/packmate/latest)](https://hub.docker.com/r/dublok/packmate)

</div>

## 🚀 Features
- **Fast & Efficient**: Optimized for handling large datasets with minimal resource usage
- **Docker-Ready**: Purpose-built for Docker volume backups
- **Configurable Compression**: Choose between no compression for speed or maximum compression for space savings
- **Multiple Output Formats**: Support for both JSON and human-readable outputs
- **Custom Naming**: Flexible archive naming options
- **Safe Operations**: Read-only source handling and comprehensive error reporting

## 📋 Usage

### ⚡ Quick Start
```bash
docker run --rm \
  -v /your/source/path:/source:ro \
  -v /your/backup/path:/output \
  dublok/packmate:latest \
  --path /source \
  --output /output \
  --name "backup-$(date +%Y%m%d)"
```

### 🎯 Parameters
| Parameter | Description | Required | Default |
|-----------|-------------|----------|---------|
| `--path` | Source path to archive | Yes | - |
| `--output` | Output directory for archive | Yes | - |
| `--name` | Custom name for archive file | No | Base64 encoded path |
| `--compression` | Compression level | No | -2 |
| `--format` | Output format (json/text) | No | json |

### 📊 Compression Levels
| Level | Description | Use Case |
|-------|-------------|----------|
| `-2` | No compression (default) | Fastest, best for already compressed data |
| `1` | Best speed | Good balance for compressible data |
| `9` | Best compression | Smallest file size, slower compression |

## 💡 Examples

### Basic Backup (No Compression)
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

## 🔒 Security Features
- Source volumes mounted read-only (`ro`)
- No root privileges required
- Minimal container footprint
- Alpine-based secure base image

## 🏗️ Building from Source
```bash
# Build the Docker image
docker build -t packmate:latest -f src/Dockerfile src

# Optional: Tag and push to your registry
docker tag packmate:latest your-registry/packmate:latest
docker push your-registry/packmate:latest
```

## ⚙️ Technical Specifications
- Go 1.23+
- Alpine Linux base
- No external runtime dependencies
- Efficient buffering for large files
- Minimal memory footprint
- Multi-stage Docker builds

## 📫 Contact & Support

<div align="center">

[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-blue.svg)](https://www.linkedin.com/in/ercindedeoglu/)
[![GitHub](https://img.shields.io/github/followers/ErcinDedeoglu?label=Follow&style=social)](https://github.com/ErcinDedeoglu)

If you find this project useful, please consider giving it a ⭐

</div>

---

<div align="center">

Made with ❤️ by [Ercin Dedeoglu](https://github.com/ErcinDedeoglu)

</div>
