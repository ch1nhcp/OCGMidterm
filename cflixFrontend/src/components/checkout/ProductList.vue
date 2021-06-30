<template>
  <div class="wrap-cart-product">
    <div class="product-list">
      <!-- info cart -->
      <p class="info">
        <span class="break">Home >> Shopping Cart</span>
        <span>
          <span class="item-count"> {{ totalNumProduct }}</span> items
        </span>
      </p>
      <!-- end info cart -->

      <!-- product wrapper -->
      <div class="product-item" v-for="product in products" :key="product.id">
        <div class="product-image">
          <img :src="product.image" alt="image" />
        </div>
        <div class="product-detail">
          <div class="product-name">{{ product.name }}</div>
          <div class="product-description">{{ product.description }}</div>
          <div class="product-price">${{ formatPrice(product.price) }}</div>
          <div class="product-counter">
            <!-- <div class="product-counter-mobile"> -->
            <button
              class="btn-decrease"
              @click="downNumProduct(--product.num, product.id)"
            >
              ▼
            </button>
            <input
              type="number"
              :value="product.num"
              @input="changeNumProduct($event, product.id)"
              min="1"
            />
            <button
              class="btn-increase"
              @click="upNumProduct(++product.num, product.id)"
            >
              ▲
            </button>
            <!-- </div> -->
          </div>
        </div>
        <div class="product-action">
          <button class="btn-delete" @click="openModal(product.id)">
            Delete
          </button>
        </div>
      </div>
    </div>
  </div>
  <transition name="fade">
    <CompModal v-if="isOpenModal" @confirm="confirmDeleteProduct" />
  </transition>
</template>

<script>
import CompModal from "./AlertModal.vue";

export default {
  name: "ProductList",
  components: {
    CompModal,
  },
  props: {
    products: Object,
  },
  data() {
    return {
      numProduct: 0,
      isOpenModal: false,
      idProduct: 0,
    };
  },
  computed: {
    totalNumProduct() {
      return this.products.reduce((num, product) => num + product.num, 0);
    },
  },
  methods: {
    formatPrice(value) {
      let val = (value / 1).toFixed(2).replace(".", ",");
      return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
    },

    changeNumProduct(event, idProduct) {
      let num = event.target.value;
      if (num !== "") {
        this.$emit("change-num-product", parseInt(num), idProduct);
      }
    },
    upNumProduct(num, idProduct) {
      this.$emit("up-num-product", parseInt(num), idProduct);
    },
    downNumProduct(num, idProduct) {
      if (num === 0) {
        this.$emit("down-num-product", 1, idProduct);
      } else {
        this.$emit("down-num-product", parseInt(num), idProduct);
      }
    },
    openModal(idProduct) {
      this.isOpenModal = true;
      this.idProduct = idProduct;
    },
    confirmDeleteProduct(isConfirm) {
      if (isConfirm) {
        this.$emit("delete-product", this.idProduct);
      }
      this.isOpenModal = false;
    },
    handleOnBlur(event, idProduct) {
      let num = event.target.value;
      if (num === "") {
        this.$emit("down-num-product", 1, idProduct);
      }
    },
  },
};
</script>

<style scoped>
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
input[type="number"] {
  -moz-appearance: textfield;
}
.wrap-cart-product {
  flex: 60%;
}
.product-list {
  margin: 10px 0;
}
.product-item {
  margin-right: 1rem;
  padding: 30px;
  display: flex;
  margin-bottom: 1rem;
  box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
}
.product-image {
  width: 100px;
  height: 150px;
}
img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.product-detail {
  flex: 50%;
  padding-left: 15px;
  padding-top: 15px;
}
.product-name {
  font-weight: 600;
  padding-bottom: 5px;
  font-size: 25px;
}
a {
  text-decoration: none;
}

.product-description {
  font-size: 13px;
}
.product-price {
  padding: 10px 0;
  font-weight: 600;
  font-weight: 25px;
}
.btn-delete:hover {
  opacity: 0.6;
}
.product-counter {
  flex: 10%;
  display: flex;
  align-items: center;
  margin-top: 15px;
}
input {
  width: 50px;
  border: none;
}

.product-action {
  flex: 10%;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
i {
  cursor: pointer;
}
.btn-delete {
  background: rgb(46, 45, 45);
  border: none;
  width: 80px;
  padding: 10px;
}
.info {
  display: flex;
  justify-content: space-between;
  padding: 15px 0;
  background-color: #f4f4f4;
  position: sticky;
  padding: 1rem;
}
.item-count {
  color: red;
}
</style>
