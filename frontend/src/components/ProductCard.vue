<template>
  <div
    class="group relative bg-white rounded-2xl shadow-sm hover:shadow-2xl border border-gray-100 hover:border-gray-200 transition-all duration-500 overflow-hidden"
  >
    <!-- Product Link Wrapper -->
    <router-link :to="`/p/${product.slug}`" class="block relative">
      <!-- Image Container -->
      <div class="relative overflow-hidden rounded-t-2xl">
        <!-- Product Image -->
        <img
          :src="product.image || placeholder"
          :alt="product.title"
          class="w-full aspect-square object-cover group-hover:scale-110 transition-transform duration-700"
          loading="lazy"
        />

        <!-- Image Overlay -->
        <div
          class="absolute inset-0 bg-gradient-to-t from-black/20 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"
        ></div>

        <!-- Discount Badge (if applicable) -->
        <div v-if="product.discount" class="absolute top-3 left-3 z-10">
          <span
            class="inline-flex items-center px-2.5 py-1 rounded-full text-xs font-bold bg-gradient-to-r from-red-500 to-pink-500 text-white shadow-lg"
          >
            -{{ product.discount }}%
          </span>
        </div>

        <!-- Wishlist Button -->
        <button
          class="absolute top-3 right-3 z-10 p-2 bg-white/90 backdrop-blur-sm rounded-full shadow-md hover:bg-white hover:scale-110 transition-all duration-200 opacity-0 group-hover:opacity-100"
        >
          <svg
            class="w-4 h-4 text-gray-600 hover:text-red-500 transition-colors"
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
        </button>

        <!-- Quick View Button -->
        <div
          class="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity duration-300"
        >
          <button
            class="px-4 py-2 bg-white/95 backdrop-blur-sm text-gray-800 rounded-full font-medium shadow-lg hover:bg-white hover:scale-105 transition-all duration-200 flex items-center gap-2"
          >
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
                d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
              />
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
              />
            </svg>
            <span class="text-sm">Lihat Detail</span>
          </button>
        </div>
      </div>

      <!-- Product Info -->
      <div class="p-4 space-y-3">
        <!-- Product Title -->
        <h3
          class="font-semibold text-gray-900 line-clamp-2 leading-tight group-hover:text-blue-600 transition-colors duration-200"
        >
          {{ product.title }}
        </h3>

        <!-- Rating and Reviews (if available) -->
        <div
          v-if="product.rating || product.reviews"
          class="flex items-center gap-2"
        >
          <div class="flex items-center text-yellow-400">
            <svg
              v-for="i in 5"
              :key="i"
              class="w-3.5 h-3.5"
              :class="
                i <= (product.rating || 4) ? 'fill-current' : 'fill-gray-200'
              "
              viewBox="0 0 20 20"
            >
              <path
                d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
              />
            </svg>
          </div>
          <span class="text-sm text-gray-500"
            >({{ product.reviews || 127 }})</span
          >
        </div>

        <!-- Price Section -->
        <div class="space-y-1">
          <!-- Original Price (if discounted) -->
          <div
            v-if="product.originalPrice"
            class="text-sm text-gray-400 line-through"
          >
            Rp {{ product.originalPrice.toLocaleString("id-ID") }}
          </div>

          <!-- Current Price -->
          <div class="flex items-center justify-between">
            <div class="text-lg font-bold text-gray-900">
              Rp {{ product.price.toLocaleString("id-ID") }}
            </div>

            <!-- Stock Indicator -->
            <div
              v-if="product.stock !== undefined"
              class="text-xs text-gray-500"
            >
              <span v-if="product.stock > 10" class="text-green-600"
                >Stok: {{ product.stock }}</span
              >
              <span v-else-if="product.stock > 0" class="text-orange-600"
                >Stok: {{ product.stock }}</span
              >
              <span v-else class="text-red-600">Habis</span>
            </div>
          </div>
        </div>

        <!-- Shop Info (if available) -->
        <div
          v-if="product.shop"
          class="flex items-center gap-2 text-sm text-gray-500"
        >
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
              d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5"
            />
          </svg>
          <span class="truncate">{{ product.shop }}</span>
        </div>
      </div>
    </router-link>

    <!-- Add to Cart Button -->
    <div class="p-4 pt-0">
      <button
        class="w-full group/btn relative overflow-hidden bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 text-white font-semibold py-3 px-4 rounded-xl transition-all duration-300 hover:shadow-lg hover:scale-[1.02] disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2"
        @click="$emit('add', product)"
        :disabled="product.stock === 0"
      >
        <!-- Loading/Success Animation Background -->
        <div
          class="absolute inset-0 bg-gradient-to-r from-green-500 to-emerald-600 translate-y-full group-active/btn:translate-y-0 transition-transform duration-300"
        ></div>

        <!-- Button Content -->
        <div class="relative flex items-center gap-2">
          <svg
            class="w-4 h-4 group-hover/btn:scale-110 transition-transform duration-200"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-1.5 6M7 13l-1.5 6m9.5-6h4.5"
            />
          </svg>
          <span>{{
            product.stock === 0 ? "Stok Habis" : "Tambah ke Keranjang"
          }}</span>
        </div>
      </button>
    </div>

    <!-- Hover Effect Gradient Border -->
    <div
      class="absolute inset-0 rounded-2xl bg-gradient-to-r from-blue-500/20 via-purple-500/20 to-pink-500/20 opacity-0 group-hover:opacity-100 transition-opacity duration-500 -z-10 blur-xl"
    ></div>
  </div>
</template>

<script setup>
const placeholder = "https://placehold.co/600x600?text=Produk";

defineProps({
  product: {
    type: Object,
    required: true,
  },
});

// Emit events
defineEmits(["add"]);
</script>

<style scoped>
/* Line clamp utility */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Custom scrollbar for better mobile experience */
@media (max-width: 640px) {
  .group {
    transform: translateZ(0);
  }
}

/* Improve touch targets on mobile */
@media (max-width: 768px) {
  .group:hover .opacity-0 {
    opacity: 1;
  }

  .group .group-hover\:opacity-100 {
    opacity: 1;
  }
}

/* Ensure proper spacing on different screen sizes */
.group {
  min-height: fit-content;
}

/* Animation performance optimization */
.group * {
  will-change: transform, opacity;
}

.group:hover * {
  will-change: auto;
}
</style>
