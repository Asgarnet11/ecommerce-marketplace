<template>
  <div v-if="p" class="min-h-screen bg-gray-50">
    <!-- Breadcrumb Navigation -->
    <div class="bg-white border-b">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <nav class="flex items-center space-x-2 text-sm text-gray-500">
          <router-link to="/" class="hover:text-gray-700 transition-colors">
            <svg
              class="w-4 h-4 inline mr-1"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
              />
            </svg>
            Beranda
          </router-link>
          <svg
            class="w-4 h-4"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5l7 7-7 7"
            />
          </svg>
          <span class="text-gray-700">Detail Produk</span>
        </nav>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 lg:gap-12">
        <!-- Product Image -->
        <div class="space-y-4">
          <div class="relative group">
            <img
              :src="p.image || placeholder"
              alt=""
              class="w-full aspect-square rounded-2xl object-cover bg-gray-100 shadow-lg hover:shadow-xl transition-shadow duration-300"
            />
            <!-- Zoom Icon Overlay -->
            <div
              class="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-10 rounded-2xl transition-all duration-300 flex items-center justify-center"
            >
              <svg
                class="w-8 h-8 text-white opacity-0 group-hover:opacity-100 transition-opacity duration-300"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"
                />
              </svg>
            </div>
          </div>
        </div>

        <!-- Product Info -->
        <div class="space-y-6">
          <!-- Title and Price -->
          <div class="space-y-4">
            <h1
              class="text-3xl lg:text-4xl font-bold text-gray-900 leading-tight"
            >
              {{ p.title }}
            </h1>
            <div class="flex items-baseline space-x-2">
              <span class="text-3xl lg:text-4xl font-bold text-blue-600">
                Rp {{ p.price.toLocaleString("id-ID") }}
              </span>
              <span class="text-lg text-gray-500 line-through">
                Rp {{ (p.price * 1.2).toLocaleString("id-ID") }}
              </span>
              <span
                class="bg-red-100 text-red-800 text-sm font-medium px-2.5 py-0.5 rounded-full"
              >
                -17%
              </span>
            </div>
          </div>

          <!-- Stock Status -->
          <div class="flex items-center space-x-3">
            <div class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-2"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-sm font-medium text-gray-700"
                >Stok tersedia:</span
              >
            </div>
            <div class="flex items-center">
              <span class="text-lg font-semibold text-gray-900">{{
                p.stock
              }}</span>
              <span class="text-sm text-gray-500 ml-1">unit</span>
            </div>
          </div>

          <!-- Product Features -->
          <div class="bg-gray-50 rounded-xl p-6 space-y-3">
            <h3 class="font-semibold text-gray-900 flex items-center">
              <svg
                class="w-5 h-5 text-blue-500 mr-2"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z"
                />
              </svg>
              Keuntungan Produk
            </h3>
            <ul class="space-y-2 text-sm text-gray-600">
              <li class="flex items-center">
                <svg
                  class="w-4 h-4 text-green-500 mr-2 flex-shrink-0"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 13l4 4L19 7"
                  />
                </svg>
                Kualitas premium terjamin
              </li>
              <li class="flex items-center">
                <svg
                  class="w-4 h-4 text-green-500 mr-2 flex-shrink-0"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 13l4 4L19 7"
                  />
                </svg>
                Gratis pengiriman ke seluruh Indonesia
              </li>
              <li class="flex items-center">
                <svg
                  class="w-4 h-4 text-green-500 mr-2 flex-shrink-0"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 13l4 4L19 7"
                  />
                </svg>
                Garansi 30 hari uang kembali
              </li>
            </ul>
          </div>

          <!-- Action Buttons -->
          <div class="space-y-4">
            <!-- Add to Cart Button -->
            <button
              class="w-full bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-semibold py-4 px-6 rounded-xl transition-all duration-200 transform hover:scale-[1.02] focus:ring-4 focus:ring-blue-200 shadow-lg flex items-center justify-center space-x-2"
              @click="add()"
            >
              <svg
                class="w-5 h-5"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m0 0H17M9 19a2 2 0 11-4 0 2 2 0 014 0zm8 0a2 2 0 11-4 0 2 2 0 014 0z"
                />
              </svg>
              <span>Tambah ke Keranjang</span>
            </button>

            <!-- Additional Actions -->
            <div class="grid grid-cols-2 gap-4">
              <button
                class="flex items-center justify-center px-4 py-3 border-2 border-gray-300 rounded-xl hover:border-gray-400 transition-colors duration-200"
              >
                <svg
                  class="w-5 h-5 text-gray-600 mr-2"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
                  />
                </svg>
                <span class="font-medium text-gray-700">Wishlist</span>
              </button>
              <button
                class="flex items-center justify-center px-4 py-3 border-2 border-gray-300 rounded-xl hover:border-gray-400 transition-colors duration-200"
              >
                <svg
                  class="w-5 h-5 text-gray-600 mr-2"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z"
                  />
                </svg>
                <span class="font-medium text-gray-700">Bagikan</span>
              </button>
            </div>
          </div>

          <!-- Trust Badges -->
          <div class="border-t pt-6">
            <div class="grid grid-cols-3 gap-4 text-center">
              <div class="flex flex-col items-center">
                <div
                  class="w-12 h-12 bg-blue-100 rounded-full flex items-center justify-center mb-2"
                >
                  <svg
                    class="w-6 h-6 text-blue-600"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"
                    />
                  </svg>
                </div>
                <span class="text-xs font-medium text-gray-600"
                  >Pengiriman Cepat</span
                >
              </div>
              <div class="flex flex-col items-center">
                <div
                  class="w-12 h-12 bg-green-100 rounded-full flex items-center justify-center mb-2"
                >
                  <svg
                    class="w-6 h-6 text-green-600"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"
                    />
                  </svg>
                </div>
                <span class="text-xs font-medium text-gray-600"
                  >Terpercaya</span
                >
              </div>
              <div class="flex flex-col items-center">
                <div
                  class="w-12 h-12 bg-purple-100 rounded-full flex items-center justify-center mb-2"
                >
                  <svg
                    class="w-6 h-6 text-purple-600"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"
                    />
                  </svg>
                </div>
                <span class="text-xs font-medium text-gray-600"
                  >Pembayaran Aman</span
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Loading State -->
  <div v-else class="min-h-screen bg-gray-50 flex items-center justify-center">
    <div class="text-center">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto"
      ></div>
      <p class="mt-4 text-gray-600">Memuat detail produk...</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import api from "../lib/api";
import { useCart } from "../stores/cart";

const route = useRoute();
const cart = useCart();
const p = ref(null);
const placeholder = "https://placehold.co/600x600?text=Produk";

async function load() {
  const { data } = await api.get(`/v1/products/${route.params.slug}`);
  p.value = data;
}
function add() {
  if (p.value) cart.add(p.value);
}
onMounted(load);
</script>
