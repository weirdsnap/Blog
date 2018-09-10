<template>
  <div>
    <div id="c" v-html="body" class="markdown-body">
    </div>
  </div>
</template>
<script>
import marked from 'marked'
export default {
  data: function () {
    return {
      aid: 123,
      title: '123',
      body: '321'
    }
  },
  created: function () {
    const aId = this.$route.params.id

    // [ajax] use aId send req to get the md
    let xmlhttp
    let that = this
    if (window.XMLHttpRequest) {
      //  IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
      xmlhttp = new XMLHttpRequest()
    }
    // 绑定事件处理函数
    xmlhttp.onreadystatechange = function () {
      if (xmlhttp.readyState === 4 && xmlhttp.status === 200) {
        let obj = JSON.parse(xmlhttp.responseText)
        let str = obj.content
        var rendererMD = new marked.Renderer()
        marked.setOptions({
          renderer: rendererMD,
          gfm: true,
          tables: true,
          breaks: false,
          pedantic: false,
          sanitize: false,
          smartLists: true,
          smartypants: false
        })
        // todo: add highlight functiion
        that.body = marked(str)
      }
    }
    xmlhttp.open('GET', '/test/api/getbyid/?id=' + aId, true) // when test
    // xmlhttp.open('GET', '/api/getbyid/?id=' + aId, true) // when build
    xmlhttp.send()
  }
}
</script>
<style>
h1,h2,h3,h4,h5,p {
  margin:0;
}
#c {
  overflow: auto;
  padding-left:10%;
  padding-right: 10%;
  margin: auto;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  background-color: bisque;
  margin-bottom: 5%;
  text-align: start;
}

</style>
