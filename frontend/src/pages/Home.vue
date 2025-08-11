<template>
  <!-- Hero Section -->
  <section class="relative overflow-hidden">
    <div
      class="absolute inset-0 bg-gradient-to-br from-blue-50 via-white to-teal-50"
    ></div>

    <div
      class="container relative py-16 md:py-24 grid md:grid-cols-2 gap-10 items-center"
    >
      <!-- Copy/CTA -->
      <div>
        <p
          class="inline-block px-3 py-1 rounded-full bg-blue-100 text-blue-700 text-xs font-semibold mb-3"
        >
          Marketplace Go + Vue
        </p>
        <h1 class="text-4xl md:text-5xl font-extrabold leading-tight">
          Belanja cepat & <span class="text-brand">hemat</span> di satu tempat
        </h1>
        <p class="mt-4 text-gray-600 max-w-prose">
          Temukan produk favorit dari penjual tepercaya. Harga bersaing,
          pengiriman cepat, pengalaman belanja menyenangkan.
        </p>
        <div class="mt-6 flex gap-3">
          <router-link to="/products" class="btn btn-primary"
            >Mulai Belanja</router-link
          >
          <router-link to="/shops" class="btn btn-ghost"
            >Jelajahi Toko</router-link
          >
        </div>

        <!-- Quick search -->
        <div class="mt-6 flex gap-2">
          <input
            v-model="q"
            @keyup.enter="goSearch"
            placeholder="Cari produk…"
            class="flex-1 border rounded-xl px-4 py-3"
          />
          <button class="btn" @click="goSearch">Cari</button>
        </div>
      </div>

      <!-- Visual / Showcase -->
      <div class="relative">
        <div class="card p-3">
          <img
            class="rounded-2xl object-cover w-full aspect-[4/3]"
            src="https://images.unsplash.com/photo-1515168833906-d2a3b82b302a?q=80&w=1200&auto=format&fit=crop"
            alt="Belanja nyaman"
          />
        </div>
        <div
          class="absolute -bottom-4 -left-4 hidden md:block card px-4 py-3 shadow-lg"
        >
          <p class="text-sm">
            <b>Diskon Musim Ini</b><br />Hingga 40% untuk kebutuhan rumah.
          </p>
        </div>
      </div>
    </div>
  </section>

  <!-- Category Pills -->
  <section class="container py-6">
    <div class="flex flex-wrap gap-2">
      <button
        v-for="c in topCategories"
        :key="c.slug"
        class="px-3 py-1 rounded-full border hover:bg-gray-50 text-sm"
        @click="goCategory(c.slug)"
      >
        {{ c.name }}
      </button>
    </div>
  </section>

  <!-- Produk Terbaru -->
  <section class="container py-12">
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-2xl font-bold">Produk Terbaru</h2>
      <router-link to="/products" class="text-sm text-brand hover:underline"
        >Lihat semua</router-link
      >
    </div>
    <ProductGrid :initialFetch="true" />
  </section>

  <!-- Feature Highlights -->
  <section class="bg-gray-50">
    <div class="container py-12 grid md:grid-cols-3 gap-6">
      <div class="card p-6">
        <h3 class="font-semibold mb-2">Harga Bersahabat</h3>
        <p class="text-gray-600">
          Perbandingan harga mudah, promo rutin setiap pekan.
        </p>
      </div>
      <div class="card p-6">
        <h3 class="font-semibold mb-2">Pengiriman Cepat</h3>
        <p class="text-gray-600">
          Didukung kurir tepercaya & pelacakan realtime.
        </p>
      </div>
      <div class="card p-6">
        <h3 class="font-semibold mb-2">Aman & Terpercaya</h3>
        <p class="text-gray-600">
          Pembayaran aman, perlindungan pembeli menyeluruh.
        </p>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import api from "../lib/api";
import ProductGrid from "../components/ProductGrid.vue";

const router = useRouter();
const q = ref("");
const topCategories = ref([]);

function goSearch() {
  router.push({ name: "catalog", query: { search: q.value } });
}
function goCategory(slug) {
  router.push({ name: "catalog", query: { cat: slug } });
}

onMounted(async () => {
  try {
    // Ambil tree kategori & ambil level teratas (maks 8 untuk pill)
    const { data } = await api.get("/v1/categories/tree");
    topCategories.value = Array.isArray(data) ? data.slice(0, 8) : [];
  } catch {
    /* ignore */
  }
});
</script>

<style scoped>
/* opsional tweak kecil */
</style>
