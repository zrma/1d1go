package main

import (
	"bufio"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

//goland:noinspection SpellCheckingInspection
func TestDuplicateFileHandler(t *testing.T) {
	t.Log("https://hyperskill.org/projects/240")

	//chopin_waltz7_op64_no2.mp3 / 9765504 bytes
	//vivaldi_four_seasons_winter.mp3 / 9158144 bytes
	//sia_snowman.mp3 / 4590560 bytes
	//gordon_ramsay_chicken_breast.avi / 4590560 bytes
	//smells_like_teen_spirit.mp3 / 4590560 bytes
	//numb.mp3 / 5786312 bytes
	//rick_astley_never_gonna_give_you_up.mp3 / 3422208 bytes
	//unknown.mp3 / 3422208 bytes
	//nea_some_say.mp3 / 3232056 bytes
	//voice.mp3 / 2319746 bytes

	// +---[root_folder]
	//    +---gordon_ramsay_chicken_breast.avi /4590560 bytes
	//    +---[audio]
	//    |   |
	//    |   +---voice.mp3 /2319746 bytes
	//    |   +---sia_snowman.mp3 /4590560 bytes
	//    |   +---nea_some_say.mp3 /3232056 bytes
	//    |   +---[classic]
	//    |   |   |
	//    |   |   +---unknown.mp3 /3422208 bytes
	//    |   |   +---vivaldi_four_seasons_winter.mp3 /9158144 bytes
	//    |   |   +---chopin_waltz7_op64_no2.mp3 /9765504 bytes
	//    |   +---[rock]
	//    |       |
	//    |       +---smells_like_teen_spirit.mp3 /4590560 bytes
	//    |       +---numb.mp3 /5786312 bytes
	//    +---[masterpiece]
	//        |
	//        +---rick_astley_never_gonna_give_you_up.mp3 /3422208 bytes

	assert.NoError(t, fileSystem.Mkdir("root_folder", 0755))
	assert.NoError(t, fileSystem.Mkdir("root_folder/audio", 0755))
	assert.NoError(t, fileSystem.Mkdir("root_folder/audio/classic", 0755))
	assert.NoError(t, fileSystem.Mkdir("root_folder/audio/rock", 0755))
	assert.NoError(t, fileSystem.Mkdir("root_folder/masterpiece", 0755))

	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/audio/classic/chopin_waltz7_op64_no2.mp3", []byte("7777777"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/audio/classic/vivaldi_four_seasons_winter.mp3", []byte("666666"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/audio/sia_snowman.mp3", []byte("55555"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/gordon_ramsay_chicken_breast.avi", []byte("55555"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/audio/rock/smells_like_teen_spirit.mp3", []byte("55555"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/audio/rock/numb.mp3", []byte("4444"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/masterpiece/rick_astley_never_gonna_give_you_up.mp3", []byte("333"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/audio/classic/unknown.mp3", []byte("333"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/audio/nea_some_say.mp3", []byte("22"), 0644))
	assert.NoError(t, afero.WriteFile(fileSystem, "root_folder/audio/voice.mp3", []byte("1"), 0644))

	tests := []struct {
		give string
		want string
	}{
		{
			`mp3
3
2`,
			`Enter file format:
Size sorting options:
1. Descending
2. Ascending

Enter a sorting option:
Wrong option

Enter a sorting option:
3 bytes
root_folder/audio/classic/unknown.mp3
root_folder/masterpiece/rick_astley_never_gonna_give_you_up.mp3

5 bytes
root_folder/audio/rock/smells_like_teen_spirit.mp3
root_folder/audio/sia_snowman.mp3

`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			filepathWalk = walkFunc
			defer func() { filepathWalk = filepath.Walk }()

			scanner := bufio.NewScanner(strings.NewReader(tt.give))
			buf := new(strings.Builder)
			writer := bufio.NewWriter(buf)
			duplicateFileHandler("root_folder", scanner, writer)

			err := writer.Flush()
			assert.NoError(t, err)

			got := buf.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
