# Sudoku Solver
Sudoku Solver - golang

This is the self-proclaimed World's fastest solver for [the sudoku puzzle](www.puzzle-sudoku.com) (on this specific website). It is similar to [my Masyu Solver](https://github.com/joshprzybyszewski/masyu).

To run, execute `make compete`.

## Results

Check the [Hall of Fame](https://www.puzzle-sudoku.com/hall.php?hallsize=7) for the results recorded by the puzzle server (which include network duration of submission). Below are the results of the solver as recorded on my machine.

_NOTE: Update this table with `make results`._

<resultsMarker>

_GOOS: linux_

_GOARCH: amd64_

_cpu: Intel(R) Core(TM) i5-3570 CPU @ 3.40GHz_

_Solve timeout: 10s_

|Puzzle|Min|p25|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-|-:|
|3x3 basic|65.99µs|589.56µs|1.54ms|3.9ms|4.74ms|5.39ms|28|
|3x3 easy|61.7µs|83.71µs|98.95µs|1.65ms|4.34ms|5.19ms|27|
|3x3 intermediate|394.24µs|1.53ms|3.39ms|4.1ms|5.03ms|8.79ms|27|
|3x3 advanced|276µs|1.71ms|3.46ms|4.81ms|7.83ms|9.49ms|27|
|3x3 extreme|70.95µs|582.94µs|1.24ms|3.62ms|5.58ms|11.42ms|26|
|3x3 evil|816.26µs|1.74ms|3.88ms|5ms|6.91ms|7.66ms|25|
|3x4 advanced|3.24ms|5.58ms|8.21ms|12.06ms|38.94ms|38.94ms|20|
|4x4 advanced|11.99ms|75.51ms|252.53ms|757.43ms|10s|10.01s|129|

_Last Updated: 19 Mar 23 15:56 CDT_
</resultsMarker>
