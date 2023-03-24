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
|3x3 basic|93.09µs|149.88µs|180.49µs|218.18µs|295.49µs|438.61µs|111|
|3x3 easy|73.6µs|114.97µs|137.67µs|177.87µs|249.84µs|309.54µs|110|
|3x3 intermediate|86.17µs|178.37µs|220.93µs|281.29µs|355.36µs|602.34µs|110|
|3x3 advanced|107.98µs|188.45µs|233.16µs|308.9µs|405.19µs|562.19µs|110|
|3x3 extreme|105.64µs|175.34µs|210.14µs|258.74µs|316.18µs|489.78µs|109|
|3x3 evil|83.63µs|204.25µs|259.82µs|313.29µs|407.84µs|743.03µs|132|
|3x4 advanced|192.61µs|389.93µs|481.26µs|606.39µs|948.86µs|1.52ms|125|
|4x4 advanced|438.29µs|1.15ms|1.73ms|3.77ms|11.4ms|138.65ms|176|

_Last Updated: 23 Mar 23 21:35 CDT_
</resultsMarker>
