#!/usr/bin/env bash

# ./test.sh <url>

# 0 ~ 99 点をランダムに加点
random_add_point() {
    local url=$1
    local device_id=$2
    local add_points=$(($RANDOM % 100))
    curl -s -X POST -H "Content-Type: application/json" -d "{\"device_id\":\"$device_id\", \"huit_points\":$add_points}" $url
}

get_device_id() {
    local id=$(($RANDOM % 10))
    echo "TESTDevice_0${id}"
}

while true
do
    id=`get_device_id`
    random_add_point $1 $id 
    echo $id
    sleep 1
done