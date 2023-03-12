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
|3x3 basic|725.88µs|1.71ms|2.48ms|8.03ms|30.56ms|101.4ms|21|
|3x3 easy|210.23µs|848.34µs|1.53ms|4.59ms|87.85ms|168.9ms|21|
|3x3 intermediate|292.05µs|3.72ms|15.15ms|26.89ms|49.21ms|155.12ms|21|
|3x3 advanced|390.16µs|3.12ms|12.39ms|51.47ms|120.05ms|358.07ms|21|
|3x3 extreme|505.26µs|2.51ms|4.52ms|21.92ms|133.55ms|133.55ms|20|
|3x3 evil|785.45µs|9.5ms|15.11ms|32.36ms|101.04ms|101.04ms|20|
|3x4 advanced|178.65ms|1.15s|4.04s|6.83s|14.16s|14.16s|15|
|4x4 advanced|9.91ms|128ms|512.55ms|1.77s|10.01s|10.07s|65|

_Last Updated: 11 Mar 23 16:21 CST_
</resultsMarker>
