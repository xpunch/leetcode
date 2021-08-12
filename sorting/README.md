# Sorting

## Benchmark
```
go test -benchmem -timeout=1s -bench=.
```
```
goos: windows
goarch: amd64
pkg: github.com/xpunch/leetcode/sorting
cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
```
n = 10
```
BenchmarkMergeSort-4                      140834              1280 ns/op             304 B/op          9 allocs/op
BenchmarkQuickSort-4                      322028               501 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-4                  299857               364 ns/op               0 B/op          0 allocs/op
BenchmarkCocktailShakerSort-4             400226               277 ns/op               0 B/op          0 allocs/op
BenchmarkGnomeSort-4                      453106               275 ns/op               0 B/op          0 allocs/op
BenchmarkBubbleSort-4                     639403               271 ns/op               0 B/op          0 allocs/op
BenchmarkCombSort-4                      1000000               236 ns/op               0 B/op          0 allocs/op
```
n = 100
```
BenchmarkCocktailShakerSort-4             111109              9987 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-4                  144805              9899 ns/op               0 B/op          0 allocs/op
BenchmarkBubbleSort-4                     103894              9891 ns/op               0 B/op          0 allocs/op
BenchmarkOddEvenSort-4                    122542              9717 ns/op               0 B/op          0 allocs/op
BenchmarkHeapSort-4                       135436              9635 ns/op               0 B/op          0 allocs/op
BenchmarkMergeSort-4                      142858              8664 ns/op            5600 B/op         99 allocs/op
BenchmarkGnomeSort-4                      169086              6906 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSort-4                      211614              5188 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSort-4                  233376              5272 ns/op               0 B/op          0 allocs/op
BenchmarkCountingSort-4                   253838              4556 ns/op             893 B/op          1 allocs/op
BenchmarkCombSort-4                       334795              3801 ns/op               0 B/op          0 allocs/op
BenchmarkShellSort-4                      378976              3505 ns/op               0 B/op          0 allocs/op
```
n = 1000
```
BenchmarkBubbleSort-4                       1443            896635 ns/op               0 B/op          0 allocs/op
BenchmarkCocktailShakerSort-4               1153            878525 ns/op               0 B/op          0 allocs/op
BenchmarkOddEvenSort-4                      1411            748136 ns/op               0 B/op          0 allocs/op
BenchmarkSelectionSort-4                    1654            663429 ns/op               0 B/op          0 allocs/op
BenchmarkGnomeSort-4                        1903            560118 ns/op               0 B/op          0 allocs/op
BenchmarkHeapSort-4                         1999            534433 ns/op               0 B/op          0 allocs/op
BenchmarkInsertionSort-4                    2900            350726 ns/op               0 B/op          0 allocs/op
BenchmarkCountingSort-4                     5814            208233 ns/op            8192 B/op          1 allocs/op
BenchmarkMergeSort-4                       10779            107495 ns/op           81344 B/op        999 allocs/op
BenchmarkShellSort-4                       20060             54798 ns/op               0 B/op          0 allocs/op
BenchmarkCombSort-4                        20582             56393 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSort-4                       16076             71482 ns/op               0 B/op          0 allocs/op
```
n = 10000
```
BenchmarkBubbleSort-4                         10         117581780 ns/op            8192 B/op          0 allocs/op
BenchmarkCocktailShakerSort-4                 10         101005500 ns/op            8192 B/op          0 allocs/op
BenchmarkOddEvenSort-4                        15          69540267 ns/op            5461 B/op          0 allocs/op
BenchmarkGnomeSort-4                          18          56270789 ns/op            4551 B/op          0 allocs/op
BenchmarkSelectionSort-4                      20          57770445 ns/op            4096 B/op          0 allocs/op
BenchmarkHeapSort-4                           24          43944612 ns/op            3413 B/op          0 allocs/op
BenchmarkInsertionSort-4                      30          34791890 ns/op            2730 B/op          0 allocs/op
BenchmarkCountingSort-4                       56          18875711 ns/op           83382 B/op          1 allocs/op
BenchmarkMergeSort-4                         764           1610069 ns/op         1112811 B/op       9999 allocs/op
BenchmarkShellSort-4                        1131            956105 ns/op              72 B/op          0 allocs/op
BenchmarkCombSort-4                         1314            829850 ns/op              62 B/op          0 allocs/op
BenchmarkQuickSort-4                        1497            723335 ns/op              54 B/op          0 allocs/op
```