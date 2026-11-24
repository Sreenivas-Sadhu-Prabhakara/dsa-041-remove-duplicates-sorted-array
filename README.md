# Remove Duplicates from Sorted Array

> Striver SDE Sheet 041 · Day 7 LL & Arrays — reader/writer pointers compact unique values to the front in place. Go, stdlib only, tested.

## Problem

Given a sorted array, remove the duplicates in place so each value appears once, and return the count of unique values. The prefix of that length must hold the deduplicated array — no second array allowed.

## Approach

The two-pointer compaction idiom: a read pointer scans every element, a write pointer marks where the next *new* value belongs. Because the input is sorted, "is this a new value?" is a single comparison against the last kept element. Every value moves at most once. This same reader/writer shape handles a whole family of in-place filters (remove element, keep at most two of each, move zeroes) — worth internalizing once.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
