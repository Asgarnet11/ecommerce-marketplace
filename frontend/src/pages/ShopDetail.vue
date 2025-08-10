<script setup>
import { ref, onMounted, watch } from "vue";
import { useRoute } from "vue-router";
import api from "../lib/api";
import ProductCard from "../components/ProductCard.vue";
import { useCart } from "../stores/cart";

const route = useRoute();
const shop = ref(null);
const products = ref([]);
const cart = useCart();

// paging & sort opsional
const page = ref(1),
  per = ref(20),
  total = ref(0);
const sort = ref(""); // '', newest, price_asc, price_desc
const search = ref("");

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
  cart.add(p);
}

onMounted(async () => {
  await loadShop();
  await loadProducts();
});
watch([page, sort, search], loadProducts);
</script>

<template>
  <div v-if="shop">
    <div class="flex items-center gap-4 mb-6">
      <div class="w-16 h-16 bg-gray-200 rounded-full"></div>
      <div>
        <h1 class="text-2xl font-bold">{{ shop.name }}</h1>
        <p class="text-sm text-gray-600">Produk: {{ shop.product_count }}</p>
      </div>
    </div>

    <div class="flex flex-wrap gap-2 items-center mb-4">
      <input
        v-model="search"
        placeholder="Cari di toko ini..."
        class="border rounded px-2 py-1"
      />
      <select v-model="sort" class="border rounded px-2 py-1">
        <option value="">Default</option>
        <option value="newest">Terbaru</option>
        <option value="price_asc">Harga ↑</option>
        <option value="price_desc">Harga ↓</option>
      </select>
      <span class="ml-auto text-sm text-gray-600">Total: {{ total }}</span>
    </div>

    <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <ProductCard
        v-for="p in products"
        :key="p.id"
        :product="p"
        @add="addToCart"
      />
    </div>

    <div class="flex items-center gap-3 justify-end mt-4">
      <button
        class="border rounded px-3 py-1"
        :disabled="page <= 1"
        @click="page--"
      >
        Prev
      </button>
      <span class="text-sm">Page {{ page }}</span>
      <button
        class="border rounded px-3 py-1"
        :disabled="page * per >= total"
        @click="page++"
      >
        Next
      </button>
    </div>
  </div>
</template>
