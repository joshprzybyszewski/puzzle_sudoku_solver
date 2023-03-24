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
|3x3 basic|93.64µs|162.22µs|195.51µs|268.97µs|370.16µs|474.42µs|103|
|3x3 easy|86.75µs|113.69µs|149.13µs|215.02µs|315.83µs|403.96µs|102|
|3x3 intermediate|115.73µs|214.17µs|266.29µs|341.07µs|491.51µs|578.59µs|102|
|3x3 advanced|94.3µs|208.15µs|278.42µs|345.48µs|468.05µs|1.78ms|102|
|3x3 extreme|111.35µs|182.59µs|231.71µs|286.87µs|373.91µs|623.09µs|101|
|3x3 evil|71.97µs|219.15µs|271.96µs|338.21µs|487.77µs|854.77µs|119|
|3x4 advanced|258.61µs|462.36µs|693.48µs|1.25ms|3.35ms|6.76ms|112|
|4x4 advanced|2.49ms|17.56ms|36.18ms|100.33ms|1.58s|2.03s|163|

_Last Updated: 23 Mar 23 20:17 CDT_
</resultsMarker>
