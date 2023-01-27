[![Go Report Card](https://goreportcard.com/badge/github.com/stianfro/azjson)](https://goreportcard.com/report/github.com/stianfro/azjson) [![Go Reference](https://pkg.go.dev/badge/github.com/stianfro/azjson.svg)](https://pkg.go.dev/github.com/stianfro/azjson) ![Github Actions](https://github.com/stianfro/azjson/actions/workflows/go.yml/badge.svg)

# azjson

Simple library for working with APIs with Azure AD authentication.

# Examples

```go
type response struct {
  ID   string `json:"id"`
  Name string `json:"name"`
}

func main() {
  var responseJSON response
  token := azcore.AccessToken{}

  res, err := GetJSON("https://example.com", token)
  if err != nil {
  	fmt.Println(err)
  }

  jsonErr := json.Unmarshal([]byte(companyRes), &responseJSON)
  if jsonErr != nil {
  	fmt.Println(jsonErr)
  }
}
```
