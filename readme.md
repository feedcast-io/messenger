# Feedcast Messenger

Feedcast Messenger allow communication between microservices. It relies on [Google PubSub](https://cloud.google.com/pubsub/docs/overview?hl=fr).

## Messages queues

| Queue name         | Message format                                                | Description                                                                                               |
|--------------------|---------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| export-comparator  | `int32`                                                       | Launch feed export to [feedcast.shopping](https://www.feedcast.shopping) comparator based on feed ID      |
| feed-synchro | `int32`                                                       | Launch feed synchro. Check on Facebook & Merchant Center if a product feed is created and if url is valid |
| feed-generation | `{feed: int32, hash: string}`                             | Lauch feed XML generation                                                                                 |
| mailer | `{to: string, template: string, language: string, data: any}` | Send email                                                                                                |

