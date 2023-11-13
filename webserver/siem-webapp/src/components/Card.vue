<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>{{ desc }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'Card',
  props: {
    msg: String,
    desc: String
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.hello {
  background-color: darkgrey;
  width: 20%;
  margin: 0 auto;
  color: white;
}
</style>
