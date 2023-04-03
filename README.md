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
|3x3 basic|77.18µs|132.21µs|160.85µs|201.6µs|297.11µs|534.04µs|278|
|3x3 easy|77.9µs|128.31µs|145.99µs|172.21µs|269.87µs|818.07µs|277|
|3x3 intermediate|65.04µs|145.49µs|207.12µs|263.94µs|343.03µs|480.32µs|277|
|3x3 advanced|74.79µs|196.98µs|237.72µs|284.93µs|382.73µs|546.83µs|277|
|3x3 extreme|78.06µs|183.17µs|217.91µs|257.5µs|354.78µs|463.91µs|276|
|3x3 evil|84.31µs|202.14µs|249.11µs|310.82µs|451.79µs|655.89µs|304|
|3x4 advanced|116.87µs|255.45µs|344.53µs|443.54µs|717.54µs|3.13ms|297|
|4x4 advanced|313.08µs|637.64µs|837.44µs|1.16ms|2.76ms|31.81ms|347|
|jigsaw 5x5 easy|45.67µs|85.57µs|96.12µs|121.08µs|175.48µs|200.6µs|144|
|jigsaw 5x5 intermediate|46.5µs|129.23µs|155.46µs|179.25µs|206.7µs|561.1µs|144|
|jigsaw 5x5 advanced|95.25µs|149.59µs|168.95µs|189.49µs|217.37µs|266.63µs|144|
|jigsaw 7x7 easy|54.95µs|100.57µs|119.37µs|184.76µs|236.12µs|275.54µs|144|
|jigsaw 7x7 intermediate|85.98µs|165.42µs|199.04µs|240.5µs|332.17µs|429.45µs|144|
|jigsaw 7x7 advanced|114.24µs|183.01µs|219.38µs|257.63µs|334.61µs|486.73µs|143|
|jigsaw 9x9 easy|82.68µs|141.88µs|175.15µs|231.17µs|324.93µs|420.13µs|142|
|jigsaw 9x9 intermediate|115.63µs|244.93µs|319.59µs|401.91µs|562.59µs|713.12µs|142|
|jigsaw 9x9 advanced|174.41µs|300.63µs|386.9µs|462.66µs|647.92µs|1.66ms|142|

_Last Updated: 03 Apr 23 06:07 CDT_
</resultsMarker>
