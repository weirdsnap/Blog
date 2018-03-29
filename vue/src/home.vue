<template>
  <div id="name" v-on:click="$emit('onNameCLick')">
    <span>Snap's Blog</span>
  </div>
</template>
<script>
// export default {
//   methods: {
//     jumpDetails: function () {
//       console.log('jump')
//     }
//   }
// }
</script>
<style>
#name {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 5.5em;
}
#name span {
  cursor: pointer;
}
</style>
