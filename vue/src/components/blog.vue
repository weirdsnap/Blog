<template>
  <div>
    <!-- todo: turn the aid to the article  {{message}}  -->
    <div>title</div>
    <div id="c">
    contentscontentscontentscontentscontentscontentscontents
    contentscontentscontentscontentscontentscontentscontents
    contentscontentscontentscontentscontentscontentscontents
    contentscontentscontentscontentscontentscontentscontents
    contentscontentscontentscontentscontentscontentscontents
    contentscontentscontentscontentscontents
    </div>
  </div>
</template>
<script>
export default {
  props: ['message']
}
</script>
<style>
#c {
  width: 600px;
  overflow: auto;
}
</style>
