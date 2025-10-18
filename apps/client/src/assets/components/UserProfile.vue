<script lang="ts" setup>
import {inject, onMounted, ref} from "vue";
import {Store} from "@tauri-apps/plugin-store";
import {User} from "../../core/models/user.ts";
import {Nullable} from "../../core/types/nullable.ts";
import {invoke} from "@tauri-apps/api/core";
import {refreshStore, RouteNames} from "../../main.ts";
import {useRouter} from "vue-router";

const store = inject<Store>("store")!;
const router = useRouter();

const user = ref<Nullable<User>>(null);

onMounted(async () => {
  const u = await store.get<User>("user");

  if (u) {
    user.value = u;
  }
});

const onClickLogout = async () => {
  await invoke("logout");
  await refreshStore();
  void router.push(RouteNames.Login);
}

</script>

<template>
  <div class="user-profile">
    <p>{{ user?.id }}</p>
    <p>{{ user?.username }}</p>
    <button @click="onClickLogout">Logout</button>
  </div>
</template>

<style scoped>
.user-profile {
  display: flex;
  gap: 0.2rem;
}
</style>