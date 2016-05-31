package main

import (
	"github.com/bububa/cjiebago"
	"os"
	"path"
)

var (
	GOPATH    = os.Getenv("GOPATH")
	DICT      = path.Join(GOPATH, "src/github.com/bububa/cjiebago/dict/jieba.dict.utf8")
	HMM_MODEL = path.Join(GOPATH, "src/github.com/bububa/cjiebago/dict/hmm_model.utf8")
	IDF       = path.Join(GOPATH, "src/github.com/bububa/cjiebago/dict/idf.utf8")
	STOPWORDS = path.Join(GOPATH, "src/github.com/bububa/cjiebago/dict/stop_words.utf8")
)

func main() {
	jieba := cjiebago.NewJieba(DICT, HMM_MODEL, "")
	defer jieba.Close()
	sentence := "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	times := 10000
	for times > 0 {
		jieba.CutForSearch(sentence, true)
		times -= 1
	}
}
