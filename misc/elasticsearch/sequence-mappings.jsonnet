// Source: http://www.lovelysystems.com/blog/es-iid-generator/

{
  settings: {
    number_of_shards: 1,
    auto_expand_replicas: "0-all",
  },
  mappings: {
    build: {
      _all: {
        enabled: false,
      },
      dynamic: "strict",
      properties: {
        iid: {
          type: "keyword",
          index: false,
        }
      },
    },
  },
}

