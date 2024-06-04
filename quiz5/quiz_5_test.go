package quiz5

import "testing"

func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"Test1", 123, "123"},
		{"Test2", 12345, "12.3K"},
		{"Test3", 1234567890, "1.23B"},
		{"Test4", 1234567890123, "1.23T"},
		{"Test5", 123.45, "123.45"},
		{"Test6", 1234.56, "1.23K"},
		{"Test7", 123456789123456, "123T"},
		{"Test8", 1234567891234567, "1234567891234567"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Format(test.input)
			if result != test.expected {
				t.Errorf("Format(%f) returned %s, expected %s", test.input, result, test.expected)
			}
		})
	}
}

func TestPadingConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"Test1", 123, "123"},
		{"Test2", 12345, "12.3K"},
		{"Test3", 1234567890, "1.23B"},
		{"Test4", 1234567890123, "1.23T"},
		{"Test5", 123.45, "123.45"},
		{"Test6", 1234.56, "1.23K"},
		{"Test7", 123456789123456, "123T"},
		{"Test8", 1234567891234567, "1234567891234567"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PadingConvert(test.input)
			if result != test.expected {
				t.Errorf("PadingConvert(%f) returned %s, expected %s", test.input, result, test.expected)
			}
		})
	}
}

func TestFormatNumberByMath(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"Test1", 123, "123"},
		{"Test2", 12345, "12.3K"},
		{"Test3", 1234567890, "1.23B"},
		{"Test4", 1234567890123, "1.23T"},
		{"Test5", 123.45, "123.45"},
		{"Test6", 1234.56, "1.23K"},
		{"Test7", 123456789123456, "123T"},
		{"Test8", 1234567891234567, "1234567891234567"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FormatNumberByMath(test.input)
			if result != test.expected {
				t.Errorf("FormatNumberByMath(%f) returned %s, expected %s", test.input, result, test.expected)
			}
		})
	}
}
