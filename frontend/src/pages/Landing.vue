<template>
  <!-- Hero Section -->
  <section class="relative overflow-hidden min-h-[80vh] flex items-center">
    <!-- Animated Background -->
    <div
      class="absolute inset-0 bg-gradient-to-br from-blue-50 via-white to-purple-50"
    >
      <!-- Floating Elements -->
      <div
        class="absolute top-20 left-10 w-20 h-20 bg-blue-200/30 rounded-full blur-xl animate-pulse"
      ></div>
      <div
        class="absolute top-40 right-20 w-32 h-32 bg-purple-200/20 rounded-full blur-2xl animate-pulse delay-1000"
      ></div>
      <div
        class="absolute bottom-20 left-1/4 w-16 h-16 bg-teal-200/40 rounded-full blur-lg animate-pulse delay-2000"
      ></div>
    </div>

    <!-- Grid Pattern Overlay -->
    <div
      class="absolute inset-0 opacity-[0.02]"
      style="
        background-image: radial-gradient(circle, #000 1px, transparent 1px);
        background-size: 50px 50px;
      "
    ></div>

    <div class="container relative py-16 md:py-20 lg:py-24">
      <div class="grid lg:grid-cols-2 gap-12 lg:gap-16 items-center">
        <!-- Left Content -->
        <div class="space-y-8 text-center lg:text-left">
          <!-- Badge -->
          <div
            class="inline-flex items-center px-4 py-2 bg-gradient-to-r from-blue-100 to-purple-100 rounded-full text-sm font-medium text-gray-700 mb-4"
          >
            <svg
              class="w-4 h-4 mr-2 text-green-500"
              fill="currentColor"
              viewBox="0 0 20 20"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
              />
            </svg>
            Platform Marketplace Terpercaya
          </div>

          <!-- Main Heading -->
          <div>
            <h1
              class="text-4xl md:text-5xl lg:text-6xl font-extrabold leading-tight tracking-tight"
            >
              <span class="block">Belanja cepat &</span>
              <span
                class="block bg-gradient-to-r from-blue-600 via-purple-600 to-teal-600 bg-clip-text text-transparent"
              >
                hemat
              </span>
              <span class="block">di satu tempat</span>
            </h1>
          </div>

          <!-- Description -->
          <p class="text-lg md:text-xl text-gray-600 max-w-2xl leading-relaxed">
            Temukan jutaan produk dari penjual terpercaya. Harga bersaing,
            pengiriman cepat, pengalaman belanja yang menyenangkan untuk semua
            kebutuhan Anda.
          </p>

          <!-- CTA Buttons -->
          <div
            class="flex flex-col sm:flex-row gap-4 justify-center lg:justify-start"
          >
            <router-link
              to="/products"
              class="group px-8 py-4 bg-gradient-to-r from-blue-500 to-purple-600 text-white rounded-2xl font-semibold text-lg shadow-xl hover:shadow-2xl hover:scale-105 transition-all duration-300 flex items-center justify-center gap-2"
            >
              Mulai Belanja
              <svg
                class="w-5 h-5 group-hover:translate-x-1 transition-transform"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M17 8l4 4m0 0l-4 4m4-4H3"
                />
              </svg>
            </router-link>
            <router-link
              to="/shops"
              class="px-8 py-4 bg-white/80 backdrop-blur text-gray-700 rounded-2xl font-semibold text-lg border-2 border-gray-200 hover:border-purple-300 hover:bg-white transition-all duration-300 flex items-center justify-center gap-2"
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
                  d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1"
                />
              </svg>
              Jelajahi Toko
            </router-link>
          </div>

          <!-- Stats -->
          <div class="grid grid-cols-3 gap-8 pt-8">
            <div class="text-center">
              <div class="text-2xl md:text-3xl font-bold text-gray-900">
                1M+
              </div>
              <div class="text-sm text-gray-600">Produk</div>
            </div>
            <div class="text-center">
              <div class="text-2xl md:text-3xl font-bold text-gray-900">
                50K+
              </div>
              <div class="text-sm text-gray-600">Toko</div>
            </div>
            <div class="text-center">
              <div class="text-2xl md:text-3xl font-bold text-gray-900">
                99%
              </div>
              <div class="text-sm text-gray-600">Puas</div>
            </div>
          </div>
        </div>

        <!-- Right Content -->
        <div class="relative">
          <!-- Main Image Card -->
          <div class="relative group">
            <div
              class="absolute inset-0 bg-gradient-to-r from-blue-500/20 to-purple-500/20 rounded-3xl blur-2xl group-hover:blur-3xl transition-all duration-300"
            ></div>
            <div
              class="relative bg-white/80 backdrop-blur-sm p-4 md:p-6 rounded-3xl shadow-2xl border border-white/50"
            >
              <img
                class="rounded-2xl object-cover w-full aspect-[4/3] group-hover:scale-[1.02] transition-transform duration-500"
                src="https://images.unsplash.com/photo-1515168833906-d2a3b82b302a?q=80&w=1200&auto=format&fit=crop"
                alt="Hero Shopping Experience"
                loading="eager"
              />
              <!-- Image Overlay -->
              <div
                class="absolute inset-4 md:inset-6 rounded-2xl bg-gradient-to-t from-black/20 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"
              ></div>
            </div>
          </div>

          <!-- Floating Discount Card -->
          <div
            class="absolute -bottom-6 -left-4 md:-bottom-8 md:-left-6 bg-white rounded-2xl shadow-2xl p-4 md:p-6 max-w-xs transform hover:scale-105 transition-all duration-300 border border-orange-100"
          >
            <div class="flex items-start gap-3">
              <div
                class="p-2 bg-gradient-to-r from-orange-400 to-red-400 rounded-xl"
              >
                <svg
                  class="w-5 h-5 text-white"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    fill-rule="evenodd"
                    d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z"
                  />
                </svg>
              </div>
              <div>
                <p class="font-bold text-gray-900 text-sm md:text-base">
                  Flash Sale!
                </p>
                <p class="text-xs md:text-sm text-gray-600 leading-relaxed">
                  Diskon hingga
                  <span class="font-bold text-red-500">70%</span> untuk kategori
                  elektronik
                </p>
              </div>
            </div>
          </div>

          <!-- Floating Rating Card -->
          <div
            class="absolute -top-4 -right-4 md:-top-6 md:-right-6 bg-white rounded-2xl shadow-xl p-3 md:p-4 transform hover:scale-105 transition-all duration-300 border border-green-100"
          >
            <div class="flex items-center gap-2">
              <div class="flex text-yellow-400">
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                  <path
                    d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
                  />
                </svg>
                <span class="text-xs font-bold text-gray-900 ml-1">4.9</span>
              </div>
              <span class="text-xs text-gray-600">Rating</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- Products Section -->
  <section class="py-16 md:py-20 bg-gradient-to-b from-white to-gray-50">
    <div class="container">
      <div class="text-center mb-12">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
          Produk
          <span
            class="bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent"
            >Terbaru</span
          >
        </h2>
        <p class="text-gray-600 text-lg max-w-2xl mx-auto">
          Temukan koleksi produk terbaru dari berbagai kategori pilihan terbaik
          untuk Anda
        </p>
      </div>
      <ProductGrid :initialFetch="true" />
    </div>
  </section>

  <!-- Features Section -->
  <section
    class="py-16 md:py-20 bg-gradient-to-br from-gray-50 via-blue-50/30 to-purple-50/30 relative overflow-hidden"
  >
    <!-- Background Pattern -->
    <div class="absolute inset-0 opacity-5">
      <div
        class="absolute inset-0"
        style="
          background-image: radial-gradient(
            circle at 25% 25%,
            #3b82f6 2px,
            transparent 2px
          );
          background-size: 60px 60px;
        "
      ></div>
    </div>

    <div class="container relative">
      <div class="text-center mb-16">
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
          Kenapa Pilih
          <span
            class="bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent"
            >Kami?</span
          >
        </h2>
        <p class="text-gray-600 text-lg max-w-2xl mx-auto">
          Pengalaman berbelanja terbaik dengan layanan prima untuk kepuasan Anda
        </p>
      </div>

      <div class="grid md:grid-cols-3 gap-8 lg:gap-12">
        <!-- Feature 1 -->
        <div class="group">
          <div
            class="bg-white/80 backdrop-blur-sm rounded-3xl p-8 shadow-xl border border-white/50 hover:shadow-2xl hover:-translate-y-2 transition-all duration-500"
          >
            <div class="mb-6">
              <div
                class="inline-flex p-4 bg-gradient-to-r from-green-400 to-blue-500 rounded-2xl"
              >
                <svg
                  class="w-8 h-8 text-white"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"
                  />
                </svg>
              </div>
            </div>
            <h3
              class="text-xl font-bold text-gray-900 mb-4 group-hover:text-blue-600 transition-colors"
            >
              Harga Bersahabat
            </h3>
            <p class="text-gray-600 leading-relaxed">
              Dapatkan harga terbaik dengan sistem perbandingan otomatis dan
              promo menarik setiap hari untuk semua kategori produk.
            </p>
          </div>
        </div>

        <!-- Feature 2 -->
        <div class="group">
          <div
            class="bg-white/80 backdrop-blur-sm rounded-3xl p-8 shadow-xl border border-white/50 hover:shadow-2xl hover:-translate-y-2 transition-all duration-500"
          >
            <div class="mb-6">
              <div
                class="inline-flex p-4 bg-gradient-to-r from-purple-400 to-pink-500 rounded-2xl"
              >
                <svg
                  class="w-8 h-8 text-white"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13 10V3L4 14h7v7l9-11h-7z"
                  />
                </svg>
              </div>
            </div>
            <h3
              class="text-xl font-bold text-gray-900 mb-4 group-hover:text-purple-600 transition-colors"
            >
              Pengiriman Cepat
            </h3>
            <p class="text-gray-600 leading-relaxed">
              Didukung jaringan kurir terpercaya dengan tracking realtime dan
              jaminan barang sampai tepat waktu ke tangan Anda.
            </p>
          </div>
        </div>

        <!-- Feature 3 -->
        <div class="group">
          <div
            class="bg-white/80 backdrop-blur-sm rounded-3xl p-8 shadow-xl border border-white/50 hover:shadow-2xl hover:-translate-y-2 transition-all duration-500"
          >
            <div class="mb-6">
              <div
                class="inline-flex p-4 bg-gradient-to-r from-orange-400 to-red-500 rounded-2xl"
              >
                <svg
                  class="w-8 h-8 text-white"
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
            </div>
            <h3
              class="text-xl font-bold text-gray-900 mb-4 group-hover:text-red-600 transition-colors"
            >
              Aman & Terpercaya
            </h3>
            <p class="text-gray-600 leading-relaxed">
              Sistem pembayaran terenkripsi dengan perlindungan pembeli 100% dan
              garansi uang kembali jika tidak puas.
            </p>
          </div>
        </div>
      </div>

      <!-- Additional Trust Indicators -->
      <div class="mt-16 text-center">
        <p class="text-gray-500 mb-8">
          Dipercaya oleh jutaan pelanggan di seluruh Indonesia
        </p>
        <div
          class="flex flex-wrap justify-center items-center gap-8 opacity-60"
        >
          <div class="flex items-center gap-2 text-gray-400">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
              />
            </svg>
            <span class="font-medium">SSL Secure</span>
          </div>
          <div class="flex items-center gap-2 text-gray-400">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M6.267 3.455a3.066 3.066 0 001.745-.723 3.066 3.066 0 013.976 0 3.066 3.066 0 001.745.723 3.066 3.066 0 012.812 2.812c.051.643.304 1.254.723 1.745a3.066 3.066 0 010 3.976 3.066 3.066 0 00-.723 1.745 3.066 3.066 0 01-2.812 2.812 3.066 3.066 0 00-1.745.723 3.066 3.066 0 01-3.976 0 3.066 3.066 0 00-1.745-.723 3.066 3.066 0 01-2.812-2.812 3.066 3.066 0 00-.723-1.745 3.066 3.066 0 010-3.976 3.066 3.066 0 00.723-1.745 3.066 3.066 0 012.812-2.812zm7.44 5.252a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
              />
            </svg>
            <span class="font-medium">Verified Sellers</span>
          </div>
          <div class="flex items-center gap-2 text-gray-400">
            <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20">
              <path
                d="M4 4a2 2 0 00-2 2v1h16V6a2 2 0 00-2-2H4zM18 9H2v5a2 2 0 002 2h12a2 2 0 002-2V9zM4 13a1 1 0 011-1h1a1 1 0 110 2H5a1 1 0 01-1-1zm5-1a1 1 0 100 2h1a1 1 0 100-2H9z"
              />
            </svg>
            <span class="font-medium">Safe Payment</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import ProductGrid from "../components/ProductGrid.vue";
</script>

<style scoped>
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding-left: 1rem;
  padding-right: 1rem;
}

@media (min-width: 768px) {
  .container {
    padding-left: 2rem;
    padding-right: 2rem;
  }
}

/* Animation keyframes */
@keyframes pulse {
  0%,
  100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}

.animate-pulse {
  animation: pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.delay-1000 {
  animation-delay: 1s;
}

.delay-2000 {
  animation-delay: 2s;
}

/* Responsive improvements */
@media (max-width: 640px) {
  .container {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}
</style>
