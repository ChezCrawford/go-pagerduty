package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type WebhookSubscriptionList struct {
	Meta
}

func WebhookSubscriptionListCommand() (cli.Command, error) {
	return &WebhookSubscriptionList{}, nil
}

func (c *WebhookSubscriptionList) Help() string {
	helpText := `
	webhook-subscription list

	Options:

	` + c.Meta.Help()
	return strings.TrimSpace(helpText)
}

func (c *WebhookSubscriptionList) Run(args []string) int {
	flags := c.Meta.FlagSet("webhook-subscription list")
	flags.Usage = func() { fmt.Println(c.Help()) }

	if err := flags.Parse(args); err != nil {
		log.Errorln(err)
		return -1
	}

	client := c.Meta.Client()
	data, err := client.ListWebhookSubscriptionsWithContext(context.Background())

	if err != nil {
		log.Error(err)
		return -1
	}

	for i, webhookSubscription := range data.WebhookSubscriptions {
		fmt.Println("Entry: ", i+1)
		data, err := yaml.Marshal(webhookSubscription)
		if err != nil {
			log.Error(err)
			return -1
		}
		fmt.Println(string(data))
	}

	return 0
}

func (c *WebhookSubscriptionList) Synopsis() string {
	return "List webhook subscriptions"
}
