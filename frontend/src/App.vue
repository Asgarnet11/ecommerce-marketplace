<template>
  <div class="min-h-screen">
    <header class="sticky top-0 bg-white border-b">
      <div class="max-w-6xl mx-auto px-4 py-3 flex items-center gap-4">
        <router-link class="font-bold text-xl" to="/">Market</router-link>
        <input
          v-model="q"
          @keyup.enter="go"
          placeholder="Cari produk..."
          class="flex-1 border rounded px-3 py-2"
        />
        <router-link class="px-3 py-2 rounded bg-gray-900 text-white" to="/cart"
          >Keranjang</router-link
        >
      </div>
    </header>
    <main class="max-w-6xl mx-auto px-4 py-6">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();
const q = ref("");
function go() {
  router.push({ name: "products", query: { search: q.value } });
}
</script>
