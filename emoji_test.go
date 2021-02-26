package goutils

import "testing"

func TestEmojiFilter(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "表情测试1",
			args: args{content: "🍎左右两边是苹果🍎"},
			want: "金果果",
		},
		{
			name: "表情测试2",
			args: args{content: "右边是叶子🍀"},
			want: "一叶轻舟",
		},
		{
			name: "表情测试3",
			args: args{content: "🇨🇳 左边是标签"},
			want: "背后有",
		},
		{
			name: "空格测试",
			args: args{content: " 左边是空格"},
			want: "背后有",
		},
		{
			name: "空格测试2",
			args: args{content: " 背  后🍀 有"},
			want: "背后有",
		},
		{
			name: "繁体测试",
			args: args{content: "繁華年間花開盡"},
			want: "繁華年間花開盡",
		},
		{
			name: "汉字测试",
			args: args{content: "李富贵"},
			want: "李富贵",
		},
		{
			name: "英文测试",
			args: args{content: "TomAndLisa"},
			want: "TomAndLisa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EmojiFilter(tt.args.content); got != tt.want {
				t.Errorf("EmojiFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
