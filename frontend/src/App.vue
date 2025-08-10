<template>
  <div class="min-h-screen">
    <header class="sticky top-0 bg-white border-b z-10">
      <div class="max-w-6xl mx-auto px-4 py-3 flex items-center gap-4">
        <router-link class="font-bold text-xl" to="/">Market</router-link>
        <input
          v-model="q"
          @keyup.enter="goSearch"
          placeholder="Cari produk..."
          class="flex-1 border rounded px-3 py-2"
        />
        <router-link class="px-3 py-2 rounded bg-gray-900 text-white" to="/cart"
          >Keranjang</router-link
        >
      </div>
    </header>

    <main
      class="max-w-6xl mx-auto px-4 py-6 grid md:grid-cols-[240px,1fr] gap-6"
    >
      <aside class="hidden md:block">
        <h3 class="font-semibold mb-2">Kategori</h3>
        <CategoryTree />
      </aside>
      <section>
        <router-view />
      </section>
    </main>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import CategoryTree from "./components/CategoryTree.vue";

const router = useRouter();
const q = ref("");
function goSearch() {
  router.push({ name: "products", query: { search: q.value } });
}
</script>
