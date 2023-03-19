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

_Solve timeout: 2s_

|Puzzle|Min|p25|Median|p75|p95|max|sample size|
|-|-|-|-|-|-|-|-:|
|3x3 basic|102.69µs|631.26µs|1.31ms|3.36ms|4.42ms|4.7ms|33|
|3x3 easy|84.71µs|129.16µs|175.82µs|1.59ms|3.51ms|3.89ms|32|
|3x3 intermediate|589.9µs|1.47ms|3.44ms|4.35ms|5.68ms|8.89ms|32|
|3x3 advanced|490.8µs|1.93ms|3.5ms|4.39ms|5.3ms|6.27ms|32|
|3x3 extreme|137.95µs|683.23µs|1.62ms|4.4ms|6.34ms|8.05ms|31|
|3x3 evil|453.1µs|1.87ms|3.31ms|4.87ms|5.98ms|6.11ms|35|
|3x4 advanced|4.22ms|5.36ms|6.48ms|9.28ms|12.02ms|19.74ms|28|
|4x4 advanced|11.61ms|40.13ms|96.85ms|338.15ms|2s|2.02s|139|

_Last Updated: 19 Mar 23 17:12 CDT_
</resultsMarker>
