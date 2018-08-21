FORMAT: 1A
# Snap's Blog API
this is snap's personal blog. [weirdsnap.top]

# Group Articles
Resources related to articles in the API.

## Articles Collection [/api]

### List All Articles's info [GET]

/api/getall

+ Response 200 (application/json)
[{
  body: {
    data: {
      articles: [{
        aid: 2018001,
        title: "title1",
        content: "longlongago..."
      },{
        aid: 2018002,
        title: "title2",
        content: "longlongago..."
      }]
    }
  }
}]


### Get One Article By Id [GET]

/api/getbyid/?id=2018001

+ Response 200 (application/json)
[{
  body: {
    data: {
      articles: {
        aid: 2018001,
        title: "title1",
        content: "longlongago..."
      }
    }
  }
}]