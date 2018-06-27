<template>
  <div id="app">
    <div class= "tabs opacity70">
      <ul>
        <li>
          <!-- 使用router-link指向定义的path -->
          <router-link :to="{name: 'Snap'}">
            Snap
          </router-link>
        </li>
        <li>
          <router-link :to="{name: 'Home'}">
            Home
          </router-link>
        </li>
        <li>
          <router-link :to="{name: 'BlogList'}">
            Blogs
          </router-link>
        </li>
        <li>
          <a href="https://github.com/weirdsnap">
            Github
          </a>
        </li>
      </ul>
    </div>
    <div class="content opacity70">
      <router-view v-on:onBlogSelect="jumpBlog($event)" v-bind:message="aid" v-on:onNameCLick="jumplist()"></router-view>
    </div>
  </div>
</template>

<script>
import router from './router'

export default {
  name: 'App',
  data () {
    return {
      aid: ''
    }
  },
  methods: {
    jumpBlog (item) {
      this.aid = item.aid
      router.push({name: 'blog', params: {id: item.aid}})
    },
    jumplist () {
      router.push({name: 'BlogList'})
    }
  }
}
</script>

<style>
html {
  height: 100%;
}

body {
  background: url("./assets/background1.jpg");
  background-position-x: 40%;
  height: 100%;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-size: 10px;
  text-align: center;
  color: black;
  margin: 0;
}

@media screen and (min-device-width: 600px){
  body {
    background-size:100%;
    background-repeat: no-repeat;
    background-attachment: fixed;
    font-size: 16px;
  }
}

#app {
  margin: 0px;
  height: 100%;
}

.tabs {
  height: 10%;
  background-color: black;
}

.content {
  height: 90%;
  background-color: bisque;
}

ul {
  height:100%;
  display: flex;
  margin: 0;
  justify-content: flex-end;
}

li {
  height: 100%;
  margin-left: 2%;
  margin-right: 2%;
  list-style: none;
  border: 1px;
  display: flex;
  justify-content: center;
  align-items: center;
}

a {
  height: 100%;
  color: whitesmoke;
  text-decoration: none;
  font-size: 2rem;
  display: flex;
  justify-content: center;
  align-items: center;
}

.opacity60 {
    filter: alpha(opacity : 60);
    opacity:0.6;
    -moz-opacity: 0.6;
    -khtml-opacity:0.6;
}

.opacity70 {
    filter: alpha(opacity : 70);
    opacity: 0.7;
    -moz-opacity: 0.7;
    -khtml-opacity: 0.7;
}

.opacity90 {
    filter: alpha(opacity : 90);
    opacity:0.9;
    -moz-opacity: 0.9;
    -khtml-opacity:0.9;
}
</style>
