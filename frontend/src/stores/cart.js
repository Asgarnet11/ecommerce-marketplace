import { defineStore } from "pinia";

export const useCart = defineStore("cart", {
  state: () => ({ items: [] }),
  getters: {
    subtotal: (s) => s.items.reduce((a, it) => a + it.price * it.qty, 0),
  },
  actions: {
    add(p) {
      const f = this.items.find((it) => it.id === p.id);
      if (f) f.qty += 1;
      else
        this.items.push({
          id: p.id,
          title: p.title,
          price: p.price,
          qty: 1,
          slug: p.slug,
        });
    },
    remove(id) {
      this.items = this.items.filter((it) => it.id !== id);
    },
    update(id, qty) {
      const it = this.items.find((i) => i.id === id);
      if (it) it.qty = qty;
    },
  },
});
