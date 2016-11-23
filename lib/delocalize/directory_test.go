package delocalize

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

type (
	TestDeleter struct {
		t *testing.T
	}
)

func (t *TestDeleter) Add(path string) {
	if !strings.HasSuffix(path, ".localized") {
		t.t.Errorf("this is not localized file: %s", path)
	}
}

func TestDirectoryDispatcher(t *testing.T) {
	table := []struct {
		path string
	}{
		{path: "./"},
	}

	for _, e := range table {
		t.Run(fmt.Sprintf("research %s", e.path), func(t *testing.T) {

			d := &TestDeleter{
				t: t,
			}
			dd := NewDirectoryDispatcher(5, 5, d)

			path, err := filepath.Abs(e.path)
			if err != nil {
				t.Fatal(err)
			}

			dd.Add(path)
			dd.Start()
			dd.Wait()
		})
	}
}
