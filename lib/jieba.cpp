extern "C" {
    #include "jieba.h"
}

#include "cppjieba/Jieba.hpp"
#include "cppjieba/KeywordExtractor.hpp"

using namespace std;

extern "C" {

char** getWords(const char* sentence, vector<string> words, size_t len) {
  char** result = (char**)malloc(sizeof(char*) * (words.size() + 1));

  for (size_t i = 0; i < words.size(); i++) {
    char* word = (char*)malloc(words[i].size() + 1);
    memset(word, 0, words[i].size() + 1);
    strcpy(word, words[i].c_str());
    *(result + i) = word;
  }
  *(result + words.size()) =  NULL;
  return result;
}

Jieba NewJieba(const char* dict_path, const char* hmm_path, const char* user_dict) {
  Jieba handle = (Jieba)(new cppjieba::Jieba(dict_path, hmm_path, user_dict));
  return handle;
}

void FreeJieba(Jieba handle) {
  cppjieba::Jieba* x = (cppjieba::Jieba*)handle;
  delete x;
}

char** Cut(Jieba handle, const char* sentence, size_t len, int hmm) {
  cppjieba::Jieba* x = (cppjieba::Jieba*)handle;
  vector<string> words;
  string s(sentence, len);
  if (hmm == 1) {
    x->Cut(s, words, true);
  }else{
    x->Cut(s, words, false);
  }
  char** ret = getWords(sentence, words, len);
  words.clear();
  vector<string>().swap(words);
  return ret;
}

char** CutAll(Jieba handle, const char* sentence, size_t len) {
  cppjieba::Jieba* x = (cppjieba::Jieba*)handle;
  vector<string> words;
  string s(sentence, len);
  x->CutAll(s, words);
  char** ret = getWords(sentence, words, len);
  words.clear();
  vector<string>().swap(words);
  return ret;
}

char** CutForSearch(Jieba handle, const char* sentence, size_t len, int hmm) {
  cppjieba::Jieba* x = (cppjieba::Jieba*)handle;
  vector<string> words;
  string s(sentence, len);
  if (hmm == 1) {
    x->CutForSearch(s, words, true);
  }else{
    x->CutForSearch(s, words, false);
  }
  char** ret = getWords(sentence, words, len);
  words.clear();
  vector<string>().swap(words);
  return ret;
}

void FreeWords(char** words) {
  char** p = words;
  while(p != NULL && *p != NULL) {
    free(*p);
    p++;
  }

  if (words != NULL) {
    free(words);
  }
}

CJiebaTag** getTags(const char* sentence, vector<pair<string, string> > tagres, size_t len) {
  CJiebaTag** result = (CJiebaTag**)malloc(sizeof(CJiebaTag*) * (tagres.size() + 1));

  for (size_t i = 0; i < tagres.size(); i++) {
    CJiebaTag* tags_i = (CJiebaTag*)malloc(sizeof(CJiebaTag));

    tags_i->word = (char*)malloc(tagres[i].first.size() + 1);
    memset(tags_i->word, 0, tagres[i].first.size() + 1);
    strcpy(tags_i->word, tagres[i].first.c_str());

    tags_i->tag = (char*)malloc(tagres[i].second.size() + 1);
    memset(tags_i->tag, 0, tagres[i].second.size() + 1);
    strcpy(tags_i->tag, tagres[i].second.c_str());

    *(result + i) = tags_i;
  }
  *(result + tagres.size()) =  NULL;

  return result;
}

int InsertUserWord(Jieba handle, const char* word)
{
    cppjieba::Jieba* x = (cppjieba::Jieba*)handle;
    return x->InsertUserWord(string(word), "u") ? 1 : 0;
}

CJiebaTag** Tag(Jieba handle, const char* sentence, size_t len) {
  cppjieba::Jieba* x = (cppjieba::Jieba*)handle;
  vector<pair<string, string> > tagres;
  string s(sentence, len);
  x->Tag(s, tagres);

  return getTags(sentence, tagres, len);
}

void FreeTags(CJiebaTag** tags) {
  CJiebaTag** p = tags;
  while(p != NULL && *p != NULL) {
    if ((*p)->word != NULL) {
      free((*p)->word);
    }
    if ((*p)->tag != NULL) {
      free((*p)->tag);
    }

    free(*p);
    p++;
  }

  if (tags != NULL) {
    free(tags);
  }
}

char* GetTagWord(CJiebaTag* tags)
{
    if (tags == NULL) {
      return NULL;
    }

    return tags->word;
}

char* GetTagAsStr(CJiebaTag* tags)
{
    if (tags == NULL) {
      return NULL;
    }

    return tags->tag;
}

Extractor NewExtractor(const char* dict_path,
      const char* hmm_path,
      const char* idf_path,
      const char* stop_word_path,
      const char* user_dict_path) {
  Extractor handle = (Extractor)(new cppjieba::KeywordExtractor(dict_path, 
          hmm_path, 
          idf_path,
          stop_word_path,
          user_dict_path));
  return handle;
}

void FreeExtractor(Extractor handle) {
  cppjieba::KeywordExtractor* x = (cppjieba::KeywordExtractor*)handle;
  delete x;
}

char** Extract(Extractor handle, const char* sentence, size_t len, size_t topn) {
  cppjieba::KeywordExtractor* x = (cppjieba::KeywordExtractor*)handle;
  vector<cppjieba::KeywordExtractor::Word> words;
  x->Extract(sentence, words, topn);
  char** result = (char**)malloc(sizeof(char*) * (words.size() + 1));
  for (size_t i = 0; i < words.size(); i++) {
    char* word = (char*)malloc(words[i].word.size() + 1);
    strcpy(word, words[i].word.c_str());

    *(result + i) = word;
  }
  *(result + words.size()) =  NULL;
  words.clear();
  vector<cppjieba::KeywordExtractor::Word>().swap(words);
  return result;
}

} // extern "C"