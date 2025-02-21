package main

import (
	"archive/zip"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ArchiveResult struct {
	Path             string `json:"path"`
	ArchivePath      string `json:"archivePath"`
	Status           string `json:"status"`
	Error            string `json:"error,omitempty"`
	CompressionLevel int    `json:"compressionLevel"`
}

func processArchive(path string, outputDir string, customName string, compressionLevel int) (*ArchiveResult, error) {
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("error creating output directory: %v", err)
	}
	path = filepath.Clean(path)

	// Determine the archive filename
	var archiveFilename string
	if customName != "" {
		archiveFilename = customName + ".zip"
	} else {
		base64Path := base64.StdEncoding.EncodeToString([]byte(path))
		archiveFilename = base64Path + ".zip"
	}

	archivePath := filepath.Join(outputDir, archiveFilename)
	result := &ArchiveResult{
		Path:             path,
		ArchivePath:      archivePath,
		CompressionLevel: compressionLevel, // Add this line
	}

	// Check if the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		result.Status = "Failed"
		result.Error = fmt.Sprintf("path does not exist: %s", path)
		return result, nil
	}

	err := createArchive(path, archivePath, compressionLevel) // Pass compressionLevel here
	if err != nil {
		os.Remove(archivePath)
		result.Status = "Failed"
		result.Error = err.Error()
	} else {
		result.Status = "Success"
	}

	return result, nil
}

func createArchive(srcDir, archiveFile string, compressionLevel int) error {
	file, err := os.Create(archiveFile)
	if err != nil {
		return fmt.Errorf("failed to create archive file: %w", err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	// Set compression level based on parameter
	if compressionLevel == -2 {
		// No compression
		zipWriter.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
			return &NoCompression{out}, nil
		})
	} else {
		// Use specified compression level
		zipWriter.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
			return flate.NewWriter(out, compressionLevel)
		})
	}

	const bufferSize = 4 * 1024 * 1024 // 4MB buffer
	buffer := make([]byte, bufferSize)

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return fmt.Errorf("failed to create zip header: %w", err)
		}

		header.Name, err = filepath.Rel(srcDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("failed to write zip header: %w", err)
		}

		if !info.IsDir() {
			data, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("failed to open file: %w", err)
			}
			defer data.Close()

			// Use buffered copy
			_, err = io.CopyBuffer(writer, data, buffer)
			if err != nil {
				return fmt.Errorf("failed to copy file content: %w", err)
			}
		}

		return nil
	})

	return err
}

// NoCompression implements a writer that doesn't compress data
type NoCompression struct {
	w io.Writer
}

func (w *NoCompression) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func (w *NoCompression) Close() error {
	return nil
}

func main() {
	pathFlag := flag.String("path", "", "Path to archive")
	outputFlag := flag.String("output", "", "Output directory for archive file")
	nameFlag := flag.String("name", "", "Custom name for the archive file (optional)")
	outputFormat := flag.String("format", "json", "Output format: 'json' or 'text'")
	compressionFlag := flag.Int("compression", -2,
		"Compression level (-2: none (default), 1: best speed, 9: best compression)")
	flag.Parse()
	flag.Parse()

	if *pathFlag == "" || *outputFlag == "" {
		fmt.Println("Error: Both --path and --output are required")
		return
	}

	result, err := processArchive(*pathFlag, *outputFlag, *nameFlag, *compressionFlag)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

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
