
// RetweetTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    
    "encoding/json"
    "errors"
    "github.com/apioo/sdkgen-go"
    "io"
    "net/http"
    "net/url"
    
)

type RetweetTag struct {
    internal *sdkgen.TagAbstract
}



// GetAll 
func (client *RetweetTag) GetAll(tweetId string, expansions string, maxResults int, fields Fields) (TweetCollection, error) {
    pathParams := make(map[string]interface{})
    pathParams["tweet_id"] = tweetId

    queryParams := make(map[string]interface{})
    queryParams["expansions"] = expansions
    queryParams["max_results"] = maxResults
    queryParams["fields"] = fields

    var queryStructNames []string
    append(queryStructNames, '0')

    u, err := url.Parse(client.internal.Parser.Url("/2/tweets/:tweet_id/retweets", pathParams))
    if err != nil {
        return TweetCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return TweetCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return TweetCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return TweetCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response TweetCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return TweetCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return TweetCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewRetweetTag(httpClient *http.Client, parser *sdkgen.Parser) *RetweetTag {
	return &RetweetTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
