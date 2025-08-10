<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Toko</h1>
    <div v-if="loading">Loading...</div>
    <div v-else class="grid md:grid-cols-3 gap-4">
      <div
        v-for="s in shops"
        :key="s.id"
        class="bg-white rounded-xl shadow p-4"
      >
        <router-link
          :to="`/shops/${s.slug}`"
          class="font-semibold hover:underline"
          >{{ s.name }}</router-link
        >
        <div class="text-sm text-gray-500">Produk: {{ s.product_count }}</div>
      </div>
    </div>
    <div class="flex items-center gap-3 justify-end mt-4">
      <button
        class="border rounded px-3 py-1"
        :disabled="page <= 1"
        @click="
          page--;
          load();
        "
      >
        Prev
      </button>
      <span class="text-sm">Page {{ page }}</span>
      <button
        class="border rounded px-3 py-1"
        :disabled="page * per >= total"
        @click="
          page++;
          load();
        "
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
  total = ref(0);
const loading = ref(false);

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
