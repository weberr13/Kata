#!/usr/bin/env ruby

def merge_sort(items, start, ending, less = lambda {|l, r| l < r})
    if !items.respond_to?("[]")
        return items
    end
    if !items.respond_to?("size")
        return items
    end
    if start+1 == ending
        return items
    end
    if start == ending
        return items
    end
    middle = start + ( (ending - start) / 2)
    merge_sort(items, start, middle, less)
    merge_sort(items, middle, ending, less)
    merge(items, start, middle, ending, less)
    return items
end

def merge(items, start, middle, ending, less)
    tmp = []
    starti = start
    middlei = middle
    while starti < middle && middlei < ending 
        if less.call(items[starti], items[middlei])
            tmp.push(items[starti])
            starti = starti + 1
        else
            tmp.push(items[middlei])
            middlei = middlei + 1
        end
    end
    while starti < middle 
        tmp.push(items[starti])
        starti = starti + 1
    end
    while middlei < ending
        tmp.push(items[middlei])
        middlei = middlei + 1
    end
    items[start..ending-1] = tmp
end

if __FILE__ == $0
    puts merge_sort("foo", 0, 0) 
    a = ["bar"]
    puts merge_sort(a, 0, a.size)
    a.push("ace")
    puts merge_sort(a, 0, a.size)
    a = ["s", "a", "t", "b", "z"]
    puts merge_sort(a, 0, a.size)
    puts merge_sort(a, 0, a.size, lambda {|l, r| l > r})
end