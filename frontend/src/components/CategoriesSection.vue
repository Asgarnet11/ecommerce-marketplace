<template>
  <section class="py-16 bg-gray-50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Section Header -->
      <div class="text-center mb-12">
        <h2 class="text-3xl lg:text-4xl font-bold text-gray-800 mb-4">
          Kategori Populer
        </h2>
        <p class="text-gray-600 max-w-2xl mx-auto">
          Jelajahi berbagai kategori produk pilihan dengan penawaran terbaik
        </p>
      </div>

      <!-- Categories Grid -->
      <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-6">
        <div
          v-for="category in categories"
          :key="category.id"
          class="group cursor-pointer"
          @click="selectCategory(category)"
        >
          <div
            class="bg-white rounded-2xl p-6 shadow-sm hover:shadow-lg transition-all duration-300 transform hover:-translate-y-2 text-center"
          >
            <!-- Icon Container -->
            <div
              class="w-16 h-16 mx-auto mb-4 rounded-full flex items-center justify-center group-hover:scale-110 transition-transform duration-300"
              :class="`bg-gradient-to-r ${category.gradient}`"
            >
              <i :class="category.icon" class="text-2xl text-white"></i>
            </div>

            <!-- Category Name -->
            <h3
              class="font-semibold text-gray-800 mb-2 group-hover:text-green-600 transition-colors"
            >
              {{ category.name }}
            </h3>

            <!-- Product Count -->
            <p class="text-sm text-gray-500">{{ category.count }} produk</p>

            <!-- Hover Effect Badge -->
            <div
              class="mt-3 opacity-0 group-hover:opacity-100 transition-opacity duration-300"
            >
              <span
                class="inline-block bg-green-100 text-green-800 text-xs px-2 py-1 rounded-full"
              >
                Lihat Semua
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- View All Categories Button -->
      <div class="text-center mt-12">
        <button
          class="bg-green-600 hover:bg-green-700 text-white font-semibold px-8 py-3 rounded-lg transition-colors duration-300 inline-flex items-center"
        >
          Lihat Semua Kategori
          <svg
            class="w-5 h-5 ml-2"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5l7 7-7 7"
            ></path>
          </svg>
        </button>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from "vue";

const categories = ref([
  {
    id: 1,
    name: "Fashion",
    count: "15.2k",
    icon: "👕",
    gradient: "from-pink-500 to-rose-500",
  },
  {
    id: 2,
    name: "Elektronik",
    count: "8.7k",
    icon: "📱",
    gradient: "from-blue-500 to-cyan-500",
  },
  {
    id: 3,
    name: "Makanan",
    count: "12.1k",
    icon: "🍕",
    gradient: "from-orange-500 to-yellow-500",
  },
  {
    id: 4,
    name: "Kesehatan",
    count: "5.3k",
    icon: "💊",
    gradient: "from-green-500 to-emerald-500",
  },
  {
    id: 5,
    name: "Olahraga",
    count: "7.8k",
    icon: "⚽",
    gradient: "from-purple-500 to-indigo-500",
  },
  {
    id: 6,
    name: "Buku",
    count: "9.4k",
    icon: "📚",
    gradient: "from-red-500 to-pink-500",
  },
  {
    id: 7,
    name: "Rumah",
    count: "11.6k",
    icon: "🏠",
    gradient: "from-teal-500 to-cyan-500",
  },
  {
    id: 8,
    name: "Otomotif",
    count: "4.9k",
    icon: "🚗",
    gradient: "from-gray-600 to-gray-800",
  },
  {
    id: 9,
    name: "Kecantikan",
    count: "13.2k",
    icon: "💄",
    gradient: "from-pink-400 to-purple-500",
  },
  {
    id: 10,
    name: "Mainan",
    count: "6.7k",
    icon: "🧸",
    gradient: "from-yellow-400 to-orange-500",
  },
  {
    id: 11,
    name: "Gaming",
    count: "8.1k",
    icon: "🎮",
    gradient: "from-indigo-500 to-purple-600",
  },
  {
    id: 12,
    name: "Lainnya",
    count: "3.5k",
    icon: "📦",
    gradient: "from-gray-500 to-gray-600",
  },
]);

const selectCategory = (category) => {
  console.log("Selected category:", category);
  // Handle category selection logic here
};
</script>

<style scoped>
/* Custom animations for category cards */
.group:hover .bg-gradient-to-r {
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%,
  100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}
</style>
