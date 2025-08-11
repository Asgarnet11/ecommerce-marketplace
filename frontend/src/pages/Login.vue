<template>
  <div class="container py-16 max-w-md">
    <h1 class="text-2xl font-bold mb-4">Masuk</h1>
    <form class="space-y-3" @submit.prevent="submit">
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
        placeholder="Password"
        class="border rounded-xl px-3 py-2 w-full"
        required
      />
      <button class="btn btn-primary w-full" :disabled="loading">
        {{ loading ? "Masuk..." : "Masuk" }}
      </button>
      <p class="text-sm text-gray-600">
        Belum punya akun?
        <router-link class="text-brand underline" to="/register"
          >Daftar</router-link
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

const email = ref("");
const password = ref("");
const loading = ref(false);
const auth = useAuth();
const router = useRouter();

async function submit() {
  loading.value = true;
  try {
    const { data } = await api.post("/v1/auth/login", {
      email: email.value,
      password: password.value,
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
