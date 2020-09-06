# Sugolver
SuGolVer is a Sudoku Golang Solver.

A detailed artcile on the algorithm used is available [here](http://thomas-joly.com/index.php/2020/08/31/sugolver/).

## Build Sugolver

```sh
make sugolver
```

## Run the unit tests

The unit tests are in the file `solver_test.go`

```sh
make unit-test
```

## Run the tests

All the tests are located in the folder `test/grids`.

A test is composed by 2 lines:
* the initial values (with the char '.' as unknown value)
* the expected solution

``` sh
$ cat test/grids/easy1.grid
....75.6.78...9......43...9.64...79..53............483......3..9....6..2....54...
349175268781629534526438179164283795853947621297561483418792356975316842632854917
```

To run all the tests, run the following command:

```sh
make test
```

## Sugolver usage

```sh
$ ./sugolver.exe -h
Usage of D:\git\sugolver\sugolver.exe:
  -dump string
        dump the grid [solution, pretty]
  -grid string
        grid to solve <1..34.5...6...7[....]>
  -stats
        dump the statistics
  -uniqueness
        enable uniqueness
  -exclusivity
        enable exclusivity
   -parity
        enable parity
```

The only mandatory option is `--grid`
```sh
 ❯❯❯ ./sugolver   \
      --grid "....75.6.78...9......43...9.64...79..53............483......3..9....6..2....54..." \
      --dump=pretty
-------------------------
| 3 4 9 | 1 7 5 | 2 6 8 |
| 7 8 1 | 6 2 9 | 5 3 4 |
| 5 2 6 | 4 3 8 | 1 7 9 |
-------------------------
| 1 6 4 | 2 8 3 | 7 9 5 |
| 8 5 3 | 9 4 7 | 6 2 1 |
| 2 9 7 | 5 6 1 | 4 8 3 |
-------------------------
| 4 1 8 | 7 9 2 | 3 5 6 |
| 9 7 5 | 3 1 6 | 8 4 2 |
| 6 3 2 | 8 5 4 | 9 1 7 |
-------------------------
```
