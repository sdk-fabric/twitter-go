
// UserTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/apioo/sdkgen-go"
    "io"
    "net/http"
    "net/url"
    
)

type UserTag struct {
    internal *sdkgen.TagAbstract
}



// GetTimeline Allows you to retrieve a collection of the most recent Tweets and Retweets posted by you and users you follow. This endpoint can return every Tweet created on a timeline over the last 7 days as well as the most recent 800 regardless of creation date.
func (client *UserTag) GetTimeline(userId string, exclude string, expansions string, pagination Pagination, fields Fields) (TweetCollection, error) {
    pathParams := make(map[string]interface{})
    pathParams["user_id"] = userId

    queryParams := make(map[string]interface{})
    queryParams["exclude"] = exclude
    queryParams["expansions"] = expansions
    queryParams["pagination"] = pagination
    queryParams["fields"] = fields

    var queryStructNames []string
    append(queryStructNames, '0')
    append(queryStructNames, '1')

    u, err := url.Parse(client.internal.Parser.Url("/2/users/:user_id/timelines/reverse_chronological", pathParams))
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

// GetLikedTweets Tweets liked by a user
func (client *UserTag) GetLikedTweets(userId string, expansions string, maxResults int, paginationToken string, fields Fields) (TweetCollection, error) {
    pathParams := make(map[string]interface{})
    pathParams["user_id"] = userId

    queryParams := make(map[string]interface{})
    queryParams["expansions"] = expansions
    queryParams["max_results"] = maxResults
    queryParams["pagination_token"] = paginationToken
    queryParams["fields"] = fields

    var queryStructNames []string
    append(queryStructNames, '0')

    u, err := url.Parse(client.internal.Parser.Url("/2/users/:user_id/liked_tweets", pathParams))
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

// RemoveLike Allows a user or authenticated user ID to unlike a Tweet.
func (client *UserTag) RemoveLike(userId string, tweetId string) (LikeResponse, error) {
    pathParams := make(map[string]interface{})
    pathParams["user_id"] = userId
    pathParams["tweet_id"] = tweetId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/2/users/:user_id/likes/:tweet_id", pathParams))
    if err != nil {
        return LikeResponse{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("DELETE", u.String(), nil)
    if err != nil {
        return LikeResponse{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return LikeResponse{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return LikeResponse{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response LikeResponse
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return LikeResponse{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return LikeResponse{}, errors.New("the server returned an unknown status code")
    }
}

// CreateLike Causes the user ID identified in the path parameter to Like the target Tweet.
func (client *UserTag) CreateLike(userId string, payload SingleTweet) (LikeResponse, error) {
    pathParams := make(map[string]interface{})
    pathParams["user_id"] = userId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/2/users/:user_id/likes", pathParams))
    if err != nil {
        return LikeResponse{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return LikeResponse{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return LikeResponse{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return LikeResponse{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return LikeResponse{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response LikeResponse
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return LikeResponse{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return LikeResponse{}, errors.New("the server returned an unknown status code")
    }
}

// FindByName Returns a variety of information about one or more users specified by their usernames.
func (client *UserTag) FindByName(usernames string, expansions string, fields Fields) (UserCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["usernames"] = usernames
    queryParams["expansions"] = expansions
    queryParams["fields"] = fields

    var queryStructNames []string
    append(queryStructNames, '0')

    u, err := url.Parse(client.internal.Parser.Url("/2/users/by", pathParams))
    if err != nil {
        return UserCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return UserCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return UserCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return UserCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response UserCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return UserCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        default:
            return UserCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewUserTag(httpClient *http.Client, parser *sdkgen.Parser) *UserTag {
	return &UserTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
