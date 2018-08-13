<template>
  <div id="snap">
      <sidebar id="cebian">
        <div id="self">
          <img src="../src/assets/head.jpg" alt="">
          <span>Snap</span>
        </div>
        <div v-for="skill in skills" v-bind:key="skill.index">
          <span>{{skill.name}}</span>
          <span>{{skill.begintime}}-至今</span>
          <progress v-bind:value="skill.degree" max="100"></progress>
        </div>
      </sidebar>
      <div>

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
          begintime: '2017',
          degree: '50'
        },
        {
          name: 'c/cpp',
          begintime: '2015',
          degree: '44'
        },
        {
          name: 'vue',
          begintime: '2018',
          degree: '33'
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

  progress
  {
    width: 168px;
    height: 5px;
  }
  progress::-webkit-progress-bar
  {
    background-color:#d7d7d7;
  }
  progress::-webkit-progress-value
  {
    background-color:orange;
  }
</style>
