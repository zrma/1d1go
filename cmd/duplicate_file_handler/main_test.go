package main

import (
	"bufio"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

//goland:noinspection SpellCheckingInspection
func TestDuplicateFileHandler(t *testing.T) {
	t.Log("https://hyperskill.org/projects/240")

	// +---[root_folder]
	//    |
	//    +---wall.png
	//    +---pass.txt
	//    +---[docs]
	//    |   |
	//    |   +---project.py
	//    |   +---calc.xls
	//    |   +---tutorial.mp4
	//    |   +---[res]
	//    |       |
	//    |       +---data.json
	//    |   +---[output]
	//    |       |
	//    |       +---result.json
	//    +---[masterpiece]
	//        |
	//        +---rick_astley_never_gonna_give_you_up.mp3

	assert.NoError(t, fileSystem.Mkdir("root_folder", 0755))
	assert.NoError(t, fileSystem.Mkdir("root_folder/docs", 0755))
	assert.NoError(t, fileSystem.Mkdir("root_folder/docs/res", 0755))
	assert.NoError(t, fileSystem.Mkdir("root_folder/docs/output", 0755))
	assert.NoError(t, fileSystem.Mkdir("root_folder/masterpiece", 0755))

	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/wall.png", []byte("wall.png"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/pass.txt", []byte("pass.txt"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/docs/project.py", []byte("project.py"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/docs/calc.xls", []byte("calc.xls"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/docs/tutorial.mp4", []byte("tutorial.mp4"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/docs/res/data.json", []byte("data.json"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/docs/output/result.json", []byte("result.json"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/masterpiece/rick_astley_never_gonna_give_you_up.mp3", []byte("rick_astley_never_gonna_give_you_up.mp3"), 0644))

	const want = `root_folder/wall.png
root_folder/pass.txt
root_folder/docs/project.py
root_folder/docs/calc.xls
root_folder/docs/tutorial.mp4
root_folder/docs/res/data.json
root_folder/docs/output/result.json
root_folder/masterpiece/rick_astley_never_gonna_give_you_up.mp3
`

	filepathWalk = walkFunc
	defer func() { filepathWalk = filepath.Walk }()

	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	duplicateFileHandler("root_folder", writer)

	err := writer.Flush()
	assert.NoError(t, err)

	got := buf.String()
	assert.ElementsMatch(t, strings.Split(want, "\n"), strings.Split(got, "\n"))
}
