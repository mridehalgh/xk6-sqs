import sqs from 'k6/x/sqs';

const client = sqs.newClient()

export default function () {
    const params = {
        DelaySeconds: 0,
        MessageAttributes: {
            "Title": {
                DataType: "String",
                StringValue: "The Whistler"
            },
            "Author": {
                DataType: "String",
                StringValue: "John Grisham"
            },
            "WeeksOn": {
                DataType: "Number",
                StringValue: "6"
            }
        },
        MessageBody: "Information about current NY Times fiction bestseller for week of 12/11/2016.",
        // MessageDeduplicationId: "TheWhistler",  // Required for FIFO queues
        // MessageGroupId: "Group1",  // Required for FIFO queues
        QueueUrl: "http://localhost:4566/000000000000/dummy-k6-queue"
    };

    sqs.send(client, params)
}
