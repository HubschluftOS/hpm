package cmd

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func UntarGzFile(filename string) error {
	fmt.Printf("[1/1] untarring %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("[0/1] failed to open file: %s", err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("[0/1] failed to create gzip reader: %s", err)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("[0/1] error reading tar header: %s", err)
		}

		targetPath := filepath.Join(".", header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return fmt.Errorf("[0/1] failed to create directory: %s", err)
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return fmt.Errorf("[0/1] failed to create parent directories: %s", err)
			}

			outFile, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("[0/1] failed to create file: %s", err)
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, tarReader); err != nil {
				return fmt.Errorf("[0/1] error extracting file %s: %s", header.Name, err)
			}
			fmt.Printf("[1/1] extracted: %s\n", header.Name)
		default:
			fmt.Printf("[1/1] unsupported tar entry type: %d\n", header.Typeflag)
		}
	}

	fmt.Printf("[1/1] extraction completed\n")
	return nil
}
