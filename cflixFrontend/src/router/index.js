import { createRouter, createWebHistory } from "vue-router";
import Body from "../components/Body.vue";
import About from "../components/About.vue";
import ProductsCollection from "../components/ProductsCollection.vue";
import CheckoutMain from "../components/checkout/CheckoutMain.vue";
import ProductsDetail from "../components/ProductDetail.vue";

const routes = [
  { path: "/", component: Body },
  { path: "/films", component: ProductsCollection },
  { path: "/about", component: About },
  { path: "/checkout", component: CheckoutMain },
  { path: "/productsdetail", component: ProductsDetail },
];
 
const router = createRouter({
  history: createWebHistory(),
  scrollBehavior() {
    return { top: 0 };
  },
  routes,
});

export default router;
