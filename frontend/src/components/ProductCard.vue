<template>
  <div
    class="bg-white rounded-2xl shadow-sm hover:shadow-xl transition-all duration-300 overflow-hidden group border border-gray-100 hover:border-green-200"
  >
    <!-- Product Link -->
    <router-link :to="`/p/${product.slug}`" class="block">
      <!-- Image Container -->
      <div
        class="relative aspect-square bg-gray-100 rounded-t-2xl overflow-hidden"
      >
        <!-- Product Image -->
        <img
          v-if="product.image"
          :src="product.image"
          :alt="product.title"
          class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
          loading="lazy"
        />
        <!-- Placeholder if no image -->
        <div
          v-else
          class="w-full h-full bg-gradient-to-br from-gray-200 to-gray-300 flex items-center justify-center"
        >
          <svg
            class="w-16 h-16 text-gray-400"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="1.5"
              d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
            />
          </svg>
        </div>

        <!-- Badges -->
        <div class="absolute top-3 left-3 flex flex-col gap-1">
          <!-- Discount Badge -->
          <div
            v-if="product.discount"
            class="bg-red-500 text-white text-xs font-bold px-2 py-1 rounded-full"
          >
            -{{ product.discount }}%
          </div>

          <!-- New Badge -->
          <div
            v-if="product.isNew"
            class="bg-green-500 text-white text-xs font-bold px-2 py-1 rounded-full"
          >
            BARU
          </div>

          <!-- Popular Badge -->
          <div
            v-if="product.isPopular"
            class="bg-purple-500 text-white text-xs font-bold px-2 py-1 rounded-full"
          >
            POPULER
          </div>
        </div>

        <!-- Wishlist Button -->
        <button
          @click.prevent="toggleWishlist"
          class="absolute top-3 right-3 w-8 h-8 bg-white bg-opacity-80 hover:bg-opacity-100 rounded-full flex items-center justify-center transition-all duration-300 transform hover:scale-110"
          :class="{
            'text-red-500': isInWishlist,
            'text-gray-400 hover:text-red-500': !isInWishlist,
          }"
        >
          <svg
            class="w-4 h-4"
            :fill="isInWishlist ? 'currentColor' : 'none'"
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
      </div>

      <!-- Product Info -->
      <div class="p-4">
        <!-- Store/Brand -->
        <div
          v-if="product.store"
          class="flex items-center text-xs text-gray-500 mb-2"
        >
          <svg
            class="w-3 h-3 mr-1"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-4m-5 0H3m2 0h4m0 0v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"
            />
          </svg>
          {{ product.store }}
        </div>

        <!-- Title -->
        <h3
          class="font-medium text-gray-800 line-clamp-2 group-hover:text-green-600 transition-colors duration-200 mb-2 text-sm leading-tight"
        >
          {{ product.title }}
        </h3>

        <!-- Rating -->
        <div v-if="product.rating" class="flex items-center mb-3">
          <div class="flex items-center mr-2">
            <svg
              v-for="star in 5"
              :key="star"
              class="w-3 h-3"
              :class="
                star <= product.rating
                  ? 'text-yellow-400 fill-current'
                  : 'text-gray-300'
              "
              viewBox="0 0 24 24"
            >
              <path
                d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"
              />
            </svg>
          </div>
          <span class="text-xs text-gray-500">
            {{ product.rating }} ({{ product.reviewCount || 0 }})
          </span>
        </div>

        <!-- Price Section -->
        <div class="space-y-1">
          <!-- Original Price (if discounted) -->
          <div
            v-if="
              product.originalPrice && product.originalPrice > product.price
            "
            class="text-xs text-gray-400 line-through"
          >
            Rp {{ product.originalPrice.toLocaleString("id-ID") }}
          </div>

          <!-- Current Price -->
          <div class="font-bold text-lg text-gray-900">
            Rp {{ product.price.toLocaleString("id-ID") }}
          </div>
        </div>

        <!-- Stock Info -->
        <div v-if="product.stock !== undefined" class="mt-2">
          <div
            v-if="product.stock > 10"
            class="text-xs text-green-600 font-medium"
          >
            Stok tersedia
          </div>
          <div
            v-else-if="product.stock > 0"
            class="text-xs text-orange-600 font-medium"
          >
            Stok terbatas ({{ product.stock }})
          </div>
          <div v-else class="text-xs text-red-600 font-medium">Stok habis</div>
        </div>

        <!-- Shipping Info -->
        <div
          v-if="product.freeShipping"
          class="mt-2 flex items-center text-xs text-green-600"
        >
          <svg
            class="w-3 h-3 mr-1"
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
          Gratis Ongkir
        </div>
      </div>
    </router-link>

    <!-- Action Buttons -->
    <div class="px-4 pb-4 space-y-2">
      <!-- Add to Cart Button -->
      <button
        @click="handleAddToCart"
        :disabled="isOutOfStock || isLoading"
        class="w-full py-2.5 rounded-xl font-medium text-sm transition-all duration-300 transform hover:scale-[0.98] disabled:cursor-not-allowed"
        :class="{
          'bg-green-600 hover:bg-green-700 text-white shadow-lg hover:shadow-xl':
            !isOutOfStock && !isLoading,
          'bg-gray-300 text-gray-500': isOutOfStock,
          'bg-green-400 text-white cursor-not-allowed': isLoading,
        }"
      >
        <span v-if="isLoading" class="flex items-center justify-center">
          <svg
            class="animate-spin -ml-1 mr-2 h-4 w-4 text-white"
            fill="none"
            viewBox="0 0 24 24"
          >
            <circle
              class="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              stroke-width="4"
            ></circle>
            <path
              class="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
            ></path>
          </svg>
          Menambahkan...
        </span>
        <span v-else-if="isOutOfStock"> Stok Habis </span>
        <span v-else class="flex items-center justify-center">
          <svg
            class="w-4 h-4 mr-2"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m0 0h8m-8 0a2 2 0 100 4 2 2 0 000-4zm8 0a2 2 0 100 4 2 2 0 000-4z"
            />
          </svg>
          Tambah ke Keranjang
        </span>
      </button>

      <!-- Quick Actions -->
      <div class="flex gap-2">
        <!-- Buy Now Button -->
        <button
          @click="handleBuyNow"
          :disabled="isOutOfStock"
          class="flex-1 py-2 border border-green-600 text-green-600 hover:bg-green-50 rounded-xl font-medium text-sm transition-all duration-300 disabled:border-gray-300 disabled:text-gray-400 disabled:cursor-not-allowed"
        >
          Beli Sekarang
        </button>

        <!-- Quick View Button -->
        <button
          @click="handleQuickView"
          class="px-3 py-2 border border-gray-300 text-gray-600 hover:bg-gray-50 rounded-xl transition-all duration-300"
          title="Quick View"
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
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";

// Props
const props = defineProps({
  product: {
    type: Object,
    required: true,
  },
});

// Emits
const emit = defineEmits(["add", "wishlist", "quickView", "buyNow"]);

// State
const isLoading = ref(false);
const isInWishlist = ref(false);

// Computed
const isOutOfStock = computed(() => {
  return props.product.stock !== undefined && props.product.stock === 0;
});

// Methods
const handleAddToCart = async () => {
  if (isOutOfStock.value) return;

  isLoading.value = true;
  try {
    // Simulate loading time
    await new Promise((resolve) => setTimeout(resolve, 500));
    emit("add", props.product);
  } finally {
    isLoading.value = false;
  }
};

const toggleWishlist = () => {
  isInWishlist.value = !isInWishlist.value;
  emit("wishlist", {
    product: props.product,
    isAdded: isInWishlist.value,
  });
};

const handleBuyNow = () => {
  if (isOutOfStock.value) return;
  emit("buyNow", props.product);
};

const handleQuickView = () => {
  emit("quickView", props.product);
};
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
