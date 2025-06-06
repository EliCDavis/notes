package notes

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type FolderContentEntry struct {
	FileName string `json:"fileName"`
	Content  string `json:"content"`
}

type FolderContents struct {
	Entries []FolderContentEntry `json:"entries"`
	Folder  string               `json:"folder"`
}

func (fc FolderContents) initializeEntry(parentFolder string, entry Entry) error {
	folderPath := path.Join(parentFolder, fc.Folder, entry.Path)
	err := os.MkdirAll(folderPath, os.ModeDir)
	if err != nil {
		return fmt.Errorf("unable to create log's directory: %w", err)
	}

	for _, fce := range fc.Entries {
		filePath := path.Join(folderPath, fce.FileName)
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("unable to create folder contents file: %w", err)
		}
		defer file.Close()

		_, err = file.WriteString(fce.Content)
		if err != nil {
			return err
		}
	}

	return openURL(path.Join(folderPath, fc.Entries[0].FileName))
}

func (fc FolderContents) writeProjectEntryMarkdown(writer *bufio.Writer, options ProjectCompileOptions, parentFolder string, entry Entry) error {
	fmt.Fprintf(writer, "*Created: %s*\n\n", entry.Created)

	for _, fce := range fc.Entries {
		entryPath := filepath.Join(parentFolder, fc.Folder, entry.Path, fce.FileName)
		if !fileExists(entryPath) {
			continue
		}

		if !options.UseMarkdownItIncludeExtension {
			logData, err := os.ReadFile(entryPath)
			if err != nil {
				return err
			}
			writer.Write(logData)
		} else {
			fmt.Fprintf(writer, ":[File](%s)\n", entryPath)
		}
		fmt.Fprint(writer, "\n")
	}

	return nil
}
