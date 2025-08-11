<template>
  <div>
    <div v-if="loading">Loading...</div>
    <div v-else class="grid grid-cols-2 md:grid-cols-4 gap-4">
      <ProductCard v-for="p in items" :key="p.id" :product="p" @add="add" />
    </div>
    <div class="flex gap-3 justify-end mt-4">
      <button
        class="btn"
        :disabled="page <= 1"
        @click="
          page--;
          fetcher();
        "
      >
        Prev
      </button>
      <span class="text-sm">Page {{ page }}</span>
      <button
        class="btn"
        :disabled="page * per >= total"
        @click="
          page++;
          fetcher();
        "
      >
        Next
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import api from "../lib/api";
import ProductCard from "./ProductCard.vue";
import { useCartApi } from "../stores/cartApi";

const props = defineProps({
  query: { type: Object, default: () => ({}) },
  initialFetch: { type: Boolean, default: false },
});

const items = ref([]),
  total = ref(0),
  page = ref(1),
  per = ref(20),
  loading = ref(false);
const cart = useCartApi();

async function fetcher() {
  loading.value = true;
  try {
    const { data } = await api.get("/v1/products", {
      params: { ...props.query, page: page.value, per_page: per.value },
    });
    items.value = data.data || [];
    total.value = data.total || 0;
  } finally {
    loading.value = false;
  }
}
function add(p) {
  cart.add(p.id, 1);
}

watch(
  () => props.query,
  () => {
    page.value = 1;
    fetcher();
  },
  { deep: true }
);
onMounted(() => {
  if (props.initialFetch) fetcher();
});
</script>
