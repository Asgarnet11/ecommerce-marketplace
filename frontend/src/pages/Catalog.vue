<template>
  <div class="container py-8">
    <div class="flex flex-wrap items-center gap-2 mb-4">
      <input
        v-model="q"
        placeholder="Cari produk..."
        class="border rounded-xl px-3 py-2"
        @keyup.enter="apply"
      />
      <select v-model="sort" class="border rounded-xl px-3 py-2">
        <option value="">Default</option>
        <option value="newest">Terbaru</option>
        <option value="price_asc">Harga ↑</option>
        <option value="price_desc">Harga ↓</option>
      </select>
      <button class="btn" @click="apply">Terapkan</button>
    </div>
    <ProductGrid
      :query="{ search: params.search, cat: params.cat, sort: params.sort }"
      :initialFetch="true"
    />
  </div>
</template>
<script setup>
import { reactive, ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import ProductGrid from "../components/ProductGrid.vue";
const route = useRoute();
const router = useRouter();
const params = reactive({
  search: route.query.search || "",
  cat: route.query.cat || "",
  sort: route.query.sort || "",
});
const q = ref(params.search);
const sort = ref(params.sort);
function apply() {
  router.push({
    name: "catalog",
    query: { search: q.value, cat: params.cat, sort: sort.value },
  });
}
watch(
  () => route.query,
  () => {
    params.search = route.query.search || "";
    params.cat = route.query.cat || "";
    params.sort = route.query.sort || "";
  }
);
</script>
