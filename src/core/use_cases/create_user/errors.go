package createuser

import (
	"fmt"
	"strings"
)

type MissingAttributesError struct {
	MissingAttributes []string
}

func (e *MissingAttributesError) AppendMissingAttribute(name string) {
	e.MissingAttributes = append(e.MissingAttributes, name)
}

func (e *MissingAttributesError) AnyMissingAttribute() bool {
	return len(e.MissingAttributes) > 0
}

func (e *MissingAttributesError) Error() string {
	return fmt.Sprintf("Missing attributes: %s", strings.Join(e.MissingAttributes, ", "))
}
