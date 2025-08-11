<template>
  <div class="container py-16 max-w-md">
    <h1 class="text-2xl font-bold mb-4">Daftar</h1>
    <form class="space-y-3" @submit.prevent="submit">
      <input
        v-model="name"
        placeholder="Nama"
        class="border rounded-xl px-3 py-2 w-full"
        required
      />
      <input
        v-model="email"
        type="email"
        placeholder="Email"
        class="border rounded-xl px-3 py-2 w-full"
        required
      />
      <input
        v-model="password"
        type="password"
        placeholder="Password (min 6)"
        class="border rounded-xl px-3 py-2 w-full"
        required
      />
      <select v-model="role" class="border rounded-xl px-3 py-2 w-full">
        <option value="buyer">Buyer</option>
        <option value="seller">Seller</option>
      </select>
      <button class="btn btn-primary w-full" :disabled="loading">
        {{ loading ? "Mendaftar..." : "Daftar" }}
      </button>
      <p class="text-sm text-gray-600">
        Sudah punya akun?
        <router-link class="text-brand underline" to="/login"
          >Masuk</router-link
        >
      </p>
    </form>
  </div>
</template>

<script setup>
import { ref } from "vue";
import api from "../lib/api";
import { useAuth } from "../stores/auth";
import { useRouter } from "vue-router";

const name = ref("");
const email = ref("");
const password = ref("");
const role = ref("buyer");
const loading = ref(false);
const auth = useAuth();
const router = useRouter();

async function submit() {
  loading.value = true;
  try {
    const { data } = await api.post("/v1/auth/register", {
      name: name.value,
      email: email.value,
      password: password.value,
      role: role.value,
    });
    auth.setSession(data.access_token, data.user);
    router.push({ name: "landing" });
  } catch (e) {
    alert(e?.response?.data?.error || e.message);
  } finally {
    loading.value = false;
  }
}
</script>
