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
|3x3 basic|102.68µs|195.92µs|251.79µs|304.04µs|433.83µs|1.31ms|108|
|3x3 easy|88.45µs|145.45µs|172.8µs|223.36µs|336.34µs|510.57µs|107|
|3x3 intermediate|140.21µs|238.17µs|304.15µs|364.6µs|471.41µs|911.14µs|107|
|3x3 advanced|118.58µs|247.84µs|304.56µs|395.36µs|585.44µs|2.33ms|107|
|3x3 extreme|100.89µs|222.94µs|276.87µs|354.04µs|481.62µs|667.76µs|106|
|3x3 evil|164.74µs|280.22µs|328µs|395.76µs|530.58µs|686.52µs|129|
|3x4 advanced|369.57µs|622.16µs|886.67µs|1.37ms|3.04ms|4.13ms|122|
|4x4 advanced|719.17µs|14.96ms|31.99ms|83.84ms|815.24ms|2.01s|173|

_Last Updated: 23 Mar 23 21:15 CDT_
</resultsMarker>
