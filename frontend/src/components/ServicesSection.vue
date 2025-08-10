<template>
  <section class="py-16 bg-white">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Section Header -->
      <div class="text-center mb-16">
        <h2 class="text-3xl lg:text-4xl font-bold text-gray-800 mb-4">
          Kenapa Pilih Kami?
        </h2>
        <p class="text-gray-600 max-w-3xl mx-auto text-lg">
          Kami memberikan pengalaman berbelanja terbaik dengan layanan yang
          dapat diandalkan
        </p>
      </div>

      <!-- Services Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8 mb-16">
        <div
          v-for="service in services"
          :key="service.id"
          class="text-center group hover:transform hover:scale-105 transition-all duration-300"
        >
          <!-- Icon Container -->
          <div
            class="w-20 h-20 mx-auto mb-6 rounded-full flex items-center justify-center group-hover:animate-bounce"
            :class="`bg-gradient-to-r ${service.gradient}`"
          >
            <svg
              class="w-10 h-10 text-white"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                :d="service.iconPath"
              />
            </svg>
          </div>

          <!-- Service Info -->
          <h3
            class="text-xl font-bold text-gray-800 mb-3 group-hover:text-green-600 transition-colors"
          >
            {{ service.title }}
          </h3>
          <p class="text-gray-600 leading-relaxed">
            {{ service.description }}
          </p>
        </div>
      </div>

      <!-- Stats Section -->
      <div
        class="bg-gradient-to-r from-green-50 to-blue-50 rounded-3xl p-8 lg:p-12"
      >
        <div class="grid grid-cols-2 lg:grid-cols-4 gap-8">
          <div v-for="stat in stats" :key="stat.id" class="text-center">
            <div
              class="text-3xl lg:text-4xl font-bold mb-2"
              :class="stat.color"
            >
              {{ stat.value }}
            </div>
            <div class="text-gray-600 font-medium">
              {{ stat.label }}
            </div>
          </div>
        </div>
      </div>

      <!-- Testimonial Section -->
      <div class="mt-16">
        <h3 class="text-2xl font-bold text-center text-gray-800 mb-12">
          Apa Kata Pelanggan Kami
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div
            v-for="testimonial in testimonials"
            :key="testimonial.id"
            class="bg-white rounded-2xl p-6 shadow-lg border border-gray-100 hover:shadow-xl transition-shadow duration-300"
          >
            <!-- Stars -->
            <div class="flex mb-4">
              <span
                v-for="star in 5"
                :key="star"
                class="text-yellow-400 text-xl"
                >★</span
              >
            </div>

            <!-- Review -->
            <p class="text-gray-600 mb-6 italic">"{{ testimonial.review }}"</p>

            <!-- User Info -->
            <div class="flex items-center">
              <img
                :src="testimonial.avatar"
                :alt="testimonial.name"
                class="w-12 h-12 rounded-full object-cover mr-4"
              />
              <div>
                <div class="font-semibold text-gray-800">
                  {{ testimonial.name }}
                </div>
                <div class="text-sm text-gray-500">
                  {{ testimonial.location }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from "vue";

const services = ref([
  {
    id: 1,
    title: "Pengiriman Cepat",
    description:
      "Pengiriman same day untuk area tertentu dan pengiriman reguler ke seluruh Indonesia",
    gradient: "from-blue-500 to-cyan-500",
    iconPath: "M13 10V3L4 14h7v7l9-11h-7z",
  },
  {
    id: 2,
    title: "Pembayaran Aman",
    description:
      "Berbagai metode pembayaran yang aman dan terpercaya untuk kemudahan transaksi",
    gradient: "from-green-500 to-emerald-500",
    iconPath:
      "M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z",
  },
  {
    id: 3,
    title: "Garansi Kualitas",
    description:
      "Jaminan kualitas produk dengan kebijakan return yang mudah dan tanpa ribet",
    gradient: "from-purple-500 to-indigo-500",
    iconPath:
      "M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z",
  },
  {
    id: 4,
    title: "Customer Service 24/7",
    description:
      "Tim customer service siap membantu Anda kapan saja melalui berbagai saluran komunikasi",
    gradient: "from-pink-500 to-rose-500",
    iconPath:
      "M18.364 5.636l-3.536 3.536m0 5.656l3.536 3.536M9.172 9.172L5.636 5.636m3.536 9.192L5.636 18.364M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-5 0a4 4 0 11-8 0 4 4 0 018 0z",
  },
]);

const stats = ref([
  {
    id: 1,
    value: "500K+",
    label: "Pelanggan Puas",
    color: "text-blue-600",
  },
  {
    id: 2,
    value: "100K+",
    label: "Produk Tersedia",
    color: "text-green-600",
  },
  {
    id: 3,
    value: "50+",
    label: "Kota Terjangkau",
    color: "text-purple-600",
  },
  {
    id: 4,
    value: "4.8⭐",
    label: "Rating Aplikasi",
    color: "text-yellow-600",
  },
]);

const testimonials = ref([
  {
    id: 1,
    name: "Sarah Wijaya",
    location: "Jakarta",
    review:
      "Pelayanan sangat memuaskan! Barang sampai cepat dan sesuai dengan deskripsi. Pasti akan belanja lagi disini.",
    avatar:
      "https://images.unsplash.com/photo-1494790108755-2616b612b786?ixlib=rb-4.0.3&auto=format&fit=crop&w=150&q=80",
  },
  {
    id: 2,
    name: "Budi Santoso",
    location: "Surabaya",
    review:
      "Harga kompetitif dan kualitas produk bagus. Customer service juga sangat responsif membantu.",
    avatar:
      "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-4.0.3&auto=format&fit=crop&w=150&q=80",
  },
  {
    id: 3,
    name: "Andi Pratama",
    location: "Kendari",
    review:
      "Pengiriman ke Kendari sangat cepat! Tidak menyangka bisa se-efisien ini. Terima kasih!",
    avatar:
      "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?ixlib=rb-4.0.3&auto=format&fit=crop&w=150&q=80",
  },
]);
</script>
