repomix --no-file-summary --no-security-check --include "src/Dockerfile,src/archiver/archiver.go" --output "repopack.yml"
repomix --no-file-summary --no-security-check --include "README.md,src/Dockerfile,src/archiver/archiver.go" --output "repopack.yml"

go build -o ../../bin/archiver ./archiver.go
go run ./src/archiver/archiver.go -source=/home/ercin/go -output=./output -compression=-2 -name=myarchive.zip -format=json
go run ./src/archiver/archiver.go -source=/home/ercin/go -output=./output -compression=-2 -name=myarchive.zip -format=json




docker build --progress=plain --no-cache -t dublok/packmate:latest -f src/Dockerfile src

docker run --rm \
  -v /home/ercin/go:/source \
  -v /home/ercin/github/ercindedeoglu/packmate/output:/output \
  dublok/packmate:latest \
  -source=/source \
  -output=/output \
  -compression=0 \
  -method=copy \
  -multithreading=true \
  -extra="-ms=off"


  docker run --rm \
  -v /home/ercin/github/ercindedeoglu/packmate/output:/source \
  -v /home/ercin/github/ercindedeoglu/packmate/x:/output \
  dublok/packmate:latest \
  -source=/source \
  -output=/output \
  -compression=0 \
  -method=copy \
  -multithreading=true \
  -extra="-ms=off"

docker run --rm \
  -v /home/ercin/go:/source \
  -v /home/ercin/github/ercindedeoglu/packmate/output:/output \
  dublok/packmate:latest \
  -source=/source \
  -output=/output \
  -method=lzma2 \
  -password=secret \
  -header-encryption \  
  -compression=9 \
  -volume-size=100m

  docker run --rm \
  -v /home/ercin/go:/source \
  -v /home/ercin/github/ercindedeoglu/packmate/output:/output \
  dublok/packmate:latest \
  -source=/source \
  -output=/output \
  -compression=0 \
  -method=copy \
  -multithreading=true \
  -extra="-ms=off"

docker run --rm \
  -v /home/ercin/go:/source:ro \
  -v /home/ercin/github/ercindedeoglu/packmate/output:/output \-
  dublok/packmate:latest \
  --name "backup-$(date +%Y%m%d)"

docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/github/ercindedeoglu/packmate/output:/output \
  dublok/packmate:latest \
  --name "backup-$(date +%Y%m%d)"

docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_logs/_data:/source:ro \
  -v /home/ercin/github/ercindedeoglu/packmate/output:/output \
  dublok/packmate:latest \
  --name "backup-$(date +%Y%m%d)"

  docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/test-output:/output \
  dublok/packmate:latest \
  --name "my-backup" \
  --format text

# No compression (-2)
docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/test-output:/output \
  archiver:latest \
  --name "backup-no-compression" \
  --compression -2 \
  --format json

# Best speed (1)
docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/test-output:/output \
    dublok/packmate:latest \
  --name "backup-fast" \
  --compression 1 \
  --format json

# Best compression (9)
docker run --rm \
  -v /var/lib/docker/volumes/docker-volume-test_test1_data/_data:/source:ro \
  -v /home/ercin/test-output:/output \
  dublok/packmate:latest \
  --name "backup-max-compression" \
  --compression 9 \
  --format json

  sudo chown -R ercin:ercin /home/ercin/github/ercindedeoglu/packmate/output

  docker run --rm   -v /home/ercin/go:/source:ro   -v /home/ercin/github/ercindedeoglu/packmate/output:/output   dublok/packmate:latest   --name "backup-$(date +%Y%m%d)"

  docker run --rm   -v /home/ercin/github/ercindedeoglu/packmate/output:/source:ro   -v /home/ercin/github/ercindedeoglu/packmate/x:/output   dublok/packmate:latest   --name "backup-$(date +%Y%m%d)"