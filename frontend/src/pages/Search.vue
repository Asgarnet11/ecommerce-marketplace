<template>
  <div class="container py-8">
    <div class="flex items-center gap-2 mb-6">
      <input
        v-model="q"
        @keyup.enter="apply"
        placeholder="Cari produk atau toko…"
        class="border rounded-xl px-3 py-2 flex-1"
      />
      <button class="btn" @click="apply">Cari</button>
    </div>

    <div v-if="loading">Mencari…</div>

    <template v-else>
      <section v-if="shops.length">
        <h2 class="text-xl font-bold mb-3">Toko</h2>
        <div class="grid sm:grid-cols-2 md:grid-cols-3 gap-4 mb-8">
          <router-link
            v-for="s in shops"
            :key="s.id"
            :to="`/shops/${s.slug}`"
            class="card p-4 hover:shadow-md transition"
          >
            <div class="font-semibold">{{ s.name }}</div>
            <div class="text-sm text-gray-500">Kunjungi toko →</div>
          </router-link>
        </div>
      </section>

      <section>
        <div class="flex items-center justify-between mb-3">
          <h2 class="text-xl font-bold">Produk</h2>
          <router-link
            :to="{ name: 'catalog', query: { search: origQ } }"
            class="text-sm text-brand underline"
            >Lihat semua hasil</router-link
          >
        </div>
        <div v-if="products.length === 0" class="text-gray-500">
          Tidak ada produk ditemukan.
        </div>
        <div v-else class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <ProductCard
            v-for="p in products"
            :key="p.id"
            :product="p"
            @add="add"
          />
        </div>
      </section>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import api from "../lib/api";
import ProductCard from "../components/ProductCard.vue";
import { useCartApi } from "../stores/cartApi";

const route = useRoute(),
  router = useRouter();
const cart = useCartApi();

const q = ref(String(route.query.q || ""));
const origQ = ref(q.value);
const products = ref([]),
  shops = ref([]),
  loading = ref(false);

function apply() {
  router.push({ name: "search", query: { q: q.value.trim() } });
}

async function fetcher() {
  const s = String(route.query.q || "");
  origQ.value = s;
  if (!s) {
    products.value = [];
    shops.value = [];
    return;
  }
  loading.value = true;
  try {
    const { data } = await api.get("/v1/search", {
      params: { q: s, limit_products: 16, limit_shops: 6 },
    });
    products.value = data.products || [];
    shops.value = data.shops || [];
  } finally {
    loading.value = false;
  }
}
function add(p) {
  cart.add(p.id, 1);
}

onMounted(fetcher);
watch(() => route.query.q, fetcher);
</script>
