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
|3x3 basic|66.96µs|119.82µs|138.97µs|180.27µs|281.48µs|465.14µs|143|
|3x3 easy|66.76µs|98.88µs|117.51µs|136.32µs|220.02µs|479.62µs|142|
|3x3 intermediate|74.73µs|128.99µs|167.58µs|220.18µs|295.23µs|457.24µs|142|
|3x3 advanced|98.04µs|168.88µs|208.72µs|285.26µs|399.66µs|800.94µs|142|
|3x3 extreme|108.27µs|154.29µs|195.24µs|234.53µs|311.87µs|514.19µs|141|
|3x3 evil|83.29µs|201.45µs|248.33µs|291.22µs|444.39µs|519.76µs|169|
|3x4 advanced|157.94µs|258.27µs|339.62µs|438.61µs|680.5µs|1.21ms|162|
|4x4 advanced|279.49µs|648.53µs|842.93µs|1.16ms|2.28ms|10.45ms|212|

_Last Updated: 25 Mar 23 14:25 CDT_
</resultsMarker>
