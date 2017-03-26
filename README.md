# http2sqs
Server that dumps HTTP POST content to SQS

### Run with Docker
```
docker run -e SQS_QUEUE_URL=<https://queue.amazonaws.com/SOMENUMBER/QUEUE_NAME> cxmcc/http2sqs
```
