package go_basic

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeManager(t *testing.T) {
	for i, tt := range []struct {
		manager *Manager
		marshal func(any) ([]byte, error)
		err     error
		want    string
	}{
		{
			manager: &Manager{
				FullName:       "Jack Oliver",
				Position:       "CEO",
				Age:            44,
				YearsInCompany: 0,
			},
			want: `{"full_name":"Jack Oliver","position":"CEO","age":44}`,
		},
		{
			manager: &Manager{
				FullName: `Jack Oliver
`,
				Age:            44,
				YearsInCompany: 8,
			},
			want: `{"full_name":"Jack Oliver\n","age":44,"years_in_company":8}`,
		},
		{
			manager: &Manager{
				FullName:       "abc",
				Position:       "de",
				Age:            123,
				YearsInCompany: 456,
			},
			marshal: func(got any) ([]byte, error) {
				m, ok := got.(*Manager)
				assert.True(t, ok)
				assert.Equal(t, "abc", m.FullName)
				assert.Equal(t, "de", m.Position)
				assert.EqualValues(t, 123, m.Age)
				assert.EqualValues(t, 456, m.YearsInCompany)
				return nil, fmt.Errorf("foobar")
			},
			err: fmt.Errorf("foobar"),
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if tt.marshal != nil {
				marshal = tt.marshal
			}
			t.Cleanup(func() {
				marshal = json.Marshal
			})
			reader, err := EncodeManager(tt.manager)
			assert.Equal(t, tt.err, err)

			if err == nil {
				got, err2 := io.ReadAll(reader)
				assert.NoError(t, err2)
				assert.Equal(t, tt.want, string(got))
			}
		})
	}
}
