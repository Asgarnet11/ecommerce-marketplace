<template>
  <footer class="bg-gray-900 text-white">
    <!-- Main Footer -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
        <!-- Company Info -->
        <div class="lg:col-span-1">
          <div class="flex items-center mb-6">
            <div
              class="w-10 h-10 bg-green-400 rounded-lg flex items-center justify-center mr-3"
            >
              <span class="text-gray-900 font-bold text-lg">T</span>
            </div>
            <span class="text-2xl font-bold text-green-400">TokoDM</span>
          </div>
          <p class="text-gray-300 mb-6 leading-relaxed">
            Platform e-commerce terpercaya yang menghadirkan ribuan produk
            berkualitas dengan harga terbaik untuk seluruh Indonesia.
          </p>

          <!-- Social Media -->
          <div class="flex space-x-4">
            <a
              v-for="social in socialMedia"
              :key="social.name"
              :href="social.url"
              class="w-10 h-10 bg-gray-800 rounded-full flex items-center justify-center hover:bg-green-600 transition-colors duration-300"
              :aria-label="social.name"
            >
              <span class="text-white font-semibold text-sm">{{
                social.icon
              }}</span>
            </a>
          </div>
        </div>

        <!-- Quick Links -->
        <div>
          <h3 class="text-lg font-bold mb-6 text-green-400">Tautan Cepat</h3>
          <ul class="space-y-3">
            <li v-for="link in quickLinks" :key="link.name">
              <a
                :href="link.url"
                class="text-gray-300 hover:text-green-400 transition-colors duration-300 flex items-center group"
              >
                <svg
                  class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity duration-300"
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
                {{ link.name }}
              </a>
            </li>
          </ul>
        </div>

        <!-- Customer Service -->
        <div>
          <h3 class="text-lg font-bold mb-6 text-green-400">
            Layanan Pelanggan
          </h3>
          <ul class="space-y-3">
            <li v-for="service in customerServices" :key="service.name">
              <a
                :href="service.url"
                class="text-gray-300 hover:text-green-400 transition-colors duration-300 flex items-center group"
              >
                <svg
                  class="w-4 h-4 mr-2 opacity-0 group-hover:opacity-100 transition-opacity duration-300"
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
                {{ service.name }}
              </a>
            </li>
          </ul>
        </div>

        <!-- Contact Info -->
        <div>
          <h3 class="text-lg font-bold mb-6 text-green-400">Hubungi Kami</h3>

          <!-- Contact Details -->
          <div class="space-y-4 mb-6">
            <div class="flex items-start">
              <svg
                class="w-5 h-5 mt-1 mr-3 text-green-400 flex-shrink-0"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
                />
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"
                />
              </svg>
              <div>
                <p class="text-gray-300">Jl. Haluoleo No. 123</p>
                <p class="text-gray-300">Kendari, Sulawesi Tenggara 93232</p>
              </div>
            </div>

            <div class="flex items-center">
              <svg
                class="w-5 h-5 mr-3 text-green-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"
                />
              </svg>
              <a
                href="tel:+6281234567890"
                class="text-gray-300 hover:text-green-400 transition-colors"
              >
                +62 812-3456-7890
              </a>
            </div>

            <div class="flex items-center">
              <svg
                class="w-5 h-5 mr-3 text-green-400"
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
              <a
                href="mailto:info@tokodm.com"
                class="text-gray-300 hover:text-green-400 transition-colors"
              >
                info@tokodm.com
              </a>
            </div>

            <div class="flex items-center">
              <svg
                class="w-5 h-5 mr-3 text-green-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                />
              </svg>
              <span class="text-gray-300">Senin - Minggu: 08:00 - 22:00</span>
            </div>
          </div>

          <!-- Download App -->
          <div>
            <p class="text-sm text-gray-400 mb-3">Download Aplikasi</p>
            <div class="flex flex-col space-y-2">
              <a
                href="#"
                class="inline-flex items-center bg-gray-800 hover:bg-gray-700 rounded-lg px-4 py-2 transition-colors duration-300"
              >
                <span class="text-green-400 font-bold mr-2">📱</span>
                <span class="text-white text-sm">Google Play</span>
              </a>
              <a
                href="#"
                class="inline-flex items-center bg-gray-800 hover:bg-gray-700 rounded-lg px-4 py-2 transition-colors duration-300"
              >
                <span class="text-white font-bold mr-2">🍎</span>
                <span class="text-white text-sm">App Store</span>
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Payment Methods & Certifications -->
    <div class="border-t border-gray-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 items-center">
          <!-- Payment Methods -->
          <div>
            <h4 class="text-sm font-semibold text-gray-400 mb-4">
              Metode Pembayaran
            </h4>
            <div class="flex flex-wrap gap-3">
              <div
                v-for="payment in paymentMethods"
                :key="payment.name"
                class="bg-white rounded-lg p-2 flex items-center justify-center w-16 h-10"
              >
                <span class="text-xs font-bold text-gray-800">{{
                  payment.name
                }}</span>
              </div>
            </div>
          </div>

          <!-- Certifications -->
          <div>
            <h4 class="text-sm font-semibold text-gray-400 mb-4">
              Keamanan & Sertifikasi
            </h4>
            <div class="flex flex-wrap gap-3">
              <div
                v-for="cert in certifications"
                :key="cert.name"
                class="bg-white rounded-lg p-2 flex items-center justify-center h-10 px-3"
              >
                <span class="text-xs font-bold text-gray-800">{{
                  cert.name
                }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Bottom Footer -->
    <div class="border-t border-gray-800 bg-gray-950">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div class="flex flex-col md:flex-row items-center justify-between">
          <p class="text-gray-400 text-sm mb-4 md:mb-0">
            © {{ currentYear }} TokoDM. Semua hak dilindungi undang-undang.
          </p>

          <!-- Legal Links -->
          <div class="flex flex-wrap items-center space-x-6 text-sm">
            <a
              href="#"
              class="text-gray-400 hover:text-green-400 transition-colors duration-300"
            >
              Syarat & Ketentuan
            </a>
            <a
              href="#"
              class="text-gray-400 hover:text-green-400 transition-colors duration-300"
            >
              Kebijakan Privasi
            </a>
            <a
              href="#"
              class="text-gray-400 hover:text-green-400 transition-colors duration-300"
            >
              Cookie Policy
            </a>
            <a
              href="#"
              class="text-gray-400 hover:text-green-400 transition-colors duration-300"
            >
              Panduan Penggunaan
            </a>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { ref, computed } from "vue";

const currentYear = computed(() => new Date().getFullYear());

const socialMedia = ref([
  { name: "Facebook", icon: "FB", url: "#" },
  { name: "Instagram", icon: "IG", url: "#" },
  { name: "Twitter", icon: "TW", url: "#" },
  { name: "YouTube", icon: "YT", url: "#" },
  { name: "TikTok", icon: "TT", url: "#" },
]);

const quickLinks = ref([
  { name: "Tentang Kami", url: "#" },
  { name: "Karir", url: "#" },
  { name: "Blog", url: "#" },
  { name: "Affiliate Program", url: "#" },
  { name: "Menjadi Seller", url: "#" },
  { name: "Promo Hari Ini", url: "#" },
]);

const customerServices = ref([
  { name: "Pusat Bantuan", url: "#" },
  { name: "Cara Berbelanja", url: "#" },
  { name: "Kebijakan Return", url: "#" },
  { name: "Metode Pembayaran", url: "#" },
  { name: "Lacak Pesanan", url: "#" },
  { name: "Hubungi CS", url: "#" },
]);

const paymentMethods = ref([
  { name: "VISA" },
  { name: "MCARD" },
  { name: "BCA" },
  { name: "MNDR" },
  { name: "GPAY" },
  { name: "OVO" },
]);

const certifications = ref([{ name: "SSL" }, { name: "ISO" }, { name: "PCI" }]);
</script>
