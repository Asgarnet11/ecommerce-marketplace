import { createRouter, createWebHistory } from "vue-router";

const Landing = () => import("../pages/Landing.vue");
const Catalog = () => import("../pages/Catalog.vue");
const Product = () => import("../pages/ProductDetail.vue");
const Shops = () => import("../pages/Shops.vue");
const ShopDetail = () => import("../pages/ShopDetail.vue");
const Cart = () => import("../pages/Cart.vue");
const Checkout = () => import("../pages/Checkout.vue");
const Login = () => import("../pages/Login.vue");
const Register = () => import("../pages/Register.vue");
const Orders = () => import("../pages/Orders.vue");
const OrderDetail = () => import("../pages/OrderDetail.vue");
const Search = () => import("../pages/Search.vue");

const routes = [
  { path: "/", name: "landing", component: Landing, meta: { public: true } },
  {
    path: "/products",
    name: "catalog",
    component: Catalog,
    meta: { public: true },
  },
  {
    path: "/p/:slug",
    name: "product",
    component: Product,
    meta: { public: true },
  },
  { path: "/shops", name: "shops", component: Shops, meta: { public: true } },
  {
    path: "/shops/:slug",
    name: "shop",
    component: ShopDetail,
    meta: { public: true },
  },
  { path: "/cart", name: "cart", component: Cart },
  { path: "/checkout", name: "checkout", component: Checkout },
  { path: "/login", name: "login", component: Login, meta: { public: true } },
  {
    path: "/register",
    name: "register",
    component: Register,
    meta: { public: true },
  },
  { path: "/orders", name: "orders", component: Orders },
  { path: "/orders/:code", name: "order-detail", component: OrderDetail },
  {
    path: "/search",
    name: "search",
    component: Search,
    meta: { public: true },
  },
];

const router = createRouter({ history: createWebHistory(), routes });

export default router;
