# Set

Sets generation from numbers sequence. The user can provide some numbers and the package should print (or just generate) sets from these numbers. Size of the sets is defined by size of the sequence. 0 < m < N, where m is number from set, N is number from the sequence with the same index. The sets can be ordered randomly. For example,

```
Input: 1, 2
Output: [0 0] [0 1]

Input: 1, 2, 3
Output:  [0 0 2] [0 0 0] [0 0 1] [0 1 1] [0 1 0] [0 1 2]
```

This is simple realisation of the task. We generate sets from the input numbers recursively. The user of the package has 2 options: to use Print(w, sequence) for results printing or Generate(sequence) to own processing of the results.

To run tests (with logs):

```go
go test -race -bench=. -v
```

This package has tests:

* Example* - to show how the package can be used
* Test* - to test validation of input data, to test generated data
* Benchmark* - to show simple benchmarks to future improves
