package sequential

import (
	"baseTechnical/src/pkg/utils"
	"reflect"
	"testing"
)

func TestSequential(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []utils.IP
		want2 []utils.URL
	}{
		{
			name:  "Test Sequential Processing",
			args:  args{"../../testData/test.log"},
			want:  11,
			want1: []utils.IP{{Ip: "177.71.128.21", Count: 10}, {Ip: "72.44.32.10", Count: 7}, {Ip: "168.41.191.40", Count: 4}},
			want2: []utils.URL{{URL: "/intranet-analytics/", Count: 10}, {URL: "/translations/", Count: 5}, {URL: "/docs/manage-websites/", Count: 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Sequential(tt.args.filename)
			if got != tt.want {
				t.Errorf("Sequential() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Sequential() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Sequential() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func BenchmarkSequential(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sequential("../../testData/test.log")
	}
}
