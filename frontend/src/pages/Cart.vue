<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Keranjang</h1>
    <div v-if="items.length === 0">Keranjang kosong.</div>
    <div v-else class="space-y-3">
      <div
        v-for="it in items"
        :key="it.id"
        class="bg-white p-3 rounded flex items-center justify-between"
      >
        <div>
          <div class="font-medium">{{ it.title }}</div>
          <div>
            Rp {{ it.price.toLocaleString("id-ID") }} ×
            <input
              class="border rounded w-16 text-center ml-2"
              type="number"
              v-model.number="it.qty"
              min="1"
            />
          </div>
        </div>
        <button @click="remove(it.id)">Hapus</button>
      </div>
      <div class="text-right font-semibold">
        Subtotal: Rp {{ subtotal.toLocaleString("id-ID") }}
      </div>
    </div>
  </div>
</template>
<script setup>
import { storeToRefs } from "pinia";
import { useCart } from "../stores/cart";
const cart = useCart();
const { items, subtotal } = storeToRefs(cart);
function remove(id) {
  cart.remove(id);
}
</script>
