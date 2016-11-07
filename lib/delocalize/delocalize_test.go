package delocalize

import "testing"

func TestIsLocalizedFile(t *testing.T) {
	table := []struct {
		path            string
		isLocalizedFile bool
	}{
		{"~/user/Documents/.localized", true},
		{"/root/.bash_profile", false},
	}

	for _, e := range table {
		result := IsLocalizedFile(e.path)
		if result != e.isLocalizedFile {
			t.Errorf("path (%s) expected localized file (%t) but (%t)",
				e.path, e.isLocalizedFile, result,
			)
		}
	}
}
