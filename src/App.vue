<template>
  <div id="app">
    <img alt="Vue logo" src="@/assets/logo.png" />
    <LuckyNumber :greeting="greeting" />
  </div>
</template>

<script>
import LuckyNumber from "@/components/LuckyNumber";

export default {
  name: "App",
  components: {
    LuckyNumber
  },
  data() {
    return {
      greeting: ""
    }
  },
  mounted() {
    const greetings = [
      "Yo dude!",
      "Hey there!",
      "Whassup homey?",
      "What's shakin'?",
      "Yo man!"
    ];

    // get random index and set greeting
    const i = Math.floor(Math.random() * greetings.length)
    this.greeting = greetings[i]
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
