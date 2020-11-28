curl -X GET \
  http://localhost:5001/moviesearch \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 8a5fa573-ff84-44f2-9798-a4744216f734' \
  -H 'cache-control: no-cache' \
  -d '{
	"searchword": "Batman",
	"pagination": "2"
}'