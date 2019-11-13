#!/usr/bin/env bash

aws s3 cp api/swagger.yml s3://${S3_BUCKET_NAME_ARTIFACT}/swagger.yml

sam package \
  --template-file ./template.yml \
  --s3-bucket "${S3_BUCKET_NAME_ARTIFACT}" \
  --output-template-file ./packaged.yml && \
sam deploy \
  --template-file ./packaged.yml \
  --stack-name "${STACK_NAME}" \
  --capabilities CAPABILITY_IAM \
  --no-fail-on-empty-changeset \
  --parameter-overrides \
    ArtifactBucket=${S3_BUCKET_NAME_ARTIFACT} \
    RekognitionCollectionID=${REKOGNITION_COLLECTION_ID}
