package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Row
	}{
		{
			name: "single row",
			input: `----------------------------------------
rowkey
  family:column @ 2021/10/01-00:00:00.000000
    "value"
`,
			want: []Row{
				{
					Key: "rowkey",
					Columns: map[string][]Cell{
						"family:column": {
							{
								Value:     `"value"`,
								Timestamp: "2021/10/01-00:00:00.000000",
							},
						},
					},
				},
			},
		},
	}

	parser := New()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parser.Parse(test.input)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
