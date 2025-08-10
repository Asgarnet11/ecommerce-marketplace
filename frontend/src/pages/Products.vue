<template>
  <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
    <ProductCard v-for="p in products" :key="p.id" :product="p" @add="add" />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import api from "../lib/api";
import ProductCard from "../components/ProductCard.vue";
import { useCart } from "../stores/cart";
const products = ref([]);
const cart = useCart();
async function load() {
  const { data } = await api.get("/v1/products");
  products.value = data.data || [];
}
function add(p) {
  cart.add(p);
}
onMounted(load);
</script>
