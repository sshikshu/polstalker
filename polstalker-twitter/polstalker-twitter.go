package main

import (
	"flag"
	"fmt"
	"net/url"
	"strconv"

	"bitbucket.org/prvn30/polstalker/Color"
	"bitbucket.org/prvn30/polstalker/SortedMap"

	"github.com/ChimeraCoder/anaconda"
)

var apiKey, apiSecret, accessToken, accessSecret string = "bv1qvBQoFGjgJnZjwOE71Mswz", "IS6hYepcGtCGDsdso7pEKqZ3F4VX2cx6UXeGSasvHGbivSWdpz", "3278398670-llhQyRUhqKm8A2HwoQ2YF7uB2dSFgX5EEOa8OZy", "i1yHnolKY10dEvFaz3VZjCitgN2Y5Jq58fSYkmAkoXe2R"

type Tweet struct {
	Author    string
	Date      string
	Country   string
	Place     string
	PlaceType string
	Location  [][]float64
	Locatable bool
	Hashtags  []string
	Media     []string
	Mentions  []string
	Text      string
	Uri       []string
}

func printSummary(user anaconda.User, color bool) {
	fmt.Print(" ", user.Name)
	fmt.Printf("(%s)", Color.Blue(user.ScreenName, color))
	fmt.Print("     id:", Color.Blue(user.IdStr, color))
	fmt.Print("\n")
	fmt.Print(" ", user.Name)
	if user.GeoEnabled {
		fmt.Printf(" has geolocation %s\n", Color.Green("enabled", color))
	} else {
		fmt.Printf(" has geolocation %s\n", Color.Magenta("disabled", color))
	}
	locale := ""
	location := user.Location
	if "" != location {
		locale += "Location:  " + Color.Green(location, color) + ";     "
	}
	timezone := user.TimeZone
	if "" != timezone {
		locale += "Timezone: " + Color.Green(timezone, color) + ";     "
	}
	language := user.Lang
	if "" != language {
		locale += "Language: " + Color.Green(language, color) + ";"
	}
	if "" != locale {
		fmt.Printf(" %s\n", locale)
	}
	fmt.Printf(" Created: %s\n", Color.Green(user.CreatedAt, color))
	fmt.Printf(" URL: %s\n", Color.Yellow(user.URL, color))
	fmt.Printf(" Profile Image: %s\n", Color.Yellow(user.ProfileImageURL, color))
	fmt.Printf(" %v favorites;     %v friends;     %v followers;\n", user.FriendsCount, user.FavouritesCount, user.FollowersCount)
	fmt.Printf(" Member of %v public lists\n", user.ListedCount)
	numberOfTweets := user.StatusesCount
	fmt.Printf(" Number of Tweets: %v\n", numberOfTweets)
	if numberOfTweets > 0 {
		fmt.Printf(" Most Recent Tweet: %s\n", user.Status.Text)
	}
}

func anacondaTweetsToTweets(tweets []anaconda.Tweet) []Tweet {
	tweetList := make([]Tweet, 0, len(tweets))
	for _, tweet := range tweets {
		tweetStruct := Tweet{
			Author:   tweet.User.ScreenName,
			Date:     tweet.CreatedAt,
			Mentions: make([]string, 0, len(tweet.Entities.User_mentions)),
			Hashtags: make([]string, 0, len(tweet.Entities.Hashtags)),
			Media:    make([]string, 0, len(tweet.Entities.Media)),
			Text:     tweet.Text,
			Uri:      make([]string, 0, len(tweet.Entities.Urls))}
		tweetStruct.Locatable = (0 != len(tweet.Place.BoundingBox.Coordinates))
		if tweetStruct.Locatable {
			tweetStruct.Location = make([][]float64, len(tweet.Place.BoundingBox.Coordinates[0]))
			for i, coordinates := range tweet.Place.BoundingBox.Coordinates[0] {
				tweetStruct.Location[i] = make([]float64, 2)
				tweetStruct.Location[i][0] = coordinates[0]
				tweetStruct.Location[i][1] = coordinates[1]
			}
			tweetStruct.Country = tweet.Place.Country
			tweetStruct.Place = tweet.Place.FullName
			tweetStruct.PlaceType = tweet.Place.PlaceType
		}
		for _, userMention := range tweet.Entities.User_mentions {
			tweetStruct.Mentions = append(tweetStruct.Mentions, userMention.Screen_name)
		}
		for _, hashtag := range tweet.Entities.Hashtags {
			tweetStruct.Hashtags = append(tweetStruct.Hashtags, hashtag.Text)
		}
		for _, medium := range tweet.Entities.Media {
			tweetStruct.Media = append(tweetStruct.Media, medium.Expanded_url)
		}
		for _, uri := range tweet.Entities.Urls {
			tweetStruct.Uri = append(tweetStruct.Uri, uri.Expanded_url)
		}

		tweetList = append(tweetList, tweetStruct)
	}
	return tweetList
}

func getHashtagsFrequency(tweets []Tweet) map[string]int {
	hashtags := make(map[string]int)
	for _, tweet := range tweets {
		for _, hashtag := range tweet.Hashtags {
			hashtags[hashtag] += 1
		}
	}
	return hashtags
}

func getMediaFrequency(tweets []Tweet) map[string]int {
	uriList := make(map[string]int)
	for _, tweet := range tweets {
		for _, uri := range tweet.Media {
			uriList[uri] += 1
		}
	}
	return uriList
}

func getMentionsFrequency(tweets []Tweet) map[string]int {
	mentions := make(map[string]int)
	for _, tweet := range tweets {
		for _, user := range tweet.Mentions {
			mentions[user] += 1
		}
	}
	return mentions
}

func getUriFrequency(tweets []Tweet) map[string]int {
	uriList := make(map[string]int)
	for _, tweet := range tweets {
		for _, uri := range tweet.Uri {
			uriList[uri] += 1
		}
	}
	return uriList
}

func getGeolocationFreq(tweets []Tweet) map[string]int {
	locationList := make(map[string]int)
	for _, tweet := range tweets {
		if tweet.Locatable {
			locationList[fmt.Sprintf("(%s)%s,%s", tweet.PlaceType, tweet.Place, tweet.Country)] += 1
		}
	}
	return locationList
}

func printGeolocationByDate(tweets []Tweet, color bool) {
	for _, tweet := range tweets {
		if tweet.Locatable {
			fmt.Printf(" %s  (%s)%s,%s within %v\n", tweet.Date, Color.Cyan(tweet.PlaceType, color), Color.Green(tweet.Place, color), Color.Green(tweet.Country, color), tweet.Location)
		}
	}
}

func printHashtagsByDate(tweets []Tweet, color bool) {
	for _, tweet := range tweets {
		for _, hashtag := range tweet.Hashtags {
			fmt.Println("", tweet.Date, Color.Green(hashtag, color))
		}
	}
}

func printMediaByDate(tweets []Tweet, color bool) {
	for _, tweet := range tweets {
		for _, media := range tweet.Media {
			fmt.Println("", tweet.Date, Color.Green(media, color))
		}
	}
}

func printUriByDate(tweets []Tweet, color bool) {
	for _, tweet := range tweets {
		for _, uri := range tweet.Uri {
			fmt.Println("", tweet.Date, Color.Green(uri, color))
		}
	}
}

func printMentionsByDate(tweets []Tweet, color bool) {
	for _, tweet := range tweets {
		for _, mention := range tweet.Mentions {
			fmt.Println("", tweet.Date, Color.Green(mention, color))
		}
	}
}

func printEntitiesByFrequency(entities map[string]int, prefix string, color bool) {
	for _, key := range SortedMap.SortedKeys(entities) {
		fmt.Printf(" %-9v  %s\n", entities[key], Color.Green(prefix+key, color))
	}
}

func main() {
	namePtr := flag.String("name", "prvn_30", "target's twitter handle")
	count := flag.Int("count", 100, "number of tweets to analyze")
	outputSummary := flag.Bool("summary", true, "summary of the target's account")
	outputGeolocationDate := flag.Bool("loc", false, "target's locations arranged by date")
	outputGeolocationFreq := flag.Bool("loc-freq", false, "target's locations arranged by frequency")
	outputHashtagFreq := flag.Bool("hash-freq", false, "targets's hashtags arranged by frequency")
	outputHashtagDate := flag.Bool("hash", false, "targets's hashtags arranged by date")
	outputMediaFreq := flag.Bool("media-freq", false, "targets's media arranged by frequency")
	outputMediaDate := flag.Bool("media", false, "targets's media arranged by date")
	outputMentionFreq := flag.Bool("mention-freq", false, "targets's user mentions arranged by frequency")
	outputMentionDate := flag.Bool("mention", false, "targets's user mentions arranged by date")
	outputUriFreq := flag.Bool("url-freq", false, "targets's urls arranged by frequency")
	outputUriDate := flag.Bool("url", false, "targets's urls arranged by date")
	outputNoColor := flag.Bool("nocolor", false, "plain output")
	flag.Parse()

	color := !*outputNoColor

	anaconda.SetConsumerKey(apiKey)
	anaconda.SetConsumerSecret(apiSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)


	if *outputSummary {
		user, errGetUser := api.GetUsersShow(*namePtr, nil)
		if errGetUser != nil {
			fmt.Println(errGetUser)
		} else {
			printSummary(user, color)
		}
	}

	if *outputGeolocationFreq || *outputGeolocationDate || *outputHashtagFreq || *outputHashtagDate || *outputMediaFreq || *outputMediaDate || *outputMentionFreq || *outputMentionDate || *outputUriFreq || *outputUriDate {
		v := url.Values{}
		v.Set("screen_name", *namePtr)
		v.Set("count", strconv.Itoa(*count))
		tweetsUnparsed, errTweets := api.GetUserTimeline(v)

		if nil != errTweets {
			fmt.Println(errTweets)
		} else {
			tweets := anacondaTweetsToTweets(tweetsUnparsed)
			if *outputGeolocationDate {
				fmt.Println("\n", Color.Yellow("Date                            Location", color))
				printGeolocationByDate(tweets, color)
			}
			if *outputGeolocationFreq {
				fmt.Println("\n", Color.Yellow("Frequency  Location", color))
				locations := getGeolocationFreq(tweets)
				printEntitiesByFrequency(locations, "", color)
			}
			if *outputHashtagDate {
				fmt.Println("\n", Color.Yellow("Date                           Hashtag", color))
				printHashtagsByDate(tweets, color)
			}
			if *outputHashtagFreq {
				fmt.Println("\n", Color.Yellow("Frequency  Hashtag", color))
				hashtags := getHashtagsFrequency(tweets)
				printEntitiesByFrequency(hashtags, "#", color)
			}
			if *outputMediaDate {
				fmt.Println("\n", Color.Yellow("Date                           Media", color))
				printMediaByDate(tweets, color)
			}
			if *outputMediaFreq {
				fmt.Println("\n", Color.Yellow("Frequency Media", color))
				media := getMediaFrequency(tweets)
				printEntitiesByFrequency(media, "", color)
			}
			if *outputMentionDate {
				fmt.Println("\n", Color.Yellow("Date                           Users Mentioned", color))
				printMentionsByDate(tweets, color)
			}
			if *outputMentionFreq {
				fmt.Println("\n", Color.Yellow("Frequency  Users Mentioned", color))
				mentions := getMentionsFrequency(tweets)
				printEntitiesByFrequency(mentions, "@", color)
			}
			if *outputUriDate {
				fmt.Println("\n", Color.Yellow("Date                           URL", color))
				printMediaByDate(tweets, color)
			}
			if *outputUriFreq {
				fmt.Println("\n", Color.Yellow("Frequency  URL", color))
				uri := getUriFrequency(tweets)
				printEntitiesByFrequency(uri, "", color)
			}
		}
	}
}
