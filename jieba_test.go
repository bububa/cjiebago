//
// jieba_test.go
//
package cjiebago

import (
	"reflect"
	"testing"
	"os"
	"path"
)

var (
	GOPATH = os.Getenv("GOPATH")
	DICT = path.Join(GOPATH, "src/github.com/bububa/cjiebago/dict/jieba.dict.utf8")
	HMM_MODEL = path.Join(GOPATH, "src/github.com/bububa/cjiebago/dict/hmm_model.utf8")
)

func TestCut(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		sentence string
		hmm      bool
		// Expected results.
		want    []string
		wantErr bool
	}{
		{
			name:     "Cut with hmm test",
			sentence: "小明硕士毕业于中国科学院计算所，后在日本京都大学深造",
			hmm:      true,
			want:     []string{"小明", "硕士", "毕业", "于", "中国科学院", "计算所", "，", "后", "在", "日本京都大学", "深造"},
		},
		{
			name:     "Cut without hmm test",
			sentence: "小明硕士毕业于中国科学院计算所，后在日本京都大学深造",
			hmm:      false,
			want:     []string{"小", "明", "硕士", "毕业", "于", "中国科学院", "计算所", "，", "后", "在", "日本京都大学", "深造"},
		},
	}
	jieba := NewJieba(DICT, HMM_MODEL, "")
	defer jieba.Close()
	for _, tt := range tests {
		got := jieba.Cut(tt.sentence, tt.hmm)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Cut() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCutAll(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		sentence string
		// Expected results.
		want    []string
		wantErr bool
	}{
		{
			name:     "CutAll test",
			sentence: "小明硕士毕业于中国科学院计算所，后在日本京都大学深造",
			want:     []string{"小", "明", "硕士", "毕业", "于", "中国", "中国科学院", "科学", "科学院", "学院", "计算", "计算所", "，", "后", "在", "日本", "日本京都大学", "京都", "京都大学", "大学", "深造"},
		},
	}
	jieba := NewJieba(DICT, HMM_MODEL, "")
	defer jieba.Close()
	for _, tt := range tests {
		got := jieba.CutAll(tt.sentence)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. CutAll() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCutForSearch(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		sentence string
		hmm      bool
		// Expected results.
		want    []string
		wantErr bool
	}{
		{
			name:     "CutForSearch with hmm test",
			sentence: "小明硕士毕业于中国科学院计算所，后在日本京都大学深造",
			hmm:      true,
			want:     []string{"小明", "硕士", "毕业", "于", "中国", "科学", "学院", "科学院", "中国科学院", "计算", "计算所", "，", "后", "在", "日本", "京都", "大学", "日本京都大学", "深造"},
		},
		{
			name:     "CutForSearch without hmm test",
			sentence: "小明硕士毕业于中国科学院计算所，后在日本京都大学深造",
			hmm:      false,
			want:     []string{"小", "明", "硕士", "毕业", "于", "中国", "科学", "学院", "科学院", "中国科学院", "计算", "计算所", "，", "后", "在", "日本", "京都", "大学", "日本京都大学", "深造"},
		},
	}
	jieba := NewJieba(DICT, HMM_MODEL, "")
	defer jieba.Close()
	for _, tt := range tests {
		got := jieba.CutForSearch(tt.sentence, tt.hmm)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. CutForSearch() = %v, want %v", tt.name, got, tt.want)
		}
	}
}