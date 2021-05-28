fib:
	psql -d postgres -f fib.sql
	go run .
