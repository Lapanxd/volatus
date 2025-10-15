<script setup lang="ts">
import { ref } from "vue";
import { invoke } from "@tauri-apps/api/core";

const token = ref("");


async function login(){
  try {
    token.value = await invoke("login", { username: "Lapanxd", password: "Password123@" });
  } catch (error) {
    console.log(error);
  }

}
</script>

<template>
  <button @click="login">Login</button>
  {{ token}}
</template>

<style scoped></style>