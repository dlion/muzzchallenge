#!/bin/sh

# Start DynamoDB Local in the background
java -jar DynamoDBLocal.jar -sharedDb -dbPath ./data &

echo "Waiting for DynamoDB Local to start..."
until curl -s http://localhost:8000; do
  sleep 1
done

echo "DynamoDB Local started."

aws dynamodb create-table --region local --endpoint-url "http://localhost:8000" --cli-input-json file:///init/swipe_table.json
echo "DynamoDB tables created."

echo "Listing tables"
aws dynamodb list-tables --endpoint-url http://localhost:8000 --region local

# Keep the script running
tail -f /dev/null
