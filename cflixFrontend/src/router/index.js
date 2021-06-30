import { createRouter, createWebHistory } from "vue-router";
import Body from "../components/Body.vue";
import About from "../components/About.vue";
import ProductsCollection from "../components/ProductsCollection.vue";
import CheckoutMain from "../components/checkout/CheckoutMain.vue";
import ProductDetail from "../components/ProductDetail.vue";
import News from "../components/News.vue";

const routes = [
  { path: "/", component: Body },
  { path: "/films", component: ProductsCollection },
  { path: "/about", component: About },
  { path: "/checkout", component: CheckoutMain },
  { path: "/product/:id", component: ProductDetail },
  { path: "/news", component: News },
];

const router = createRouter({
  history: createWebHistory(),
  scrollBehavior() {
    return { top: 0 };
  },
  routes,
});

export default router;
