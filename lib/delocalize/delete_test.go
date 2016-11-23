package delocalize

import (
	"fmt"
	"testing"
)

func TestDeleteDispatcher(t *testing.T) {
	table := []struct {
		path []string
	}{
		{path: []string{
			"./test/.localized",
		}},
	}

	for _, e := range table {
		t.Run(fmt.Sprintf("delete: %s", e.path), func(t *testing.T) {
			dd := NewDeleteDispatcher(5, 10, DeleteModeDebugPrint)

			for _, p := range e.path {
				dd.Add(p)
			}

			dd.Start()
			dd.Wait()
		})
	}
}
