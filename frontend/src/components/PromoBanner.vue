<template>
  <section
    class="py-16 bg-gradient-to-r from-purple-600 via-pink-600 to-red-600 relative overflow-hidden"
  >
    <!-- Background Pattern -->
    <div class="absolute inset-0 opacity-10">
      <div
        class="absolute top-0 left-0 w-40 h-40 bg-white rounded-full -translate-x-20 -translate-y-20"
      ></div>
      <div
        class="absolute top-1/2 right-0 w-60 h-60 bg-white rounded-full translate-x-30 -translate-y-30"
      ></div>
      <div
        class="absolute bottom-0 left-1/3 w-32 h-32 bg-white rounded-full translate-y-16"
      ></div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative">
      <!-- Flash Sale Banner -->
      <div class="text-center mb-12">
        <div
          class="inline-flex items-center bg-yellow-400 text-black px-6 py-2 rounded-full font-bold text-lg mb-4 animate-bounce"
        >
          ⚡ FLASH SALE ⚡
        </div>
        <h2 class="text-4xl lg:text-6xl font-bold text-white mb-4">
          Diskon Hingga
          <span class="text-yellow-300 animate-pulse">70%</span>
        </h2>
        <p class="text-xl text-white opacity-90 mb-8">
          Penawaran terbatas! Berakhir dalam:
        </p>

        <!-- Countdown Timer -->
        <div class="flex justify-center space-x-4 mb-8">
          <div
            class="bg-white bg-opacity-20 rounded-lg p-4 text-white text-center min-w-[80px]"
          >
            <div class="text-2xl lg:text-3xl font-bold">
              {{ countdown.hours }}
            </div>
            <div class="text-sm opacity-75">Jam</div>
          </div>
          <div
            class="bg-white bg-opacity-20 rounded-lg p-4 text-white text-center min-w-[80px]"
          >
            <div class="text-2xl lg:text-3xl font-bold">
              {{ countdown.minutes }}
            </div>
            <div class="text-sm opacity-75">Menit</div>
          </div>
          <div
            class="bg-white bg-opacity-20 rounded-lg p-4 text-white text-center min-w-[80px]"
          >
            <div class="text-2xl lg:text-3xl font-bold">
              {{ countdown.seconds }}
            </div>
            <div class="text-sm opacity-75">Detik</div>
          </div>
        </div>
      </div>

      <!-- Promo Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
        <div
          v-for="promo in promos"
          :key="promo.id"
          class="bg-white rounded-2xl overflow-hidden shadow-lg hover:shadow-2xl transition-all duration-300 transform hover:scale-105"
        >
          <div class="relative">
            <img
              :src="promo.image"
              :alt="promo.title"
              class="w-full h-48 object-cover"
            />
            <div
              class="absolute top-4 left-4 bg-red-500 text-white px-3 py-1 rounded-full text-sm font-bold"
            >
              -{{ promo.discount }}%
            </div>
            <div
              class="absolute bottom-4 right-4 bg-black bg-opacity-50 text-white px-2 py-1 rounded text-xs"
            >
              {{ promo.timeLeft }}
            </div>
          </div>
          <div class="p-6">
            <h3 class="font-bold text-lg text-gray-800 mb-2">
              {{ promo.title }}
            </h3>
            <p class="text-gray-600 text-sm mb-4">{{ promo.description }}</p>
            <div class="flex items-center justify-between">
              <div>
                <span class="text-gray-400 line-through text-sm"
                  >Rp {{ promo.originalPrice }}</span
                >
                <span class="text-red-600 font-bold text-lg ml-2"
                  >Rp {{ promo.salePrice }}</span
                >
              </div>
              <button
                class="bg-gradient-to-r from-pink-500 to-red-500 text-white px-4 py-2 rounded-lg hover:from-pink-600 hover:to-red-600 transition-all duration-300 text-sm font-semibold"
              >
                Beli Sekarang
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- CTA Section -->
      <div class="text-center">
        <button
          class="bg-yellow-400 hover:bg-yellow-300 text-black font-bold px-12 py-4 rounded-full text-xl transition-all duration-300 transform hover:scale-105 shadow-lg"
        >
          Lihat Semua Promo
        </button>
      </div>
    </div>

    <!-- Floating Elements -->
    <div class="absolute top-20 left-10 animate-float">
      <div class="w-4 h-4 bg-yellow-300 rounded-full"></div>
    </div>
    <div class="absolute top-40 right-20 animate-float-delayed">
      <div class="w-6 h-6 bg-pink-300 rounded-full"></div>
    </div>
    <div class="absolute bottom-20 left-1/4 animate-float">
      <div class="w-3 h-3 bg-white rounded-full"></div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";

// Countdown timer
const countdown = ref({
  hours: 12,
  minutes: 34,
  seconds: 56,
});

// Promo data
const promos = ref([
  {
    id: 1,
    title: "Fashion Week Sale",
    description: "Koleksi pakaian tren terbaru dengan diskon fantastis",
    discount: 60,
    originalPrice: "299.000",
    salePrice: "119.000",
    timeLeft: "5j 23m",
    image:
      "https://images.unsplash.com/photo-1441986300917-64674bd600d8?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
  },
  {
    id: 2,
    title: "Elektronik Super Sale",
    description: "Gadget terbaru dengan harga yang tidak bisa ditolak",
    discount: 45,
    originalPrice: "2.999.000",
    salePrice: "1.649.000",
    timeLeft: "3j 45m",
    image:
      "https://images.unsplash.com/photo-1526738549149-8e07eca6c147?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
  },
  {
    id: 3,
    title: "Home & Living",
    description: "Dekorasi rumah impian dengan harga spesial",
    discount: 70,
    originalPrice: "899.000",
    salePrice: "269.000",
    timeLeft: "7j 12m",
    image:
      "https://images.unsplash.com/photo-1586023492125-27b2c045efd7?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80",
  },
]);

let countdownInterval;

onMounted(() => {
  // Start countdown timer
  countdownInterval = setInterval(() => {
    if (countdown.value.seconds > 0) {
      countdown.value.seconds--;
    } else if (countdown.value.minutes > 0) {
      countdown.value.minutes--;
      countdown.value.seconds = 59;
    } else if (countdown.value.hours > 0) {
      countdown.value.hours--;
      countdown.value.minutes = 59;
      countdown.value.seconds = 59;
    }
  }, 1000);
});

onUnmounted(() => {
  if (countdownInterval) {
    clearInterval(countdownInterval);
  }
});
</script>

<style scoped>
@keyframes float {
  0%,
  100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-20px);
  }
}

@keyframes float-delayed {
  0%,
  100% {
    transform: translateY(-10px);
  }
  50% {
    transform: translateY(-30px);
  }
}

.animate-float {
  animation: float 3s ease-in-out infinite;
}

.animate-float-delayed {
  animation: float-delayed 4s ease-in-out infinite;
}
</style>
