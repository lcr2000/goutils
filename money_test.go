package goutils

import "testing"

func TestFenToYuanFloat(t *testing.T) {
	type args struct {
		money int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "case1",
			args: args{money: 100},
			want: 1,
		},
		{
			name: "case2",
			args: args{1},
			want: 0.01,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FenToYuanFloat(tt.args.money); got != tt.want {
				t.Errorf("FenToYuan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFenToYuanString(t *testing.T) {
	type args struct {
		money int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{money: 100},
			want: "1",
		},
		{
			name: "case2",
			args: args{money: 1},
			want: "0.01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FenToYuanString(tt.args.money); got != tt.want {
				t.Errorf("FenToYuanString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYuanToFen(t *testing.T) {
	type args struct {
		price float64
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{
		{
			name:    "test1",
			args:    args{price: 11.84},
			wantRes: 1184,
		},
		{
			name:    "test2",
			args:    args{price: 11.888},
			wantRes: 1188,
		},
		{
			name:    "test3",
			args:    args{price: 11.88},
			wantRes: 1188,
		},
		{
			name:    "test4",
			args:    args{price: 11.00},
			wantRes: 1100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := YuanToFen(tt.args.price); gotRes != tt.wantRes {
				t.Errorf("YuanToFen() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestYuanToFenRounding(t *testing.T) {
	type args struct {
		price float64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{price: 11.00},
			want: 1100,
		},
		{
			name: "test2",
			args: args{price: 11.45},
			want: 1145,
		},
		{
			name: "test3",
			args: args{price: 11.454},
			want: 1145,
		},
		{
			name: "test3",
			args: args{price: 11.455},
			want: 1146,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YuanToFenRounding(tt.args.price); got != tt.want {
				t.Errorf("YuanToFen2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscountCalculation(t *testing.T) {
	type args struct {
		realMoney float64
		tagPrice  float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				realMoney: 100,
				tagPrice:  200,
			},
			want: "5.0",
		},
		{
			name: "case2",
			args: args{
				realMoney: 100,
				tagPrice:  100,
			},
			want: "10.0",
		},
		{
			name: "case3",
			args: args{
				realMoney: 120,
				tagPrice:  100,
			},
			want: "12.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DiscountCalculation(tt.args.realMoney, tt.args.tagPrice)
			t.Log(got)
		})
	}
}

func TestDiscountCalculationInt(t *testing.T) {
	type args struct {
		realMoney int
		tagPrice  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				realMoney: 100,
				tagPrice:  200,
			},
			want: "5.0",
		},
		{
			name: "case2",
			args: args{
				realMoney: 100,
				tagPrice:  100,
			},
			want: "10.0",
		},
		{
			name: "case3",
			args: args{
				realMoney: 120,
				tagPrice:  100,
			},
			want: "12.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiscountCalculationInt(tt.args.realMoney, tt.args.tagPrice); got != tt.want {
				t.Errorf("DiscountCalculationInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
