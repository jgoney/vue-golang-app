<template>
  <div class="hello">
    <h1>{{ greeting }}</h1>
    <h2>Your lucky number is {{ luckyNumber }}</h2>
    <button @click="updateLuckyNumber">Click here to get a new lucky number</button>
  </div>
</template>

<script>
export default {
  name: "LuckyNumber",
  props: {
    greeting: {
      type: String,
      default: "Hey there"
    }
  },
  data() {
    return {
      luckyNumber: 0
    };
  },
  methods: {
    updateLuckyNumber() {
      console.log("update")
      fetch("/api/lucky")
        .then(resp => resp.json())
        .then(data => {
          this.luckyNumber = data.number;
        });
    }
  },
  mounted() {
    this.updateLuckyNumber();
  }
};
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
