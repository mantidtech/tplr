package tplr

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetFileReader provides unit test coverage for GetFileReader()
func TestGetFileReader(t *testing.T) {
	type Args struct {
		filename string
	}

	tests := []struct {
		name         string
		args         Args
		wantIoReader io.Reader
		wantContent  []byte
		wantError    bool
	}{
		{
			name: "stdin",
			args: Args{
				filename: "-",
			},
			wantIoReader: os.Stdin,
			wantError:    false,
		},
		{
			name: "file exists",
			args: Args{
				filename: "testdata/a_file.txt",
			},
			wantContent: []byte("Just a plain text file\n"),
			wantError:   false,
		},
		{
			name: "file doesn't exist",
			args: Args{
				filename: "NO SUCH FILE",
			},
			wantError: true,
		},
		{
			name: "file is not a plain file",
			args: Args{
				filename: "testdata",
			},
			wantError: true,
		},
		{
			name: "file is not readable",
			args: Args{
				filename: "testdata/secret",
			},
			wantError: true,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotIoReader, gotError := GetFileReader(tt.args.filename)

			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}

			if tt.wantContent != nil {
				d, err := ioutil.ReadAll(gotIoReader)
				require.NoError(t, err)
				assert.Equal(t, tt.wantContent, d)
			}

			if tt.wantIoReader != nil {
				assert.Equal(t, tt.wantIoReader, gotIoReader)
			}
		})
	}
}

// TestGetFileWriter provides unit test coverage for GetFileWriter()
func TestGetFileWriter(t *testing.T) {
	type Args struct {
		filename string
		force    bool
	}

	tests := []struct {
		name         string
		args         Args
		wantIoWriter io.Writer
		wantError    bool
	}{
		{
			name: "stdin",
			args: Args{
				filename: "-",
			},
			wantIoWriter: os.Stdout,
			wantError:    false,
		},
		{
			name: "file is not writable",
			args: Args{
				filename: "testdata/secret",
				force:    true,
			},
			wantError: true,
		},
		{
			name: "file exists, but don't force overwrite",
			args: Args{
				filename: "testdata/a_file.txt",
				force:    false,
			},
			wantError: true,
		},
		{
			name: "file exists, and force overwrite",
			args: Args{
				filename: "testdata/empty",
				force:    true,
			},
			wantError: false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotIoWriter, gotError := GetFileWriter(tt.args.filename, tt.args.force)

			if tt.wantError {
				require.Error(t, gotError)
			} else {
				require.NoError(t, gotError)
			}

			if tt.wantIoWriter != nil {
				assert.Equal(t, tt.wantIoWriter, gotIoWriter)
			}
		})
	}
}

// TestFileExists provides unit test coverage for FileExists()
func TestFileExists(t *testing.T) {
	type Args struct {
		filename string
	}

	tests := []struct {
		name string
		args Args
		want bool
	}{
		{
			name: "file exists",
			args: Args{
				filename: "testdata/empty",
			},
			want: true,
		},
		{
			name: "file doesn't exist",
			args: Args{
				filename: "NO SUCH FILE",
			},
			want: false,
		},
		{
			name: "file exists but is not readable",
			args: Args{
				filename: "testdata/secret",
			},
			want: true,
		},
		{
			name: "file exists but is not a plain file",
			args: Args{
				filename: "testdata",
			},
			want: false,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := FileExists(tt.args.filename)
			assert.Equal(t, tt.want, got)
		})
	}
}
