# CAFA API go Wrapper

This is an SDK for existing microservice called CAFA. It's APIs are deployed and this code enabled easy integration with them.
### Example
For example look at `example_client/main.go` - there are all possible requests used.

For example with special access: [Doc](https://erply.atlassian.net/wiki/spaces/CAFA/blog/2020/04/08/403210241/Special+access+without+the+session+key+HMAC), look at `example_special_access_client/main.go`

### Swagger
To get swagger documentation of the deployed CAFA API - please request getServiceEndpoints ERPLY API endpoint.
More info here: https://manual.erply.com/eng/erply-api/which-api-to-use

### Import 
Import as a private library "gl.nimi24.com/back-office/cafa-go-wrapper/pkg/cafa"
### Author
* david.zingerman@erply.com

### Versions
* v1.0.1 - added special access example client, key generation function