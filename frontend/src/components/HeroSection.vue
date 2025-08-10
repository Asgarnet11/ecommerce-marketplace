<template>
  <section
    class="relative bg-gradient-to-r from-green-600 to-blue-600 overflow-hidden"
  >
    <div class="absolute inset-0 bg-black opacity-20"></div>
    <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16 lg:py-24">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 items-center">
        <!-- Content -->
        <div class="text-white">
          <h1 class="text-4xl lg:text-6xl font-bold mb-4">
            Belanja Mudah,
            <span class="text-yellow-300">Hemat Setiap Hari</span>
          </h1>
          <p class="text-lg lg:text-xl mb-8 text-gray-100">
            Temukan ribuan produk berkualitas dengan harga terbaik. Pengiriman
            cepat ke seluruh Indonesia.
          </p>

          <!-- CTA Buttons -->
          <div class="flex flex-col sm:flex-row gap-4 mb-8">
            <button
              class="bg-yellow-500 hover:bg-yellow-400 text-black font-semibold px-8 py-3 rounded-lg transition-colors duration-300 transform hover:scale-105"
            >
              Mulai Belanja
            </button>
            <button
              class="border-2 border-white text-white hover:bg-white hover:text-green-600 font-semibold px-8 py-3 rounded-lg transition-all duration-300"
            >
              Download App
            </button>
          </div>

          <!-- Features -->
          <div class="grid grid-cols-3 gap-4 text-center">
            <div class="bg-white bg-opacity-20 rounded-lg p-3">
              <div class="text-yellow-300 text-2xl font-bold">24/7</div>
              <div class="text-sm">Layanan</div>
            </div>
            <div class="bg-white bg-opacity-20 rounded-lg p-3">
              <div class="text-yellow-300 text-2xl font-bold">100K+</div>
              <div class="text-sm">Produk</div>
            </div>
            <div class="bg-white bg-opacity-20 rounded-lg p-3">
              <div class="text-yellow-300 text-2xl font-bold">⭐4.8</div>
              <div class="text-sm">Rating</div>
            </div>
          </div>
        </div>

        <!-- Image/Illustration -->
        <div class="relative">
          <!-- Main Hero Image -->
          <div
            class="relative bg-white rounded-2xl shadow-2xl p-6 transform rotate-3 hover:rotate-0 transition-transform duration-500"
          >
            <img
              src="https://images.unsplash.com/photo-1556742049-0cfed4f6a45d?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80"
              alt="Shopping"
              class="w-full h-64 lg:h-80 object-cover rounded-lg"
            />
            <!-- Floating Elements -->
            <div
              class="absolute -top-4 -right-4 bg-red-500 text-white px-3 py-1 rounded-full text-sm font-bold animate-bounce"
            >
              SALE 50%
            </div>
            <div
              class="absolute -bottom-4 -left-4 bg-yellow-500 text-black px-3 py-1 rounded-full text-sm font-bold"
            >
              Free Shipping
            </div>
          </div>

          <!-- Background Decorations -->
          <div
            class="absolute -z-10 top-10 right-10 w-32 h-32 bg-yellow-300 rounded-full opacity-20 animate-pulse"
          ></div>
          <div
            class="absolute -z-10 bottom-10 left-10 w-24 h-24 bg-pink-300 rounded-full opacity-20 animate-pulse"
          ></div>
        </div>
      </div>
    </div>

    <!-- Bottom Wave -->
    <div class="absolute bottom-0 w-full">
      <svg
        viewBox="0 0 1440 120"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M0 120L60 105C120 90 240 60 360 45C480 30 600 30 720 37.5C840 45 960 60 1080 67.5C1200 75 1320 75 1380 75L1440 75V120H1380C1320 120 1200 120 1080 120C960 120 840 120 720 120C600 120 480 120 360 120C240 120 120 120 60 120H0Z"
          fill="#f9fafb"
        />
      </svg>
    </div>
  </section>
</template>

<script setup>
// Hero section component - no additional logic needed for now
</script>
