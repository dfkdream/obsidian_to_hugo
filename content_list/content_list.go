package contentList

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
	ObsidianIdentifier string
	HugoIdentifier     string
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

			if strings.HasSuffix(path, ".md") {
				// if markdown file
				t, err := getTimeFromFile(path)
				if err != nil {
					return err
				}

				if permalink == "" {
					result = append(result, Content{
						ObsidianIdentifier: strings.TrimSuffix(relPath, ".md"),
						HugoIdentifier:     "/" + strings.TrimSuffix(relPath, ".md") + "/",
					})
				} else {
					result = append(result, Content{
						ObsidianIdentifier: strings.TrimSuffix(relPath, ".md"),
						HugoIdentifier:     expandPermalink(permalink, t, strings.TrimSuffix(d.Name(), ".md")) + "/",
					})

				}

			} else {
				result = append(result, Content{
					ObsidianIdentifier: relPath,
					HugoIdentifier:     "/" + relPath,
				})
			}

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

func getTimeFromFile(path string) (time.Time, error) {
	f, err := os.Open(path)
	if err != nil {
		return time.Time{}, err
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

	t, err := time.Parse("2006-01-02 15:04:05 -0700", frontMap["date"])
	if err == nil {
		return t, nil
	}

	stat, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}

	return stat.ModTime(), err
}
