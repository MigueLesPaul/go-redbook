package obsidianfrontmatter

import (
	"bufio"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// ReadFrontMatter reads a markdown file and extracts the YAML frontmatter as a map.
func ReadFrontMatter(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var frontmatterLines []string
	inFrontMatter := false
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		if lineNum == 1 {
			if strings.TrimSpace(line) != "---" {
				return nil, errors.New("frontmatter must start with '---'")
			}
			inFrontMatter = true
			continue
		}
		if inFrontMatter {
			if strings.TrimSpace(line) == "---" {
				// End of frontmatter
				break
			}
			frontmatterLines = append(frontmatterLines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error Leyendo archivo %v", err)
		return nil, err
	}

	frontmatterContent := strings.Join(frontmatterLines, "\n")
	var frontmatter map[string]interface{}
	err = yaml.Unmarshal([]byte(frontmatterContent), &frontmatter)
	if err != nil {
		return nil, err
	}

	return frontmatter, nil
}

// LoadFrontMattersFromDir lists markdown files and loads their frontmatters.
func LoadFrontMattersFromDir(dir string) ([]map[string]interface{}, error) {
	var frontmatters []map[string]interface{}

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".md" || filepath.Ext(path) == ".markdown" {
			fm, err := ReadFrontMatter(path)
			if err == nil {
				frontmatters = append(frontmatters, fm)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return frontmatters, nil
}
