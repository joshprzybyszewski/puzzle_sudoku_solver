# Sudoku Solver
Sudoku Solver - golang

This is the self-proclaimed World's fastest solver for [the sudoku puzzle](www.puzzle-sudoku.com) (on this specific website). It is similar to [my Masyu Solver](https://github.com/joshprzybyszewski/masyu).

To run, execute `make compete`.

## Results

Check the [Hall of Fame](https://www.puzzle-sudoku.com/hall.php?hallsize=9) for the results recorded by the puzzle server (which include network duration of submission). Below are the results of the solver as recorded on my machine.

_NOTE: Update this table with `make results`._

<resultsMarker>

_GOOS: linux_

_GOARCH: amd64_

_cpu: Intel(R) Core(TM) i5-3570 CPU @ 3.40GHz_

_Solve timeout: 10s_

|Puzzle|Min|p25|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-|-:|
|3x3 basic|641.61µs|1.39ms|2.8ms|8.39ms|31.55ms|103.91ms|21|
|3x3 easy|142.82µs|963.36µs|2.37ms|7.08ms|91.02ms|165.52ms|21|
|3x3 intermediate|258.91µs|3.84ms|14.61ms|28.27ms|48.27ms|155.18ms|21|
|3x3 advanced|418.73µs|3.41ms|12.52ms|51.67ms|121.17ms|357.64ms|21|
|3x3 extreme|458.22µs|2.5ms|4.26ms|25.12ms|131.65ms|131.65ms|20|
|3x3 evil|828.8µs|9.26ms|16.96ms|32.64ms|103.64ms|103.64ms|20|
|3x4 advanced|178.25ms|1.14s|4.04s|6.91s|14.13s|14.13s|15|
|4x4 advanced|8.61ms|132.44ms|485.56ms|1.75s|10s|10s|65|

_Last Updated: 11 Mar 23 16:14 CST_
</resultsMarker>
