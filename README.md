# Sudoku Solver
Sudoku Solver - golang

This is the self-proclaimed World's fastest solver for [the sudoku puzzle](https://www.puzzle-sudoku.com) (on this specific website). It is similar to [my Masyu Solver](https://github.com/joshprzybyszewski/masyu).

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
|3x3 basic|86.71µs|573.16µs|1.68ms|3.83ms|4.82ms|7.32ms|97|
|3x3 easy|89.33µs|135.03µs|173.2µs|1.67ms|4.47ms|4.77ms|96|
|3x3 intermediate|519.68µs|1.58ms|2.9ms|4.16ms|5.32ms|8.42ms|96|
|3x3 advanced|102.91µs|1.57ms|3.31ms|4.41ms|5.96ms|9.29ms|96|
|3x3 extreme|92.94µs|716.05µs|1.84ms|4.24ms|6.02ms|8.06ms|95|
|3x3 evil|157.33µs|2ms|3.74ms|4.75ms|6.58ms|10.06ms|98|
|3x4 advanced|2.93ms|5.49ms|6.79ms|10.21ms|13.51ms|28ms|91|
|4x4 advanced|12.87ms|39.06ms|90.94ms|232.02ms|2s|2.01s|142|

_Last Updated: 22 Mar 23 06:36 CDT_
</resultsMarker>
