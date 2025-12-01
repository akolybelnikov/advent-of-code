package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Input(year, day int, session string, path string) error {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	destPath := filepath.Join(path, "inputs", fmt.Sprintf("day%02d.txt", day))
	fmt.Println(destPath)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", "session="+session)
	// Advent of Code may reject requests without a proper User-Agent — set one
	req.Header.Set("User-Agent", "aoc-cli/1.0 (+https://github.com/akolybelnikov)")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		// try to read a small snippet of the body for helpful diagnostic info
		buf := make([]byte, 1024)
		n, _ := resp.Body.Read(buf)
		snippet := strings.TrimSpace(string(buf[:n]))
		if snippet != "" {
			return fmt.Errorf("failed to download %s: %s (body: %.200s)", url, resp.Status, snippet)
		}
		return fmt.Errorf("failed to download %s: %s", url, resp.Status)
	}

	err = os.MkdirAll(filepath.Dir(destPath), 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(destPath)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = io.Copy(file, resp.Body)
	return err
}
