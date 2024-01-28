#!/bin/bash

# Specify the number of iterations
num_iterations=10

# Counter variable
i=1

# URL to be used in the curl request
url="http://localhost:8080/hello"

while [ $i -le $num_iterations ]; do
    echo "Iteration $i"
    curl -X GET "$url"
    echo "----------------------------------"
    ((i++))
done
