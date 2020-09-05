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
$ ./sugolver  --grid \
    "....75.6.78...9......43...9.64...79..53............483......3..9....6..2....54..."
```