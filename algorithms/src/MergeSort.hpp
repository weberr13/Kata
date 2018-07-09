#pragma once
#include <functional>
#include <iterator>
#include <iostream>

template<class T>
void printRange(std::vector<T>& v, size_t begin, size_t end) {
    std::cout << "Range: [";
    for (size_t i = begin ; i < end ; i++) {
        std::cout << v[i] << ",";
    }
    std::cout << "]" << std::endl;
}

template<class T,  class C = std::less<T> >
void merge(std::vector<T>& v, size_t begin, size_t middle, size_t end) {
    auto ilh = begin;
    auto irh = middle;
    auto c = C();
    std::vector<T> tmp;

    while (ilh < middle && irh < end) {
        if (c(v[irh], v[ilh])) {
            tmp.push_back(v[irh++]);
        } else {
            tmp.push_back(v[ilh++]);
        }
    }
    for (; ilh < middle ;) {
        tmp.push_back(v[ilh++]);
    }
    for (; irh < end ;) {
        tmp.push_back(v[irh++]);
    }
    for (size_t i = 0; i < tmp.size(); i++) {
        v[i+begin] = tmp[i];
    }

}



template<class T, class C = std::less<T> >
void MergeSort(std::vector<T>& v, size_t begin, size_t end) {
    if (begin + 1 == end) {
        return; // fully partitioned
    }

    auto halfway = (end-begin)/2;  
    auto middle = begin+halfway;

    MergeSort<T,C>(v, begin, middle);
    MergeSort<T,C>(v, middle, end);
    merge<T,C>(v, begin, middle, end);
}



