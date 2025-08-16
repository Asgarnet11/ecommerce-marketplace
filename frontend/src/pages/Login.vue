<template>
  <div
    class="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 relative overflow-hidden"
  >
    <!-- Background Elements -->
    <div
      class="absolute top-20 left-10 w-32 h-32 bg-blue-200/20 rounded-full blur-2xl animate-pulse"
    ></div>
    <div
      class="absolute bottom-20 right-10 w-40 h-40 bg-purple-200/20 rounded-full blur-2xl animate-pulse delay-1000"
    ></div>
    <div
      class="absolute top-1/2 left-1/4 w-24 h-24 bg-teal-200/30 rounded-full blur-xl animate-pulse delay-2000"
    ></div>

    <!-- Grid Pattern -->
    <div
      class="absolute inset-0 opacity-[0.02]"
      style="
        background-image: radial-gradient(circle, #000 1px, transparent 1px);
        background-size: 50px 50px;
      "
    ></div>

    <div class="relative flex items-center justify-center min-h-screen p-4">
      <div class="w-full max-w-md">
        <!-- Login Card -->
        <div
          class="bg-white/80 backdrop-blur-xl rounded-3xl shadow-2xl border border-white/50 p-8 md:p-10"
        >
          <!-- Header -->
          <div class="text-center mb-8">
            <!-- Logo/Icon -->
            <div
              class="inline-flex p-4 bg-gradient-to-r from-blue-500 to-purple-600 rounded-2xl mb-4"
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
                  d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                />
              </svg>
            </div>

            <!-- Title -->
            <h1 class="text-3xl font-bold text-gray-900 mb-2">
              Selamat Datang
            </h1>
            <p class="text-gray-600">
              Masuk ke akun Anda untuk melanjutkan belanja
            </p>
          </div>

          <!-- Login Form -->
          <form class="space-y-6" @submit.prevent="submit">
            <!-- Email Input -->
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700">
                Email
              </label>
              <div class="relative">
                <div
                  class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none"
                >
                  <svg
                    class="w-5 h-5 text-gray-400"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207"
                    />
                  </svg>
                </div>
                <input
                  v-model="email"
                  type="email"
                  placeholder="Masukkan email Anda"
                  class="w-full pl-12 pr-4 py-4 bg-gray-50/50 border border-gray-200 rounded-2xl focus:bg-white focus:border-blue-500 focus:ring-4 focus:ring-blue-500/10 transition-all duration-200 text-gray-900 placeholder-gray-500"
                  required
                />
              </div>
            </div>

            <!-- Password Input -->
            <div class="space-y-2">
              <label class="block text-sm font-semibold text-gray-700">
                Password
              </label>
              <div class="relative">
                <div
                  class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none"
                >
                  <svg
                    class="w-5 h-5 text-gray-400"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                    />
                  </svg>
                </div>
                <input
                  v-model="password"
                  type="password"
                  placeholder="Masukkan password Anda"
                  class="w-full pl-12 pr-4 py-4 bg-gray-50/50 border border-gray-200 rounded-2xl focus:bg-white focus:border-blue-500 focus:ring-4 focus:ring-blue-500/10 transition-all duration-200 text-gray-900 placeholder-gray-500"
                  required
                />
              </div>
            </div>

            <!-- Forgot Password -->
            <div class="text-right">
              <button
                type="button"
                class="text-sm text-blue-600 hover:text-blue-700 font-medium hover:underline transition-colors"
              >
                Lupa password?
              </button>
            </div>

            <!-- Submit Button -->
            <button
              type="submit"
              class="group relative w-full bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 text-white font-semibold py-4 px-6 rounded-2xl transition-all duration-300 hover:shadow-xl hover:scale-[1.02] disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none flex items-center justify-center gap-3"
              :disabled="loading"
            >
              <!-- Loading Spinner -->
              <div
                v-if="loading"
                class="animate-spin rounded-full h-5 w-5 border-2 border-white border-t-transparent"
              ></div>

              <!-- Login Icon -->
              <svg
                v-else
                class="w-5 h-5 group-hover:translate-x-1 transition-transform duration-200"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"
                />
              </svg>

              <span>{{ loading ? "Memproses..." : "Masuk" }}</span>
            </button>

            <!-- Social Login Divider -->
            <div class="relative my-8">
              <div class="absolute inset-0 flex items-center">
                <div class="w-full border-t border-gray-200"></div>
              </div>
              <div class="relative flex justify-center text-sm">
                <span class="px-4 bg-white text-gray-500 font-medium"
                  >atau masuk dengan</span
                >
              </div>
            </div>

            <!-- Social Login Buttons -->
            <div class="grid grid-cols-2 gap-4">
              <button
                type="button"
                class="flex items-center justify-center gap-3 px-4 py-3 bg-white border border-gray-200 rounded-xl hover:bg-gray-50 hover:border-gray-300 transition-all duration-200 font-medium text-gray-700"
              >
                <svg class="w-5 h-5" viewBox="0 0 24 24">
                  <path
                    fill="#4285F4"
                    d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                  />
                  <path
                    fill="#34A853"
                    d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                  />
                  <path
                    fill="#FBBC05"
                    d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                  />
                  <path
                    fill="#EA4335"
                    d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                  />
                </svg>
                <span class="hidden sm:inline">Google</span>
              </button>

              <button
                type="button"
                class="flex items-center justify-center gap-3 px-4 py-3 bg-white border border-gray-200 rounded-xl hover:bg-gray-50 hover:border-gray-300 transition-all duration-200 font-medium text-gray-700"
              >
                <svg class="w-5 h-5" fill="#1877F2" viewBox="0 0 24 24">
                  <path
                    d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"
                  />
                </svg>
                <span class="hidden sm:inline">Facebook</span>
              </button>
            </div>
          </form>

          <!-- Register Link -->
          <div class="mt-8 text-center">
            <p class="text-gray-600">
              Belum punya akun?
              <router-link
                to="/register"
                class="text-blue-600 hover:text-blue-700 font-semibold hover:underline transition-colors ml-1"
              >
                Daftar sekarang
              </router-link>
            </p>
          </div>
        </div>

        <!-- Footer -->
        <div class="text-center mt-8">
          <p class="text-sm text-gray-500">
            Dengan masuk, Anda menyetujui
            <button class="text-blue-600 hover:underline">
              Syarat & Ketentuan
            </button>
            dan
            <button class="text-blue-600 hover:underline">
              Kebijakan Privasi
            </button>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import api from "../lib/api";
import { useAuth } from "../stores/auth";
import { useRouter } from "vue-router";

const email = ref("");
const password = ref("");
const loading = ref(false);
const auth = useAuth();
const router = useRouter();

async function submit() {
  loading.value = true;
  try {
    const { data } = await api.post("/v1/auth/login", {
      email: email.value,
      password: password.value,
    });
    auth.setSession(data.access_token, data.user);
    router.push({ name: "landing" });
  } catch (e) {
    alert(e?.response?.data?.error || e.message);
  } finally {
    loading.value = false;
  }
}
</script>

<style scoped>
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

/* Custom focus styles */
.focus\:ring-4:focus {
  box-shadow: 0 0 0 4px var(--tw-ring-color);
}

/* Responsive improvements */
@media (max-width: 640px) {
  .backdrop-blur-xl {
    backdrop-filter: blur(12px);
  }
}

/* Button press effect */
.group:active {
  transform: scale(0.98);
}

/* Loading spinner animation */
@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.animate-spin {
  animation: spin 1s linear infinite;
}
</style>
