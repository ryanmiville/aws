# aws
common abstractions for interacting with AWS in Go
## DynamoDB
The `dynamodb` package provides an `Iterator` to easily iterate over the results of a PartiQL query.
```go
sess := session.Must(session.NewSession())
client := dynamodb.New(sess)
query := `SELECT * FROM "people" WHERE "Name" = 'Ryan'`
iter := dynamodb.NewIter(client, query)
for iter.Next(ctx) {
	var p Person
	if err := iter.Document(&p); err != nil {
		return err
	}
    //do stuff with p
}
if iter.Err() != nil {
	return iter.Err()
}
```
`Scan` is a helper function for an `Iterator` that scans a table
```go
scanner := dynamodb.Scan(client, "people")
for scanner.Next(ctx) {
	var p Person
	if err := scanner.Document(&p); err != nil {
		return err
	}
    //do stuff with p
}
if scanner.Err() != nil {
	return scanner.Err()
}
```
## S3
The `s3` package provides a `Bucket` abstraction to allow for easy reading and writing to a bucket.

The `NewWriter` method returns an `io.WriteCloser`. The writer MUST be closed for the upload to complete.
```go
sess := session.Must(session.NewSession())
client := s3.New(sess)
bucket := s3.NewBucket(client, "people-bucket")

w := bucket.NewWriter(ctx, "key/to/person.json")
p := newPerson()
if err := json.NewEncoder(w).Encode(&p); err != nil {
    return err
}
// remember to close w
if err := w.Close(); err {
    return err
}
```
Likewise, The `NewReader` method returns an `io.ReadCloser`. The reader must always be closed as well.
```go
r, err := bucket.NewReader(ctx, "key/to/person.json")
if err != nil {
	return err
}
var p Person
if err := json.NewDecoder(r).Decode(&p); err != nil {
	return err
}
// remember to close r
if err := r.Close(); err {
	return err
}
```
