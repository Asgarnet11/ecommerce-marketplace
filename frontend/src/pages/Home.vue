<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navbar -->
    <Navbar />

    <!-- Hero Section -->
    <HeroSection />

    <!-- Categories Section -->
    <CategoriesSection />

    <!-- Featured Products -->
    <section class="py-12 px-4 max-w-7xl mx-auto">
      <h2 class="text-3xl font-bold text-gray-800 mb-8 text-center">
        Produk Unggulan
      </h2>
      <ProductsGrid />
    </section>

    <!-- Promo Banner -->
    <PromoBanner />

    <!-- Services Section -->
    <ServicesSection />

    <!-- Newsletter -->
    <NewsletterSection />

    <!-- Footer -->
    <Footer />
  </div>
</template>

<script setup>
import Navbar from "../components/Navbar.vue";
import HeroSection from "../components/HeroSection.vue";
import CategoriesSection from "../components/CategoriesSection.vue";
import ProductsGrid from "../components/ProductsGrid.vue";
import PromoBanner from "../components/PromoBanner.vue";
import ServicesSection from "../components/ServicesSection.vue";
import NewsletterSection from "../components/NewsletterSection.vue";
import Footer from "../components/Footer.vue";
</script>
