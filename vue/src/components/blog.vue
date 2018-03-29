<template>
  <div>
    <!-- todo: turn the aid to the article  {{message}}  -->
    <div> {{title}} </div>
    <div id="c" v-html="body">
    </div>
  </div>
</template>
<script>
import marked from 'marked'
export default {
  props: ['message'],
  data: function () {
    return {
      title: '123',
      body: '321'
    }
  },
  created: function (message) {
    // todo : ajax / use message send req to get the md
    let xmlhttp
    let that = this
    if (window.XMLHttpRequest) {
      //  IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
      xmlhttp = new XMLHttpRequest()
    }
    // 绑定事件处理函数
    xmlhttp.onreadystatechange = function () {
      if (xmlhttp.readyState === 4 && xmlhttp.status === 200) {
        alert(xmlhttp.responseText)
        let obj = JSON.parse(xmlhttp.responseText)
        let str = obj.content
        // let str = marked('# asd\n## 1233412')
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
        that.body = marked(str)
      }
    }
    xmlhttp.open('GET', '/api/?id=' + message, true)
    xmlhttp.send()
  }
}
</script>
<style>
#c {
  width: 600px;
  margin: auto;
  overflow: auto;
}
</style>
