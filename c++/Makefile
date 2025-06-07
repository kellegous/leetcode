OBJS := 0001-two-sum.o \
	0002-add-two-numbers.o \
	0003-longest-substring-without-repeating-characters.o \
	0004-median-of-two-sorted-arrays.o \
	0005-longest-palindromic-substring.o \
	0006-zigzag-conversion.o \
	0007-reverse-integer.o \
	0008-string-to-integer-atoi.o \
	0009-palindrome-number.o \
	0010-regular-expression-matching.o \
	0011-container-with-most-water.o \
	0012-integer-to-roman.o \
	0013-roman-to-integer.o \
	0014-longest-common-prefix.o \
	0015-3sum.o

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