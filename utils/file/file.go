package file

import (
	"io/ioutil"
	"os"
)

// ReadFile reads the content of a file and returns it as a string.
func ReadFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile writes the given content to a file.
func WriteFile(filePath string, content string) error {
	return ioutil.WriteFile(filePath, []byte(content), 0644)
}

// AppendToFile appends the given content to a file.
func AppendToFile(filePath string, content string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFile deletes the specified file.
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// CreateFile creates a new file with the given content.
func CreateFile(filePath string, content string) error {
	return ioutil.WriteFile(filePath, []byte(content), 0644)
}

// FileExists checks if a file exists at the specified path.
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// IsDirectory checks if the specified path is a directory.
func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// CreateDirectory creates a new directory at the specified path.
func CreateDirectory(dirPath string) error {
	return os.MkdirAll(dirPath, os.ModePerm)
}
