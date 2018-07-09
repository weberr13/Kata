#pragma once
#include <functional>
#include <iostream>

template<class V>
struct KeyReturn {
    V value;
    bool ok;
};

template<class K>
struct KeyStore {
    K key;
    bool found;
};

template<class K, class V, class hasher = std::hash<K> >
class HashTable {
public:
    HashTable() : mTableSize(4096) {
        KeyStore<K> empty;
        empty.found = false;
        mStore.resize(mTableSize);
        mKeys.resize(mTableSize, empty);
    }
    virtual ~HashTable() {}
    KeyReturn<V> Get(const K&);
    void Put(const K&, const V&);
private:
    size_t mTableSize;
    std::vector<KeyStore<K> > mKeys;
    std::vector<V> mStore;

    void rebalance();
};

template<class K, class V, class hasher>
KeyReturn<V> HashTable<K,V,hasher>::Get(const K& k) {
    auto index = hasher{}(k)%mTableSize;
    KeyReturn<V> ret;
    ret.ok = false;

    if (!mKeys.at(index).found) {
        return ret;
    }
    ret.value = mStore[index];
    ret.ok = true;
    return ret;
}

template<class K, class V, class hasher>
void HashTable<K,V,hasher>::Put(const K& k, const V& v) {
    auto index = hasher{}(k)%mTableSize;

    if (!mKeys.at(index).found) {
        mStore[index] = v;
        mKeys[index].key = k;
        mKeys[index].found = true;
        return;
    }
    if (mKeys[index].key == k) {
        mStore[index] = v;
        return;
    }
    rebalance();
    Put(k, v);
}

template<class K, class V, class hasher>
void HashTable<K,V,hasher>::rebalance() {
    std::vector<KeyStore<K>> newKeys;
    std::vector<V> newStore;
    size_t newLimit = mTableSize * 2;
    KeyStore<K> empty;
    empty.found = false;
    
    newKeys.resize(newLimit, empty);
    newStore.resize(newLimit);

    for (auto k : mKeys) {
        auto oldIndex = hasher{}(k.key)%mTableSize;
        auto newIndex = hasher{}(k.key)%newLimit;

        newKeys[newIndex] = k;
        newStore[newIndex] = mStore[oldIndex];
    }
    mStore.swap(newStore);
    mKeys.swap(newKeys);
    mTableSize = newLimit;
}