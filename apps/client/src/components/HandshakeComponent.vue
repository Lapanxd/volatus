<script lang="ts" setup>
import {invoke} from "@tauri-apps/api/core";
import {ref} from "vue";
import {PendingHandshakes} from "../core/models/pending-handshakes.ts";

const handshakes = ref<PendingHandshakes[]>([]);

const getHandshakes = async () => {
  handshakes.value = await invoke("get_pending_handshakes");
}

const responseHandshake = async (session_id: string, accepted: boolean) => {
  try {
    await invoke("handshake_response", {sessionId: session_id, accepted, sdpAnswer: "wip_sdp_answer"});
  } catch (error) {
    console.error(error);
  }
}
</script>

<template>
  <button @click="getHandshakes">Get pending handshakes</button>
  <div v-for="handshake in handshakes">
    {{ handshake }}
    <button @click="responseHandshake(handshake.session_id, true)">Accept</button>
    <button @click="responseHandshake(handshake.session_id, false)">Refuse</button>
  </div>
</template>

<style scoped>

</style>