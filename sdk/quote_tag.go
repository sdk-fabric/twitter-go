
// QuoteTag automatically generated by SDKgen please do not edit this file manually
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

type QuoteTag struct {
    internal *sdkgen.TagAbstract
}



// GetAll Returns Quote Tweets for a Tweet specified by the requested Tweet ID.
func (client *QuoteTag) GetAll(tweetId string, exclude string, expansions string, maxResults int, paginationToken string, fields Fields) (TweetCollectionResponse, error) {
    pathParams := make(map[string]interface{})
    pathParams["tweet_id"] = tweetId

    queryParams := make(map[string]interface{})
    queryParams["exclude"] = exclude
    queryParams["expansions"] = expansions
    queryParams["max_results"] = maxResults
    queryParams["pagination_token"] = paginationToken
    queryParams["fields"] = fields

    var queryStructNames []string
    append(queryStructNames, '0'),

    u, err := url.Parse(client.internal.Parser.Url("/2/tweets/:tweet_id/quote_tweets", pathParams))
    if err != nil {
        return TweetCollectionResponse{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return TweetCollectionResponse{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return TweetCollectionResponse{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return TweetCollectionResponse{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response TweetCollectionResponse
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return TweetCollectionResponse{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return TweetCollectionResponse{}, errors.New("the server returned an unknown status code")
    }
}



func NewQuoteTag(httpClient *http.Client, parser *sdkgen.Parser) *QuoteTag {
	return &QuoteTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
