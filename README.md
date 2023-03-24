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
|3x3 basic|87.94µs|162.87µs|237.33µs|277.17µs|381.67µs|979.47µs|110|
|3x3 easy|87.09µs|125.06µs|158.2µs|218.25µs|331.52µs|531.13µs|109|
|3x3 intermediate|143.13µs|218.85µs|286.19µs|335.87µs|487.96µs|1.39ms|109|
|3x3 advanced|146.46µs|231.08µs|289.43µs|365.67µs|511.77µs|660.19µs|109|
|3x3 extreme|113.7µs|199.58µs|260.42µs|313.49µs|463.47µs|830.02µs|108|
|3x3 evil|127.02µs|251.05µs|308.38µs|391.88µs|545.8µs|968.41µs|131|
|3x4 advanced|256.97µs|544.72µs|668.52µs|949.65µs|2.58ms|4.14ms|124|
|4x4 advanced|789.27µs|10.91ms|20.15ms|56.21ms|419.94ms|2.01s|175|

_Last Updated: 23 Mar 23 21:29 CDT_
</resultsMarker>
