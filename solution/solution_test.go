package solution

import (
	"github.com/kyokomi/emoji"
	"testing"
)

func TestSolution(t *testing.T) {
	if GetMessage() != emoji.Sprint("Hello :world_map:!") {
		t.Errorf("Wrong String")
	}
}
