<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <div v-for="item in list">
      <Card :msg=item />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Card from '@/components/Card.vue'; // @ is an alias to /src

export default defineComponent({
  name: 'HomeView',
  components: {
    Card,
  },
  data() {
    return {
      count:1,
      list: ["l", "a", "b", "u", "y"]
    }
  },
});

</script>
