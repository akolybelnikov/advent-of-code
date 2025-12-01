// Package cmd
/*
Copyright © 2025 Andrei Kolybelnikov <a.kolybelnikov@gmail.com>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/akolybelnikov/aoc-cli/internal"
	"github.com/akolybelnikov/aoc-cli/internal/auth"
	"github.com/akolybelnikov/aoc-cli/internal/download"

	"github.com/spf13/cobra"
)

var downloadPath string
var language string
var templatesPath string
var skipDownload bool

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap a solution for a specific day",
	Long: `Bootstrap a solution for a specific day. Requires a valid session and a path to the project root.
Downloaded input will be stored in the /inputs directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if downloadYear == 0 {
			downloadYear = time.Now().Year()
		}
		if day == 0 {
			day = time.Now().Day()
		}
		if day < 1 || day > 25 {
			fmt.Println("Invalid day. Please choose a day between 1 and 25.")
			return
		}

		if downloadPath == "" {
			fmt.Printf("Please provide a path to the project root.\n")
			return
		}

		// Validate language
		if language != "go" && language != "elixir" && language != "rust" {
			fmt.Println("Invalid language. Please choose 'go', 'elixir', or 'rust'.")
			return
		}

		var dayFolder string
		if language == "go" {
			cmdPath := filepath.Join(downloadPath, "cmd")
			err := os.MkdirAll(cmdPath, os.ModePerm)
			if err != nil {
				fmt.Printf("Failed to create directory at %s: %v\n", downloadPath, err)
				return
			}
			dayFolder = filepath.Join(downloadPath, "cmd", fmt.Sprintf("day%02d", day))
		} else if language == "elixir" {
			// For Elixir, create lib and test directories
			libPath := filepath.Join(downloadPath, "lib")
			testPath := filepath.Join(downloadPath, "test")
			err := os.MkdirAll(libPath, os.ModePerm)
			if err != nil {
				fmt.Printf("Failed to create lib directory at %s: %v\n", downloadPath, err)
				return
			}
			err = os.MkdirAll(testPath, os.ModePerm)
			if err != nil {
				fmt.Printf("Failed to create test directory at %s: %v\n", downloadPath, err)
				return
			}
			dayFolder = downloadPath
		} else if language == "rust" {
			// For Rust, create src/bin/dayXX directory
			binPath := filepath.Join(downloadPath, "src", "bin", fmt.Sprintf("day%02d", day))
			err := os.MkdirAll(binPath, os.ModePerm)
			if err != nil {
				fmt.Printf("Failed to create bin directory at %s: %v\n", downloadPath, err)
				return
			}
			dayFolder = binPath
		}

		err := copyTemplate(dayFolder)
		if err != nil {
			fmt.Printf("Failed to copy template: %v\n", err)
			return
		}

		fmt.Printf("Package for Day %02d created successfully!\n", day)

		// If user requested to skip downloading inputs, we're done
		if skipDownload {
			fmt.Println("Skipping input download as requested (--no-download).")
			return
		}

		session, err := auth.GetSession()
		if err != nil {
			fmt.Println("Invalid or expired session. Please run auth to update your session.")
			return
		}

		// If session file read succeeded but returned empty content, prompt the user to run auth
		if strings.TrimSpace(session) == "" {
			fmt.Println("No session token found. Please run 'aoc-cli auth' to save your session token.")
			return
		}

		err = auth.ValidateSession(session, downloadYear)
		if err != nil {
			fmt.Printf("Invalid session: %v. Please run 'aoc-cli auth' to refresh your session.\n", err)
			return
		}

		err = download.Input(downloadYear, day, session, downloadPath)
		if err != nil {
			fmt.Printf("Failed to download the input: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Input for Day %02d downloaded successfully!\n", day)
	},
}

func init() {
	bootstrapCmd.Flags().IntVarP(&day, "day", "d", 0, "Day of Advent of Code (1-25)")
	bootstrapCmd.Flags().IntVarP(&downloadYear, "year", "y", 0, "Advent of Code year (default: current year)")
	bootstrapCmd.Flags().StringVarP(&downloadPath, "path", "p", "", "Custom path for downloading files")
	bootstrapCmd.Flags().StringVarP(&language, "lang", "l", "go", "Language for the solution (go, elixir, rust)")
	// allow overriding where templates are loaded from (root of templates folder)
	bootstrapCmd.Flags().StringVar(&templatesPath, "templates-path", "", "Custom templates root path (overrides embedded templates)")
	bootstrapCmd.Flags().BoolVar(&skipDownload, "no-download", false, "Skip downloading the puzzle input after bootstrapping templates")
	rootCmd.AddCommand(bootstrapCmd)
}

// copyTemplate copies files from templatePath to targetPath
func copyTemplate(targetPath string) error {
	dayStr := fmt.Sprintf("%02d", day)
	// embedded FS uses forward slashes
	embeddedDir := path.Join("templates", language)

	// decide which fs to use: prefer an explicit templatesPath if provided, otherwise embedded
	var tplFS fs.FS = internal.Templates
	var templateDir string = embeddedDir

	if templatesPath != "" {
		// templatesPath is expected to point to the templates root (so language subdir exists beneath it)
		tplFS = os.DirFS(templatesPath)
		templateDir = language
	}

	// Verify that the chosen fs actually contains the templates; try fallbacks if not
	if _, err := fs.ReadDir(tplFS, templateDir); err != nil {
		// first fallback: embedded templates
		if tplFS != internal.Templates {
			if _, err2 := fs.ReadDir(internal.Templates, embeddedDir); err2 == nil {
				tplFS = internal.Templates
				templateDir = embeddedDir
			} else if _, err3 := fs.ReadDir(os.DirFS("."), embeddedDir); err3 == nil {
				// second fallback: current working directory's templates folder
				tplFS = os.DirFS(".")
				templateDir = embeddedDir
			} else {
				return fmt.Errorf("templates directory not found in provided path, embedded assets, or current directory")
			}
		} else {
			// tplFS is embedded but templates not present there; try current directory
			if _, err3 := fs.ReadDir(os.DirFS("."), embeddedDir); err3 == nil {
				tplFS = os.DirFS(".")
				templateDir = embeddedDir
			} else {
				return fmt.Errorf("templates directory '%s' not found in embedded assets or current directory", embeddedDir)
			}
		}
	}

	return fs.WalkDir(tplFS, templateDir, func(fpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the template directory itself
		if fpath == templateDir {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		// Read file content from the selected template FS
		content, err := fs.ReadFile(tplFS, fpath)
		if err != nil {
			return err
		}

		// Replace placeholders
		updatedContent := strings.ReplaceAll(string(content), "{{DAY}}", dayStr)

		var destPath string
		fileName := filepath.Base(fpath)

		if language == "go" {
			updatedContent = strings.ReplaceAll(updatedContent, "package templates", "package main")

			// Handle file renaming logic for Go
			if strings.HasSuffix(fileName, "solution.go") {
				fileName = fmt.Sprintf("day%s.go", dayStr)
			} else if strings.HasSuffix(fileName, "test.go") {
				fileName = fmt.Sprintf("day%s_test.go", dayStr)
			}
			destPath = filepath.Join(targetPath, fileName)
		} else if language == "elixir" {
			// Handle file renaming logic for Elixir
			if strings.HasSuffix(fileName, "solution.ex") {
				fileName = fmt.Sprintf("day%s.ex", dayStr)
				destPath = filepath.Join(targetPath, "lib", fileName)
			} else if strings.HasSuffix(fileName, "test.exs") {
				fileName = fmt.Sprintf("day%s_test.exs", dayStr)
				destPath = filepath.Join(targetPath, "test", fileName)
			}
		} else if language == "rust" {
			// Handle file renaming logic for Rust
			if strings.HasSuffix(fileName, "main.rs") {
				destPath = filepath.Join(targetPath, "main.rs")
			}
		}

		// Write the updated content to the destination file
		return os.WriteFile(destPath, []byte(updatedContent), os.ModePerm)
	})
}
