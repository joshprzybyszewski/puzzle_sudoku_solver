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
|3x3 basic|79.61µs|120.91µs|145.07µs|184.96µs|301.5µs|425.3µs|147|
|3x3 easy|72.62µs|112.86µs|129.58µs|157.59µs|241.66µs|296.25µs|146|
|3x3 intermediate|76.48µs|151.95µs|201.79µs|258.95µs|366.03µs|731.64µs|146|
|3x3 advanced|92.29µs|190.8µs|245.35µs|313.45µs|417.67µs|719.62µs|146|
|3x3 extreme|73.71µs|181.34µs|225.2µs|260.81µs|322.32µs|510.65µs|145|
|3x3 evil|86.46µs|206.89µs|245.9µs|297.45µs|410.54µs|788.06µs|173|
|3x4 advanced|126.99µs|245.03µs|339.77µs|432.4µs|647.09µs|960.42µs|166|
|4x4 advanced|273.23µs|613.79µs|805.7µs|1.08ms|2.5ms|14.87ms|216|

_Last Updated: 02 Apr 23 15:13 CDT_
</resultsMarker>
