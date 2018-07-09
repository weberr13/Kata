#include "gtest/gtest.h"
#include <iostream>

int main(int argc, char *argv[]) {
    testing::InitGoogleTest(&argc, argv);
    srand(time(NULL));

    int return_val = RUN_ALL_TESTS();
    std::cout << "FINISHED TESTS " << std::endl;
    return return_val;
}