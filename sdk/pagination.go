
// pagination automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk
type Pagination struct {
    StartTime string `json:"start_time"`
    EndTime string `json:"end_time"`
    SinceId string `json:"since_id"`
    UntilId string `json:"until_id"`
    MaxResults int `json:"max_results"`
    PaginationToken string `json:"pagination_token"`
}