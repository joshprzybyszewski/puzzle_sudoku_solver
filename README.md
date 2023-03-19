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
|3x3 basic|746.1µs|1.99ms|4.92ms|11.45ms|34.73ms|105.78ms|24|
|3x3 easy|145.01µs|941.68µs|1.59ms|5.11ms|90.93ms|166.62ms|24|
|3x3 intermediate|296.47µs|2.5ms|14.46ms|27.43ms|52.52ms|155.77ms|24|
|3x3 advanced|359.47µs|2.93ms|10.41ms|50.35ms|121.76ms|361.41ms|24|
|3x3 extreme|593.31µs|2.01ms|4.41ms|15.69ms|63.02ms|132.48ms|23|
|3x3 evil|853.59µs|8.31ms|14.62ms|29.38ms|48.51ms|101.69ms|22|
|3x4 advanced|178.23ms|1.18s|3.03s|6.84s|14.28s|14.28s|17|
|4x4 advanced|8.45ms|50.1ms|184.06ms|576.76ms|7.41s|10.01s|118|

_Last Updated: 18 Mar 23 21:17 CDT_
</resultsMarker>
