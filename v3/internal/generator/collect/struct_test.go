package collect

import "testing"

func TestFieldNameFromTag(t *testing.T) {
	tests := []struct {
		name     string
		tag      string
		wantName string
		wantSkip bool
	}{
		{name: "empty tag", tag: ""},
		{name: "bare name", tag: "snake_case", wantName: "snake_case"},
		{name: "bare name with omitempty", tag: "snake_case,omitempty", wantName: "snake_case"},
		{name: "bare dash skips", tag: "-", wantSkip: true},
		{name: "bare options only", tag: ",omitempty"},
		{
			name:     "protobuf with json= wins over name=",
			tag:      "bytes,1,opt,name=commit_short_hash,json=commitShortHash,proto3",
			wantName: "commitShortHash",
		},
		{
			name:     "protobuf with only name=",
			tag:      "bytes,1,opt,name=version,proto3",
			wantName: "version",
		},
		{
			name:     "protobuf enum tag",
			tag:      "varint,1,opt,name=status,proto3,enum=foo.Bar",
			wantName: "status",
		},
		{
			name:     "yaml-style with inline option",
			tag:      "field_name,inline",
			wantName: "field_name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotSkip := fieldNameFromTag(tt.tag)
			if gotName != tt.wantName {
				t.Errorf("name = %q, want %q", gotName, tt.wantName)
			}
			if gotSkip != tt.wantSkip {
				t.Errorf("skip = %v, want %v", gotSkip, tt.wantSkip)
			}
		})
	}
}
