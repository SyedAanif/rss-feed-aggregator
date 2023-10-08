package main

import (
	"time"

	"github.com/SyedAanif/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct{
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string 	`json:"name"`
	APIKey 	  string	`json:"api_key"`		
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string 	`json:"name"`
	Url       string	`json:"url"`
	UserID    uuid.UUID	`json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID	`json:"user_id"`
	FeedID    uuid.UUID	`json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string 	  `json:"title"`
	Description *string   `json:"description"` // pointer to string marshalls to exact value if it is null
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

// DTO converter
func databaseUserToUser(dbUser database.User) User{
	return User{
		ID: 		dbUser.ID,
		CreatedAt: 	dbUser.CreatedAt,
		UpdatedAt: 	dbUser.UpdatedAt,
		Name: 		dbUser.Name,
		APIKey: 	dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed{
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed{
	feeds :=  []Feed{}
	for _, dbFeed := range dbFeeds{
		feeds = append(feeds,databaseFeedToFeed(dbFeed))
	}
	return feeds
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow{
	return FeedFollow{
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
	}
}

func databaseFeedsFollowToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow{
	feedFollows := []FeedFollow{}

	for _, dbFeedFollow := range dbFeedFollows{
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}

func databasePostToPost(dbPost database.Post) Post{
	var description *string // pointer string
	if dbPost.Description.Valid { // not null
		description = &dbPost.Description.String // address value
	}
	return Post{
		ID: dbPost.ID,
		CreatedAt: dbPost.CreatedAt,
		UpdatedAt: dbPost.UpdatedAt,
		Title: dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url: dbPost.Url,
		FeedID: dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post{
	posts := []Post{}

	for _, post := range dbPosts{
		posts = append(posts,databasePostToPost(post))
	}
	return posts
}