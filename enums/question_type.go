package enums

import (
	"fmt"
	"io"
	"strconv"
)

type QuestionType string

const (
	MultipleChoiceSingleSelection   QuestionType = "MULTIPLE_CHOICE_SINGLE_SELECTION"
	MultipleChoiceMultipleSelection QuestionType = "MULTIPLE_CHOICE_MULTIPLE_SELECTION"
	FreeResponse                    QuestionType = "FREE_RESPONSE"
	TextTagging                     QuestionType = "TEXT_TAGGING"
	Ranking                         QuestionType = "RANKING"
	Chatbot                         QuestionType = "CHATBOT"
)

// Values provides list valid values for Enum.
func (QuestionType) Values() (kinds []string) {
	for _, s := range []QuestionType{MultipleChoiceSingleSelection, MultipleChoiceMultipleSelection, FreeResponse, TextTagging, Ranking, Chatbot} {
		kinds = append(kinds, string(s))
	}
	return
}

// MarshalGQL implements graphql.Marshaler interface.
func (_g QuestionType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(string(_g)))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (_g *QuestionType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*_g = QuestionType(str)

	switch *_g {
	case MultipleChoiceSingleSelection, MultipleChoiceMultipleSelection, FreeResponse, TextTagging, Ranking, Chatbot:
		return nil
	default:
		return fmt.Errorf("%s is not a valid Gender", str)
	}
}
