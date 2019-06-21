#!/bin/bash

source scripts/echo-color.sh

aws s3api list-buckets

aws s3api create-bucket \
  --bucket ${S3_BUCKET_NAME} \
  --create-bucket-configuration LocationConstraint=ap-northeast-1

aws s3api put-bucket-lifecycle-configuration \
   --bucket ${S3_BUCKET_NAME} \
   --lifecycle-configuration file://scripts/lifecycle.json \

echoSuccess "Created ${S3_BUCKET_NAME}."
