<template>
  <div id="snap">
      <div id="cebian">
        <div id="self">
          <img src="../src/assets/head.jpg" alt="">
          <span>Snap</span>
        </div>
        <div v-for="skill in skills" v-bind:key="skill.index">
          <span>{{skill.name}}</span>
          <span>{{skill.begintime}}至今</span>
        </div>
      </div>
      <div>
        <button v-on:click="addBlog()">添加博客</button>
      </div>
  </div>
</template>
<script>
export default {
  data () {
    return {
      skills: [
        {
          name: '前端开发',
          begintime: '2017'
        },
        {
          name: 'c/cpp',
          begintime: '2015'
        }
      ]
    }
  },
  methods: {
    addBlog () {
      //
    }
  }
}
</script>
<style>
  #snap {
    display:flex;
    flex-direction:row;
    font-weight:600;
  }
  #cebian{
    width:20%;
  }
  #self {
    font-size:2rem;
  }
  #self img{
    margin:20% 20% 0% 20%;
    width:60%;
    height:60%;
  }
</style>
