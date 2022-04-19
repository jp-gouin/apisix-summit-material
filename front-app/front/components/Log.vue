<template>
  <div
    class="flex flex-col bg-white shadow rounded-xl w-full mr-2 h-auto bg-opacity-75 min-heigth"
  >
    <div class="dot-loader">
      <div class="dot"></div>
      <div class="dot"></div>
      <div class="dot"></div>
    </div>
    <h2 class="subtitle">Results</h2>
    <div class="text-left ml-4" v-for="(message, index) in messages" :key="index">
      <span class="text-green-600">{{ message.action }}</span> :
      <span class="text-gray-800">{{ message.result }}</span>
      <br />
    </div>
  </div>
</template>
<script>
export default {
  props: {
    messages: Array,
  },
};
</script>
<style scoped>
.dot-loader {
  -webkit-box-align: center;
  align-items: center;
  margin-left: 1rem;
  display: -webkit-box;
  display: flex;
  height: 1.5em;
  -webkit-box-pack: justify;
  justify-content: space-between;
  width: 2.5em;
}
.dot-loader .dot {
  -webkit-animation: bounce 0.5s infinite alternate;
  animation: bounce 0.5s infinite alternate;
  -webkit-animation-fill-mode: both;
  animation-fill-mode: both;
  -webkit-animation-timing-function: ease;
  animation-timing-function: ease;
  background-color: rgb(92, 92, 92);
  border-radius: 50%;
  height: 0.5em;
  width: 0.5em;
}
.dot-loader .dot:nth-of-type(2) {
  -webkit-animation-delay: 0.1s;
  animation-delay: 0.1s;
}
.dot-loader .dot:nth-of-type(3) {
  -webkit-animation-delay: 0.2s;
  animation-delay: 0.2s;
}

@-webkit-keyframes bounce {
  from {
    -webkit-transform: translateY(0.2em);
    transform: translateY(0.2em);
  }
  to {
    -webkit-transform: translateY(-0.2em);
    transform: translateY(-0.2em);
  }
}

@keyframes bounce {
  from {
    -webkit-transform: translateY(0.2em);
    transform: translateY(0.2em);
  }
  to {
    -webkit-transform: translateY(-0.2em);
    transform: translateY(-0.2em);
  }
}
</style>