// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package gencontroller

import (
	"errors"

	"github.com/mattermost/mattermost-load-test-ng/defaults"
)

// Config holds information about the data to be generated by the
// GenController.
type Config struct {
	// The target number of teams to be created.
	NumTeams int64 `default:"2" validate:"range:[0,]"`
	// The target number of direct messages to be created.
	NumChannelsDM int64 `default:"20" validate:"range:[0,]"`
	// The target number of group messages to be created.
	NumChannelsGM int64 `default:"20" validate:"range:[0,]"`
	// The target number of private channels to be created.
	NumChannelsPrivate int64 `default:"20" validate:"range:[0,]"`
	// The target number of public channels to be created.
	NumChannelsPublic int64 `default:"20" validate:"range:[0,]"`
	// The target number of posts to be created.
	NumPosts int64 `default:"1000" validate:"range:[0,]"`
	// The target number of reactions to be created.
	NumReactions int64 `default:"200" validate:"range:[0,]"`
	// The target number of post reminders to be created.
	NumPostReminders int64 `default:"200" validate:"range:[0,]"`
	// The target number of sidebar categories to be created.
	NumSidebarCategories int64 `default:"10" validate:"range:[0,]"`
	// The target number of threads to follow.
	NumFollowedThreads int64 `default:"10" validate:"range:[0,]"`
	// The percentage of replies to be created.
	PercentReplies float64 `default:"0.5" validate:"range:[0,1]"`
	// The percentage of replies that should be in long threads.
	PercentRepliesInLongThreads float64 `default:"0.05" validate:"range:[0,1]"`
	// The percentage of post that are marked as urgent.
	PercentUrgentPosts float64 `default:"0.001" validate:"range:[0,1]"`

	// Indicates the distribution of chanel members within channels.
	ChannelMembersDistribution []ChannelMemberDistribution
}

// ChannelMemberDistribution holds the member limit and what percentage of channels will have the limit.
// Probability indicates how frequently to choose a range. PercentChannels indicates the
// what percent of channels will contain the change.
//
// For example, a probability of 0.8 and 0.5 of percentChannels with an array of
// length 10 would mean that indices [0,5) would be chosen at a 80% probability chance.
// And then the Memberlimit indicates the max member count for that range.
type ChannelMemberDistribution struct {
	MemberLimit     int64
	PercentChannels float64 `validate:"range:[0,1]"`
	Probability     float64 `validate:"range:[0,1]"`
}

// ReadConfig reads the configuration file from the given string. If the string
// is empty, it will return a config with default values.
func ReadConfig(configFilePath string) (*Config, error) {
	var cfg Config

	if err := defaults.ReadFromJSON(configFilePath, "./config/gencontroller.json", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// IsValid reports whether a given gencontroller.Config is valid or not.
// Returns an error if the validation fails.
func (c *Config) IsValid() error {
	totalPercent := 0.0
	for _, item := range c.ChannelMembersDistribution {
		totalPercent += item.Probability
	}

	if totalPercent != 1.0 {
		return errors.New("sum of probabilities for channel distribution should be equal to 1")
	}

	totalPercent = 0.0
	for _, item := range c.ChannelMembersDistribution {
		totalPercent += item.PercentChannels
	}

	if totalPercent != 1.0 {
		return errors.New("sum of percentages for channel member distribution should be equal to 1")
	}

	return nil
}

func (c *Config) NumTotalChannels() int64 {
	return c.NumChannelsDM + c.NumChannelsGM + c.NumChannelsPrivate + c.NumChannelsPublic
}
