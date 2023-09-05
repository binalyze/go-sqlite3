package stats_test

import (
	"testing"

	. "github.com/mattn/go-sqlite3/binalyze/tests"
)

func TestStats(t *testing.T) {
	testCases := []struct {
		name     string
		query    string
		expected interface{}
	}{
		{
			name:     "stats stddev",
			query:    `SELECT stddev(value) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(1),
		},
		{
			name:     "stats stddev_samp",
			query:    `SELECT stddev_samp(value) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(1),
		},
		{
			name:     "stats stddev_pop",
			query:    `SELECT round(stddev_pop(value), 5) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(0.8165),
		},
		{
			name:     "stats variance",
			query:    `SELECT variance(value) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(1),
		},
		{
			name:     "stats var_samp",
			query:    `SELECT var_samp(value) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(1),
		},
		{
			name:     "stats var_pop",
			query:    `SELECT round(var_pop(value), 5) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(0.66667),
		},
		{
			name:     "stats median",
			query:    `SELECT median(value) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(2),
		},
		{
			name:     "stats median",
			query:    `SELECT median(value) FROM (SELECT 5 value UNION ALL SELECT 3 UNION ALL SELECT 3)`,
			expected: float64(3),
		},
		{
			name:     "stats percentile",
			query:    `SELECT percentile(value,25) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(1.5),
		},
		{
			name:     "stats percentile_25",
			query:    `SELECT percentile_25(value) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(1.5),
		},
		{
			name:     "stats percentile_75",
			query:    `SELECT percentile_75(value) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(2.5),
		},
		{
			name:     "stats percentile_90",
			query:    `SELECT percentile_90(value) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(2.8),
		},
		{
			name:     "stats percentile_95",
			query:    `SELECT round(percentile_95(value), 5) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(2.9),
		},
		{
			name:     "stats percentile_99",
			query:    `SELECT round(percentile_99(value), 5) FROM (SELECT 1 value UNION ALL SELECT 2 UNION ALL SELECT 3)`,
			expected: float64(2.98),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			TestOne(t, tc.query, tc.expected)
		})
	}
}
