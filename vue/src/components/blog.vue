<template>
  <div>
    <!-- todo: turn the aid to the article  {{message}}  -->
    <div>{{title}}</div>
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
      title: '',
      body: ''
    }
  },
  created: function (message) {
    // todo : use message send req to get the md
    let xmlhttp
    if (window.XMLHttpRequest) {
      //  IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
      xmlhttp = new XMLHttpRequest()
    } else {
      // IE6, IE5 浏览器执行代码
      // xmlhttp=new ActiveXObject('Microsoft.XMLHTTP')
    }
    // 绑定事件处理函数
    xmlhttp.onreadystatechange = function () {
      if (xmlhttp.readyState === 4 && xmlhttp.status === 200) {
      // document.getElementById('myDiv').innerHTML=xmlhttp.responseText
        alert(xmlhttp.responseText)
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
