<template>
  <div class="hello">
    <router-link to="/api">Home</router-link>
    <button @click="doreq()">request</button>
    <h1>{{ response }}</h1>
  </div>
</template>

<script>
// import axios from 'axios'
export default {
  name: 'HelloWorld',
  data() {
    return {
      response: ""
    }
  },
  methods: {
    async doreq(){
      var host = process.env.VUE_APP_BE_HOST
      var port = process.env.VUE_APP_BE_PORT
      console.log(`http://${host}:${port}/`)
      // axios.get(`http://${host}:${port}/`).then((res) => this.response = res.data)
      var res = await fetch(`http://${host}:${port}/`)
      res.json().then((data) => this.response = data)
    }
  },
}
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
</style>
