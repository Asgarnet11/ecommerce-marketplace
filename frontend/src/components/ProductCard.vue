<template>
  <div class="bg-white rounded-xl shadow p-3 hover:shadow-md transition">
    <router-link :to="`/p/${product.slug}`">
      <div class="aspect-square bg-gray-100 rounded mb-3"></div>
      <h3 class="font-medium line-clamp-2">{{ product.title }}</h3>
      <div class="mt-1 font-semibold">
        Rp {{ product.price.toLocaleString("id-ID") }}
      </div>
    </router-link>
    <button
      class="mt-3 w-full border rounded py-2"
      @click="$emit('add', product)"
    >
      Tambah ke Keranjang
    </button>
  </div>
</template>
<script setup>
defineProps({ product: { type: Object, required: true } });
</script>
