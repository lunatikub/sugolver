# Sugolver
SuGolVer is a Sudoku Golang Solver.

Detailed documentation available [here](http://thomas-joly.com/index.php/2020/08/31/sugolver/).

## Build Sugolver

```sh
make sugolver
```

## Run the unit tests

```sh
make unit-test
```

## Run the tests

```sh
make test
```

## Sugolver usage

```sh
$ ./sugolver -h
Usage of sugolver:
  -grid string
        grid to solve <1..34.5...6...7[....]>
  -solution
        dump the solution
  -stats
        dump the statistics
```