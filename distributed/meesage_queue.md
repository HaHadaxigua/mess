# Core of Message Queue

1. message
2. queue
3. distributed
4. fault tolerance

# 1. Message

Each message is processed only once by a single consumer.

Producer sends messages to the queue, then consumer will consume it in async.

Then MQ need to lock the message to prevent someone processes this message replicated.

After consumer processing the message, MQ need to delete the message from queue, to prevent
someone re-processing it.

## key words

### 1. idempotent

Push or Pull Message might be failed, and we need to re-push/re-pull message.

### 2. persistence

Based on AWS s3.

### 3. API

API TO MAINTAIN MESSAGES.

