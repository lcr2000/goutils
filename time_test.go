package goutils

import (
	"testing"
	"time"
)

const defaultFormat = "2006-01-02 15:04:05"

func TestTimeFormatter(t *testing.T) {
	type args struct {
		i      interface{}
		format string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				i:      time.Now().Unix(),
				format: defaultFormat,
			},
		},
		{
			name: "case2",
			args: args{
				i:      time.Now(),
				format: defaultFormat,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeFormatter(tt.args.i, tt.args.format)
			t.Log(got)
		})
	}
}

func TestGetTodayStartTime(t *testing.T) {
	t.Log(GetTodayStartTime(defaultLayout))
}

var (
	today     = time.Now().Unix()
	tomorrow  = time.Now().AddDate(0, 0, 1).Unix()
	yesterday = time.Now().AddDate(0, 0, -1).Unix()
)

func TestIsTomorrow(t *testing.T) {
	type args struct {
		t int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				t: today,
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				t: tomorrow,
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				t: yesterday,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTomorrow(tt.args.t); got != tt.want {
				t.Errorf("IsTomorrow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsToday(t *testing.T) {
	type args struct {
		t int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				t: today,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				t: tomorrow,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				t: yesterday,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsToday(tt.args.t); got != tt.want {
				t.Errorf("IsToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsYesterday(t *testing.T) {
	type args struct {
		t int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				t: today,
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				t: tomorrow,
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				t: yesterday,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsYesterday(tt.args.t); got != tt.want {
				t.Errorf("IsYesterday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeAgo(t *testing.T) {
	m, _ := time.ParseDuration("-1m")
	h, _ := time.ParseDuration("-1h")

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{t: time.Now().AddDate(0, 0, -2)},
		},
		{
			name: "case2",
			args: args{t: time.Now().Add(m)},
		},
		{
			name: "case3",
			args: args{t: time.Now().Add(h)},
		},
		{
			name: "case4",
			args: args{t: time.Now()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeAgo(tt.args.t)
			t.Log(got)
		})
	}
}

func TestIsLeap(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{year: 2021},
			want: false,
		},
		{
			name: "case1",
			args: args{year: 2020},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeap(tt.args.year); got != tt.want {
				t.Errorf("IsLeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
