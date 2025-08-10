<template>
  <div v-if="shop">
    <div class="flex items-center gap-4 mb-6">
      <div class="w-16 h-16 bg-gray-200 rounded-full"></div>
      <div>
        <h1 class="text-2xl font-bold">{{ shop.name }}</h1>
        <p class="text-sm text-gray-600">Produk: {{ shop.product_count }}</p>
      </div>
    </div>

    <h2 class="text-xl font-semibold mb-3">Produk</h2>
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <ProductCard
        v-for="p in products"
        :key="p.id"
        :product="p"
        @add="addToCart"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import api from "../lib/api";
import ProductCard from "../components/ProductCard.vue";
import { useCart } from "../stores/cart";

const route = useRoute();
const shop = ref(null);
const products = ref([]);
const cart = useCart();

async function load() {
  const slug = route.params.slug;
  const shopRes = await api.get(`/v1/shops/${slug}`);
  shop.value = shopRes.data;

  // ambil produk shop ini (kita pakai endpoint /v1/products dengan filter q=shop_slug sementara)
  const { data } = await api.get("/v1/products", {
    params: { search: shop.value.name },
  });
  products.value = data.data || [];
}
function addToCart(p) {
  cart.add(p);
}
onMounted(load);
</script>
