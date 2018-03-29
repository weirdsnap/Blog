<template>
  <div>
    <!-- todo: turn the aid to the article  {{message}}  -->
    <div> {{title}} </div>
    <div id="c">
      {{body}}
    </div>
  </div>
</template>
<script>
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
        let str = JSON.parse(xmlhttp.responseText)
        that.bogy += str
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
  overflow: auto;
}
</style>
