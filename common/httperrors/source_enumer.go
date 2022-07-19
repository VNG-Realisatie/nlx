// Code generated by "enumer -type=Source -linecomment -json"; DO NOT EDIT.

package httperrors

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _SourceName = "inwayoutway"

var _SourceIndex = [...]uint8{0, 5, 11}

const _SourceLowerName = "inwayoutway"

func (i Source) String() string {
	i -= 1
	if i < 0 || i >= Source(len(_SourceIndex)-1) {
		return fmt.Sprintf("Source(%d)", i+1)
	}
	return _SourceName[_SourceIndex[i]:_SourceIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _SourceNoOp() {
	var x [1]struct{}
	_ = x[Inway-(1)]
	_ = x[Outway-(2)]
}

var _SourceValues = []Source{Inway, Outway}

var _SourceNameToValueMap = map[string]Source{
	_SourceName[0:5]:       Inway,
	_SourceLowerName[0:5]:  Inway,
	_SourceName[5:11]:      Outway,
	_SourceLowerName[5:11]: Outway,
}

var _SourceNames = []string{
	_SourceName[0:5],
	_SourceName[5:11],
}

// SourceString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func SourceString(s string) (Source, error) {
	if val, ok := _SourceNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _SourceNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Source values", s)
}

// SourceValues returns all values of the enum
func SourceValues() []Source {
	return _SourceValues
}

// SourceStrings returns a slice of all String values of the enum
func SourceStrings() []string {
	strs := make([]string, len(_SourceNames))
	copy(strs, _SourceNames)
	return strs
}

// IsASource returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Source) IsASource() bool {
	for _, v := range _SourceValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Source
func (i Source) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Source
func (i *Source) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Source should be a string, got %s", data)
	}

	var err error
	*i, err = SourceString(s)
	return err
}
