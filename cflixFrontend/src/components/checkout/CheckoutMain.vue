<template>
  <div class="container cart">
    <!-- <h1>Giỏ hàng</h1> -->
    <div class="none-cart" v-if="!products.length">
      Nothing in your cart.
      <button @click="resetPage">Shopping now</button>
    </div>
    <ProductList
      v-if="products.length"
      :products="products"
      @change-num-product="changeNumProduct"
      @up-num-product="upNumProduct"
      @down-num-product="downNumProduct"
      @delete-product="deleteProductById"
    />
    <Checkout :products="products" v-if="products.length" />
  </div>
</template>

<script>
import Checkout from "./Checkout.vue";
import ProductList from "./ProductList.vue";
export default {
  name: "App",
  components: {
    Checkout,
    ProductList,
  },
  data() {
    return {
      products: [
        {
          id: 1,
          name: "Coco",
          price: 28,
          description: "Nội dung mô tả sản phẩm",
          num: 2,
          image: "https://i.imgur.com/Ujqq4tk.jpg",
        },
        {
          id: 2,
          name: "End Game",
          price: 82,
          num: 3,
          description: "Nội dung mô tả sản phẩm",
          image: "https://i.imgur.com/Ujqq4tk.jpg",
        },
        {
          id: 3,
          name: "Wonder Woman",
          price: 20,
          num: 4,
          description: "Nội dung mô tả sản phẩm",
          image: "https://i.imgur.com/Ujqq4tk.jpg",
        },
      ],
      productsClone: [],
    };
  },
  methods: {
    changeNumProduct(num, idProduct) {
      this.products = this.products.map((product) => {
        if (product.id === idProduct) {
          product.num = num;
        }
        return product;
      });
    },
    upNumProduct(num, idProduct) {
      this.changeNumProduct(num, idProduct);
    },
    downNumProduct(num, idProduct) {
      this.changeNumProduct(num, idProduct);
    },
    deleteProductById(idProduct) {
      this.products = this.products.filter(
        (product) => product.id !== idProduct
      );
    },
    resetPage() {
      this.products = this.productsClone.map((product) => ({ ...product }));
    },
  },
  created() {
    this.productsClone = this.products.map((product) => ({ ...product }));
  },
};
</script>

<style scoped>
#CheckoutMain {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
body {
  background-color: #fff;
}
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}
.container cart{
  width: 60%;
}
.cart {
  display: flex;
  flex-wrap: wrap;
  position: relative;
}

h1 {
  font-size: 20px;
  flex: 100%;
}
input,
button {
  padding: 5px 10px;
  border: none;
  outline: none;
  text-align: center;
}
input {
  border-top: 1px solid rgb(145, 145, 145);
  border-bottom: 1px solid rgb(145, 145, 145);
}
button {
  background-color: #1F9AFF;
  color: white;
  border-top: 1px solid #1F9AFF;
  border-bottom: 1px solid #1F9AFF;
  cursor: pointer;
  transition: all 0.2s ease-in-out;
}
button:hover {
  opacity: 0.8;
}
.none-cart {
  width: 100%;
  height: 80vh;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
.none-cart button {
  margin-top: 20px;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
