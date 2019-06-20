#!/bin/bash

aws s3api list-buckets

aws s3api create-bucket \
  --bucket ${ARTIFACT_BUCKET_NAME} \
  --create-bucket-configuration LocationConstraint=ap-northeast-1
aws s3api put-bucket-lifecycle-configuration \
   --bucket ${ARTIFACT_BUCKET_NAME} \
   --lifecycle-configuration file://scripts/lifecycle.json

echo "created ${ARTIFACT_BUCKET_NAME}"
