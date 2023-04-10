run: 
	go run cmd/server/api.go

test-bench:
	autocannon -m POST -b '{"name": "John Doe","email": "password","password": "password"}' -H "Content-Type: application/json" -a 500 -c 10 http://localhost:3000/users