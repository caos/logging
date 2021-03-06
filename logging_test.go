package logging

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntryFields(t *testing.T) {
	tests := []struct {
		name           string
		entry          *Entry
		expectedFields logrus.Fields
	}{
		{
			"without error",
			Log("UTILS-B7l7"),
			logrus.Fields{"logID": "UTILS-B7l7"},
		},
		{
			"with error",
			Log("UTILS-Ld9V").WithError(fmt.Errorf("im an error")),
			logrus.Fields{
				"logID": "UTILS-Ld9V",
				"error": fmt.Errorf("im an error"),
			},
		},
		{
			"on error",
			Log("UTILS-Ld9V").OnError(fmt.Errorf("im an error")),
			logrus.Fields{
				"logID": "UTILS-Ld9V",
				"error": fmt.Errorf("im an error"),
			},
		},
		{
			"on error without",
			Log("UTILS-Ld9V").OnError(nil),
			logrus.Fields{
				"logID": "UTILS-Ld9V",
			},
		},
		{
			"with fields",
			LogWithFields("LOGGI-5kk6z", "field1", 134, "field2", "asdlkfj"),
			logrus.Fields{
				"logID":  "LOGGI-5kk6z",
				"field1": 134,
				"field2": "asdlkfj",
			},
		},
		{
			"with field",
			LogWithFields("LOGGI-5kk6z").WithField("field1", 134),
			logrus.Fields{
				"logID":  "LOGGI-5kk6z",
				"field1": 134,
			},
		},
		{
			"fields odd",
			LogWithFields("LOGGI-xWzy4", "kevin"),
			logrus.Fields{
				"logID":     "LOGGI-xWzy4",
				"oddFields": 1,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.entry.Debug()
			for key, expectedValue := range test.entry.Data {
				value, ok := test.expectedFields[key]
				if !ok {
					t.Errorf("\"%s\" was not expected", key)
				}
				if !reflect.DeepEqual(expectedValue, value) {
					t.Errorf("wrong value for \"%s\": expected %T.%v, got %T.%v", key, expectedValue, expectedValue, value, value)
				}
			}
		})
	}
}
