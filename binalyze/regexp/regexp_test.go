package regexp_test

import (
	"testing"

	. "github.com/mattn/go-sqlite3/binalyze/tests"
)

func TestRegexp(t *testing.T) {
	testCases := []struct {
		name     string
		query    string
		expected interface{}
	}{
		{
			name:     "regexp stmt",
			query:    `SELECT 1 WHERE '1234' regexp '1'`,
			expected: int64(1),
		},
		{
			name:     "regexp stmt",
			query:    `SELECT false WHERE 'abc' regexp 'a.c'`,
			expected: int64(0),
		},
		{
			name:     "regexp stmt no result",
			query:    `SELECT 1 WHERE '1234' regexp '5'`,
			expected: NoResult,
		},
		{
			name:     "regexp_like",
			query:    `SELECT regexp_like('abc', 'a.c')`,
			expected: int64(1),
		},
		{
			name:     "regexp_like no match",
			query:    `SELECT regexp_like('abc', 'a.d')`,
			expected: int64(0),
		},
		{
			name:     "regexp_substr",
			query:    `SELECT regexp_substr('abcdef', 'b.d') = 'bcd'`,
			expected: int64(1),
		},
		{
			name:     "regexp_substr",
			query:    `SELECT regexp_substr('abcdef', 'b.d')`,
			expected: "bcd",
		},
		{
			name:     "regexp_capture",
			query:    `SELECT regexp_capture('abcdef', 'b(.)d')`,
			expected: "bcd",
		},
		{
			name:     "regexp_capture",
			query:    `SELECT regexp_capture('abcdef', 'b(.)d', 1)`,
			expected: "c",
		},
		{
			name:     "regexp_replace",
			query:    `SELECT regexp_replace('1234', '[24]', 'x')`,
			expected: "1x3x",
		},
		{
			name:     "regexp_replace",
			query:    `SELECT regexp_replace('abcdef', 'b.d', '...')`,
			expected: "a...ef",
		},
		{
			name:     "regexp_replace no match",
			query:    `SELECT regexp_replace('1234', '[5]', 'x')`,
			expected: "1234",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			TestOne(t, tc.query, tc.expected)
		})
	}
}
