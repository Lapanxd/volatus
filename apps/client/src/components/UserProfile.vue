<script lang="ts" setup>
import {inject, onMounted, ref} from "vue";
import {Store} from "@tauri-apps/plugin-store";
import {User} from "../core/models/user.ts";
import {Nullable} from "../core/types/nullable.ts";
import LogoutButton from "./LogoutButton.vue";

const store = inject<Store>("store")!;

const user = ref<Nullable<User>>(null);

onMounted(async () => {
  const u = await store.get<User>("user");

  if (u) {
    user.value = u;
  }
});


</script>

<template>
  <div class="user-profile">
    <LogoutButton/>
    <img :src="`https://ui-avatars.com/api/?name=${user?.username}`" alt="User's avatar"/>
  </div>
</template>

<style scoped>
.user-profile {
  display: flex;
  gap: 0.5rem;
}

img {
  border-radius: 50%;
  height: 30px;
  width: 30px;
}
</style>