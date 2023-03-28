package utils

import (
	"baseTechnical/src/utils"
	"reflect"
	"testing"
)

func TestProcessRow(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name:  "should output only IP address and URL from input line",
			args:  args{line: `177.71.128.21 - - [10/Jul/2018:22:21:28 +0200] "GET /intranet-analytics/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (X11; U; Linux x86_64; fr-FR) AppleWebKit/534.7 (KHTML, like Gecko) Epiphany/2.30.6 Safari/534.7`},
			want:  "177.71.128.21",
			want1: "/intranet-analytics/",
		},
		{
			name:  "should output empty strings if input is not long enough to access required fields",
			args:  args{line: `this is an erroneous line`},
			want:  "",
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := utils.ProcessRow(tt.args.line)
			if got != tt.want {
				t.Errorf("ProcessRow() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ProcessRow() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProcessIPActivity(t *testing.T) {
	type args struct {
		ipActivity map[string]int
	}
	tests := []struct {
		name string
		args args
		want []utils.IP
	}{
		{
			name: "should process IP activity as expected",
			args: args{ipActivity: map[string]int{"155.121.21.34": 4, "177.40.21.254": 2, "139.21.44.109": 3, "177.71.128.21": 5}},
			want: []utils.IP{{Ip: "177.71.128.21", Count: 5}, {Ip: "155.121.21.34", Count: 4}, {Ip: "139.21.44.109", Count: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ProcessIPActivity(tt.args.ipActivity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessIPActivity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessURLVisits(t *testing.T) {
	type args struct {
		urlVisits map[string]int
	}
	tests := []struct {
		name string
		args args
		want []utils.URL
	}{
		{
			name: "should process URL visits as expected",
			args: args{urlVisits: map[string]int{"/siteone": 50, "/sitetwo": 34, "/sitethree": 11, "/sitefour": 24}},
			want: []utils.URL{{URL: "/siteone", Count: 50}, {URL: "/sitetwo", Count: 34}, {URL: "/sitefour", Count: 24}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ProcessURLVisits(tt.args.urlVisits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessURLVisits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateURL(t *testing.T) {
	type args struct {
		urlString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "URL Validation 1",
			args: args{urlString: "https://www.example.com/"},
			want: true,
		},
		{
			name: "URL Validation 2",
			args: args{urlString: "www.example.com"},
			want: false,
		},
		{
			name: "URL Validation 3",
			args: args{urlString: "/example/data"},
			want: true,
		},
		{
			name: "URL Validation 4",
			args: args{urlString: "example/data"},
			want: false,
		},
		{
			name: "URL Validation 5",
			args: args{urlString: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ValidateURL(tt.args.urlString); got != tt.want {
				t.Errorf("ValidateURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateIPAddress(t *testing.T) {
	type args struct {
		ipAddress string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "IP Validation 1",
			args: args{ipAddress: "111.111.111.111"},
			want: true,
		},
		{
			name: "IP Validation 2",
			args: args{ipAddress: "obviously not an IP address"},
			want: false,
		},
		{
			name: "IP Validation 3",
			args: args{ipAddress: "111.111.111.0/255"},
			want: false,
		},
		{
			name: "IP Validation 4",
			args: args{ipAddress: "0.0.0.0"},
			want: true,
		},
		{
			name: "IP Validation 5",
			args: args{ipAddress: "999.333.666.111"},
			want: false,
		},
		{
			name: "IP Validation 6",
			args: args{ipAddress: "50.112.00.11"},
			want: false,
		},
		{
			name: "IP Validation 7",
			args: args{ipAddress: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ValidateIPAddress(tt.args.ipAddress); got != tt.want {
				t.Errorf("ValidateIPv4Address() = %v, want %v", got, tt.want)
			}
		})
	}
}
