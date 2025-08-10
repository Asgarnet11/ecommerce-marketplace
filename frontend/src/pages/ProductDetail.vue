<template>
  <div v-if="p">
    <div class="grid md:grid-cols-2 gap-6">
      <div class="bg-gray-100 aspect-square rounded"></div>
      <div>
        <h1 class="text-2xl font-bold">{{ p.title }}</h1>
        <div class="text-xl font-semibold mt-2">
          Rp {{ p.price.toLocaleString("id-ID") }}
        </div>
        <p class="mt-4 text-gray-600">Stok: {{ p.stock }}</p>
        <button
          class="mt-4 px-4 py-2 rounded bg-gray-900 text-white"
          @click="add()"
        >
          Tambah ke Keranjang
        </button>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import api from "../lib/api";
import { useCart } from "../stores/cart";
const route = useRoute();
const cart = useCart();
const p = ref(null);
async function load() {
  const { data } = await api.get(`/v1/products/${route.params.slug}`);
  p.value = data;
}
function add() {
  cart.add(p.value);
}
onMounted(load);
</script>
