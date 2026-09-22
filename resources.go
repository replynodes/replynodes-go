// Code generated from the canonical GET operation registry. DO NOT EDIT BY HAND.

package replynodes

import "context"

// AppStoreService exposes the canonical AppStore operations.
type AppStoreService struct{ client *Client }

// BrandService exposes the canonical Brand operations.
type BrandService struct{ client *Client }

// FomoService exposes the canonical Fomo operations.
type FomoService struct{ client *Client }

// GoogleService exposes the canonical Google operations.
type GoogleService struct{ client *Client }

// GoogleMapsService exposes the canonical GoogleMaps operations.
type GoogleMapsService struct{ client *Client }

// GooglePlayService exposes the canonical GooglePlay operations.
type GooglePlayService struct{ client *Client }

// GoogleShoppingService exposes the canonical GoogleShopping operations.
type GoogleShoppingService struct{ client *Client }

// HackerNewsService exposes the canonical HackerNews operations.
type HackerNewsService struct{ client *Client }

// InstagramService exposes the canonical Instagram operations.
type InstagramService struct{ client *Client }

// RedditService exposes the canonical Reddit operations.
type RedditService struct{ client *Client }

// TiktokService exposes the canonical Tiktok operations.
type TiktokService struct{ client *Client }

// WebService exposes the canonical Web operations.
type WebService struct{ client *Client }

// YoutubeService exposes the canonical Youtube operations.
type YoutubeService struct{ client *Client }

// AppStoreAppParams contains the path and query parameters for appStoreApp.
type AppStoreAppParams struct {
	ID      string `query:"id"`
	AppID   string `query:"appId"`
	Country string `query:"country"`
	Lang    string `query:"lang"`
	Ratings bool   `query:"ratings"`
}

// AppStoreDeveloperParams contains the path and query parameters for appStoreDeveloper.
type AppStoreDeveloperParams struct {
	DevID   string `query:"devId,required"`
	Country string `query:"country"`
	Lang    string `query:"lang"`
}

// AppStoreListParams contains the path and query parameters for appStoreList.
type AppStoreListParams struct {
	Collection string `query:"collection"`
	Category   string `query:"category"`
	Country    string `query:"country"`
	Lang       string `query:"lang"`
	Num        int32  `query:"num"`
	FullDetail bool   `query:"fullDetail"`
}

// AppStorePrivacyParams contains the path and query parameters for appStorePrivacy.
type AppStorePrivacyParams struct {
	ID string `query:"id,required"`
}

// AppStoreRatingsParams contains the path and query parameters for appStoreRatings.
type AppStoreRatingsParams struct {
	ID      string `query:"id"`
	Country string `query:"country"`
	AppID   string `query:"appId"`
}

// AppStoreReviewsParams contains the path and query parameters for appStoreReviews.
type AppStoreReviewsParams struct {
	ID      string `query:"id"`
	AppID   string `query:"appId"`
	Country string `query:"country"`
	Page    int32  `query:"page"`
	Sort    string `query:"sort"`
}

// AppStoreSearchParams contains the path and query parameters for appStoreSearch.
type AppStoreSearchParams struct {
	Term    string `query:"term,required"`
	Num     int32  `query:"num"`
	Page    int32  `query:"page"`
	Country string `query:"country"`
	Lang    string `query:"lang"`
	IdsOnly bool   `query:"idsOnly"`
}

// AppStoreSimilarParams contains the path and query parameters for appStoreSimilar.
type AppStoreSimilarParams struct {
	ID    string `query:"id"`
	AppID string `query:"appId"`
}

// AppStoreSuggestParams contains the path and query parameters for appStoreSuggest.
type AppStoreSuggestParams struct {
	Term string `query:"term,required"`
}

// BrandFontsParams contains the path and query parameters for brandFonts.
type BrandFontsParams struct {
	Domain string `query:"domain"`
	URL    string `query:"url"`
}

// BrandRetrieveParams contains the path and query parameters for brandRetrieve.
type BrandRetrieveParams struct {
	Domain string `query:"domain"`
	URL    string `query:"url"`
}

// BrandSearchParams contains the path and query parameters for brandSearch.
type BrandSearchParams struct {
	Query string `query:"query,required"`
	Limit int32  `query:"limit"`
}

// BrandStyleguideParams contains the path and query parameters for brandStyleguide.
type BrandStyleguideParams struct {
	Domain string `query:"domain"`
	URL    string `query:"url"`
}

// FomoAlertsParams contains the path and query parameters for fomoAlerts.
type FomoAlertsParams struct {
	Type  string `query:"type"`
	Chain string `query:"chain"`
	Since string `query:"since"`
	Limit int32  `query:"limit"`
}

// FomoLeaderboardParams contains the path and query parameters for fomoLeaderboard.
type FomoLeaderboardParams struct {
	Window string `path:"window,required"`
	Chain  string `query:"chain"`
	Limit  int32  `query:"limit"`
}

// FomoNotificationsParams contains the path and query parameters for fomoNotifications.
type FomoNotificationsParams struct {
	Type  string `query:"type"`
	Chain string `query:"chain"`
	Since string `query:"since"`
	Limit int32  `query:"limit"`
}

// FomoSearchParams contains the path and query parameters for fomoSearch.
type FomoSearchParams struct {
	Q     string `query:"q,required"`
	Type  string `query:"type"`
	Chain string `query:"chain"`
	Limit int32  `query:"limit"`
}

// FomoThesisParams contains the path and query parameters for fomoThesis.
type FomoThesisParams struct {
	Chain string `query:"chain"`
	Limit int32  `query:"limit"`
}

// FomoThesisByTokenParams contains the path and query parameters for fomoThesisByToken.
type FomoThesisByTokenParams struct {
	Mint    string `path:"mint,required"`
	Network string `query:"network"`
	Limit   int32  `query:"limit"`
}

// FomoThesisByUserParams contains the path and query parameters for fomoThesisByUser.
type FomoThesisByUserParams struct {
	ID    string `path:"id,required"`
	Chain string `query:"chain"`
	Sort  string `query:"sort"`
	Limit int32  `query:"limit"`
}

// FomoThesisByUserTokenParams contains the path and query parameters for fomoThesisByUserToken.
type FomoThesisByUserTokenParams struct {
	ID      string `path:"id,required"`
	Address string `path:"address,required"`
	Limit   int32  `query:"limit"`
}

// FomoTokenHoldersParams contains the path and query parameters for fomoTokenHolders.
type FomoTokenHoldersParams struct {
	Address string `path:"address,required"`
	Limit   int32  `query:"limit"`
}

// FomoTokensGraduatedParams contains the path and query parameters for fomoTokensGraduated.
type FomoTokensGraduatedParams struct {
	Chain string `query:"chain"`
	Limit int32  `query:"limit"`
}

// FomoTokensMostHeldParams contains the path and query parameters for fomoTokensMostHeld.
type FomoTokensMostHeldParams struct {
	Chain string `query:"chain"`
	Limit int32  `query:"limit"`
}

// FomoTokensTrendingParams contains the path and query parameters for fomoTokensTrending.
type FomoTokensTrendingParams struct {
	Chain string `query:"chain"`
	Limit int32  `query:"limit"`
}

// FomoTradeParams contains the path and query parameters for fomoTrade.
type FomoTradeParams struct {
	TradeID string `path:"trade_id,required"`
}

// FomoUserBalancesParams contains the path and query parameters for fomoUserBalances.
type FomoUserBalancesParams struct {
	Handle string `path:"handle,required"`
}

// FomoUserProfileParams contains the path and query parameters for fomoUserProfile.
type FomoUserProfileParams struct {
	Handle string `path:"handle,required"`
}

// FomoUserTradesParams contains the path and query parameters for fomoUserTrades.
type FomoUserTradesParams struct {
	Handle string `path:"handle,required"`
	Limit  int32  `query:"limit"`
}

// GoogleMapsPlaceDetailsParams contains the path and query parameters for googleMapsPlaceDetails.
type GoogleMapsPlaceDetailsParams struct {
	PlaceID string `path:"place_id,required"`
}

// GoogleMapsPlaceReviewsParams contains the path and query parameters for googleMapsPlaceReviews.
type GoogleMapsPlaceReviewsParams struct {
	PlaceID string `path:"place_id,required"`
	Limit   int32  `query:"limit"`
}

// GoogleMapsSearchPlacesParams contains the path and query parameters for googleMapsSearchPlaces.
type GoogleMapsSearchPlacesParams struct {
	Query  string `query:"query,required"`
	Lat    string `query:"lat"`
	Lng    string `query:"lng"`
	Radius string `query:"radius"`
}

// GooglePlayAppDetailsParams contains the path and query parameters for googlePlayAppDetails.
type GooglePlayAppDetailsParams struct {
	ID       string `path:"id,required"`
	Country  string `query:"country"`
	Language string `query:"language"`
}

// GooglePlayAvailabilityParams contains the path and query parameters for googlePlayAvailability.
type GooglePlayAvailabilityParams struct {
	ID        string `path:"id,required"`
	Country   string `query:"country"`
	Countries string `query:"countries"`
	Language  string `query:"language"`
}

// GooglePlayCategoriesParams contains the path and query parameters for googlePlayCategories.
type GooglePlayCategoriesParams struct {
	Country  string `query:"country"`
	Language string `query:"language"`
}

// GooglePlayCategoryAppsParams contains the path and query parameters for googlePlayCategoryApps.
type GooglePlayCategoryAppsParams struct {
	Category string `path:"category,required"`
	Country  string `query:"country"`
	Language string `query:"language"`
	Limit    int32  `query:"limit"`
}

// GooglePlayDataSafetyParams contains the path and query parameters for googlePlayDataSafety.
type GooglePlayDataSafetyParams struct {
	ID       string `path:"id,required"`
	Country  string `query:"country"`
	Language string `query:"language"`
}

// GooglePlayDeveloperParams contains the path and query parameters for googlePlayDeveloper.
type GooglePlayDeveloperParams struct {
	DevID    string `query:"dev_id,required"`
	Country  string `query:"country"`
	Language string `query:"language"`
	Limit    int32  `query:"limit"`
}

// GooglePlayPermissionsParams contains the path and query parameters for googlePlayPermissions.
type GooglePlayPermissionsParams struct {
	ID       string `path:"id,required"`
	Country  string `query:"country"`
	Language string `query:"language"`
}

// GooglePlayReviewsParams contains the path and query parameters for googlePlayReviews.
type GooglePlayReviewsParams struct {
	ID        string `path:"id,required"`
	Country   string `query:"country"`
	Language  string `query:"language"`
	Limit     int32  `query:"limit"`
	Sort      int32  `query:"sort"`
	Score     int32  `query:"score"`
	NextToken string `query:"next_token"`
}

// GooglePlaySearchParams contains the path and query parameters for googlePlaySearch.
type GooglePlaySearchParams struct {
	Term     string `query:"term,required"`
	Country  string `query:"country"`
	Language string `query:"language"`
	Limit    int32  `query:"limit"`
}

// GooglePlaySimilarAppsParams contains the path and query parameters for googlePlaySimilarApps.
type GooglePlaySimilarAppsParams struct {
	ID       string `path:"id,required"`
	Country  string `query:"country"`
	Language string `query:"language"`
	Limit    int32  `query:"limit"`
}

// GooglePlaySuggestParams contains the path and query parameters for googlePlaySuggest.
type GooglePlaySuggestParams struct {
	Term     string `query:"term,required"`
	Country  string `query:"country"`
	Language string `query:"language"`
	Limit    int32  `query:"limit"`
}

// GoogleSearchParams contains the path and query parameters for googleSearch.
type GoogleSearchParams struct {
	Text    string `query:"text,required"`
	Engines string `query:"engines"`
	Lang    string `query:"lang"`
	Region  string `query:"region"`
	Date    string `query:"date"`
	Site    string `query:"site"`
	Limit   int32  `query:"limit"`
	Start   int32  `query:"start"`
	Cursor  string `query:"cursor"`
}

// GoogleShoppingProductOffersParams contains the path and query parameters for googleShoppingProductOffers.
type GoogleShoppingProductOffersParams struct {
	ProductID string `query:"product_id"`
	Text      string `query:"text,required"`
	Lang      string `query:"lang"`
	Country   string `query:"country"`
	Domain    string `query:"domain"`
	Limit     int32  `query:"limit"`
}

// GoogleShoppingSearchParams contains the path and query parameters for googleShoppingSearch.
type GoogleShoppingSearchParams struct {
	Text    string `query:"text,required"`
	Lang    string `query:"lang"`
	Country string `query:"country"`
	Domain  string `query:"domain"`
	Limit   int32  `query:"limit"`
	Start   int32  `query:"start"`
}

// HackerNewsItemParams contains the path and query parameters for hackerNewsItem.
type HackerNewsItemParams struct {
	ID    int32 `path:"id,required"`
	Depth int32 `query:"depth"`
}

// HackerNewsSearchParams contains the path and query parameters for hackerNewsSearch.
type HackerNewsSearchParams struct {
	Q     string `query:"q,required"`
	Tags  string `query:"tags"`
	Limit int32  `query:"limit"`
	Page  int32  `query:"page"`
}

// HackerNewsStoriesAskParams contains the path and query parameters for hackerNewsStoriesAsk.
type HackerNewsStoriesAskParams struct {
	Limit int32 `query:"limit"`
	Page  int32 `query:"page"`
}

// HackerNewsStoriesBestParams contains the path and query parameters for hackerNewsStoriesBest.
type HackerNewsStoriesBestParams struct {
	Limit int32 `query:"limit"`
	Page  int32 `query:"page"`
}

// HackerNewsStoriesJobParams contains the path and query parameters for hackerNewsStoriesJob.
type HackerNewsStoriesJobParams struct {
	Limit int32 `query:"limit"`
	Page  int32 `query:"page"`
}

// HackerNewsStoriesNewParams contains the path and query parameters for hackerNewsStoriesNew.
type HackerNewsStoriesNewParams struct {
	Limit int32 `query:"limit"`
	Page  int32 `query:"page"`
}

// HackerNewsStoriesShowParams contains the path and query parameters for hackerNewsStoriesShow.
type HackerNewsStoriesShowParams struct {
	Limit int32 `query:"limit"`
	Page  int32 `query:"page"`
}

// HackerNewsStoriesTopParams contains the path and query parameters for hackerNewsStoriesTop.
type HackerNewsStoriesTopParams struct {
	Limit int32 `query:"limit"`
	Page  int32 `query:"page"`
}

// HackerNewsUserParams contains the path and query parameters for hackerNewsUser.
type HackerNewsUserParams struct {
	Handle string `path:"handle,required"`
}

// InstagramPostsParams contains the path and query parameters for instagramPosts.
type InstagramPostsParams struct {
	Username string `query:"username"`
	Limit    int32  `query:"limit"`
	Cursor   string `query:"cursor"`
}

// InstagramProfileParams contains the path and query parameters for instagramProfile.
type InstagramProfileParams struct {
	Username string `query:"username"`
}

// RedditPostByIDParams contains the path and query parameters for redditPostById.
type RedditPostByIDParams struct {
	ID      string `path:"id,required"`
	Context string `query:"context"`
}

// RedditPostByPermalinkParams contains the path and query parameters for redditPostByPermalink.
type RedditPostByPermalinkParams struct {
	Permalink string `query:"permalink"`
	Context   string `query:"context"`
}

// RedditSearchParams contains the path and query parameters for redditSearch.
type RedditSearchParams struct {
	Query      string `query:"query,required"`
	Sort       string `query:"sort"`
	Limit      int32  `query:"limit"`
	TimeFilter string `query:"time_filter"`
}

// RedditSubredditPostsParams contains the path and query parameters for redditSubredditPosts.
type RedditSubredditPostsParams struct {
	Subreddit  string `path:"subreddit,required"`
	Category   string `query:"category"`
	TimeFilter string `query:"time_filter"`
	Sort       string `query:"sort"`
	Limit      int32  `query:"limit"`
}

// RedditUserActivityParams contains the path and query parameters for redditUserActivity.
type RedditUserActivityParams struct {
	Username string `path:"username,required"`
	Limit    int32  `query:"limit"`
	Sort     string `query:"sort"`
}

// RedditUserPostsParams contains the path and query parameters for redditUserPosts.
type RedditUserPostsParams struct {
	Username   string `path:"username,required"`
	Category   string `query:"category"`
	TimeFilter string `query:"time_filter"`
	Sort       string `query:"sort"`
	Limit      int32  `query:"limit"`
}

// TiktokPostParams contains the path and query parameters for tiktokPost.
type TiktokPostParams struct {
	ID string `path:"id,required"`
}

// TiktokUserParams contains the path and query parameters for tiktokUser.
type TiktokUserParams struct {
	Handle string `path:"handle,required"`
}

// TiktokUserPostsParams contains the path and query parameters for tiktokUserPosts.
type TiktokUserPostsParams struct {
	Handle string `path:"handle,required"`
	Count  int32  `query:"count"`
	Cursor string `query:"cursor"`
}

// WebBrandParams contains the path and query parameters for webBrand.
type WebBrandParams struct {
	URL string `query:"url,required"`
}

// WebCrawlParams contains the path and query parameters for webCrawl.
type WebCrawlParams struct {
	URL      string `query:"url,required"`
	MaxPages int32  `query:"max_pages"`
	MaxDepth int32  `query:"max_depth"`
}

// WebMapParams contains the path and query parameters for webMap.
type WebMapParams struct {
	URL string `query:"url,required"`
}

// WebScrapeParams contains the path and query parameters for webScrape.
type WebScrapeParams struct {
	URL              string `query:"url,required"`
	IncludeSelectors string `query:"include_selectors"`
	ExcludeSelectors string `query:"exclude_selectors"`
}

// YoutubeChannelParams contains the path and query parameters for youtubeChannel.
type YoutubeChannelParams struct {
	ID       string `path:"id,required"`
	Language string `query:"language"`
}

// YoutubeCommentsParams contains the path and query parameters for youtubeComments.
type YoutubeCommentsParams struct {
	ID    string `path:"id,required"`
	Limit int32  `query:"limit"`
}

// YoutubePlaylistParams contains the path and query parameters for youtubePlaylist.
type YoutubePlaylistParams struct {
	ID       string `path:"id,required"`
	Language string `query:"language"`
}

// YoutubeRelatedParams contains the path and query parameters for youtubeRelated.
type YoutubeRelatedParams struct {
	ID       string `path:"id,required"`
	Language string `query:"language"`
}

// YoutubeSearchParams contains the path and query parameters for youtubeSearch.
type YoutubeSearchParams struct {
	Term     string `query:"term,required"`
	Language string `query:"language"`
	Limit    int32  `query:"limit"`
}

// YoutubeTranscriptParams contains the path and query parameters for youtubeTranscript.
type YoutubeTranscriptParams struct {
	ID       string `path:"id,required"`
	Language string `query:"language"`
}

// YoutubeVideoParams contains the path and query parameters for youtubeVideo.
type YoutubeVideoParams struct {
	ID       string `path:"id,required"`
	Language string `query:"language"`
}

// App calls appStoreApp.
func (s *AppStoreService) App(ctx context.Context, params AppStoreAppParams) (*Response, error) {
	return s.client.do(ctx, "appStoreApp", "/v1/appstore/app", params)
}

// Developer calls appStoreDeveloper.
func (s *AppStoreService) Developer(ctx context.Context, params AppStoreDeveloperParams) (*Response, error) {
	return s.client.do(ctx, "appStoreDeveloper", "/v1/appstore/developer", params)
}

// List calls appStoreList.
func (s *AppStoreService) List(ctx context.Context, params AppStoreListParams) (*Response, error) {
	return s.client.do(ctx, "appStoreList", "/v1/appstore/list", params)
}

// Privacy calls appStorePrivacy.
func (s *AppStoreService) Privacy(ctx context.Context, params AppStorePrivacyParams) (*Response, error) {
	return s.client.do(ctx, "appStorePrivacy", "/v1/appstore/privacy", params)
}

// Ratings calls appStoreRatings.
func (s *AppStoreService) Ratings(ctx context.Context, params AppStoreRatingsParams) (*Response, error) {
	return s.client.do(ctx, "appStoreRatings", "/v1/appstore/ratings", params)
}

// Reviews calls appStoreReviews.
func (s *AppStoreService) Reviews(ctx context.Context, params AppStoreReviewsParams) (*Response, error) {
	return s.client.do(ctx, "appStoreReviews", "/v1/appstore/reviews", params)
}

// Search calls appStoreSearch.
func (s *AppStoreService) Search(ctx context.Context, params AppStoreSearchParams) (*Response, error) {
	return s.client.do(ctx, "appStoreSearch", "/v1/appstore/search", params)
}

// Similar calls appStoreSimilar.
func (s *AppStoreService) Similar(ctx context.Context, params AppStoreSimilarParams) (*Response, error) {
	return s.client.do(ctx, "appStoreSimilar", "/v1/appstore/similar", params)
}

// Suggest calls appStoreSuggest.
func (s *AppStoreService) Suggest(ctx context.Context, params AppStoreSuggestParams) (*Response, error) {
	return s.client.do(ctx, "appStoreSuggest", "/v1/appstore/suggest", params)
}

// Fonts calls brandFonts.
func (s *BrandService) Fonts(ctx context.Context, params BrandFontsParams) (*Response, error) {
	return s.client.do(ctx, "brandFonts", "/v1/brand/fonts", params)
}

// Retrieve calls brandRetrieve.
func (s *BrandService) Retrieve(ctx context.Context, params BrandRetrieveParams) (*Response, error) {
	return s.client.do(ctx, "brandRetrieve", "/v1/brand/retrieve", params)
}

// Search calls brandSearch.
func (s *BrandService) Search(ctx context.Context, params BrandSearchParams) (*Response, error) {
	return s.client.do(ctx, "brandSearch", "/v1/brand/search", params)
}

// Styleguide calls brandStyleguide.
func (s *BrandService) Styleguide(ctx context.Context, params BrandStyleguideParams) (*Response, error) {
	return s.client.do(ctx, "brandStyleguide", "/v1/brand/styleguide", params)
}

// Alerts calls fomoAlerts.
func (s *FomoService) Alerts(ctx context.Context, params FomoAlertsParams) (*Response, error) {
	return s.client.do(ctx, "fomoAlerts", "/v1/fomo/alerts", params)
}

// Leaderboard calls fomoLeaderboard.
func (s *FomoService) Leaderboard(ctx context.Context, params FomoLeaderboardParams) (*Response, error) {
	return s.client.do(ctx, "fomoLeaderboard", "/v1/fomo/leaderboard/{window}", params)
}

// Notifications calls fomoNotifications.
func (s *FomoService) Notifications(ctx context.Context, params FomoNotificationsParams) (*Response, error) {
	return s.client.do(ctx, "fomoNotifications", "/v1/fomo/notifications", params)
}

// Search calls fomoSearch.
func (s *FomoService) Search(ctx context.Context, params FomoSearchParams) (*Response, error) {
	return s.client.do(ctx, "fomoSearch", "/v1/fomo/search", params)
}

// Thesis calls fomoThesis.
func (s *FomoService) Thesis(ctx context.Context, params FomoThesisParams) (*Response, error) {
	return s.client.do(ctx, "fomoThesis", "/v1/fomo/thesis", params)
}

// ThesisByToken calls fomoThesisByToken.
func (s *FomoService) ThesisByToken(ctx context.Context, params FomoThesisByTokenParams) (*Response, error) {
	return s.client.do(ctx, "fomoThesisByToken", "/v1/fomo/thesis/token/{mint}", params)
}

// ThesisByUser calls fomoThesisByUser.
func (s *FomoService) ThesisByUser(ctx context.Context, params FomoThesisByUserParams) (*Response, error) {
	return s.client.do(ctx, "fomoThesisByUser", "/v1/fomo/thesis/user/{id}", params)
}

// ThesisByUserToken calls fomoThesisByUserToken.
func (s *FomoService) ThesisByUserToken(ctx context.Context, params FomoThesisByUserTokenParams) (*Response, error) {
	return s.client.do(ctx, "fomoThesisByUserToken", "/v1/fomo/thesis/user/{id}/token/{address}", params)
}

// TokenHolders calls fomoTokenHolders.
func (s *FomoService) TokenHolders(ctx context.Context, params FomoTokenHoldersParams) (*Response, error) {
	return s.client.do(ctx, "fomoTokenHolders", "/v1/fomo/tokens/{address}/holders", params)
}

// TokensGraduated calls fomoTokensGraduated.
func (s *FomoService) TokensGraduated(ctx context.Context, params FomoTokensGraduatedParams) (*Response, error) {
	return s.client.do(ctx, "fomoTokensGraduated", "/v1/fomo/tokens/graduated", params)
}

// TokensMostHeld calls fomoTokensMostHeld.
func (s *FomoService) TokensMostHeld(ctx context.Context, params FomoTokensMostHeldParams) (*Response, error) {
	return s.client.do(ctx, "fomoTokensMostHeld", "/v1/fomo/tokens/most-held", params)
}

// TokensTrending calls fomoTokensTrending.
func (s *FomoService) TokensTrending(ctx context.Context, params FomoTokensTrendingParams) (*Response, error) {
	return s.client.do(ctx, "fomoTokensTrending", "/v1/fomo/tokens/trending", params)
}

// Trade calls fomoTrade.
func (s *FomoService) Trade(ctx context.Context, params FomoTradeParams) (*Response, error) {
	return s.client.do(ctx, "fomoTrade", "/v1/fomo/trades/{trade_id}", params)
}

// UserBalances calls fomoUserBalances.
func (s *FomoService) UserBalances(ctx context.Context, params FomoUserBalancesParams) (*Response, error) {
	return s.client.do(ctx, "fomoUserBalances", "/v1/fomo/users/{handle}/balances", params)
}

// UserProfile calls fomoUserProfile.
func (s *FomoService) UserProfile(ctx context.Context, params FomoUserProfileParams) (*Response, error) {
	return s.client.do(ctx, "fomoUserProfile", "/v1/fomo/users/{handle}", params)
}

// UserTrades calls fomoUserTrades.
func (s *FomoService) UserTrades(ctx context.Context, params FomoUserTradesParams) (*Response, error) {
	return s.client.do(ctx, "fomoUserTrades", "/v1/fomo/users/{handle}/trades", params)
}

// Search calls googleSearch.
func (s *GoogleService) Search(ctx context.Context, params GoogleSearchParams) (*Response, error) {
	return s.client.do(ctx, "googleSearch", "/v1/web/search", params)
}

// PlaceDetails calls googleMapsPlaceDetails.
func (s *GoogleMapsService) PlaceDetails(ctx context.Context, params GoogleMapsPlaceDetailsParams) (*Response, error) {
	return s.client.do(ctx, "googleMapsPlaceDetails", "/v1/googlemaps/details/{place_id}", params)
}

// PlaceReviews calls googleMapsPlaceReviews.
func (s *GoogleMapsService) PlaceReviews(ctx context.Context, params GoogleMapsPlaceReviewsParams) (*Response, error) {
	return s.client.do(ctx, "googleMapsPlaceReviews", "/v1/googlemaps/reviews/{place_id}", params)
}

// SearchPlaces calls googleMapsSearchPlaces.
func (s *GoogleMapsService) SearchPlaces(ctx context.Context, params GoogleMapsSearchPlacesParams) (*Response, error) {
	return s.client.do(ctx, "googleMapsSearchPlaces", "/v1/googlemaps/search", params)
}

// AppDetails calls googlePlayAppDetails.
func (s *GooglePlayService) AppDetails(ctx context.Context, params GooglePlayAppDetailsParams) (*Response, error) {
	return s.client.do(ctx, "googlePlayAppDetails", "/v1/googleplay/app_details/{id}", params)
}

// Availability calls googlePlayAvailability.
func (s *GooglePlayService) Availability(ctx context.Context, params GooglePlayAvailabilityParams) (*Response, error) {
	return s.client.do(ctx, "googlePlayAvailability", "/v1/googleplay/availability/{id}", params)
}

// Categories calls googlePlayCategories.
func (s *GooglePlayService) Categories(ctx context.Context, params GooglePlayCategoriesParams) (*Response, error) {
	return s.client.do(ctx, "googlePlayCategories", "/v1/googleplay/categories", params)
}

// CategoryApps calls googlePlayCategoryApps.
func (s *GooglePlayService) CategoryApps(ctx context.Context, params GooglePlayCategoryAppsParams) (*Response, error) {
	return s.client.do(ctx, "googlePlayCategoryApps", "/v1/googleplay/category_apps/{category}", params)
}

// DataSafety calls googlePlayDataSafety.
func (s *GooglePlayService) DataSafety(ctx context.Context, params GooglePlayDataSafetyParams) (*Response, error) {
	return s.client.do(ctx, "googlePlayDataSafety", "/v1/googleplay/data_safety/{id}", params)
}

// Developer calls googlePlayDeveloper.
func (s *GooglePlayService) Developer(ctx context.Context, params GooglePlayDeveloperParams) (*Response, error) {
	return s.client.do(ctx, "googlePlayDeveloper", "/v1/googleplay/developer", params)
}

// Permissions calls googlePlayPermissions.
func (s *GooglePlayService) Permissions(ctx context.Context, params GooglePlayPermissionsParams) (*Response, error) {
	return s.client.do(ctx, "googlePlayPermissions", "/v1/googleplay/permissions/{id}", params)
}

// Reviews calls googlePlayReviews.
func (s *GooglePlayService) Reviews(ctx context.Context, params GooglePlayReviewsParams) (*Response, error) {
	return s.client.do(ctx, "googlePlayReviews", "/v1/googleplay/reviews/{id}", params)
}

// Search calls googlePlaySearch.
func (s *GooglePlayService) Search(ctx context.Context, params GooglePlaySearchParams) (*Response, error) {
	return s.client.do(ctx, "googlePlaySearch", "/v1/googleplay/search", params)
}

// SimilarApps calls googlePlaySimilarApps.
func (s *GooglePlayService) SimilarApps(ctx context.Context, params GooglePlaySimilarAppsParams) (*Response, error) {
	return s.client.do(ctx, "googlePlaySimilarApps", "/v1/googleplay/similar_apps/{id}", params)
}

// Suggest calls googlePlaySuggest.
func (s *GooglePlayService) Suggest(ctx context.Context, params GooglePlaySuggestParams) (*Response, error) {
	return s.client.do(ctx, "googlePlaySuggest", "/v1/googleplay/suggest", params)
}

// ProductOffers calls googleShoppingProductOffers.
func (s *GoogleShoppingService) ProductOffers(ctx context.Context, params GoogleShoppingProductOffersParams) (*Response, error) {
	return s.client.do(ctx, "googleShoppingProductOffers", "/v1/googleshopping/product_offers", params)
}

// Search calls googleShoppingSearch.
func (s *GoogleShoppingService) Search(ctx context.Context, params GoogleShoppingSearchParams) (*Response, error) {
	return s.client.do(ctx, "googleShoppingSearch", "/v1/googleshopping/search", params)
}

// Item calls hackerNewsItem.
func (s *HackerNewsService) Item(ctx context.Context, params HackerNewsItemParams) (*Response, error) {
	return s.client.do(ctx, "hackerNewsItem", "/v1/hackernews/item/{id}", params)
}

// Search calls hackerNewsSearch.
func (s *HackerNewsService) Search(ctx context.Context, params HackerNewsSearchParams) (*Response, error) {
	return s.client.do(ctx, "hackerNewsSearch", "/v1/hackernews/search", params)
}

// StoriesAsk calls hackerNewsStoriesAsk.
func (s *HackerNewsService) StoriesAsk(ctx context.Context, params HackerNewsStoriesAskParams) (*Response, error) {
	return s.client.do(ctx, "hackerNewsStoriesAsk", "/v1/hackernews/stories_ask", params)
}

// StoriesBest calls hackerNewsStoriesBest.
func (s *HackerNewsService) StoriesBest(ctx context.Context, params HackerNewsStoriesBestParams) (*Response, error) {
	return s.client.do(ctx, "hackerNewsStoriesBest", "/v1/hackernews/stories_best", params)
}

// StoriesJob calls hackerNewsStoriesJob.
func (s *HackerNewsService) StoriesJob(ctx context.Context, params HackerNewsStoriesJobParams) (*Response, error) {
	return s.client.do(ctx, "hackerNewsStoriesJob", "/v1/hackernews/stories_job", params)
}

// StoriesNew calls hackerNewsStoriesNew.
func (s *HackerNewsService) StoriesNew(ctx context.Context, params HackerNewsStoriesNewParams) (*Response, error) {
	return s.client.do(ctx, "hackerNewsStoriesNew", "/v1/hackernews/stories_new", params)
}

// StoriesShow calls hackerNewsStoriesShow.
func (s *HackerNewsService) StoriesShow(ctx context.Context, params HackerNewsStoriesShowParams) (*Response, error) {
	return s.client.do(ctx, "hackerNewsStoriesShow", "/v1/hackernews/stories_show", params)
}

// StoriesTop calls hackerNewsStoriesTop.
func (s *HackerNewsService) StoriesTop(ctx context.Context, params HackerNewsStoriesTopParams) (*Response, error) {
	return s.client.do(ctx, "hackerNewsStoriesTop", "/v1/hackernews/stories_top", params)
}

// User calls hackerNewsUser.
func (s *HackerNewsService) User(ctx context.Context, params HackerNewsUserParams) (*Response, error) {
	return s.client.do(ctx, "hackerNewsUser", "/v1/hackernews/user/{handle}", params)
}

// Posts calls instagramPosts.
func (s *InstagramService) Posts(ctx context.Context, params InstagramPostsParams) (*Response, error) {
	return s.client.do(ctx, "instagramPosts", "/v1/instagram/posts", params)
}

// Profile calls instagramProfile.
func (s *InstagramService) Profile(ctx context.Context, params InstagramProfileParams) (*Response, error) {
	return s.client.do(ctx, "instagramProfile", "/v1/instagram/profile", params)
}

// PostByID calls redditPostById.
func (s *RedditService) PostByID(ctx context.Context, params RedditPostByIDParams) (*Response, error) {
	return s.client.do(ctx, "redditPostById", "/v1/reddit/post_by_id/{id}", params)
}

// PostByPermalink calls redditPostByPermalink.
func (s *RedditService) PostByPermalink(ctx context.Context, params RedditPostByPermalinkParams) (*Response, error) {
	return s.client.do(ctx, "redditPostByPermalink", "/v1/reddit/post_by_permalink", params)
}

// Search calls redditSearch.
func (s *RedditService) Search(ctx context.Context, params RedditSearchParams) (*Response, error) {
	return s.client.do(ctx, "redditSearch", "/v1/reddit/search_posts", params)
}

// SubredditPosts calls redditSubredditPosts.
func (s *RedditService) SubredditPosts(ctx context.Context, params RedditSubredditPostsParams) (*Response, error) {
	return s.client.do(ctx, "redditSubredditPosts", "/v1/reddit/subreddit_posts/{subreddit}", params)
}

// UserActivity calls redditUserActivity.
func (s *RedditService) UserActivity(ctx context.Context, params RedditUserActivityParams) (*Response, error) {
	return s.client.do(ctx, "redditUserActivity", "/v1/reddit/user_activity/{username}", params)
}

// UserPosts calls redditUserPosts.
func (s *RedditService) UserPosts(ctx context.Context, params RedditUserPostsParams) (*Response, error) {
	return s.client.do(ctx, "redditUserPosts", "/v1/reddit/user_posts/{username}", params)
}

// Post calls tiktokPost.
func (s *TiktokService) Post(ctx context.Context, params TiktokPostParams) (*Response, error) {
	return s.client.do(ctx, "tiktokPost", "/v1/tiktok/post/{id}", params)
}

// User calls tiktokUser.
func (s *TiktokService) User(ctx context.Context, params TiktokUserParams) (*Response, error) {
	return s.client.do(ctx, "tiktokUser", "/v1/tiktok/user/{handle}", params)
}

// UserPosts calls tiktokUserPosts.
func (s *TiktokService) UserPosts(ctx context.Context, params TiktokUserPostsParams) (*Response, error) {
	return s.client.do(ctx, "tiktokUserPosts", "/v1/tiktok/user_posts/{handle}", params)
}

// Brand calls webBrand.
func (s *WebService) Brand(ctx context.Context, params WebBrandParams) (*Response, error) {
	return s.client.do(ctx, "webBrand", "/v1/webcontext/brand", params)
}

// Crawl calls webCrawl.
func (s *WebService) Crawl(ctx context.Context, params WebCrawlParams) (*Response, error) {
	return s.client.do(ctx, "webCrawl", "/v1/webcontext/crawl", params)
}

// Map calls webMap.
func (s *WebService) Map(ctx context.Context, params WebMapParams) (*Response, error) {
	return s.client.do(ctx, "webMap", "/v1/webcontext/map", params)
}

// Scrape calls webScrape.
func (s *WebService) Scrape(ctx context.Context, params WebScrapeParams) (*Response, error) {
	return s.client.do(ctx, "webScrape", "/v1/webcontext/scrape", params)
}

// Search is the intentional web.search compatibility alias for googleSearch.
func (s *WebService) Search(ctx context.Context, params GoogleSearchParams) (*Response, error) {
	return s.client.do(ctx, "googleSearch", "/v1/web/search", params)
}

// Channel calls youtubeChannel.
func (s *YoutubeService) Channel(ctx context.Context, params YoutubeChannelParams) (*Response, error) {
	return s.client.do(ctx, "youtubeChannel", "/v1/youtube/channel/{id}", params)
}

// Comments calls youtubeComments.
func (s *YoutubeService) Comments(ctx context.Context, params YoutubeCommentsParams) (*Response, error) {
	return s.client.do(ctx, "youtubeComments", "/v1/youtube/comments/{id}", params)
}

// Playlist calls youtubePlaylist.
func (s *YoutubeService) Playlist(ctx context.Context, params YoutubePlaylistParams) (*Response, error) {
	return s.client.do(ctx, "youtubePlaylist", "/v1/youtube/playlist/{id}", params)
}

// Related calls youtubeRelated.
func (s *YoutubeService) Related(ctx context.Context, params YoutubeRelatedParams) (*Response, error) {
	return s.client.do(ctx, "youtubeRelated", "/v1/youtube/related/{id}", params)
}

// Search calls youtubeSearch.
func (s *YoutubeService) Search(ctx context.Context, params YoutubeSearchParams) (*Response, error) {
	return s.client.do(ctx, "youtubeSearch", "/v1/youtube/search", params)
}

// Transcript calls youtubeTranscript.
func (s *YoutubeService) Transcript(ctx context.Context, params YoutubeTranscriptParams) (*Response, error) {
	return s.client.do(ctx, "youtubeTranscript", "/v1/youtube/transcript/{id}", params)
}

// Video calls youtubeVideo.
func (s *YoutubeService) Video(ctx context.Context, params YoutubeVideoParams) (*Response, error) {
	return s.client.do(ctx, "youtubeVideo", "/v1/youtube/video/{id}", params)
}

// PublicOperationRegistry maps stable resource/action names to canonical operation IDs.
// web.search is an intentional compatibility alias for google.search.
var PublicOperationRegistry = map[string]map[string]string{
	"app_store": {
		"app":       "appStoreApp",
		"developer": "appStoreDeveloper",
		"list":      "appStoreList",
		"privacy":   "appStorePrivacy",
		"ratings":   "appStoreRatings",
		"reviews":   "appStoreReviews",
		"search":    "appStoreSearch",
		"similar":   "appStoreSimilar",
		"suggest":   "appStoreSuggest",
	},
	"brand": {
		"fonts":      "brandFonts",
		"retrieve":   "brandRetrieve",
		"search":     "brandSearch",
		"styleguide": "brandStyleguide",
	},
	"fomo": {
		"alerts":            "fomoAlerts",
		"leaderboard":       "fomoLeaderboard",
		"notifications":     "fomoNotifications",
		"search":            "fomoSearch",
		"thesis":            "fomoThesis",
		"thesisByToken":     "fomoThesisByToken",
		"thesisByUser":      "fomoThesisByUser",
		"thesisByUserToken": "fomoThesisByUserToken",
		"tokenHolders":      "fomoTokenHolders",
		"tokensGraduated":   "fomoTokensGraduated",
		"tokensMostHeld":    "fomoTokensMostHeld",
		"tokensTrending":    "fomoTokensTrending",
		"trade":             "fomoTrade",
		"userBalances":      "fomoUserBalances",
		"userProfile":       "fomoUserProfile",
		"userTrades":        "fomoUserTrades",
	},
	"google": {
		"search": "googleSearch",
	},
	"google_maps": {
		"placeDetails": "googleMapsPlaceDetails",
		"placeReviews": "googleMapsPlaceReviews",
		"searchPlaces": "googleMapsSearchPlaces",
	},
	"google_play": {
		"appDetails":   "googlePlayAppDetails",
		"availability": "googlePlayAvailability",
		"categories":   "googlePlayCategories",
		"categoryApps": "googlePlayCategoryApps",
		"dataSafety":   "googlePlayDataSafety",
		"developer":    "googlePlayDeveloper",
		"permissions":  "googlePlayPermissions",
		"reviews":      "googlePlayReviews",
		"search":       "googlePlaySearch",
		"similarApps":  "googlePlaySimilarApps",
		"suggest":      "googlePlaySuggest",
	},
	"google_shopping": {
		"productOffers": "googleShoppingProductOffers",
		"search":        "googleShoppingSearch",
	},
	"hacker_news": {
		"item":        "hackerNewsItem",
		"search":      "hackerNewsSearch",
		"storiesAsk":  "hackerNewsStoriesAsk",
		"storiesBest": "hackerNewsStoriesBest",
		"storiesJob":  "hackerNewsStoriesJob",
		"storiesNew":  "hackerNewsStoriesNew",
		"storiesShow": "hackerNewsStoriesShow",
		"storiesTop":  "hackerNewsStoriesTop",
		"user":        "hackerNewsUser",
	},
	"instagram": {
		"posts":   "instagramPosts",
		"profile": "instagramProfile",
	},
	"reddit": {
		"postByID":        "redditPostById",
		"postByPermalink": "redditPostByPermalink",
		"search":          "redditSearch",
		"subredditPosts":  "redditSubredditPosts",
		"userActivity":    "redditUserActivity",
		"userPosts":       "redditUserPosts",
	},
	"tiktok": {
		"post":      "tiktokPost",
		"user":      "tiktokUser",
		"userPosts": "tiktokUserPosts",
	},
	"web": {
		"brand":  "webBrand",
		"crawl":  "webCrawl",
		"map":    "webMap",
		"scrape": "webScrape",
		"search": "googleSearch",
	},
	"youtube": {
		"channel":    "youtubeChannel",
		"comments":   "youtubeComments",
		"playlist":   "youtubePlaylist",
		"related":    "youtubeRelated",
		"search":     "youtubeSearch",
		"transcript": "youtubeTranscript",
		"video":      "youtubeVideo",
	},
}
