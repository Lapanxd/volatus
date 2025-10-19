<script lang="ts" setup>
import UserProfile from "../components/UserProfile.vue";
import HandshakeComponent from "../components/HandshakeComponent.vue";
</script>

<template>
  <div class="home-page">
    <div class="nav">
      <div class="divider"></div>
      <UserProfile/>
    </div>
    <div class="content">
      <HandshakeComponent/>
    </div>
  </div>

</template>

<style scoped>
.home-page {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
}

.nav {
  display: flex;
  justify-content: end;
  padding: 0.5rem;
  box-sizing: border-box;
}

.content {
  padding: 0.5rem;
  box-sizing: border-box;
}
</style>