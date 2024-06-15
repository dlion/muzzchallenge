FROM ubuntu:latest

# Install OpenJDK 11 and other dependencies
RUN apt-get update && apt-get install -y openjdk-11-jre curl unzip

# Set working directory
WORKDIR /home/dynamodblocal

# Download and extract DynamoDB Local
RUN curl -L -o dynamodb_local_latest.tar.gz https://s3.us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.tar.gz \
    && tar -xvzf dynamodb_local_latest.tar.gz \
    && rm dynamodb_local_latest.tar.gz

# Install AWS CLI v2 for ARM to run it on M series
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip" -o "awscliv2.zip" \
&& unzip awscliv2.zip \
&& ./aws/install \
&& rm -rf awscliv2.zip aws

# Install AWS CLI v3 for Intel
# RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" \
#     && unzip awscliv2.zip \
#     && ./aws/install \
#     && rm -rf awscliv2.zip aws

# Copy initialization scripts
COPY init /init

# Set entrypoint for container
ENTRYPOINT ["/bin/sh", "/init/init.sh"]

# Default command to start DynamoDB Local
CMD ["-jar", "DynamoDBLocal.jar", "-sharedDb", "-dbPath", "./data"]
