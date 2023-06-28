#!/bin/bash

# Define the function to invoke the stress test endpoint
invoke_stress_test() {
  while true; do
    curl -sS "http://localhost:8080/stress?number=99999999999999"
    sleep 0.1  # Adjust the sleep duration as needed
  done
}

# Set the number of simultaneous sessions
num_sessions=20

# Run the stress test in the background for each session
for ((i=1; i<=num_sessions; i++)); do
  invoke_stress_test &
done

# Wait for all background processes to finish
wait