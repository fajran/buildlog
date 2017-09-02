// POST to elasticsearch:9200/sequence/build/1/_update?_source=iid

local increment = 1;
local initialValue = 1;

{
  lang: "groovy",
  script: "ctx._source.iid += " + increment,
  upsert: {
    iid: initialValue,
  },
}

