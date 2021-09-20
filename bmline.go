package bmline

// Body lines and elements
import (
	"github.com/mmirko/bmmeta"
)

type BasmBody struct {
	*bmmeta.BasmMeta
	lines []*BasmLine
}

type BasmLine struct {
	*bmmeta.BasmMeta
	operand  *BasmElement
	elements []*BasmElement
}

type BasmElement struct {
	*bmmeta.BasmMeta
	string
}
