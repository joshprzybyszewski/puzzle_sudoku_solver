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
|3x3 basic|103.31µs|161.54µs|215.71µs|270.02µs|396.77µs|705.6µs|109|
|3x3 easy|77.35µs|132.62µs|159.1µs|214.2µs|319.18µs|430.32µs|108|
|3x3 intermediate|118.94µs|229.18µs|281.84µs|364.3µs|457.26µs|977.67µs|108|
|3x3 advanced|122.66µs|206.82µs|280.98µs|359.29µs|495.43µs|701.83µs|108|
|3x3 extreme|127.27µs|193.53µs|247.7µs|304.11µs|414.93µs|1.63ms|107|
|3x3 evil|126.85µs|243.08µs|302.1µs|394.28µs|538.98µs|1.02ms|130|
|3x4 advanced|324.06µs|593.85µs|865.11µs|1.41ms|3.3ms|7.1ms|123|
|4x4 advanced|972.39µs|17.36ms|35.1ms|81.09ms|845.03ms|2.01s|174|

_Last Updated: 23 Mar 23 21:21 CDT_
</resultsMarker>
