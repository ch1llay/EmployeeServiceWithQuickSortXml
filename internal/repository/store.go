package repository

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type SqlReader struct {
	ScriptPath string
}

func GetSqlReader(scriptPath string) *SqlReader {
	return &SqlReader{ScriptPath: scriptPath}
}

func (s *SqlReader) GetSqlFromFile(scriptName string) string {
	path := filepath.Join(s.ScriptPath, scriptName, ".sql")
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		panic(fmt.Sprintf("файл %s не читается", path))
	}
	return string(c)

}
