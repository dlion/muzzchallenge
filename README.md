# Muzz Coding Challenge

## Test

The `server/server_test.go` file contains integration tests that use `testcontainers` to spawn a `dynamodb-local` instance and test different scenarios against it.

> Be sure to have Docker up and running

`go test ./...`

## How to run

### Run DynamoDB

1. `docker-compose up`

### Run the server

1. `go run main.go` (by default it listens to the port `37857`)

### Parameters
* `p`: Service Port (Default: `37857`)
* `dynamoEndpoint`: DynamoDB Endpoint (Default: `http://localhost:8000`)
* `dynamoAccessKey`: DynamoDB Accesskey (Default: `dummy`)
* `dynamoSecretAccessKey`: DynamoDB SecretAccessKey (Default: `dummy`)
* `dynamoRegion`: DynamoDB Region (Default: `local`)

### Thoughts during the way

* I'm using DynamoDB to be as closer as possible to the real implementation
* I started implementing the swipe feature to be able to add new items into the db
* I also added a condition on this key to make it idempotent. It means that if I receive multiple requests with the same pk, it doesn't overwrite the items in the db with new values.
* In the LikedYou functionality I'm using the scanning to get all the people (actorId) who liked a certain account (recipientId) who liked me (like = true), with a gender filter (`actor_gender`), using the timestamp in descending order (timestamp)
* In the LikedYou response my assumption is that the `timestamp` is the timestamp of when other profiles encounter the `MarriageProfileId`.


## High-level explanation

* The function `PutSwipe` is used to populate the `Swipe` table. It has 2 main functionality:
    1. Populate the db with new swipe entries.
    2. Update the `likedBack` recipient's field according to the `like` in the actor's request.
* The function `LikedYou` is used to get back profiles according to the filter that has been used:
    1. `LIKE_YOU_NEW`: Returns all profiles that have the `like`field set to true, and have a `likedBack` field set to false.
    2. `LIKED_YOU_SWIPED`: Returns all profiles that have the `liked` and `likedBack` fields set to true.

## Low-level explanation

### The DB

* The db has a `Swipe` table where entries are added every time we call the `PutSwipe` function.
* The primary key is composed by the `actor_marriage_profile_id` and the `recipient_marriage_profile_id`.
* Beside the `PutSwipeRequest` request fields, we add to the db a `likedBack` field which identify if during that swipe we got a match. (the recipient gave a like back)

### PutSwipe
* Update the recipient's `likedBack` field. I.e. If the `actor` likes the `recipient`, then we set the `recipient`'s `likedBack` field as a true.
* Add the `actor` entry into the database and its `likedBack` field based on the `recipient`'s like field.

### LikedYou
* Get all the `actor`profiles who liked a specific marriage profile ID. Based on the filter applied we get who got a likedBack from our married profile or not.
* We sort all profiles in descending order
* If a limit is set it returns just a subset of all profiles.

