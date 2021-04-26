package main

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type WebhookSubscriptionShow struct {
	Meta
}

func WebhookSubscriptionShowCommand() (cli.Command, error) {
	return &WebhookSubscriptionShow{}, nil
}

func (c *WebhookSubscriptionShow) Help() string {
	helpText := `
	webhook-subscription show

	Options:

		 -id

	` + c.Meta.Help()
	return strings.TrimSpace(helpText)
}

func (c *WebhookSubscriptionShow) Run(args []string) int {
	var id string = "P3SJPP8"

	flags := c.Meta.FlagSet("webhook-subscription show")
	flags.Usage = func() { fmt.Println(c.Help()) }

	if err := flags.Parse(args); err != nil {
		log.Errorln(err)
		return -1
	}

	client := c.Meta.Client()
	webhook_subscription, err := client.GetWebhookSubscription(id)

	if err != nil {
		log.Error(err)
		return -1
	}

	data, err := yaml.Marshal(webhook_subscription)
	if err != nil {
		log.Error(err)
		return -1
	}
	fmt.Println(string(data))
	return 0
}

func (c *WebhookSubscriptionShow) Synopsis() string {
	return "Get details about an existing webhook subscription"
}
