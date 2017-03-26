# http2sqs
Server that dumps HTTP POST content to SQS

### Purpose
Can be used as a simple and robust webhook receiver.

### Run it with Docker
```
docker run -e SQS_QUEUE_URL=<https://queue.amazonaws.com/SOMENUMBER/QUEUE_NAME> cxmcc/http2sqs
```
