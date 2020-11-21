#!/bin/bash

echo "Running ng build"
npm start

echo "Copying file to /tmp/dist"
cp -r dist /tmp/
