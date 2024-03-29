package datetime

import (
	"testing"
	"time"

	"github.com/mantidtech/tplr/functions/helper"
	"github.com/stretchr/testify/assert"
)

func init() {
	// set to return a constant for testing
	helper.Now = func() time.Time {
		return time.Date(2020, 8, 29, 2, 14, 0, 133_700_000, time.UTC)
	}
}

// TestListFunctions provides unit test coverage for ListFunctions
func TestFunctions(t *testing.T) {
	fn := Functions()
	assert.Len(t, fn, 5, "weakly ensuring functions haven't been added/removed without updating tests")
}

// TestNow provides unit test coverage for Now.
func TestNow(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "now - default format",
			Template: `{{ now }}`,
			Args:     helper.TestArgs{},
			Want:     "2020-08-29T02:14:00Z",
		},
		{
			Name:     "now - supplied format",
			Template: `{{ now .F }}`,
			Args: helper.TestArgs{
				"F": "02 Jan 06 15:04 UTC",
			},
			Want: "29 Aug 20 02:14 UTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestTimeFormat provides unit test coverage for TimeFormat.
func TestTimeFormat(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "RFC3339",
			Template: `{{ timeFormat .F .T }}`,
			Args: helper.TestArgs{
				"F": "2006-01-02T15:04:05Z07:00",
				"T": time.Date(2023, 05, 18, 21, 00, 51, 0, time.UTC),
			},
			Want: "2023-05-18T21:00:51Z",
		},
		{
			Name:     "RFC822",
			Template: `{{ timeFormat .F .T }}`,
			Args: helper.TestArgs{
				"F": "02 Jan 06 15:04 UTC",
				"T": time.Date(2023, 05, 18, 21, 00, 51, 0, time.UTC),
			},
			Want: "18 May 23 21:00 UTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestTimeParse provides unit test coverage for TimeParse.
func TestTimeParse(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "RFC3339",
			Template: `{{ timeParse .F .T }}`,
			Args: helper.TestArgs{
				"F": "2006-01-02T15:04:05",
				"T": "2023-05-18T21:00:51",
			},
			Want: "2023-05-18 21:00:51 +0000 UTC",
		},
		{
			Name:     "RFC822",
			Template: `{{ timeParse .F .T }}`,
			Args: helper.TestArgs{
				"F": "2006-01-02T15:04:05Z",
				"T": "2023-05-18T21:00:51Z",
			},
			Want: "2023-05-18 21:00:51 +0000 UTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}

// TestTimeToUnix provides unit test coverage for TimeToUnix.
func TestTimeToUnix(t *testing.T) {
	type Args struct {
		ts time.Time
	}

	tests := []struct {
		name      string
		args      Args
		wantInt64 int64
	}{
		// table test data goes here
	}

	for _, tx := range tests {
		tc := tx
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			gotInt64 := TimeToUnix(tc.args.ts)
			assert.Equal(t, tc.wantInt64, gotInt64)
		})
	}
}

func TestUnixToTime(t *testing.T) {
	tests := []helper.TestSet{
		{
			Name:     "int",
			Template: `{{ unixToTime .I | timeToUnix }}`,
			Args: helper.TestArgs{
				"I": 1684407651,
			},
			Want: "1684407651",
		},
		{
			Name:     "string",
			Template: `{{ unixToTime .S }}`,
			Args: helper.TestArgs{
				"S": "1684407651",
			},
			WantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, helper.TemplateTest(tt, Functions()))
	}
}
