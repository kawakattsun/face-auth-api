#!/bin/bash

aws rekognition create-collection \
  --collection-id "${REKOGNITION_COLLECTION_ID}"
