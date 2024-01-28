
// tweet automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


type Tweet struct {
    DirectMessageDeepLink string `json:"direct_message_deep_link"`
    ForSuperFollowersOnly bool `json:"for_super_followers_only"`
    Geo TweetGeo `json:"geo"`
    Media TweetMedia `json:"media"`
    Poll TweetPoll `json:"poll"`
    QuoteTweetId string `json:"quote_tweet_id"`
    Reply TweetReply `json:"reply"`
    ReplySettings string `json:"reply_settings"`
    Text string `json:"text"`
}
