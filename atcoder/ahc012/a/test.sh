#!/bin/sh
for i in {0..99}
do
    inid=`printf %04d ${i}`
    outid=`printf %02d ${i}`
    go run main.go < in/${inid}.txt > out/${outid}.txt
done