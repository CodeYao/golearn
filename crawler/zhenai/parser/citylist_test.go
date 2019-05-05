package parser

import (
	"reflect"
	"testing"

	"github.com/golearn/crawler/fetcher"
)

func TestParserCityList(t *testing.T) {
	contests, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	type args struct {
		contents []byte
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				contents: contests,
			},
			want: []interface{}{"阿坝", "阿克苏", "阿拉善盟"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParserCityList(tt.args.contents); !reflect.DeepEqual(got.Items[0:3], tt.want) {
				t.Errorf("ParserCityList() = %v, want %v", got, tt.want)
			}
		})
	}
}
