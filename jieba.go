package cjiebago

// #cgo LDFLAGS: -L /usr/local/lib/ -L . -ljieba -lstdc++
// #cgo CFLAGS: -I ./lib/
// #include <stdlib.h>
// #include "jieba.h"
import "C"
import "unsafe"

type Jieba struct {
	pointer C.Jieba
}

type CJiebaWord C.struct_CJiebaWord

func NewJieba(dict string, hmm string, userDict string) *Jieba {
	dictC := C.CString(dict)
	hmmC := C.CString(hmm)
	userDictC := C.CString(userDict)
	defer func() {
		C.free(unsafe.Pointer(dictC))
		C.free(unsafe.Pointer(hmmC))
		C.free(unsafe.Pointer(userDictC))
	}()

	ret := C.NewJieba(dictC, hmmC, userDictC)
	return &Jieba{ret}
}

func (this *Jieba) Close() {
	C.FreeJieba(this.pointer)
}

func (this *Jieba) Cut(sentence string, hmm bool) []string {
	if len(sentence) == 0 {
		return []string{}
	}

	sizeC := C.size_t(len([]byte(sentence)))
	sentenceC := C.CString(sentence)
	defer C.free(unsafe.Pointer(sentenceC))

	var jiebaWordsC **CJiebaWord
	defer C.FreeWords((**C.struct___0)(unsafe.Pointer(jiebaWordsC)))

	if hmm {
		jiebaWordsC = (**CJiebaWord)(unsafe.Pointer(C.Cut(this.pointer, sentenceC, sizeC, 1)))
	} else {
		jiebaWordsC = (**CJiebaWord)(unsafe.Pointer(C.Cut(this.pointer, sentenceC, sizeC, 0)))
	}

	return this.convert_jieba_words(jiebaWordsC)
}

func (this *Jieba) CutAll(sentence string) []string {
	if len(sentence) == 0 {
		return []string{}
	}

	sizeC := C.size_t(len([]byte(sentence)))
	sentenceC := C.CString(sentence)
	defer C.free(unsafe.Pointer(sentenceC))

	jiebaWordsC := (**CJiebaWord)(unsafe.Pointer(C.CutAll(this.pointer, sentenceC, sizeC)))
	defer C.FreeWords((**C.struct___0)(unsafe.Pointer(jiebaWordsC)))

	return this.convert_jieba_words(jiebaWordsC)
}

func (this *Jieba) CutForSearch(sentence string, hmm bool) []string {
	if len(sentence) == 0 {
		return []string{}
	}

	sizeC := C.size_t(len([]byte(sentence)))
	sentenceC := C.CString(sentence)
	defer C.free(unsafe.Pointer(sentenceC))

	var jiebaWordsC **CJiebaWord
	defer C.FreeWords((**C.struct___0)(unsafe.Pointer(jiebaWordsC)))

	if hmm {
		jiebaWordsC = (**CJiebaWord)(unsafe.Pointer(C.CutForSearch(this.pointer, sentenceC, sizeC, 1)))
	} else {
		jiebaWordsC = (**CJiebaWord)(unsafe.Pointer(C.CutForSearch(this.pointer, sentenceC, sizeC, 0)))
	}

	return this.convert_jieba_words(jiebaWordsC)
}

func (this *Jieba) convert_jieba_words(jiebaWordsC **CJiebaWord) (words_list []string) {
	p := jiebaWordsC
	for {
		if *p == nil {
			break
		}
		words := C.GetWordAsStr((*C.struct___0)(unsafe.Pointer(*p)))
		if words == nil {
			break
		}
		words_list = append(words_list, C.GoString(words))

		p = (**CJiebaWord)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(unsafe.Sizeof(*jiebaWordsC))))
	}

	return
}
