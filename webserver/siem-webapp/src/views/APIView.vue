<template>
    <div>
        <button>Trigger API</button>
        <h1>{{ somevalue }}</h1>
    </div>
</template>
<script lang="ts">
import axios from 'axios';
var someval; 
function getApi() {
    axios.get("http://localhost:5000").then((res) => {
        console.log(res)
        someval = "cavalo";
    });
}
export default {
    
    data() {
        return {
            somevalue: ""
        }
    },
}
</script>
