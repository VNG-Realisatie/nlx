// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package transactionlog

import (
	"testing"
)

func TestParseDataSubjectHeaderHappy(t *testing.T) {
	headerString := "key1=value1 ,key2= value2 , key3 =value3,key4=value4"
	keyValuesPairs, err := ParseDataSubjectHeader(headerString)
	if err != nil {
		t.Errorf("failed to parse header %s", err)
	}

	expectedResult := make(map[string]string)
	expectedResult["key1"] = "value1"
	expectedResult["key2"] = "value2"
	expectedResult["key3"] = "value3"
	expectedResult["key4"] = "value4"

	valuesChecked := make(map[string]string)
	for key, value := range keyValuesPairs {
		if len(valuesChecked[key]) != 0 {
			continue
		}

		v, exists := expectedResult[key]
		if !exists {
			t.Errorf("parsed invalid key %s", key)
		}

		if v != value {
			t.Errorf("value: %s expected: %s", value, v)
		}

		valuesChecked[key] = value
	}

	if len(valuesChecked) != 4 {
		t.Errorf("excepted to find 4 key value pairs but found %d", len(valuesChecked))
	}
}

func TestParseDataSubjectHeaderExceptionEmptyKeyValuePair(t *testing.T) {
	headerString := "key1=value1,,key3=value3"
	_, err := ParseDataSubjectHeader(headerString)
	if err.Error() != "invalid datasubject in header, 2th subject is not a correct key=value format" {
		t.Errorf("result: %s expected: invalid datasubject in header, 2th subject is not a correct key=value format", err)
	}

}

func TestParseDataSubjectHeaderExceptionIncorrectKeyValuePair(t *testing.T) {
	headerString := "key1=value1,key2=value2,key3"
	_, err := ParseDataSubjectHeader(headerString)
	if err.Error() != "invalid datasubject in header, 3th subject is not a correct key=value format" {
		t.Errorf("result: %s expected: invalid datasubject in header, 3th subject is not a correct key=value format", err)
	}

}

func TestParseDataSubjectHeaderExceptionFailRegExKey(t *testing.T) {
	headerString := "key1=value1,key!!2=value2,key3=value3"
	_, err := ParseDataSubjectHeader(headerString)
	if err.Error() != "invalid datasubject in header, 2th subject key 'key!!2' is not valid" {
		t.Errorf("result: %s expected: invalid datasubject in header, 2th subject key 'key!!2' is not valid", err)
	}

}

func TestParseDataSubjectHeaderExceptionFailRegExValue(t *testing.T) {
	headerString := "key1=value1,key2=value2,key3=Avalue!3"
	_, err := ParseDataSubjectHeader(headerString)
	if err.Error() != "invalid datasubject in header, 3th subject value 'Avalue!3' is not valid" {
		t.Errorf("result: %s expected: invalid datasubject in header, 3th subject value 'Avalue!3' is not valid", err)
	}

}
