{
  mappings: {
    logs: {
      properties: {
        'buildId': {
          type: "long",
        },
        'key': {
          type: "keyword",
        },
        'timestamp': {
          type: "date",
          format: "epoch_millis||yyyy-MM-dd HH:mm:ss||yyyy-MM-dd HH:mm:ss.SSS||strict_date_hour_minute_second||strict_date_hour_minute_second_fractioN",
        },
        'type': {
          type: "keyword",
        },
        contentType: {
          type: "keyword",
          index: false,
        },
        contentTypeParameter: {
          type: "keyword",
          index: false,
        },
        size: {
          type: "long",
          index: false,
        },

        // field to store the payload of application/text
        text: {
          type: "text",
          index: false,
        },

        // field to store the payload of application/json
        json: {
          type: "object",
          enabled: false,
          dynamic: false,
        },

        // field to store the payload of other types
        blob: {
          type: "binary",
          doc_values: false,
          store: false,
        },
      },
    },
  },
}
