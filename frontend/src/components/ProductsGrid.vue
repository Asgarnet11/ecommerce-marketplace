<template>
  <div class="w-full">
    <!-- Loading State -->
    <div
      v-if="isLoading"
      class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4"
    >
      <div
        v-for="n in 10"
        :key="n"
        class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden animate-pulse"
      >
        <div class="aspect-square bg-gray-200"></div>
        <div class="p-4 space-y-3">
          <div class="h-4 bg-gray-200 rounded w-3/4"></div>
          <div class="h-4 bg-gray-200 rounded w-1/2"></div>
          <div class="h-6 bg-gray-200 rounded w-2/3"></div>
          <div class="h-8 bg-gray-200 rounded"></div>
        </div>
      </div>
    </div>

    <!-- Products Grid -->
    <div
      v-else-if="products.length > 0"
      class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4"
    >
      <ProductCard
        v-for="product in products"
        :key="product.id"
        :product="product"
        @add="handleAddToCart"
        @wishlist="handleWishlist"
        @quickView="handleQuickView"
        @buyNow="handleBuyNow"
      />
    </div>

    <!-- Empty State -->
    <div v-else class="text-center py-16">
      <svg
        class="mx-auto h-24 w-24 text-gray-300 mb-4"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="1"
          d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2 2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 009.586 13H7"
        />
      </svg>
      <h3 class="text-lg font-medium text-gray-800 mb-2">Tidak ada produk</h3>
      <p class="text-gray-500">Belum ada produk yang tersedia saat ini.</p>
    </div>

    <!-- Success Toast -->
    <Transition name="toast">
      <div
        v-if="showToast"
        class="fixed top-4 right-4 bg-green-600 text-white px-6 py-3 rounded-lg shadow-lg z-50 flex items-center"
      >
        <svg
          class="w-5 h-5 mr-2"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
          />
        </svg>
        {{ toastMessage }}
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import ProductCard from "./ProductCard.vue";

// Props
const props = defineProps({
  fetchProducts: {
    type: Function,
    default: null,
  },
});

// State
const products = ref([]);
const isLoading = ref(false);
const showToast = ref(false);
const toastMessage = ref("");

// Sample product data for demo
const sampleProducts = [
  {
    id: 1,
    slug: "smartphone-samsung-galaxy",
    title: "Samsung Galaxy S24 Ultra 256GB",
    price: 18999000,
    originalPrice: 21999000,
    discount: 14,
    rating: 4.8,
    reviewCount: 234,
    image:
      "https://images.unsplash.com/photo-1511707171634-5f897ff02aa9?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
    store: "Samsung Official Store",
    stock: 25,
    freeShipping: true,
    isPopular: true,
  },
  {
    id: 2,
    slug: "laptop-macbook-air",
    title: 'MacBook Air 13" M3 Chip 8GB RAM 256GB SSD',
    price: 18999000,
    rating: 4.9,
    reviewCount: 156,
    image:
      "https://images.unsplash.com/photo-1517336714731-489689fd1ca8?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
    store: "Apple Authorized Reseller",
    stock: 12,
    freeShipping: true,
    isNew: true,
  },
  {
    id: 3,
    slug: "sepatu-nike-air-max",
    title: "Nike Air Max 270 React Sneakers Pria",
    price: 1299000,
    originalPrice: 1599000,
    discount: 19,
    rating: 4.6,
    reviewCount: 89,
    image:
      "https://images.unsplash.com/photo-1542291026-7eec264c27ff?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
    store: "Nike Official Store",
    stock: 45,
    freeShipping: true,
  },
  {
    id: 4,
    slug: "tas-backpack-travel",
    title: "Backpack Travel Waterproof 40L dengan USB Port",
    price: 299000,
    originalPrice: 399000,
    discount: 25,
    rating: 4.4,
    reviewCount: 67,
    image:
      "https://images.unsplash.com/photo-1553062407-98eeb64c6a62?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
    store: "Adventure Gear",
    stock: 8,
    freeShipping: false,
  },
  {
    id: 5,
    slug: "kamera-canon-eos",
    title: "Canon EOS R50 Mirrorless Camera dengan Lensa Kit",
    price: 8999000,
    rating: 4.7,
    reviewCount: 45,
    image:
      "https://images.unsplash.com/photo-1516035069371-29a1b244cc32?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
    store: "Canon Official Store",
    stock: 15,
    freeShipping: true,
    isNew: true,
  },
  {
    id: 6,
    slug: "jam-tangan-apple-watch",
    title: "Apple Watch Series 9 GPS 45mm Sport Band",
    price: 5999000,
    originalPrice: 6999000,
    discount: 14,
    rating: 4.8,
    reviewCount: 123,
    image:
      "https://images.unsplash.com/photo-1434493789847-2f02dc6ca35d?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
    store: "Apple Authorized Reseller",
    stock: 0,
    freeShipping: true,
    isPopular: true,
  },
  {
    id: 7,
    slug: "headphone-sony-wh1000xm5",
    title: "Sony WH-1000XM5 Wireless Noise Canceling Headphones",
    price: 4999000,
    rating: 4.9,
    reviewCount: 78,
    image:
      "https://images.unsplash.com/photo-1505740420928-5e560c06d30e?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
    store: "Sony Official Store",
    stock: 20,
    freeShipping: true,
  },
  {
    id: 8,
    slug: "keyboard-mechanical-gaming",
    title: "Mechanical Gaming Keyboard RGB Cherry MX Blue",
    price: 899000,
    originalPrice: 1199000,
    discount: 25,
    rating: 4.5,
    reviewCount: 92,
    image:
      "https://images.unsplash.com/photo-1541140532154-b024d705b90a?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
    store: "Gaming Pro",
    stock: 35,
    freeShipping: false,
  },
];

// Methods
const loadProducts = async () => {
  isLoading.value = true;
  try {
    if (props.fetchProducts) {
      products.value = await props.fetchProducts();
    } else {
      // Simulate API call
      await new Promise((resolve) => setTimeout(resolve, 1000));
      products.value = sampleProducts;
    }
  } catch (error) {
    console.error("Error loading products:", error);
  } finally {
    isLoading.value = false;
  }
};

const handleAddToCart = (product) => {
  showToast.value = true;
  toastMessage.value = `${product.title} ditambahkan ke keranjang`;

  // Hide toast after 3 seconds
  setTimeout(() => {
    showToast.value = false;
  }, 3000);

  console.log("Added to cart:", product);
};

const handleWishlist = (data) => {
  const { product, isAdded } = data;
  showToast.value = true;
  toastMessage.value = isAdded
    ? `${product.title} ditambahkan ke wishlist`
    : `${product.title} dihapus dari wishlist`;

  setTimeout(() => {
    showToast.value = false;
  }, 3000);

  console.log("Wishlist action:", data);
};

const handleQuickView = (product) => {
  console.log("Quick view:", product);
  // Implement quick view modal logic here
};

const handleBuyNow = (product) => {
  console.log("Buy now:", product);
  // Redirect to checkout or implement buy now logic
};

// Lifecycle
onMounted(() => {
  loadProducts();
});
</script>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}

.toast-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(100%);
}
</style>
