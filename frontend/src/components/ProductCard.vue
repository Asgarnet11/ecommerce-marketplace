<template>
  <div class="card p-3 hover:shadow-md transition">
    <router-link :to="`/p/${product.slug}`">
      <img
        :src="product.image || placeholder"
        class="rounded-xl object-cover w-full aspect-square"
      />
      <h3 class="font-medium mt-2 line-clamp-2">{{ product.title }}</h3>
      <div class="mt-1 font-semibold">
        Rp {{ product.price.toLocaleString("id-ID") }}
      </div>
    </router-link>
    <button class="btn w-full mt-3" @click="$emit('add', product)">
      Tambah ke Keranjang
    </button>
  </div>
</template>

<script setup>
const placeholder = "https://placehold.co/600x600?text=Produk";
defineProps({ product: { type: Object, required: true } });
</script>
