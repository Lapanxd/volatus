<script lang="ts" setup>
import {invoke} from "@tauri-apps/api/core";
import {ref, watchEffect} from "vue";
import {PendingHandshakesOutput, PendingHandshakeWithUser} from "../core/models/pending-handshakes-output.ts";
import {User} from "../core/models/user.ts";

const handshakes = ref<PendingHandshakesOutput[]>([]);
const handshakesWithUsers = ref<PendingHandshakeWithUser[]>([])


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

const getUserById = async (userId: number): Promise<User> => {
  return await invoke<User>("get_user_by_id", {id: userId});
}

watchEffect(async () => {
  handshakesWithUsers.value = await Promise.all(
      handshakes.value.map(async h => ({
        sessionId: h.session_id,
        fromUser: await getUserById(h.from_user_id)
      }))
  )
})
</script>

<template>
  <div class="handshake">
    <div class="header">
      <p>Pending requests</p>
      <button @click="getHandshakes">Refresh</button>
    </div>
    <div v-for="handshake in handshakesWithUsers" class="pending-handshake">
      <div class="name">
        <img :src="`https://ui-avatars.com/api/?name=${handshake.fromUser.username}`" alt="User's avatar"/>
        <p>{{ handshake.fromUser.username }}</p>
      </div>
      <div class="actions">
        <button @click="responseHandshake(handshake.sessionId, true)">Accept</button>
        <button @click="responseHandshake(handshake.sessionId, false)">Refuse</button>
      </div>
    </div>
  </div>

</template>

<style scoped>
.handshake {
  background: #f3f3f3;
  padding: 1rem;
  border-radius: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.header {
  display: flex;
  align-items: start;
  justify-content: space-between;
}

.pending-handshake {
  display: flex;
  gap: 0.5rem;
  align-items: center;
  justify-content: space-between;
  background: #ffffff;
  padding: 0.5rem;
  border-radius: 0.8rem;
}

.pending-handshake img {
  height: 30px;
  width: 30px;
  border-radius: 50%;
}

.pending-handshake .name {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.pending-handshake .actions {
  display: flex;
  align-items: center;
  gap: 0.2rem;
}
</style>