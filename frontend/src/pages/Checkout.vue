<template>
  <div class="container py-8">
    <h1 class="text-2xl font-bold mb-4">Checkout</h1>

    <div v-if="!auth.isLogged" class="p-4 border rounded-xl">
      <p>
        Silakan
        <router-link class="text-brand underline" to="/login"
          >login</router-link
        >
        terlebih dahulu.
      </p>
    </div>

    <form v-else class="grid md:grid-cols-2 gap-6" @submit.prevent="submit">
      <div class="card p-4 space-y-3">
        <h3 class="font-semibold mb-2">Alamat Pengiriman</h3>
        <input
          v-model="addr.label"
          placeholder="Label (Rumah/Kantor)"
          class="border rounded-xl px-3 py-2 w-full"
        />
        <input
          v-model="addr.recipient"
          placeholder="Penerima"
          class="border rounded-xl px-3 py-2 w-full"
          required
        />
        <input
          v-model="addr.phone"
          placeholder="Telepon"
          class="border rounded-xl px-3 py-2 w-full"
          required
        />
        <input
          v-model="addr.line1"
          placeholder="Alamat"
          class="border rounded-xl px-3 py-2 w-full"
          required
        />
        <div class="grid grid-cols-2 gap-3">
          <input
            v-model="addr.city"
            placeholder="Kota"
            class="border rounded-xl px-3 py-2 w-full"
            required
          />
          <input
            v-model="addr.province"
            placeholder="Provinsi"
            class="border rounded-xl px-3 py-2 w-full"
            required
          />
        </div>
        <input
          v-model="addr.postal_code"
          placeholder="Kode Pos"
          class="border rounded-xl px-3 py-2 w-full"
          required
        />
        <input
          v-model="addr.notes"
          placeholder="Catatan (opsional)"
          class="border rounded-xl px-3 py-2 w-full"
        />
      </div>

      <div class="card p-4 space-y-3">
        <h3 class="font-semibold mb-2">Ringkasan</h3>
        <div class="flex justify-between">
          <span>Subtotal</span
          ><span>Rp {{ cart.subtotal.toLocaleString("id-ID") }}</span>
        </div>
        <div class="flex justify-between text-gray-500 text-sm">
          <span>Ongkir</span><span>Rp 0</span>
        </div>
        <hr />
        <div class="flex justify-between font-semibold">
          <span>Total</span
          ><span>Rp {{ cart.subtotal.toLocaleString("id-ID") }}</span>
        </div>

        <div class="grid grid-cols-2 gap-3">
          <select v-model="courier" class="border rounded-xl px-3 py-2">
            <option>Dummy</option>
          </select>
          <select v-model="service" class="border rounded-xl px-3 py-2">
            <option>REG</option>
          </select>
        </div>

        <input
          v-model="note"
          placeholder="Catatan pesanan (opsional)"
          class="border rounded-xl px-3 py-2 w-full"
        />

        <button class="btn btn-primary w-full mt-2" :disabled="placing">
          {{ placing ? "Memproses..." : "Buat Pesanan" }}
        </button>

        <div
          v-if="result"
          class="bg-green-50 border border-green-200 rounded-xl p-3 text-sm"
        >
          <div><b>Checkout ID:</b> {{ result.checkout_group_id }}</div>
          <div><b>Order Codes:</b> {{ result.order_codes.join(", ") }}</div>
        </div>
      </div>
    </form>
  </div>
</template>

<script setup>
import { reactive, ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useAuth } from "../stores/auth";
import { useCartApi } from "../stores/cartApi";
import api from "../lib/api";

const router = useRouter();
const auth = useAuth();
const cart = useCartApi();

const addr = reactive({
  label: "Rumah",
  recipient: "",
  phone: "",
  line1: "",
  city: "",
  province: "",
  postal_code: "",
  notes: "",
});
const courier = ref("Dummy");
const service = ref("REG");
const note = ref("");
const placing = ref(false);
const result = ref(null);

onMounted(async () => {
  if (!auth.isLogged) return;
  await cart.fetch();
});

async function submit() {
  if (!auth.isLogged) return router.push({ name: "login" });
  placing.value = true;
  try {
    const { data } = await api.post("/v1/checkout", {
      address: addr,
      courier: courier.value,
      service: service.value,
      note: note.value || null,
    });
    result.value = data;
    // refresh cart
    await cart.fetch();
  } catch (e) {
    alert(e?.response?.data?.error || e.message);
  } finally {
    placing.value = false;
  }
}
</script>
