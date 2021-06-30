<template>
  <Slider />
  <ProductsCarousel category="Trending" />
  <ProductsCarousel category="Best Selling" />
  <ProductsCarousel category="Latest films" />
  <ProductsCarousel category="Release soon" />
</template>

<script>
import Slider from "./Slider.vue";
import ProductsCarousel from "./ProductsCarousel.vue";

export default {
  components: {
    Slider,
    ProductsCarousel,
  },
};
</script>

<style></style>
