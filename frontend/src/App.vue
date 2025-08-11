<template>
  <div>
    <NavBar />
    <main>
      <router-view />
    </main>
    <footer class="mt-16 border-t py-10 text-sm text-gray-500">
      <div class="container">
        © {{ new Date().getFullYear() }} Market — All rights reserved.
      </div>
    </footer>
  </div>
</template>
<script setup>
import NavBar from "./components/NavBar.vue";
</script>
