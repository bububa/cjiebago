package cjiebago

// #cgo LDFLAGS: -L /usr/local/lib/ -ljieba -lstdc++
// #cgo CFLAGS: -I ./lib/
// #include <stdlib.h>
// #include "jieba.h"
import "C"
import "unsafe"

type Jieba struct {
	pointer C.Jieba
}

type Extractor struct {
	pointer C.Extractor
}

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

	var jiebaWordsC **C.char

	if hmm {
		jiebaWordsC = (**C.char)(unsafe.Pointer(C.Cut(this.pointer, sentenceC, sizeC, 1)))
	} else {
		jiebaWordsC = (**C.char)(unsafe.Pointer(C.Cut(this.pointer, sentenceC, sizeC, 0)))
	}
	defer C.FreeWords((**C.char)(unsafe.Pointer(jiebaWordsC)))
	return getWords(jiebaWordsC)
}

func (this *Jieba) CutAll(sentence string) []string {
	if len(sentence) == 0 {
		return []string{}
	}

	sizeC := C.size_t(len([]byte(sentence)))
	sentenceC := C.CString(sentence)
	defer C.free(unsafe.Pointer(sentenceC))

	jiebaWordsC := (**C.char)(unsafe.Pointer(C.CutAll(this.pointer, sentenceC, sizeC)))
	//defer freeWords(jiebaWordsC)
	defer C.FreeWords((**C.char)(unsafe.Pointer(jiebaWordsC)))
	return getWords(jiebaWordsC)
}

func (this *Jieba) CutForSearch(sentence string, hmm bool) []string {
	if len(sentence) == 0 {
		return []string{}
	}

	sizeC := C.size_t(len([]byte(sentence)))
	sentenceC := C.CString(sentence)
	defer C.free(unsafe.Pointer(sentenceC))

	var jiebaWordsC **C.char

	if hmm {
		jiebaWordsC = (**C.char)(unsafe.Pointer(C.CutForSearch(this.pointer, sentenceC, sizeC, 1)))
	} else {
		jiebaWordsC = (**C.char)(unsafe.Pointer(C.CutForSearch(this.pointer, sentenceC, sizeC, 0)))
	}
	//defer freeWords(jiebaWordsC)
	defer C.FreeWords((**C.char)(unsafe.Pointer(jiebaWordsC)))
	return getWords(jiebaWordsC)
}

func NewExtractor(dict string, hmm string, idf string, stopWords string, userDict string) *Extractor {
	dictC := C.CString(dict)
	hmmC := C.CString(hmm)
	idfC := C.CString(idf)
	stopWordsC := C.CString(stopWords)
	userDictC := C.CString(userDict)
	defer func() {
		C.free(unsafe.Pointer(dictC))
		C.free(unsafe.Pointer(hmmC))
		C.free(unsafe.Pointer(idfC))
		C.free(unsafe.Pointer(stopWordsC))
		C.free(unsafe.Pointer(userDictC))
	}()

	ret := C.NewExtractor(dictC, hmmC, idfC, stopWordsC, userDictC)
	return &Extractor{ret}
}

func (this *Extractor) Close() {
	C.FreeExtractor(this.pointer)
}

func (this *Extractor) Extract(sentence string, topn uint) []string {
	if len(sentence) == 0 {
		return []string{}
	}

	sizeC := C.size_t(len([]byte(sentence)))
	topnC := C.size_t(topn)
	sentenceC := C.CString(sentence)
	defer C.free(unsafe.Pointer(sentenceC))
	jiebaWordsC := (**C.char)(unsafe.Pointer(C.Extract(this.pointer, sentenceC, sizeC, topnC)))
	defer C.FreeWords((**C.char)(unsafe.Pointer(jiebaWordsC)))

	return getWords(jiebaWordsC)
}

func getWords(jiebaWordsC **C.char) (words_list []string) {
	p := jiebaWordsC
	ptrSize := unsafe.Sizeof(*p)
	for {
		if *p == nil {
			break
		}
		words_list = append(words_list, C.GoString((*C.char)(*p)))

		p = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(ptrSize)))
	}

	return
}

func freeWords(jiebaWordsC **C.char) {
	if jiebaWordsC == nil {
		return
	}
	p := jiebaWordsC
	ptrSize := unsafe.Sizeof(*p)
	for {
		if *p == nil {
			break
		}
		C.free(unsafe.Pointer(*p))
		p = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(ptrSize)))
	}
	C.free(unsafe.Pointer(jiebaWordsC))
}
