all: libjieba
install:
	mv lib/libjieba.a /usr/local/lib/libjieba.a
libjieba:
	g++ -v -Wall -o jieba.o -c -DLOGGING_LEVEL=LL_WARNING -I deps/ lib/jieba.cpp
	ar rs lib/libjieba.a jieba.o
	rm -f *.o
clean:
	rm -f lib/*.a *.o