<script lang="ts" setup>
import UserProfile from "../components/UserProfile.vue";
import Handshake from "../components/Handshake.vue";
import Conversations from "../components/Conversations.vue";
</script>

<template>
  <div class="home-page">
    <div class="nav">
      <UserProfile/>
    </div>
    <div class="content">
      <Handshake/>
      <Conversations/>
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
  gap: 2rem;
  padding: 1rem;
  box-sizing: border-box;
}

.content {
  height: 100%;
  padding: 0 1rem 1rem 1rem;
  box-sizing: border-box;
  display: flex;
  gap: 0.5rem;
}

.content > div {
  box-sizing: border-box;
  width: 100%;
  height: 100%;
}
</style>