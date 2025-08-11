import { defineStore } from "pinia";
import api from "../lib/api";
import { useAuth } from "./auth";

export const useCartApi = defineStore("cartApi", {
  state: () => ({
    items: [],
    subtotal: 0,
    loading: false,
    error: "",
  }),
  actions: {
    ensureAuth() {
      const auth = useAuth();
      if (!auth.isLogged) throw new Error("Silakan login terlebih dahulu.");
    },
    async fetch() {
      try {
        this.ensureAuth();
        this.loading = true;
        const { data } = await api.get("/v1/cart");
        this.items = data.items || [];
        this.subtotal = data.subtotal || 0;
      } finally {
        this.loading = false;
      }
    },
    async add(productId, qty = 1) {
      this.ensureAuth();
      await api.post("/v1/cart/items", { product_id: productId, qty });
      await this.fetch();
    },
    async update(productId, qty) {
      this.ensureAuth();
      await api.put("/v1/cart/items", { product_id: productId, qty });
      await this.fetch();
    },
    async remove(productId) {
      this.ensureAuth();
      await api.delete("/v1/cart/items", { params: { product_id: productId } });
      await this.fetch();
    },
  },
});
