package tplr

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestLoadAndGenerateFromTemplate provides unit test coverage for Load() and GenerateFromTemplate()
func TestLoadAndGenerateFromTemplate(t *testing.T) {

	tests := []struct {
		name              string
		tpl               string
		vars              map[string]interface{}
		want              string
		wantLoadError     bool
		wantGenerateError bool
	}{
		{
			name: "hello world",
			tpl:  "Hello {{.to}}!",
			vars: map[string]interface{}{
				"to": "World",
			},
			want:              "Hello World!",
			wantLoadError:     false,
			wantGenerateError: false,
		},
		{
			name: "bad template",
			tpl:  "Hello {{.to",
			vars: map[string]interface{}{
				"to": "World",
			},
			want:              "",
			wantLoadError:     true,
			wantGenerateError: false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := bytes.NewBufferString(tt.tpl)
			tp, loadError := Load(r)
			if tt.wantLoadError {
				require.Error(t, loadError)
				return
			}
			require.NoError(t, loadError)

			var got bytes.Buffer
			generateError := GenerateFromTemplate(&got, tp, tt.vars)
			if tt.wantGenerateError {
				require.Error(t, generateError)
			} else {
				require.NoError(t, generateError)
			}

			assert.Equal(t, tt.want, got.String())
		})
	}
}
