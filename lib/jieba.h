#ifdef __cplusplus
extern "C" {
#endif

#include <stdlib.h>

typedef void* Jieba;
Jieba NewJieba(const char* dict_path, const char* hmm_path, const char* user_dict);
void FreeJieba(Jieba);
int InsertUserWord(Jieba handle, const char* word);

char** Cut(Jieba handle, const char* sentence, size_t len, int hmm);
char** CutAll(Jieba handle, const char* sentence, size_t len);
char** CutForSearch(Jieba handle, const char* sentence, size_t len, int hmm);
void FreeWords(char** words);

typedef struct {
  char* word;
  char* tag;
} CJiebaTag;
CJiebaTag** Tag(Jieba handle, const char* sentence, size_t len);
void FreeTags(CJiebaTag** tags);
char* GetTagWord(CJiebaTag* tags);
char* GetTagAsStr(CJiebaTag* tags);

typedef void* Extractor;
Extractor NewExtractor(const char* dict_path,
      const char* hmm_path,
      const char* idf_path,
      const char* stop_word_path,
      const char* user_dict_path);
char** Extract(Extractor handle, const char* sentence, size_t len, size_t topn);
void FreeExtractor(Extractor handle);

#ifdef __cplusplus
}
#endif