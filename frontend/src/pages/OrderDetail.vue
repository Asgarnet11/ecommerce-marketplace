<template>
  <div class="container py-8" v-if="o">
    <div class="mb-4">
      <router-link to="/orders" class="text-sm text-brand underline"
        >← Kembali</router-link
      >
    </div>

    <div class="card p-4 mb-6">
      <div class="flex flex-wrap items-center justify-between gap-3">
        <div>
          <div class="text-sm text-gray-500">Kode Pesanan</div>
          <div class="text-xl font-bold">{{ o.code }}</div>
        </div>
        <div class="text-right">
          <div class="text-sm">
            Status: <b>{{ o.status }}</b> / {{ o.payment_status }}
          </div>
          <div class="text-sm text-gray-500">Toko: {{ o.shop_name }}</div>
        </div>
      </div>
    </div>

    <div class="grid md:grid-cols-3 gap-6">
      <div class="md:col-span-2 space-y-3">
        <div class="card p-4">
          <h3 class="font-semibold mb-2">Item</h3>
          <div class="divide-y">
            <div
              v-for="it in o.items"
              :key="it.product_id"
              class="py-2 flex items-center justify-between"
            >
              <div>
                <div class="font-medium">{{ it.title }}</div>
                <div class="text-sm text-gray-600">
                  Rp {{ it.price.toLocaleString("id-ID") }} × {{ it.qty }}
                </div>
              </div>
              <div class="font-semibold">
                Rp {{ it.subtotal.toLocaleString("id-ID") }}
              </div>
            </div>
          </div>
        </div>

        <div class="card p-4">
          <h3 class="font-semibold mb-2">Pengiriman</h3>
          <div class="text-sm">
            Kurir: <b>{{ o.shipment.courier || "-" }}</b>
            <span v-if="o.shipment.service">({{ o.shipment.service }})</span
            ><br />
            Tracking: {{ o.shipment.tracking || "-" }}<br />
            Status: {{ o.shipment.status || "-" }}
          </div>
        </div>
      </div>

      <div class="space-y-3">
        <div class="card p-4">
          <h3 class="font-semibold mb-2">Alamat</h3>
          <div class="text-sm">
            <div>
              <b>{{ o.address.recipient }}</b> ({{ o.address.phone }})
            </div>
            <div>{{ o.address.line1 }}</div>
            <div>
              {{ o.address.city }}, {{ o.address.province }}
              {{ o.address.postal_code }}
            </div>
            <div v-if="o.address.notes" class="text-gray-500 mt-1">
              {{ o.address.notes }}
            </div>
          </div>
        </div>

        <div class="card p-4">
          <h3 class="font-semibold mb-2">Pembayaran</h3>
          <div class="text-sm">
            Provider: <b>{{ o.payment.provider || "-" }}</b
            ><br />
            Metode: {{ o.payment.method || "-" }}<br />
            Status: {{ o.payment.status || "-" }}
          </div>
        </div>

        <div class="card p-4">
          <h3 class="font-semibold mb-2">Ringkasan</h3>
          <div class="flex justify-between">
            <span>Subtotal</span
            ><span>Rp {{ o.subtotal.toLocaleString("id-ID") }}</span>
          </div>
          <div class="flex justify-between text-gray-500 text-sm">
            <span>Ongkir</span
            ><span>Rp {{ o.shipping_cost.toLocaleString("id-ID") }}</span>
          </div>
          <hr class="my-2" />
          <div class="flex justify-between font-semibold">
            <span>Total</span
            ><span>Rp {{ o.total.toLocaleString("id-ID") }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import api from "../lib/api";
const route = useRoute();
const o = ref(null);
onMounted(async () => {
  const { data } = await api.get(`/v1/orders/${route.params.code}`);
  o.value = data;
});
</script>
