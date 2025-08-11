<template>
  <div class="container py-8" v-if="p">
    <div class="grid md:grid-cols-2 gap-8">
      <img
        :src="p.image || placeholder"
        class="rounded-2xl object-cover w-full aspect-square"
      />
      <div>
        <h1 class="text-3xl font-bold">{{ p.title }}</h1>
        <div class="text-2xl font-semibold mt-2">
          Rp {{ p.price.toLocaleString("id-ID") }}
        </div>
        <p class="mt-4 text-gray-600">{{ p.description }}</p>
        <div class="mt-4 text-sm text-gray-500">Stok: {{ p.stock }}</div>
        <button class="btn btn-primary mt-6" @click="add()">
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
import { useCartApi } from "../stores/cartApi";
const route = useRoute();
const p = ref(null);
const placeholder = "https://placehold.co/600x600?text=Produk";
const cart = useCartApi();
async function load() {
  const { data } = await api.get(`/v1/products/${route.params.slug}`);
  p.value = data;
}
function add() {
  if (p.value) cart.add(p.value.id, 1);
}
onMounted(load);
</script>
