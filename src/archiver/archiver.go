package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ArchiveResult struct {
	Path             string `json:"path"`
	ArchivePath      string `json:"archivePath"`
	Status           string `json:"status"`
	Error            string `json:"error,omitempty"`
	CompressionLevel int    `json:"compressionLevel"`
}

func processArchive(path, outputDir, customName string, compressionLevel int, extraFlags []string) (*ArchiveResult, error) {
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("error creating output directory: %v", err)
	}
	path = filepath.Clean(path)

	// Determine the archive filename
	var archiveFilename string
	if customName != "" {
		archiveFilename = customName
	} else {
		base64Path := base64.StdEncoding.EncodeToString([]byte(path))
		archiveFilename = base64Path
	}

	// Construct the full archive path
	archivePath := filepath.Join(outputDir, archiveFilename+".7z")
	result := &ArchiveResult{
		Path:             path,
		ArchivePath:      archivePath,
		CompressionLevel: compressionLevel,
	}

	// Check if the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		result.Status = "Failed"
		result.Error = fmt.Sprintf("path does not exist: %s", path)
		return result, nil
	}

	// Create the archive using p7zip
	err := create7zArchive(path, archivePath, compressionLevel, extraFlags)
	if err != nil {
		os.Remove(archivePath)
		result.Status = "Failed"
		result.Error = err.Error()
	} else {
		result.Status = "Success"
	}
	return result, nil
}

func create7zArchive(srcDir, archiveFile string, compressionLevel int, extraFlags []string) error {
	// Define the base command
	args := []string{"a", "-t7z"}

	// Add compression level flag
	args = append(args, fmt.Sprintf("-mx=%d", compressionLevel))

	// Add extra flags (e.g., -ms=off, -m0=copy, etc.)
	args = append(args, extraFlags...)

	// Add the archive file and source directory
	args = append(args, archiveFile, srcDir)

	// Execute the 7z command
	cmd := exec.Command("7z", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create 7z archive: %s\n%s", err, string(output))
	}
	return nil
}

func main() {
	// Define flags
	sourceFlag := flag.String("source", "/source", "Source directory to archive (default: /source)")
	outputFlag := flag.String("output", "/output", "Output directory for the archive (default: /output)")
	nameFlag := flag.String("name", "", "Custom name for the archive file (optional)")
	outputFormat := flag.String("format", "json", "Output format: 'json' or 'text'")
	compressionFlag := flag.Int("compression", 5, "Compression level (0: none, 1: best speed, 9: best compression)")
	methodFlag := flag.String("method", "", "Compression method (e.g., lzma2, ppmd)")
	passwordFlag := flag.String("password", "", "Password for encrypting the archive")
	headerEncryptionFlag := flag.Bool("header-encryption", false, "Enable header encryption")
	volumeSizeFlag := flag.String("volume-size", "", "Create multi-volume archive with specified size (e.g., 100m)")
	threadsFlag := flag.Bool("multithreading", true, "Enable multithreading")
	extraFlag := flag.String("extra", "", "Additional p7zip flags (e.g., -ms=off)")

	flag.Parse()

	// Build extra flags for p7zip
	extraFlags := []string{}
	if *methodFlag != "" {
		extraFlags = append(extraFlags, fmt.Sprintf("-m0=%s", *methodFlag))
	}
	if *passwordFlag != "" {
		extraFlags = append(extraFlags, fmt.Sprintf("-p%s", *passwordFlag))
		if *headerEncryptionFlag {
			extraFlags = append(extraFlags, "-mhe=on")
		}
	}
	if *volumeSizeFlag != "" {
		extraFlags = append(extraFlags, fmt.Sprintf("-v%s", *volumeSizeFlag))
	}
	if !*threadsFlag {
		extraFlags = append(extraFlags, "-mt=off")
	}
	if *extraFlag != "" {
		extraFlags = append(extraFlags, strings.Split(*extraFlag, " ")...)
	}

	// Process the archive
	result, err := processArchive(*sourceFlag, *outputFlag, *nameFlag, *compressionFlag, extraFlags)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Output the result
	if *outputFormat == "json" {
		jsonOutput, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Printf("\nArchive Creation Result:\n")
		fmt.Println("=======================")
		if result.Error != "" {
			fmt.Printf("\n❌ Path: %s\n", result.Path)
			fmt.Printf("   Archive: %s\n", result.ArchivePath)
			fmt.Printf("   Compression Level: %d\n", result.CompressionLevel)
			fmt.Printf("   Status: %s\n", result.Status)
			fmt.Printf("   Error: %s\n", result.Error)
		} else {
			fmt.Printf("\n✅ Path: %s\n", result.Path)
			fmt.Printf("   Archive: %s\n", result.ArchivePath)
			fmt.Printf("   Compression Level: %d\n", result.CompressionLevel)
			fmt.Printf("   Status: %s\n", result.Status)
		}
	}
}
