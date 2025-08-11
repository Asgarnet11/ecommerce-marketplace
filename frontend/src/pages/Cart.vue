<template>
  <div class="container py-8">
    <h1 class="text-2xl font-bold mb-4">Keranjang</h1>

    <div v-if="!auth.isLogged" class="p-4 border rounded-xl">
      <p>
        Silakan
        <router-link class="text-brand underline" to="/login"
          >login</router-link
        >
        untuk melihat keranjang.
      </p>
    </div>

    <div v-else>
      <div v-if="cart.loading">Loading...</div>
      <div v-else-if="cart.items.length === 0">Keranjang kosong.</div>
      <div v-else class="space-y-3">
        <div
          v-for="it in cart.items"
          :key="it.product_id"
          class="card p-3 flex items-center justify-between"
        >
          <div>
            <div class="font-medium">{{ it.title }}</div>
            <div class="text-sm text-gray-600">
              Rp {{ it.price.toLocaleString("id-ID") }} ×
              <input
                type="number"
                min="1"
                class="border rounded w-16 text-center ml-2"
                :value="it.qty"
                @change="onQty(it.product_id, $event.target.value)"
              />
            </div>
          </div>
          <div class="flex items-center gap-4">
            <div class="font-semibold">
              Rp {{ it.subtotal.toLocaleString("id-ID") }}
            </div>
            <button class="btn" @click="cart.remove(it.product_id)">
              Hapus
            </button>
          </div>
        </div>
        <div class="text-right font-semibold text-lg">
          Subtotal: Rp {{ cart.subtotal.toLocaleString("id-ID") }}
        </div>
        <div class="text-right">
          <router-link class="btn btn-primary" to="/checkout"
            >Lanjutkan ke Checkout</router-link
          >
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from "vue";
import { useCartApi } from "../stores/cartApi";
import { useAuth } from "../stores/auth";

const cart = useCartApi();
const auth = useAuth();

onMounted(() => {
  if (auth.isLogged) cart.fetch();
});
function onQty(pid, val) {
  const qty = parseInt(val, 10) || 1;
  cart.update(pid, qty);
}
</script>
