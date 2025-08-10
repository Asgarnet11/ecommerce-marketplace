<template>
  <nav>
    <ul class="space-y-1">
      <CategoryNode v-for="c in tree" :key="c.id" :node="c" @go="goCat" />
    </ul>
  </nav>
</template>

<script setup>
import { onMounted, ref } from "vue";
import api from "../lib/api";
import { useRouter } from "vue-router";

const router = useRouter();
const tree = ref([]);

async function load() {
  const { data } = await api.get("/v1/categories/tree");
  tree.value = data || [];
}
function goCat(slug) {
  router.push({ name: "products", query: { cat: slug } });
}
onMounted(load);
</script>

<script>
export default {
  components: {
    CategoryNode: {
      props: { node: { type: Object, required: true } },
      emits: ["go"],
      template: `
        <li>
          <button class="text-left w-full hover:underline" @click="$emit('go', node.slug)">
            {{ node.name }}
          </button>
          <ul v-if="node.children?.length" class="pl-4 border-l mt-1 space-y-1">
            <CategoryNode v-for="c in node.children" :key="c.id" :node="c" @go="$emit('go', $event)" />
          </ul>
        </li>
      `,
    },
  },
};
</script>
