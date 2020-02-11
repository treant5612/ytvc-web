package youtube

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var ServiceFSC *youtube.Service

func InitServiceFSC() (err error) {

	client := getClient(youtube.YoutubeForceSslScope, "client_secret.json", "ForceSslScopeToken.json")
	ServiceFSC, err = youtube.NewService(context.TODO(), option.WithHTTPClient(client))
	return err
}
