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
- **Advanced Parameters**: Support for custom compression methods, multi-volume archives, encryption, and more  
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
  --name "backup-$(date +%Y%m%d)"
```

### 🎯 Parameters
| Parameter           | Description                                                                 | Required | Default       |
|---------------------|-----------------------------------------------------------------------------|----------|---------------|
| `--source`          | Source directory to archive                                                | Yes      | `/source`     |
| `--output`          | Output directory for the archive                                           | Yes      | `/output`     |
| `--name`            | Custom name for the archive file                                           | No       | Base64 encoded path |
| `--compression`     | Compression level (0: none, 1: fastest, 9: best compression)                | No       | `5`           |
| `--method`          | Compression method (`copy`, `lzma2`, `ppmd`, etc.)                         | No       | `lzma2`       |
| `--password`        | Password for encrypting the archive                                        | No       | None          |
| `--header-encryption` | Enable header encryption                                                  | No       | `false`       |
| `--volume-size`     | Create multi-volume archive with specified size (e.g., `100m`)             | No       | None          |
| `--multithreading`  | Enable multithreading for faster compression                               | No       | `true`        |
| `--extra`           | Additional `p7zip` flags (e.g., `-ms=off`)                                 | No       | None          |
| `--format`          | Output format (`json` or `text`)                                           | No       | `json`        |

### 📊 Compression Levels
| Level | Description         | Use Case                                   |
|-------|---------------------|--------------------------------------------|
| `0`   | No compression      | Fastest, best for already compressed data |
| `1`   | Best speed          | Good balance for compressible data        |
| `5`   | Default compression | Balanced speed and compression ratio      |
| `9`   | Best compression    | Smallest file size, slower compression    |

### 🔐 Encryption Options
| Parameter           | Description                                                                 |
|---------------------|-----------------------------------------------------------------------------|
| `--password`        | Encrypt the archive with a password                                         |
| `--header-encryption` | Encrypt the archive headers for additional security                        |

### 📦 Multi-Volume Archives
Use the `--volume-size` parameter to split the archive into multiple files:
```bash
--volume-size=100m  # Creates 100 MB volumes
```

## 💡 Examples

### Basic Backup (No Compression)
```bash
docker run --rm \
  -v /var/lib/docker/volumes/myapp_data/_data:/source:ro \
  -v /backup:/output \
  dublok/packmate:latest \
  --name "backup-$(date +%Y%m%d)" \
  --compression=0 \
  --method=copy
```

### Maximum Compression Backup
```bash
docker run --rm \
  -v /var/lib/docker/volumes/myapp_data/_data:/source:ro \
  -v /backup:/output \
  dublok/packmate:latest \
  --name "backup-$(date +%Y%m%d)" \
  --compression=9 \
  --method=lzma2
```

### Encrypted Backup
```bash
docker run --rm \
  -v /var/lib/docker/volumes/myapp_data/_data:/source:ro \
  -v /backup:/output \
  dublok/packmate:latest \
  --name "backup-$(date +%Y%m%d)" \
  --compression=5 \
  --method=lzma2 \
  --password=secret \
  --header-encryption
```

### Multi-Volume Archive
```bash
docker run --rm \
  -v /var/lib/docker/volumes/myapp_data/_data:/source:ro \
  -v /backup:/output \
  dublok/packmate:latest \
  --name "backup-$(date +%Y%m%d)" \
  --compression=5 \
  --method=lzma2 \
  --volume-size=100m
```

### Fastest Archiving (No Compression, No Solid Mode)
```bash
docker run --rm \
  -v /var/lib/docker/volumes/myapp_data/_data:/source:ro \
  -v /backup:/output \
  dublok/packmate:latest \
  --name "backup-$(date +%Y%m%d)" \
  --compression=0 \
  --method=copy \
  --multithreading=true \
  --extra="-ms=off"
```

## 📤 Output Examples

### JSON Format
```json
{
  "path": "/source",
  "archivePath": "/output/backup-20250221.7z",
  "status": "Success",
  "compressionLevel": 0,
  "error": ""
}
```

### Text Format
```
Archive Creation Result:
=======================
✅ Path: /source
   Archive: /output/backup-20250221.7z
   Compression Level: 0
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
- Uses `p7zip` for `.7z` archive creation  
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
