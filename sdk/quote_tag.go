
// QuoteTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app



import (
    
    "encoding/json"
    "errors"
    "fmt"
    "github.com/apioo/sdkgen-go/v2"
    "io"
    "net/http"
    "net/url"
    
)

type QuoteTag struct {
    internal *sdkgen.TagAbstract
}



// GetAll Returns Quote Tweets for a Tweet specified by the requested Tweet ID.
func (client *QuoteTag) GetAll(tweetId string, exclude string, expansions string, maxResults int, paginationToken string, fields Fields) (*TweetCollection, error) {
    pathParams := make(map[string]interface{})
    pathParams["tweet_id"] = tweetId

    queryParams := make(map[string]interface{})
    queryParams["exclude"] = exclude
    queryParams["expansions"] = expansions
    queryParams["max_results"] = maxResults
    queryParams["pagination_token"] = paginationToken
    queryParams["fields"] = fields

    var queryStructNames []string
    append(queryStructNames, '0')

    u, err := url.Parse(client.internal.Parser.Url("/2/tweets/:tweet_id/quote_tweets", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data TweetCollection
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data Errors
        err := json.Unmarshal(respBody, &data)

        return nil, &ErrorsException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewQuoteTag(httpClient *http.Client, parser *sdkgen.Parser) *QuoteTag {
	return &QuoteTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
