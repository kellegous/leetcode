OBJS := 0001-two-sum.o \
	0002-add-two-numbers.o \
	0003-longest-substring-without-repeating-characters.o \
	0004-median-of-two-sorted-arrays.o \
	0005-longest-palindromic-substring.o

all: leetcode_tests

.PHONY: all clean setup reset test

.clangd: make-clangd.py
	python3 make-clangd.py

setup: .clangd

%.o: %.cc
	g++ -std=c++23 $(shell pkg-config --cflags gtest_main) -c -o $@ $<

leetcode_tests: $(OBJS)
	g++ -std=c++23 $(shell pkg-config --cflags --libs gtest_main) -o $@ $^

test: leetcode_tests
	./leetcode_tests

reset: clean
	rm -f .clangd

clean:
	rm -f $(OBJS) leetcode_tests