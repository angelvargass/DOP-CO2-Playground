#!/bin/bash

mkdir deployment
cp -R scripts appspec.yml deployment/
mv build/devops-playground deployment/
cd deployment
zip -r ../deployment.zip .
cd ..