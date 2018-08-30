<template>
  <div>
    <ul id="list">
      <li v-for="item in items" :key="item.index" v-on:click="$emit('onBlogSelect', item)">
        {{ item.title }}
      </li>
    </ul>
  </div>
</template>
<script>
export default {
  data () {
    return {
      items: []
      // items: [{aid: '201801', title: '一篇博客的标题'}, {aid: '201802', title: '另一篇博客的标题'}]
    }
  },
  created () {
    this.$http.get('/api/getall') // for build
    // this.$http.get('/test/api/getall') // for dev test
      .then((data) => {
        const articles = data.body.articles
        console.log(articles) // ----------------------------------- console log -----------------------------------
        articles.forEach(element => {
          // todo :
          // use data update this.items
          this.items.push({aid: element.AID, title: element.Title, class: element.Class, interview: element.Content})
        })
      })
  }
}
</script>
<style>
#list {
  display: flex;
  flex-direction: column;
  margin: 0;
  padding: 0;
  background-color: dimgray;
}

#list li {
  font-family: "Arial","Microsoft YaHei","黑体","宋体",sans-serif;
  margin: 10px;
  color:white;
  font-size: 2em;
  cursor: pointer;
  /* font-weight: 550; */
  display: flex;
  align-items: center;
}
</style>
