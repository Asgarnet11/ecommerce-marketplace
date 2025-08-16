<template>
  <header
    class="sticky top-0 z-30 bg-white/95 backdrop-blur-md border-b border-gray-200 shadow-sm"
  >
    <div class="container mx-auto px-4 lg:px-6">
      <div class="flex h-16 items-center justify-between">
        <!-- Logo Section -->
        <router-link
          to="/"
          class="flex items-center gap-3 font-bold text-xl text-gray-900 hover:text-brand transition-colors duration-200"
        >
          <div class="relative">
            <img
              src="../assets/logo.png"
              alt="Market Logo"
              class="w-10 h-10 object-contain rounded-xl bg-gradient-to-br from-brand/10 to-brand/5 p-1"
              onerror="this.style.display='none'; this.nextElementSibling.style.display='flex'"
            />
            <div
              class="hidden w-10 h-10 rounded-xl bg-gradient-to-br from-blue-500 to-purple-600 items-center justify-center"
            >
              <span class="text-white font-bold text-lg">M</span>
            </div>
          </div>
          <span class="hidden sm:block">XAN | SHOP</span>
        </router-link>

        <!-- Desktop Navigation -->
        <nav class="hidden md:flex items-center gap-8">
          <router-link
            class="text-gray-700 hover:text-brand font-medium transition-all duration-200 hover:scale-105 relative group"
            to="/products"
          >
            Produk
            <span
              class="absolute bottom-0 left-0 w-0 h-0.5 bg-brand transition-all duration-200 group-hover:w-full"
            ></span>
          </router-link>
          <router-link
            class="text-gray-700 hover:text-brand font-medium transition-all duration-200 hover:scale-105 relative group"
            to="/shops"
          >
            Toko
            <span
              class="absolute bottom-0 left-0 w-0 h-0.5 bg-brand transition-all duration-200 group-hover:w-full"
            ></span>
          </router-link>
          <form
            class="hidden md:flex items-center gap-2"
            @submit.prevent="goSearch"
          >
            <input
              v-model="q"
              placeholder="Cari produk atau toko…"
              class="border rounded-xl px-3 py-2 w-64"
            />
            <button class="btn">Cari</button>
          </form>
        </nav>

        <!-- Desktop Auth & Cart Section -->
        <div class="hidden md:flex items-center gap-3">
          <!-- Cart Button -->
          <router-link
            to="/cart"
            class="relative p-2 text-gray-600 hover:text-brand hover:bg-gray-50 rounded-lg transition-all duration-200 group"
          >
            <svg
              class="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-1.5 6M7 13l-1.5 6m9.5-6h4.5M17 13v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6"
              />
            </svg>
            <span
              class="absolute -top-1 -right-1 w-2 h-2 bg-red-500 rounded-full opacity-0 group-hover:opacity-100 transition-opacity"
            ></span>
          </router-link>
          <router-link
            v-if="auth.isLogged"
            class="text-gray-700 hover:text-brand font-medium transition-all duration-200 hover:scale-105 relative group"
            to="/orders"
          >
            Pesanan
            <span
              class="absolute bottom-0 left-0 w-0 h-0.5 bg-brand transition-all duration-200 group-hover:w-full"
            ></span>
          </router-link>

          <!-- Auth Buttons -->
          <template v-if="!auth.isLogged">
            <router-link
              to="/login"
              class="px-4 py-2 text-gray-700 hover:text-brand hover:bg-gray-50 rounded-lg font-medium transition-all duration-200"
            >
              Masuk
            </router-link>
            <router-link
              to="/register"
              class="px-6 py-2 bg-gradient-to-r from-blue-500 to-purple-600 text-white rounded-lg font-medium hover:shadow-lg hover:scale-105 transition-all duration-200"
            >
              Daftar
            </router-link>
          </template>
          <template v-else>
            <div class="flex items-center gap-3">
              <span class="text-sm text-gray-600 font-medium">
                Hi, <span class="text-brand">{{ auth.user?.name }}</span>
              </span>
              <button
                class="px-4 py-2 text-gray-700 hover:text-red-500 hover:bg-red-50 rounded-lg font-medium transition-all duration-200"
                @click="logout"
              >
                Logout
              </button>
            </div>
          </template>
        </div>

        <!-- Mobile Menu Button -->
        <button
          class="md:hidden p-2 text-gray-600 hover:text-brand hover:bg-gray-50 rounded-lg transition-all duration-200"
          @click="open = !open"
          aria-label="Toggle menu"
        >
          <svg
            class="w-6 h-6"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              :d="open ? 'M6 18L18 6M6 6l12 12' : 'M4 6h16M4 12h16M4 18h16'"
            />
          </svg>
        </button>
      </div>

      <!-- Mobile Menu -->
      <transition
        enter-active-class="transition-all duration-300 ease-out"
        enter-from-class="opacity-0 transform -translate-y-2"
        enter-to-class="opacity-100 transform translate-y-0"
        leave-active-class="transition-all duration-200 ease-in"
        leave-from-class="opacity-100 transform translate-y-0"
        leave-to-class="opacity-0 transform -translate-y-2"
      >
        <div
          v-if="open"
          class="md:hidden border-t border-gray-200 bg-white/95 backdrop-blur-md"
        >
          <div class="container mx-auto px-4 py-6">
            <!-- Mobile Navigation Links -->
            <div class="space-y-4 mb-6">
              <router-link
                class="flex items-center gap-3 p-3 text-gray-700 hover:text-brand hover:bg-gray-50 rounded-lg font-medium transition-all duration-200"
                to="/products"
                @click="open = false"
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
                    d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10"
                  />
                </svg>
                Produk
              </router-link>
              <router-link
                class="flex items-center gap-3 p-3 text-gray-700 hover:text-brand hover:bg-gray-50 rounded-lg font-medium transition-all duration-200"
                to="/shops"
                @click="open = false"
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
                Toko
              </router-link>
              <form
                class="hidden md:flex items-center gap-2"
                @submit.prevent="goSearch"
              >
                <input
                  v-model="q"
                  placeholder="Cari produk atau toko…"
                  class="border rounded-xl px-3 py-2 w-64"
                />
                <button class="btn">Cari</button>
              </form>
              <router-link
                class="flex items-center gap-3 p-3 text-gray-700 hover:text-brand hover:bg-gray-50 rounded-lg font-medium transition-all duration-200"
                to="/cart"
                @click="open = false"
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
                    d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-1.5 6M7 13l-1.5 6m9.5-6h4.5M17 13v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6"
                  />
                </svg>
                Keranjang
              </router-link>
            </div>

            <!-- Mobile Auth Section -->
            <div class="pt-4 border-t border-gray-200">
              <template v-if="!auth.isLogged">
                <div class="grid grid-cols-2 gap-3">
                  <router-link
                    to="/login"
                    class="text-center px-4 py-3 text-gray-700 hover:text-brand hover:bg-gray-50 rounded-lg font-medium transition-all duration-200 border border-gray-200"
                    @click="open = false"
                  >
                    Masuk
                  </router-link>
                  <router-link
                    to="/register"
                    class="text-center px-4 py-3 bg-gradient-to-r from-blue-500 to-purple-600 text-white rounded-lg font-medium hover:shadow-lg transition-all duration-200"
                    @click="open = false"
                  >
                    Daftar
                  </router-link>
                </div>
              </template>
              <template v-else>
                <div class="space-y-3">
                  <div class="text-center p-3 bg-gray-50 rounded-lg">
                    <span class="text-sm text-gray-600">
                      Selamat datang,
                      <span class="text-brand font-medium">{{
                        auth.user?.name
                      }}</span>
                    </span>
                  </div>
                  <button
                    class="w-full p-3 text-red-600 hover:text-white hover:bg-red-500 rounded-lg font-medium transition-all duration-200 border border-red-200"
                    @click="logout"
                  >
                    Logout
                  </button>
                </div>
              </template>
            </div>
          </div>
        </div>
      </transition>
    </div>
  </header>
</template>

<script setup>
import { ref } from "vue";
import { useAuth } from "../stores/auth";
import { useRouter } from "vue-router";

const auth = useAuth();
const open = ref(false);
const router = useRouter();

function logout() {
  auth.logout();
  open.value = false;
}
function goSearch() {
  if (q.value.trim())
    router.push({ name: "search", query: { q: q.value.trim() } });
}
</script>

<style scoped>
.container {
  max-width: 1200px;
}

/* Custom brand color - you can adjust this in your CSS variables */
.text-brand {
  color: var(--brand-color, #3b82f6);
}

.hover\:text-brand:hover {
  color: var(--brand-color, #3b82f6);
}

.bg-brand {
  background-color: var(--brand-color, #3b82f6);
}

.from-brand\/10 {
  --tw-gradient-from: rgb(59 130 246 / 0.1);
}

.to-brand\/5 {
  --tw-gradient-to: rgb(59 130 246 / 0.05);
}

/* Additional responsive improvements */
@media (max-width: 640px) {
  .container {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}
</style>
