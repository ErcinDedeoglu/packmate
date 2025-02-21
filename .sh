repomix --no-file-summary --no-security-check --include "src/**" --output "repopack.yml"

docker build -t archiver:latest -f src/Dockerfile src

docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/test-output:/output \
  archiver:latest \
  --path /source \
  --output /output \
  --name "backup-$(date +%Y%m%d)" \
  --format json

  docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/test-output:/output \
  archiver:latest \
  --path /source \
  --output /output \
  --name "my-backup" \
  --format text