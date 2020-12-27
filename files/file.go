package files

import (
	"bufio"
	"os"
)

// ReadFile reads contents of the file given by the Path and returns each line in the file as a string
func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}
