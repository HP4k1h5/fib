# fib

## setup/install

### prerequisites

- running postgres instance
  - if your local pg instance needs a password for the root user you will
    need the `$PGPASSWORD` environment variable to be set
  - if your pg instance already has a database named `fib` you will have to
    delete it before running `make`, i.e. run `DROP DATABASE fib;` in
    `psql`
- go (tested with go1.16.4)
- make (built with GNU Make 3.81)
- nothing running on port 7357 locally

### run

run make in a terminal

```
make
```

this will create a database named `fib` with a single table called `memo`,
and, for simplicity's sake, it will immediately start the api running on port
7357

you can test that it is up by visiting `http://localhost:7357/fib/33` in your
browser or running

```
curl http://localhost:7357/fib/33
```

you should get a 200 response with value 3524578

## endpoints

- ` GET /fib/{n}` pass an integer value for n 

returns the corresponding fibonacci value

ex: 

```
curl localhost:7357/fib/23
```

- `GET /memoized/{val}` pass an integer value for val

returns the number of memoized values below val

ex: 
```
curl localhost:7357/memoized/12000
```

- `DELETE /memoized` clears the memoized cache

returns 200 status code

ex: 
```
curl -X DELETE localhost:7357/memoized
```

## tests

run `go test` from the `fib` directory to run some basic tests

## benchmarks

run `go test -bench *test.go` from the `fib` directory to get an idea of
performance metrics for the `Fib` function
