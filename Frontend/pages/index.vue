<template>
  <div>
    <div class="light lightPink"></div>
    <div class="light lightBlue"></div>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.light {
  position: absolute;
  border-radius: 50%;
  opacity: 0.3;
  filter: blur(100px);
}

.lightPink {
  width: 2500px;
  height: 700px;
  background-color: rgba(194, 82, 225, 0.25);
  top: -30%;
  left: -10%;
}
.lightBlue {
  width: 2500px;
  height: 700px;
  background-color: rgba(110, 203, 245, 0.25);
  top: 25%;
  left: -15%;
}
</style>