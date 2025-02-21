repomix --no-file-summary --no-security-check --include "src/Dockerfile,src/archiver/archiver.go" --output "repopack.yml"
repomix --no-file-summary --no-security-check --include "README.md,src/Dockerfile,src/archiver/archiver.go" --output "repopack.yml"

docker build -t archiver:latest -f src/Dockerfile src

docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/test-output:/output \
  dublok/packmate:latest \
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

# No compression (-2)
docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/test-output:/output \
  archiver:latest \
  --path /source \
  --output /output \
  --name "backup-no-compression" \
  --compression -2 \
  --format json

# Best speed (1)
docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/test-output:/output \
  archiver:latest \
  --path /source \
  --output /output \
  --name "backup-fast" \
  --compression 1 \
  --format json

# Best compression (9)
docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/test-output:/output \
  archiver:latest \
  --path /source \
  --output /output \
  --name "backup-max-compression" \
  --compression 9 \
  --format json