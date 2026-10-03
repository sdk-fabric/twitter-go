
# twitter-go

This [SDK](https://github.com/sdk-fabric/twitter-go) is managed by the [SDK Fabric](https://sdk-fabric.org/) project, a global infrastructure to
automatically generate SDKs for every API.

You can find more information about this SDK at [TypeHub](https://typehub.cloud/):
https://app.typehub.cloud/d/sdkfabric/twitter

## Usage

```go
import (
	"github.com/sdk-fabric/twitter-go/sdk"
)

var client, _ = sdk.Build("[access_token]");

// Returns a variety of information about the Tweet specified by the requested ID or list of IDs.
response, err := client.Tweet().getAll("ids", "expansions", nil)

// Returns a variety of information about a single Tweet specified by the requested ID.
response, err := client.Tweet().get("tweet_id", "expansions", nil)

// Creates a Tweet on behalf of an authenticated user.
response, err := client.Tweet().create(Tweet{})

// Allows a user or authenticated user ID to delete a Tweet.
response, err := client.Tweet().delete("tweet_id")

// Hides or unhides a reply to a Tweet.
response, err := client.Tweet().hideReply("tweet_id", HideReply{})

// Allows you to get information about a Tweet’s liking users.
response, err := client.Tweet().getLikingUsers("tweet_id", "expansions", 1, "pagination_token")

// The Usage API in the Twitter API v2 allows developers to programmatically retrieve their project usage.
response, err := client.Usage().getTweets()

// Returns a variety of information about one or more users specified by the requested IDs.
response, err := client.User().getAll("ids", "expansions", nil)

// Returns a variety of information about a single user specified by the requested ID.
response, err := client.User().get("user_id", "expansions", nil)

// Allows you to retrieve a collection of the most recent Tweets and Retweets posted by you and users you follow.
response, err := client.User().getTimeline("user_id", "exclude", "expansions", nil, nil)

// Tweets liked by a user.
response, err := client.User().getLikedTweets("user_id", "expansions", 1, "pagination_token", nil)

// Allows a user or authenticated user ID to unlike a Tweet.
response, err := client.User().removeLike("user_id", "tweet_id")

// Causes the user ID identified in the path parameter to Like the target Tweet.
response, err := client.User().createLike("user_id", Single_Tweet{})

// Returns a variety of information about one or more users specified by their usernames.
response, err := client.User().findByName("usernames", "expansions", nil)

// Returns information about an authorized user.
response, err := client.User().getMe("expansions", "fields")

// Allows you to get an authenticated user's 800 most recent bookmarked Tweets.
response, err := client.Bookmark().getAll("user_id", "expansions", "pagination_token", nil)

response, err := client.Bookmark().create("user_id", Single_Tweet{})

response, err := client.Bookmark().delete("user_id", "tweet_id")

response, err := client.Search().getRecent("query", "sort_order", "expansions", nil, nil)

// Returns Quote Tweets for a Tweet specified by the requested Tweet ID.
response, err := client.Quote().getAll("tweet_id", "exclude", "expansions", 1, "pagination_token", nil)

// The Trends lookup endpoint allow developers to get the Trends for a location, specified using the where-on-earth id (WOEID).
response, err := client.Trends().getByWoeid("woeid")

// Returns the Retweets for a given Tweet ID.
response, err := client.Retweet().getAll("tweet_id", "expansions", 1, nil)
```
