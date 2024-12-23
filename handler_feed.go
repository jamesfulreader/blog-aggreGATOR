package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jamesfulreader/gator/internal/database"
)

func handlerFetch(s *state, cmd command) error {
	// this will need to be changed later on to fetch any address
	// for now hard coding https://www.wagslane.dev/index.xml
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s <name> (no arguments needed)", cmd.Name)
	}

	fetchRSS, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching RSS feed: %s", err)
	}

	fmt.Println(fetchRSS)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := s.cfg.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
