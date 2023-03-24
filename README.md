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
|3x3 basic|70.16µs|152.39µs|202.03µs|247.53µs|350.49µs|641.36µs|105|
|3x3 easy|71.33µs|112.79µs|147.84µs|191.35µs|281.66µs|371.43µs|104|
|3x3 intermediate|114.7µs|197.37µs|251.21µs|295.4µs|380.51µs|947.21µs|104|
|3x3 advanced|112.65µs|207.39µs|264.56µs|347.98µs|466.23µs|1.58ms|104|
|3x3 extreme|84.98µs|172.51µs|214.51µs|282.42µs|361.94µs|499.25µs|103|
|3x3 evil|80.41µs|239.3µs|280.46µs|354.05µs|491.93µs|1.04ms|126|
|3x4 advanced|267.73µs|535.78µs|704.94µs|1.2ms|3.92ms|5.61ms|119|
|4x4 advanced|1.49ms|14.52ms|32.72ms|97.41ms|1.13s|2.01s|170|

_Last Updated: 23 Mar 23 20:49 CDT_
</resultsMarker>
