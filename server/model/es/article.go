package es

const (
	ArticleIndex = "article"
)

const (
	ArticleMapping = `{
	"mappings":{
		"dynamic": "false",
		"properties":{
			"id": 				{ "type": "keyword" },
			"title": 			{ "type": "text" },
			"description":		{ "type": "text" },
			"content":			{ "type": "text" },
			"author":			{ "type": "text" },
			"type":				{ "type": "keyword" },
			"created_at":		{ "type": "date" },
			"updated_at":		{ "type": "date" }
		}
	},
	"settings" : {
      "index" : {
        "number_of_shards" : "1",
        "number_of_replicas" : "1"
      }
    }
}`
)
