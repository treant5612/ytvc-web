package youtubeapi

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var ServiceFSC *youtube.Service

func InitServiceFSC(clientSecretFile, tokenFile string) (err error) {
	client := getClient(youtube.YoutubeForceSslScope, clientSecretFile, tokenFile)
	ServiceFSC, err = youtube.NewService(context.TODO(), option.WithHTTPClient(client))
	return err
}
