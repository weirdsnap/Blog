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
        // let str = marked("# ES6学习\n \n ## let 命令\n \n \n ### 作用域\n \n 用法类似于var，但是所声明的变量，只在let命令所在的代码块内有效。\n let，声明的变量仅在**块级作用域**内有效。\n \n ``` lang=js\n {\nlet a = 10;\nvar b = 1;\n}\n \n a // ReferenceError: a is not defined.\n b // 1\n ```\n ### for\n 对于for循环很有意思,控制条件和循环体不是一个域\n \n ``` lang=JavaScript\n for (let i = 0; i < 3; i++) {\n   let i = 'abc';\n   console.log(i);\n }\n // abc\n // abc\n // abc \n ```\n \n ### 没有变量提升\n \n ### 暂时性死区 TDZ\n \n 就是哪怕外面有同名全局变量，但是在内部有let或者const的同名变量，那么在let前不能使用此变量。\n \n typeof不再安全\n \n ``` lang=js\n typeof x; // ReferenceError\n let x;\n ```\n \n ## 块级作用域\n \n {{todo}}\n ES5中的一些东西这里不一样了，比如IIFE，但是具体浏览器或者环境支持的程度不明确。\n \n ## const\n \n 保证一个变量的指向的内存地址不得改动，怎么来理解这句话呢，如果是简单变量比如int，那么就是他就是存在这个地址的，所有不可以改变。但是复杂的数据结构就不一样，等于规定了他只能指向这个结构但是这个结构里面是可以改动的。\n 所以无法用const使得一个对象变成不可变的。要用freeze。\n \n ``` lang=js\n const a = [];\n a.push('Hello'); // 可执行\n a.length = 0;    // 可执行\n a = ['Dave'];    // 报错\n ```\n \n ## 顶层对象\n \n window还是global，全局还是顶层？\n \n ES5中，你不去特意声明，你的全局变量就是顶层对象的属性。\n")
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
    // xmlhttp.open('GET', '/api/?id=' + aId, true) // when build
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
