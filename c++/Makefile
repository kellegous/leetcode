OBJS := 0001-two-sum.o 0002-add-two-numbers.o

all: leetcode_tests

%.o: %.cc
	g++ -std=c++23 $(shell pkg-config --cflags gtest_main) -c -o $@ $<

leetcode_tests: $(OBJS)
	g++ -std=c++23 $(shell pkg-config --cflags --libs gtest_main) -o $@ $^

clean:
	rm -f $(OBJS) leetcode_tests