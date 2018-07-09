#include "TestHashTable.hpp"
#include "HashTable.hpp"

TEST_F(TestHashTable, put_some_stuff) {
    HashTable<int, int> aTable;

    aTable.Put(1, 128);
    auto r = aTable.Get(1);
    EXPECT_TRUE(r.ok);
    EXPECT_EQ(r.value, 128);
}

TEST_F(TestHashTable, put_lots_of_stuff) {
    HashTable<int, int> aTable;
    int iterations = 10000;

    for (int i = 0 ; i < iterations ; i++) {
        int index = std::rand()%100000;
        aTable.Put(index, i);
        auto r = aTable.Get(index);
        EXPECT_TRUE(r.ok);
        EXPECT_EQ(r.value, i);
    }

}