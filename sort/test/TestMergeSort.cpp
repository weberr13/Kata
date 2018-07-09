#include "TestMergeSort.hpp"
#include "MergeSort.hpp"
#include <vector>
#include <algorithm>
#include <functional>
#include <iostream> 

TEST_F(TestMergeSort, vectorOfThings) {
    std::vector<int> someInts;
    auto targetSize = 1111111;

    for (int i = 0; i < targetSize ; i++) {
        someInts.push_back(std::rand()%1000);
    }

    MergeSort<int>(someInts, 0, someInts.size());

    EXPECT_TRUE(std::is_sorted(someInts.begin(), someInts.end()));
}