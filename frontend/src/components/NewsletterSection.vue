<template>
  <section
    class="py-16 bg-gradient-to-r from-green-600 to-blue-600 relative overflow-hidden"
  >
    <!-- Background Pattern -->
    <div class="absolute inset-0 opacity-10">
      <svg
        width="60"
        height="60"
        viewBox="0 0 60 60"
        class="absolute top-0 left-0 text-white"
      >
        <g fill="currentColor">
          <circle cx="30" cy="30" r="2" />
          <circle cx="10" cy="10" r="2" />
          <circle cx="50" cy="10" r="2" />
          <circle cx="10" cy="50" r="2" />
          <circle cx="50" cy="50" r="2" />
        </g>
      </svg>
      <div
        class="absolute inset-0"
        style="
          background-image: radial-gradient(
            circle at 25% 25%,
            rgba(255, 255, 255, 0.1) 1px,
            transparent 1px
          );
          background-size: 30px 30px;
        "
      ></div>
    </div>

    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center relative">
      <!-- Header -->
      <div class="mb-12">
        <div
          class="inline-flex items-center justify-center w-20 h-20 bg-yellow-400 rounded-full mb-6"
        >
          <svg
            class="w-10 h-10 text-gray-800"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
            />
          </svg>
        </div>

        <h2 class="text-4xl lg:text-5xl font-bold text-white mb-4">
          Dapatkan Penawaran Eksklusif!
        </h2>
        <p class="text-xl text-white opacity-90 max-w-2xl mx-auto">
          Berlangganan newsletter kami dan jadi yang pertama tahu tentang promo,
          produk baru, dan penawaran khusus lainnya.
        </p>
      </div>

      <!-- Benefits -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
        <div class="bg-white bg-opacity-20 rounded-2xl p-6 backdrop-blur-sm">
          <div
            class="w-12 h-12 bg-yellow-400 rounded-full flex items-center justify-center mx-auto mb-4"
          >
            <svg
              class="w-6 h-6 text-gray-800"
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
          <h3 class="font-bold text-white text-lg mb-2">Diskon Eksklusif</h3>
          <p class="text-white opacity-75 text-sm">
            Dapatkan kode diskon khusus subscriber hingga 20%
          </p>
        </div>

        <div class="bg-white bg-opacity-20 rounded-2xl p-6 backdrop-blur-sm">
          <div
            class="w-12 h-12 bg-pink-400 rounded-full flex items-center justify-center mx-auto mb-4"
          >
            <svg
              class="w-6 h-6 text-white"
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
          <h3 class="font-bold text-white text-lg mb-2">Info Terdepan</h3>
          <p class="text-white opacity-75 text-sm">
            Jadi yang pertama tahu tentang produk dan promo terbaru
          </p>
        </div>

        <div class="bg-white bg-opacity-20 rounded-2xl p-6 backdrop-blur-sm">
          <div
            class="w-12 h-12 bg-blue-400 rounded-full flex items-center justify-center mx-auto mb-4"
          >
            <svg
              class="w-6 h-6 text-white"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
              />
            </svg>
          </div>
          <h3 class="font-bold text-white text-lg mb-2">Tips Belanja</h3>
          <p class="text-white opacity-75 text-sm">
            Dapatkan tips dan trik berbelanja cerdas setiap minggu
          </p>
        </div>
      </div>

      <!-- Subscription Form -->
      <div
        class="bg-white rounded-3xl p-8 lg:p-12 shadow-2xl max-w-2xl mx-auto"
      >
        <form @submit.prevent="subscribe" class="space-y-6">
          <!-- Email Input -->
          <div class="flex flex-col sm:flex-row gap-4">
            <div class="flex-1 relative">
              <input
                v-model="email"
                type="email"
                placeholder="Masukkan email Anda"
                required
                class="w-full px-6 py-4 border-2 border-gray-200 rounded-2xl focus:border-green-500 focus:ring-0 focus:outline-none text-lg placeholder-gray-400"
                :class="{ 'border-red-300': emailError }"
              />
              <div
                v-if="emailError"
                class="absolute -bottom-6 left-0 text-red-500 text-sm"
              >
                {{ emailError }}
              </div>
            </div>

            <button
              type="submit"
              :disabled="isLoading"
              class="px-8 py-4 bg-gradient-to-r from-green-600 to-blue-600 text-white font-bold rounded-2xl hover:from-green-700 hover:to-blue-700 transition-all duration-300 transform hover:scale-105 disabled:opacity-50 disabled:cursor-not-allowed whitespace-nowrap"
            >
              <span v-if="!isLoading" class="flex items-center">
                Berlangganan
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
                    d="M14 5l7 7m0 0l-7 7m7-7H3"
                  />
                </svg>
              </span>
              <span v-else class="flex items-center">
                <svg
                  class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
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
                Memproses...
              </span>
            </button>
          </div>

          <!-- Success Message -->
          <div
            v-if="showSuccess"
            class="bg-green-100 border border-green-400 text-green-700 px-6 py-4 rounded-2xl"
          >
            <div class="flex items-center">
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
              Terima kasih! Anda berhasil berlangganan newsletter kami.
            </div>
          </div>

          <!-- Terms -->
          <p class="text-sm text-gray-500 text-center">
            Dengan berlangganan, Anda menyetujui
            <a href="#" class="text-green-600 hover:underline"
              >Syarat & Ketentuan</a
            >
            dan
            <a href="#" class="text-green-600 hover:underline"
              >Kebijakan Privasi</a
            >
            kami.
          </p>
        </form>

        <!-- Subscriber Count -->
        <div class="mt-8 pt-6 border-t border-gray-200">
          <p class="text-gray-600 text-center">
            <span class="font-bold text-green-600"
              >{{ subscriberCount.toLocaleString() }}+</span
            >
            orang sudah bergabung dengan newsletter kami
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from "vue";

const email = ref("");
const emailError = ref("");
const isLoading = ref(false);
const showSuccess = ref(false);
const subscriberCount = ref(25847);

const validateEmail = (email) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
};

const subscribe = async () => {
  emailError.value = "";

  if (!email.value) {
    emailError.value = "Email wajib diisi";
    return;
  }

  if (!validateEmail(email.value)) {
    emailError.value = "Format email tidak valid";
    return;
  }

  isLoading.value = true;

  try {
    // Simulate API call
    await new Promise((resolve) => setTimeout(resolve, 2000));

    showSuccess.value = true;
    email.value = "";
    subscriberCount.value += 1;

    // Hide success message after 5 seconds
    setTimeout(() => {
      showSuccess.value = false;
    }, 5000);
  } catch (error) {
    emailError.value = "Terjadi kesalahan, silakan coba lagi";
  } finally {
    isLoading.value = false;
  }
};
</script>
