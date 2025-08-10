<template>
  <div class="space-y-4">
    <div class="flex flex-wrap gap-2 items-center">
      <span class="text-sm text-gray-600">Filter:</span>
      <select v-model="sort" class="border rounded px-2 py-1">
        <option value="">Default</option>
        <option value="newest">Terbaru</option>
        <option value="price_asc">Harga ↑</option>
        <option value="price_desc">Harga ↓</option>
      </select>
      <button class="border rounded px-3 py-1" @click="load">Terapkan</button>
    </div>

    <div v-if="loading">Loading...</div>
    <div v-else class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <ProductCard
        v-for="p in products"
        :key="p.id"
        :product="p"
        @add="addToCart"
      />
    </div>

    <div class="flex items-center gap-3 justify-end">
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
import { onMounted, ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import api from "../lib/api";
import ProductCard from "../components/ProductCard.vue";
import { useCart } from "../stores/cart";

const route = useRoute();
const router = useRouter();
const cart = useCart();

const products = ref([]);
const loading = ref(false);
const total = ref(0);
const page = ref(1),
  per = ref(20);
const sort = ref(route.query.sort || "");

async function load() {
  loading.value = true;
  try {
    const params = {
      search: route.query.search || "",
      cat: route.query.cat || "",
      sort: sort.value || "",
      page: page.value,
      per_page: per.value,
    };
    const { data } = await api.get("/v1/products", { params });
    products.value = data.data || [];
    total.value = data.total || 0;
  } finally {
    loading.value = false;
  }
}
function addToCart(p) {
  cart.add(p);
}

watch(
  () => route.query,
  () => {
    page.value = 1;
    sort.value = route.query.sort || "";
    load();
  }
);
onMounted(load);
</script>
