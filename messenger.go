package messenger

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/feedcast-io/messenger/types"
	"strconv"
)

type Service struct {
	cli *pubsub.Client
}

func NewService(projectName string) *Service {
	cli, _ := pubsub.NewClient(context.Background(), projectName)

	return &Service{
		cli: cli,
	}
}

func (s *Service) SendEmail(emailTo, language string, template types.EmailTemplate, data map[string]any) error {
	body, _ := json.Marshal(map[string]any{
		"to":       emailTo,
		"template": template,
		"language": language,
		"data":     data,
	})

	_, err := s.cli.Topic("mailer").Publish(context.Background(), &pubsub.Message{
		Data: body,
	}).Get(context.Background())

	return err
}

func (s *Service) SendShoppingImport(feedId int32) error {
	_, err := s.cli.Topic("export-comparator").Publish(context.Background(), &pubsub.Message{
		Data: []byte(strconv.Itoa(int(feedId))),
	}).Get(context.Background())

	return err
}

func (s *Service) SendFeedSynchro(feedId int32) error {
	_, err := s.cli.Topic("feed-synchro").Publish(context.Background(), &pubsub.Message{
		Data: []byte(strconv.Itoa(int(feedId))),
	}).Get(context.Background())

	return err
}

func (s *Service) SendFeedGeneration(feedId int32, hash string) error {
	data, _ := json.Marshal(map[string]any{
		"feed": feedId,
		"hash": hash,
	})

	_, err := s.cli.Topic("feed-generation").Publish(context.Background(), &pubsub.Message{
		Data: data,
	}).Get(context.Background())

	return err
}
