package logging

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	type expected struct {
		wantErr   bool
		level     logrus.Level
		formatter logrus.Formatter
	}
	tests := [...]struct {
		name    string
		jsonRaw []byte
		expect  expected
	}{
		{
			"debug level json format",
			[]byte(`{"level": "debug", "formatter":{"format": "json", "data": {"dataKey":"hodor"}}}`),
			expected{false, logrus.DebugLevel, &logrus.JSONFormatter{}},
		},
		{
			"warn level text format",
			[]byte(`{"level": "warn", "formatter":{"format": "text", "data": null}}`),
			expected{false, logrus.WarnLevel, &logrus.TextFormatter{}},
		},
		{
			"warn level text format forceColor",
			[]byte(`{"level": "warn", "formatter":{"format": "text", "data": {"forceColors": true}}}`),
			expected{false, logrus.WarnLevel, &logrus.TextFormatter{ForceColors: true}},
		},
		{
			"warn level default format",
			[]byte(`{"level": "error"}`),
			expected{false, logrus.ErrorLevel, &logrus.TextFormatter{}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := Logger{}
			err := json.Unmarshal(test.jsonRaw, &config)
			if !test.expect.wantErr && err != nil {
				t.Fatalf("no error expected: %s", err)
			}
			if log.Level != test.expect.level {
				t.Errorf("expected level \"%s\" got \"%s\"", test.expect.level, config.Level)
			}
			formatterType := reflect.TypeOf(log.Formatter).Elem()
			expectedType := reflect.TypeOf(test.expect.formatter).Elem()
			if formatterType.String() != expectedType.String() {
				t.Errorf("expected formatter \"%s\" got \"%s\"", expectedType, formatterType)
			}
		})
	}
}