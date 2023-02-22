package contentList

import (
	"io/fs"
	"log"
	"obsidian_md/config"
	"path/filepath"
)

type Content struct {
	ObsidianIdentifier string
	HugoIdentifier     string
}

type permalinkStack struct {
	stack []string
}

func (p permalinkStack) push(permalink string) {
	p.stack = append(p.stack, permalink)
}
func (p permalinkStack) pop() {
	if len(p.stack) < 1 {
		return
	}
	p.stack = p.stack[:len(p.stack)-1]
}
func (p permalinkStack) peek() string {
	for i := len(p.stack) - 1; i >= 0; i-- {
		if p.stack[i] != "" {
			return p.stack[i]
		}
	}
	return ""
}

func FromDirectory(path string, config config.Config) ([]Content, error) {
	result := make([]Content, 0)

	// Create map for ignoreFiles
	ignoreMap := make(map[string]bool)
	for _, v := range config.IgnoreFiles {
		ignoreMap[v] = true
	}

	stack := permalinkStack{stack: make([]string, 0)}

	err := filepath.WalkDir(path,
		func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() && (d.Name()[0] == '.' || ignoreMap[d.Name()]) {
				return filepath.SkipDir
			}

			if d.IsDir() {
				stack.pop()
				stack.push(config.Permalinks[d.Name()])
			}

			log.Println(path, d, err, stack.peek())
			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
