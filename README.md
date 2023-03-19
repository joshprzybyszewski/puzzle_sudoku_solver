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
|3x3 basic|90.77µs|681.26µs|1.96ms|4.18ms|5.09ms|5.64ms|94|
|3x3 easy|82.64µs|145.77µs|184.37µs|1.71ms|4.71ms|5.21ms|93|
|3x3 intermediate|500.11µs|1.9ms|3.62ms|4.62ms|5.61ms|9.11ms|93|
|3x3 advanced|178.73µs|1.86ms|3.61ms|4.72ms|5.53ms|9.48ms|93|
|3x3 extreme|128.5µs|734.12µs|1.81ms|4.56ms|5.48ms|9.5ms|92|
|3x3 evil|635.71µs|2.29ms|3.91ms|4.94ms|6.7ms|9.08ms|95|
|3x4 advanced|1.96ms|5.57ms|6.71ms|9.76ms|12.48ms|19.64ms|88|
|4x4 advanced|12.87ms|39.03ms|80.31ms|242.45ms|2s|2.01s|139|

_Last Updated: 19 Mar 23 17:45 CDT_
</resultsMarker>
