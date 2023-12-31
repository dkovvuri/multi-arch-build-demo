#!/bin/bash

AMD_SERVICE=$(aws cloudformation describe-stacks --stack-name ecs-task-demo --query 'Stacks[0].Outputs[?OutputKey==`AMDService`].OutputValue' --output text --region us-east-1)

# Set the URL of the web service to test
URL="http://$AMD_SERVICE/stress?number=99999999999999"

# Set the number of requests to be made
NUM_REQUESTS=10000000

# Set the number of concurrent requests to be made
CONCURRENCY=10

# Perform the load test
for ((i=1; i<=$CONCURRENCY; i++))
do
  {
    for ((j=1; j<=$NUM_REQUESTS/$CONCURRENCY; j++))
    do
      curl -s $URL
    done
  } &
done

# Wait for all background processes to finish
wait
