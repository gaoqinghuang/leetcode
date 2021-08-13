package main

import "testing"

func TestMerge1(t *testing.T) {
    a := &Range{
        left:  1,
        right: 2,
    }
    b := &Range{
        left:  3,
        right: 4,
    }
    c, d := merge(a, b)

    if !rangeEqual(c, a) || !rangeEqual(d, b) {
        t.Errorf("交集合并失败")
    }
}

func TestMerge2(t *testing.T) {
    a := &Range{
        left:  1,
        right: 3,
    }
    b := &Range{
        left:  2,
        right: 4,
    }
    c, d := merge(a, b)

    if !rangeEqual(c, nil) || !rangeEqual(d, &Range{
        left: 1,
        right: 4,
    }) {
        t.Errorf("在左边合并失败")
    }
}
func rangeEqual(a, b *Range) bool {
    if a == nil || b == nil {
        //有一个是nil，两个都要是nil
        return b == nil && a == nil
    }

    if a.left == b.left && a.right == b.right {
        return true
    }
    return false
}
