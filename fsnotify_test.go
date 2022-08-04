//go:build !plan9
// +build !plan9

package fsnotify

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestEventString(t *testing.T) {
	tests := []struct {
		in   Event
		want string
	}{
		{Event{}, `"": `},
		{Event{"/file", 0}, `"/file": `},

		{Event{"/file", Chmod | Create},
			`"/file": CREATE|CHMOD`},
		{Event{"/file", Rename},
			`"/file": RENAME`},
		{Event{"/file", Remove},
			`"/file": REMOVE`},
		{Event{"/file", Write | Chmod},
			`"/file": WRITE|CHMOD`},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			have := tt.in.String()
			if have != tt.want {
				t.Errorf("\nhave: %q\nwant: %q", have, tt.want)
			}
		})
	}
}

func TestFindDirs(t *testing.T) {
	join := func(list ...string) string {
		return "\n\t" + strings.Join(list, "\n\t")
	}

	t.Run("finds dirs", func(t *testing.T) {
		tmp := t.TempDir()

		mkdirAll(t, tmp, "/one/two/three/four")
		cat(t, "asd", tmp, "one/two/file.txt")
		symlink(t, "/", tmp, "link")

		dirs, err := findDirs(tmp)
		if err != nil {
			t.Fatal(err)
		}

		have := join(dirs...)
		want := join([]string{
			tmp,
			filepath.Join(tmp, "one"),
			filepath.Join(tmp, "one/two"),
			filepath.Join(tmp, "one/two/three"),
			filepath.Join(tmp, "one/two/three/four"),
		}...)

		if have != want {
			t.Errorf("\nhave: %s\nwant: %s", have, want)
		}
	})

	t.Run("file", func(t *testing.T) {
		tmp := t.TempDir()
		cat(t, "asd", tmp, "file")

		dirs, err := findDirs(filepath.Join(tmp, "file"))
		if !errorContains(err, "not a directory") {
			t.Errorf("wrong error: %s", err)
		}
		if len(dirs) > 0 {
			t.Errorf("dirs contains entries: %s", dirs)
		}
	})
}
