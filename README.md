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
|3x3 basic|86.87µs|151.46µs|172.82µs|221.7µs|352.79µs|543.57µs|278|
|3x3 easy|65.55µs|132.16µs|156.31µs|176.29µs|260.2µs|340.59µs|277|
|3x3 intermediate|93.57µs|184.73µs|243.45µs|292.42µs|391.31µs|443.41µs|277|
|3x3 advanced|107.75µs|242.48µs|281.6µs|332.03µs|434.21µs|528.37µs|277|
|3x3 extreme|109.14µs|214.23µs|252.38µs|299.68µs|381.73µs|508.53µs|276|
|3x3 evil|93.42µs|244.29µs|297.07µs|351.08µs|463.54µs|597.33µs|304|
|3x4 advanced|208.77µs|312.22µs|411.21µs|497.67µs|757.49µs|971.57µs|297|
|4x4 advanced|399.4µs|743.02µs|938.23µs|1.31ms|2.72ms|31.34ms|347|
|jigsaw 5x5 easy|61.34µs|87.49µs|97.02µs|109.97µs|210.83µs|228.48µs|144|
|jigsaw 5x5 intermediate|74.66µs|155.33µs|182.02µs|203.06µs|228.47µs|259.19µs|144|
|jigsaw 5x5 advanced|89.76µs|173.54µs|185.95µs|201.11µs|234.67µs|273.11µs|144|
|jigsaw 7x7 easy|46.84µs|96.84µs|118.69µs|187.74µs|223.98µs|244.12µs|144|
|jigsaw 7x7 intermediate|88.54µs|199.53µs|220.78µs|254.91µs|336.27µs|460.35µs|144|
|jigsaw 7x7 advanced|156.94µs|202.9µs|229.18µs|260.74µs|338.62µs|428.07µs|143|
|jigsaw 9x9 easy|87.48µs|129.67µs|150.4µs|232.37µs|320.3µs|390.99µs|142|
|jigsaw 9x9 intermediate|135.81µs|300.96µs|366.54µs|448.5µs|623.56µs|861.08µs|142|
|jigsaw 9x9 advanced|147.95µs|321.11µs|399.46µs|470.11µs|679.75µs|833.15µs|142|

_Last Updated: 06 May 23 20:37 CDT_
</resultsMarker>
