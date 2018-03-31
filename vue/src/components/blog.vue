<template>
  <div>
    <div id="c" v-html="body" class="markdown-body">
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
    // [ajax] use message send req to get the md
    let xmlhttp
    let that = this
    if (window.XMLHttpRequest) {
      //  IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
      xmlhttp = new XMLHttpRequest()
    }
    // 绑定事件处理函数
    xmlhttp.onreadystatechange = function () {
      if (xmlhttp.readyState === 4 && xmlhttp.status === 200) {
        // alert(xmlhttp.responseText)
        let obj = JSON.parse(xmlhttp.responseText)
        let str = obj.content
        // let str = marked('# 大标题啊\n## 二标题啊1233412\n我这里想说几句话\n真的真的真的真的真的真的真的真的真的真的真的真的真的真的很多句话\n```lang=js\ndocument.get()\n```\n 但是呢\n 这样子\n 是不是不太好\n 其实应\n 该\n 多写\n一些的\nb\nc\na\nb\nc\n')
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
    xmlhttp.open('GET', '/api/?id=' + message, true)
    xmlhttp.send()
  }
}
</script>
<style>
h1,h2,h3,h4,h5,p {
  margin:0;
}
#c {
  width: 70%;
  margin: auto;
  overflow: auto;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}
</style>
