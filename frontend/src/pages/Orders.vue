<template>
  <div class="container py-8">
    <h1 class="text-2xl font-bold mb-4">Pesanan Saya</h1>

    <div class="flex gap-2 items-center mb-4">
      <select v-model="status" class="border rounded-xl px-3 py-2">
        <option value="">Semua Status</option>
        <option value="pending">Pending</option>
        <option value="paid">Paid</option>
        <option value="processing">Processing</option>
        <option value="shipped">Shipped</option>
        <option value="delivered">Delivered</option>
        <option value="cancelled">Cancelled</option>
        <option value="refunded">Refunded</option>
      </select>
      <button
        class="btn"
        @click="
          page = 1;
          load();
        "
      >
        Terapkan
      </button>
    </div>

    <div v-if="loading">Loading...</div>
    <div v-else-if="items.length === 0" class="text-gray-500">
      Belum ada pesanan.
    </div>
    <div v-else class="space-y-3">
      <div
        v-for="o in items"
        :key="o.code"
        class="card p-4 flex flex-col md:flex-row md:items-center md:justify-between gap-3"
      >
        <div>
          <div class="text-sm text-gray-500">
            Kode:
            <router-link class="underline" :to="`/orders/${o.code}`">{{
              o.code
            }}</router-link>
          </div>
          <div class="font-semibold">{{ o.shop_name }}</div>
          <div class="text-sm text-gray-600">{{ o.created_at }}</div>
        </div>
        <div class="flex items-center gap-6">
          <div class="text-right">
            <div class="text-sm">
              Status: <b>{{ o.status }}</b> / {{ o.payment_status }}
            </div>
            <div class="font-semibold">
              Total: Rp {{ o.total.toLocaleString("id-ID") }}
            </div>
          </div>
          <router-link class="btn" :to="`/orders/${o.code}`"
            >Detail</router-link
          >
        </div>
      </div>

      <div class="flex gap-3 justify-end">
        <button
          class="btn"
          :disabled="page <= 1"
          @click="
            page--;
            load();
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
            load();
          "
        >
          Next
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import api from "../lib/api";
const items = ref([]),
  total = ref(0),
  page = ref(1),
  per = ref(10),
  status = ref(""),
  loading = ref(false);

async function load() {
  loading.value = true;
  try {
    const { data } = await api.get("/v1/orders", {
      params: { page: page.value, per_page: per.value, status: status.value },
    });
    items.value = data.data || [];
    total.value = data.total || 0;
  } finally {
    loading.value = false;
  }
}
onMounted(load);
</script>
