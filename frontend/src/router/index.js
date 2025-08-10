import { createRouter, createWebHistory } from "vue-router";
import Home from "../pages/Home.vue";
import Products from "../pages/Products.vue";
import ProductDetail from "../pages/ProductDetail.vue";
import Cart from "../pages/Cart.vue";
import Shops from "../pages/Shops.vue";
import ShopDetail from "../pages/ShopDetail.vue";

const routes = [
  { path: "/", name: "home", component: Home },
  { path: "/products", name: "products", component: Products },
  { path: "/p/:slug", name: "product", component: ProductDetail },
  { path: "/cart", name: "cart", component: Cart },
  { path: "/shops", name: "shops", component: Shops },
  { path: "/shops/:slug", name: "shop", component: ShopDetail },
];

export default createRouter({ history: createWebHistory(), routes });
