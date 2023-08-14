package enums

import (
	"fmt"
	"io"
	"strconv"
)

type QualificationValue string

const (
	Coder   QualificationValue = "CODER"
	Medical QualificationValue = "MEDICAL"
	Artist  QualificationValue = "ARTIST"
	STEM    QualificationValue = "STEM"
	Other   QualificationValue = "OTHER"
)

// Values provides list valid values for Enum.
func (QualificationValue) Values() (kinds []string) {
	for _, s := range []QualificationValue{Coder, Medical, Artist, STEM} {
		kinds = append(kinds, string(s))
	}
	return
}

// MarshalGQL implements graphql.Marshaler interface.
func (_g QualificationValue) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(string(_g)))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (_g *QualificationValue) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*_g = QualificationValue(str)

	switch *_g {
	case Coder, Medical, Artist, STEM, Other:
		return nil
	default:
		return fmt.Errorf("%s is not a valid Gender", str)
	}
}
