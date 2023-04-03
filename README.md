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
|3x3 basic|86.21µs|132.67µs|155.19µs|211.55µs|352.95µs|406µs|148|
|3x3 easy|78.5µs|121.56µs|138.46µs|163.56µs|261.18µs|457.49µs|147|
|3x3 intermediate|83.78µs|152.78µs|203µs|270.4µs|357.07µs|783.59µs|147|
|3x3 advanced|114.45µs|213.01µs|248.7µs|297.83µs|427.55µs|505.91µs|147|
|3x3 extreme|104.18µs|197.3µs|232.56µs|275.48µs|367.05µs|654.5µs|146|
|3x3 evil|93.03µs|216.04µs|258.18µs|318.64µs|429.8µs|758.37µs|174|
|3x4 advanced|152.39µs|252.93µs|344.28µs|453.66µs|709.31µs|943.92µs|167|
|4x4 advanced|299.9µs|660.05µs|856.73µs|1.14ms|2.65ms|12.7ms|217|
|jigsaw 5x5 easy|76.36µs|95.15µs|112.01µs|144.7µs|182.48µs|182.48µs|14|
|jigsaw 5x5 intermediate|116.89µs|148.71µs|181.93µs|184.75µs|407.06µs|407.06µs|14|
|jigsaw 5x5 advanced|118.06µs|138.06µs|155.98µs|194.49µs|218.33µs|218.33µs|14|
|jigsaw 7x7 easy|73.1µs|101.12µs|112.53µs|128.19µs|179.23µs|179.23µs|14|
|jigsaw 7x7 intermediate|113.46µs|188.77µs|246.46µs|266.43µs|335.04µs|335.04µs|14|
|jigsaw 7x7 advanced|155.28µs|190.37µs|218.97µs|255.56µs|302.85µs|302.85µs|13|
|jigsaw 9x9 easy|104.01µs|127.21µs|154.17µs|254.86µs|369.35µs|369.35µs|12|
|jigsaw 9x9 intermediate|286.48µs|458.73µs|551.62µs|743.68µs|1.07ms|1.07ms|12|
|jigsaw 9x9 advanced|246.46µs|326.69µs|370.53µs|429.15µs|768.24µs|768.24µs|12|

_Last Updated: 02 Apr 23 22:01 CDT_
</resultsMarker>
