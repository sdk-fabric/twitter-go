
// TweetTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app



import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/apioo/sdkgen-go"
    "io"
    "net/http"
    "net/url"
    
)

type TweetTag struct {
    internal *sdkgen.TagAbstract
}



// GetAll Returns a variety of information about the Tweet specified by the requested ID or list of IDs.
func (client *TweetTag) GetAll(ids string, expansions string, mediaFields string, placeFields string, pollFields string, tweetFields string, userFields string) (TweetCollectionResponse, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["ids"] = ids
    queryParams["expansions"] = expansions
    queryParams["media.fields"] = mediaFields
    queryParams["place.fields"] = placeFields
    queryParams["poll.fields"] = pollFields
    queryParams["tweet.fields"] = tweetFields
    queryParams["user.fields"] = userFields

    u, err := url.Parse(client.internal.Parser.Url("/2/tweets", pathParams))
    if err != nil {
        return TweetCollectionResponse{}, err
    }

    u.RawQuery = client.internal.Parser.Query(queryParams).Encode()


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

// Get Returns a variety of information about a single Tweet specified by the requested ID.
func (client *TweetTag) Get(id string, expansions string, mediaFields string, placeFields string, pollFields string, tweetFields string, userFields string) (TweetEntityResponse, error) {
    pathParams := make(map[string]interface{})
    pathParams["id"] = id

    queryParams := make(map[string]interface{})
    queryParams["expansions"] = expansions
    queryParams["media.fields"] = mediaFields
    queryParams["place.fields"] = placeFields
    queryParams["poll.fields"] = pollFields
    queryParams["tweet.fields"] = tweetFields
    queryParams["user.fields"] = userFields

    u, err := url.Parse(client.internal.Parser.Url("/2/tweets/:id", pathParams))
    if err != nil {
        return TweetEntityResponse{}, err
    }

    u.RawQuery = client.internal.Parser.Query(queryParams).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return TweetEntityResponse{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return TweetEntityResponse{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return TweetEntityResponse{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response TweetEntityResponse
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return TweetEntityResponse{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return TweetEntityResponse{}, errors.New("the server returned an unknown status code")
    }
}

// Create Creates a Tweet on behalf of an authenticated user.
func (client *TweetTag) Create(payload Tweet) (TweetCreateResponse, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    u, err := url.Parse(client.internal.Parser.Url("/2/tweets", pathParams))
    if err != nil {
        return TweetCreateResponse{}, err
    }

    u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return TweetCreateResponse{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return TweetCreateResponse{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return TweetCreateResponse{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return TweetCreateResponse{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response TweetCreateResponse
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return TweetCreateResponse{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return TweetCreateResponse{}, errors.New("the server returned an unknown status code")
    }
}

// Delete Allows a user or authenticated user ID to delete a Tweet.
func (client *TweetTag) Delete(id string) (TweetDeleteResponse, error) {
    pathParams := make(map[string]interface{})
    pathParams["id"] = id

    queryParams := make(map[string]interface{})

    u, err := url.Parse(client.internal.Parser.Url("/2/tweets/:id", pathParams))
    if err != nil {
        return TweetDeleteResponse{}, err
    }

    u.RawQuery = client.internal.Parser.Query(queryParams).Encode()


    req, err := http.NewRequest("DELETE", u.String(), nil)
    if err != nil {
        return TweetDeleteResponse{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return TweetDeleteResponse{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return TweetDeleteResponse{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response TweetDeleteResponse
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return TweetDeleteResponse{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return TweetDeleteResponse{}, errors.New("the server returned an unknown status code")
    }
}



func NewTweetTag(httpClient *http.Client, parser *sdkgen.Parser) *TweetTag {
	return &TweetTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
