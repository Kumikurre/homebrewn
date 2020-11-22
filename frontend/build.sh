#!/bin/bash

echo "Running ng build"
npm start

echo "Copying compiled js files to /tmp/dist"
cp -r dist /tmp/
echo "Copying static files to /tmp/dist"
cp -r src/app/* /frontend/dist
