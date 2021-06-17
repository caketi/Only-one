#!/bin/bash
read -p "ee" n

for a in 	`seq 1 $n`
do 
echo -n "$a+"  
let sum+=a
 let a++
done
echo "  "
echo $sum
