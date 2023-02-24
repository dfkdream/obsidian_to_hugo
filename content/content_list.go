package content

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"obsidian_md/config"
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

func FromDirectory(root string, config config.Config) ([]Content, error) {
	result := make([]Content, 0)

	// Create map for ignoreFiles
	ignoreMap := make(map[string]bool)
	for _, v := range config.IgnoreFiles {
		ignoreMap[v] = true
	}

	err := filepath.WalkDir(root,
		func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() && (d.Name()[0] == '.' || ignoreMap[d.Name()]) {
				return filepath.SkipDir
			}

			if d.IsDir() {
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

			result = append(result, Content{
				FrontMatter: getFrontMatter(path),
				Body:        "", // TODO: Add body
				dirEntry:    d,
				fileInfo:    fi,
				permalink:   permalink,
				relPath:     relPath,
			})

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func expandPermalink(permalink string, time time.Time, filename string) string {
	result := permalink
	result = strings.ReplaceAll(result, ":2006", fmt.Sprintf("%04d", time.Year()))
	result = strings.ReplaceAll(result, ":01", fmt.Sprintf("%02d", time.Month()))
	result = strings.ReplaceAll(result, ":02", fmt.Sprintf("%02d", time.Day()))
	result = strings.ReplaceAll(result, ":filename", filename)
	return result
}

func getFrontMatter(path string) map[string]string {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}

	s := bufio.NewScanner(f)

	frontMatter := ""
	needToAdd := false
	for s.Scan() {
		if s.Text() == "---" {
			if !needToAdd {
				needToAdd = true
				continue
			} else {
				break
			}
		}

		if needToAdd {
			frontMatter += s.Text() + "\n"
		}
	}

	frontMap := make(map[string]string)

	// suppress error as time.Parse can handle it
	_ = yaml.Unmarshal([]byte(frontMatter), &frontMap)

	return frontMap
}
