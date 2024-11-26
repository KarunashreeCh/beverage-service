# beverage-service
docker build -t beverage-service .
docker run -d -p 5000:5000 --name beverage-service beverage-service


curl -X POST -H "Content-Type: application/json" \
-d '{"name": "Coke", "quantity": 10, "cost": 1.5}' \
http://localhost:5000/beverages

curl -X POST -H "Content-Type: application/json" -d "{\"name\": \"Coke\", \"quantity\": 10, \"cost\": 1.5}" http://localhost:5000/beverages


curl http://localhost:5000/beverages

curl http://localhost:5000/beverages/1

Create a Beverage
curl -X POST -H "Content-Type: application/json" \
-d '{"name": "Coke", "quantity": 10, "cost": 1.5}' \
http://localhost:5000/beverages




Get All Beverages
curl http://localhost:5000/beverages

Get a Single Beverage
curl http://localhost:5000/beverages/1

Update a Beverage
curl -X PUT -H "Content-Type: application/json" \
-d '{"name": "Pepsi", "quantity": 20}' \
http://localhost:5000/beverages/1

curl -X PUT -H "Content-Type: application/json" -d "{\"name\": \"Pepsi\", \"quantity\": 20}" http://localhost:5000/beverages/1


Delete a Beverage
curl -X DELETE http://localhost:5000/beverages/1

