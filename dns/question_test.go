package dns

import "testing"

// "\fcodecrafters\x02io\x00\x00\x01\x00\x01"

func TestQuestionBytes(t *testing.T) {
	var tq Question = loadJson(t, "question", "question").(Question)
	tcs := []struct {
		name string
		question Question
		expected int
	}{
		{
			question: tq,
			expected: 21,
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			b := tc.question.Bytes()
			if len(b) != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, len(b))
			}
		})
	}

}