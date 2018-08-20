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
// new Vue (参数)
export default {
  data () {
    return {
      items: [{aid: '201801', title: '一篇博客的标题'}, {aid: '201802', title: '另一篇博客的标题'}]
    }
  },
  created () {
    this.$http.get('/test/api/')
      .then((data) => {
        // const movies = data.body.data.movies
        // movies.forEach((item) => {
          // use reduce: array to string
          // item.actors = item.actors.reduce((acc, item) => acc + ' ' + item)
          // item.movieId = item.movie_id
        // })
        // console.log(movies)
        // this.movielist = movies
      })
    // todo: 向后台异步获取博客数据存入items，并使用nextTick进行处理。
    // console.log('list created')
  }
  // methods: {
  //   jumpToBlog (item) {
  //     console.log(item.aid)
  //   }
  // }
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
