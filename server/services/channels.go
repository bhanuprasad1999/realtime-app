package services

import "fmt"

type Channels struct {
	ID          string   `json:"id"`
	ChannelName string   `json:"channel_name"`
	Usernames   []string `json:"usernames"`
}

func EditChannel(channel_name string, username string) {

	channel := GetChannelDetails(channel_name)

	fmt.Println(channel)
	if channel.ChannelName == channel_name {
		DeleteChannel(channel_name)
		channel.Usernames = append(channel.Usernames, username)
	} else {
		channel = Channels{
			ID:          channel_name,
			ChannelName: channel_name,
			Usernames:   []string{username},
		}
	}
	var all_channels []Channels
	readJSONFile(&all_channels, "channels")
	all_channels = append(all_channels, channel)
	writeJSONFile(all_channels, "channels")
}

func GetChannelDetails(channel_name string) Channels {
	var all_channels []Channels
	readJSONFile(&all_channels, "channels")
	for _, channel := range all_channels {
		if channel.ChannelName == channel_name {
			return channel
		}
	}
	return Channels{}
}

func DeleteChannel(channel_name string) {
	fmt.Println("Delete channel")
	var all_channels []Channels
	readJSONFile(&all_channels, "channels")

	var updatedChannels []Channels
	for _, channel := range all_channels {
		if channel.ChannelName != channel_name {
			updatedChannels = append(updatedChannels, channel)
		}
	}
	writeJSONFile(updatedChannels, "channels")
}

func RemoveUserFromChannel(channel_name string, username string) {

}
