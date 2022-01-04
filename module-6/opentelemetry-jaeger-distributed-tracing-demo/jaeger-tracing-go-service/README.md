# jaeger-tracing-go-service
This is a Go project to demo Jaeger Distributed Tracing

# Get Employee by Id
```bash
curl --location --request GET 'localhost:4747/employee/ddc7562a-ed6b-4148-836e-712dd13dfb1f'
```

# Create Employee
```bash
curl --location --request POST 'localhost:4747/employee' \
--header 'Content-Type: application/json' \
--data-raw '{
	"firstName": "Chris",
	"lastName": "Ketchums",
	"occupation": "Software Engineer",
	"salaryGrade": "A10"
}'
```