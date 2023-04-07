package content

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"obsidian_to_hugo/config"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Content struct {
	FrontMatter map[string]string
	Body        string
	dirEntry    os.DirEntry
	fileInfo    os.FileInfo
	permalink   string
	relPath     string
	path        string
}

func (c Content) HugoIdentifier() string {
	if !strings.HasSuffix(c.dirEntry.Name(), ".md") {
		return "/" + c.relPath
	}

	if c.permalink == "" {
		return "/" + strings.TrimSuffix(c.relPath, ".md") + "/"
	}

	trimmedName := strings.TrimSuffix(c.dirEntry.Name(), ".md")

	t, err := time.Parse("2006-01-02 15:04:05 -0700", c.FrontMatter["date"])
	if err != nil {
		return expandPermalink(c.permalink, c.fileInfo.ModTime(), trimmedName) + "/"
	}

	return expandPermalink(c.permalink, t, trimmedName) + "/"
}

func (c Content) ObsidianIdentifier() string {
	return strings.TrimSuffix(c.relPath, ".md")
}

func (c Content) Path() string {
	return c.path
}

func FromDirectory(root string, config config.Config) ([]Content, error) {
	result := make([]Content, 0)

	// Create map for ignoreFiles
	ignoreMap := make(map[string]bool)
	for _, v := range config.IgnoreFiles {
		ignoreMap[v] = true
	}

	err := filepath.WalkDir(root,
		func(path string, d fs.DirEntry, err error) error {

			// Skip dot and ignored directories
			if d.IsDir() && (d.Name()[0] == '.' || ignoreMap[d.Name()]) {
				return filepath.SkipDir
			}

			// Omit directories and dot files
			if d.IsDir() || d.Name()[0] == '.' {
				return nil
			}

			pathComponents := strings.Split(path, string(os.PathSeparator))
			permalink := ""
			for i := len(pathComponents) - 1; i >= 0; i-- {
				permalink = config.Permalinks[pathComponents[i]]
				if permalink != "" {
					break
				}
			}

			relPath, _ := filepath.Rel(root, path)
			fi, _ := os.Stat(path)

			c, err := fromFile(path)
			c.dirEntry = d
			c.fileInfo = fi
			c.permalink = permalink
			c.relPath = relPath
			c.path = path

			result = append(result, c)

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func expandPermalink(permalink string, time time.Time, filename string) string {
	// TODO: Need to support all permalink parameters available
	result := permalink
	result = strings.ReplaceAll(result, ":2006", fmt.Sprintf("%04d", time.Year()))
	result = strings.ReplaceAll(result, ":01", fmt.Sprintf("%02d", time.Month()))
	result = strings.ReplaceAll(result, ":02", fmt.Sprintf("%02d", time.Day()))
	result = strings.ReplaceAll(result, ":filename", filename)
	return result
}

func fromFile(path string) (Content, error) {
	var result Content

	f, err := os.Open(path)
	if err != nil {
		return Content{}, err
	}

	s := bufio.NewScanner(f)

	frontMatter := ""
	body := ""

	if s.Scan() && s.Text() == "---" { //has frontMatter
		for s.Scan() && s.Text() != "---" {
			frontMatter += s.Text() + "\n"
		}
	}

	for s.Scan() {
		body += s.Text() + "\n"
	}

	frontMap := make(map[string]string)

	// suppress error as time.Parse can handle it
	_ = yaml.Unmarshal([]byte(frontMatter), &frontMap)
	result.FrontMatter = frontMap

	if len(body) > 1 {
		result.Body = body[:len(body)-1] // remove trailing newline
	}

	return result, nil
}
