<template>
  <div class="container py-8" v-if="shop">
    <div class="flex items-center gap-4 mb-6">
      <div class="w-16 h-16 rounded-full bg-gray-200"></div>
      <div>
        <h1 class="text-2xl font-bold">{{ shop.name }}</h1>
        <p class="text-gray-500 text-sm">Produk: {{ shop.product_count }}</p>
      </div>
    </div>

    <div class="flex flex-wrap gap-2 items-center mb-4">
      <input
        v-model="search"
        placeholder="Cari di toko ini..."
        class="border rounded-xl px-3 py-2"
      />
      <select v-model="sort" class="border rounded-xl px-3 py-2">
        <option value="">Default</option>
        <option value="newest">Terbaru</option>
        <option value="price_asc">Harga ↑</option>
        <option value="price_desc">Harga ↓</option>
      </select>
    </div>

    <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <ProductCard
        v-for="p in products"
        :key="p.id"
        :product="p"
        @add="addToCart"
      />
    </div>

    <div class="flex gap-3 justify-end mt-4">
      <button class="btn" :disabled="page <= 1" @click="page--, loadProducts()">
        Prev
      </button>
      <span class="text-sm">Page {{ page }}</span>
      <button
        class="btn"
        :disabled="page * per >= total"
        @click="page++, loadProducts()"
      >
        Next
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import api from "../lib/api";
import ProductCard from "../components/ProductCard.vue";
import { useCartApi } from "../stores/cartApi";

const route = useRoute();
const router = useRouter();
const shop = ref(null);
const products = ref([]);
const total = ref(0);
const page = ref(1),
  per = ref(20);
const sort = ref("");
const search = ref("");
const cart = useCartApi();

async function loadShop() {
  const { data } = await api.get(`/v1/shops/${route.params.slug}`);
  shop.value = data;
}
async function loadProducts() {
  const { data } = await api.get(`/v1/shops/${route.params.slug}/products`, {
    params: {
      page: page.value,
      per_page: per.value,
      sort: sort.value,
      search: search.value,
    },
  });
  products.value = data.data || [];
  total.value = data.total || 0;
}
function addToCart(p) {
  cart.add(p.id, 1).catch(() => router.push({ name: "login" }));
}

onMounted(async () => {
  await loadShop();
  await loadProducts();
});
watch([sort, search], loadProducts);
</script>
