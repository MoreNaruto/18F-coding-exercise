# 18F-coding-exercise

## How to create an SLCSP file

Run these commands

```shell script
cd project
go install .
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
project
```

You should now see a file named `new_slcsp.csv` in 
your `resources` path with updated rates

## How to run tests

```shell script
cd project
go test ./...
```