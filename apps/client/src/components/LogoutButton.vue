<script lang="ts" setup>
import {invoke} from "@tauri-apps/api/core";
import {refreshStore, RouteNames} from "../main.ts";
import {useRouter} from "vue-router";
import {ArrowLeftStartOnRectangleIcon} from "@heroicons/vue/16/solid";

const router = useRouter();

const onClickLogout = async () => {
  await invoke("logout");
  await refreshStore();
  void router.push(RouteNames.Login);
}
</script>

<template>
  <button @click="onClickLogout">
    <ArrowLeftStartOnRectangleIcon class="icon"/>
  </button>
</template>

<style scoped>
button {
  height: 30px;
  width: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  border: none;
  border-radius: 10px;
  background: #dedede;
  color: #000;
  cursor: pointer;
}

button:hover {
  background: #c9c9c9;
}

button:active {
  background: #b0b0b0;
}

.icon {
  height: 0.9rem;
  width: 0.9rem;
  pointer-events: none;
}
</style>