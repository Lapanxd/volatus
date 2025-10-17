<script setup lang="ts">
import { ref } from "vue";
import { invoke } from "@tauri-apps/api/core";

const token = ref("");
const user = ref("");


async function login(){
  try {
    token.value = await invoke("login", { username: "Lapanxd", password: "Password123@" });
  } catch (error) {
    console.log(error);
  }
}

async function getMe() {
  try {
    user.value = await invoke("get_me");
  } catch (error) {
    console.log(error);
  }
}
</script>

<template>
  <button @click="login">Login</button>
  <button @click="getMe">Get profile</button>
  {{ user }}
</template>

<style scoped></style>