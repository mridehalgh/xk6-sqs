# xk6-sqs

This is a [k6](https://go.k6.io/k6) extension using the [xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team or by the maintainer, and may break in the future. USE AT YOUR OWN RISK! |
|------|

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then, install [xk6](https://github.com/k6io/xk6) and build your custom k6 binary with the SQS extension:

1. Install `xk6`:

  ```shell
  $ go install github.com/k6io/xk6/cmd/xk6@latest
  ```

2. Build the binary:

  ```shell
  $ xk6 build --with github.com/mridehalgh/xk6-sqs@latest
  ```

## AWS credentials

This plugin uses the AWS SDK Go v2 default credential chain. It looks for credentials in the following order:

1. Environment variables.
   1. Static Credentials (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN`)
   2. Web Identity Token (`AWS_WEB_IDENTITY_TOKEN_FILE`)
1. Shared configuration files.
   1. SDK defaults to `credentials` file under `.aws` folder that is placed in the home folder on your computer.
   1. SDK defaults to `config` file under `.aws` folder that is placed in the home folder on your computer.
1. If your application uses an ECS task definition or RunTask API operation, IAM role for tasks.
1. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.

Source: https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specifying-credentials

## Example

```javascript
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
        QueueUrl: "QUEUE_URL"
    };

    sqs.send(client,params)
}

```
