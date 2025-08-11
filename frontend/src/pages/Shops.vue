<template>
  <div class="container py-8">
    <h1 class="text-2xl font-bold mb-4">Toko</h1>
    <div v-if="loading">Loading...</div>
    <div v-else class="grid sm:grid-cols-2 md:grid-cols-3 gap-4">
      <router-link
        v-for="s in shops"
        :key="s.id"
        :to="`/shops/${s.slug}`"
        class="card p-4 hover:shadow-md transition"
      >
        <div class="font-semibold">{{ s.name }}</div>
        <div class="text-sm text-gray-500">Produk: {{ s.product_count }}</div>
      </router-link>
    </div>
    <div class="flex gap-3 justify-end mt-4">
      <button class="btn" :disabled="page <= 1" @click="page--, load()">
        Prev
      </button>
      <span class="text-sm">Page {{ page }}</span>
      <button
        class="btn"
        :disabled="page * per >= total"
        @click="page++, load()"
      >
        Next
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import api from "../lib/api";
const shops = ref([]);
const page = ref(1),
  per = ref(20),
  total = ref(0),
  loading = ref(false);
async function load() {
  loading.value = true;
  try {
    const { data } = await api.get("/v1/shops", {
      params: { page: page.value, per_page: per.value },
    });
    shops.value = data.data || [];
    total.value = data.total || 0;
  } finally {
    loading.value = false;
  }
}
onMounted(load);
</script>
