# http2sqs
Server that dumps HTTP POST content to SQS

### Purpose
Can be used as a simple and robust webhook receiver.

### Configurations

The application accepts environemnt variables.

For the AWS credential variables, you can totally just use AMI Roles if it's on EC2.

| Variable                | Description   | Required?  |
| ----------------------- |:-------------:| ----------:|
| `SQS_QUEUE_URL`         | SQS queue URL | yes        |
| `AWS_REGION`            |               | yes        |
| `AWS_ACCESS_KEY_ID`     |               | no         |
| `AWS_SECRET_ACCESS_KEY` |               | no         |

### Hosted Docker Image

[cxmcc/http2sqs](https://hub.docker.com/r/cxmcc/http2sqs/)

### Run it with Docker
```bash
docker run --env SQS_QUEUE_URL=<https://queue.amazonaws.com/SOMENUMBER/QUEUE_NAME> \
           --env AWS_REGION=us-east-1 \
           --publish 8080:8080 \
           cxmcc/http2sqs
```
