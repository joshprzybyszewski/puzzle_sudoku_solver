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
|3x3 basic|80.46µs|140.27µs|185.97µs|246.07µs|331.94µs|488.51µs|142|
|3x3 easy|72.32µs|133.59µs|154.32µs|187.25µs|257.52µs|460.33µs|141|
|3x3 intermediate|77.77µs|184.89µs|228.36µs|277.84µs|347.77µs|612.84µs|141|
|3x3 advanced|106.24µs|197.63µs|241µs|312.04µs|446.37µs|900.11µs|141|
|3x3 extreme|112.05µs|184.25µs|222.25µs|269.2µs|372.83µs|508.92µs|140|
|3x3 evil|114.12µs|229.37µs|268.28µs|339.08µs|479.42µs|572.66µs|168|
|3x4 advanced|145.27µs|372.9µs|482.72µs|584.43µs|810.63µs|1.53ms|161|
|4x4 advanced|370.39µs|1.09ms|1.79ms|3.4ms|12.07ms|135ms|211|

_Last Updated: 23 Mar 23 21:47 CDT_
</resultsMarker>
