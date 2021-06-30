<template>
  <div class="wrap-cart-info">
    <p class="hide">Hide</p>
    <div>
      <div class="wrap-promo-code">
        <p>Promotion code</p>
        <input type="text" v-model="promoCode" />
        <button @click="applyPromoCode">Apply</button>
        <div>
          <small>{{ alertInput }}</small>
        </div>
      </div>
      <div class="wrap-checkout">
        <div class="wrap-price-subtotal">
          <span>Subtotal</span>
          <span>${{ formatPrice(subTotal) }}</span>
        </div>
        <div class="wrap-price-sale" v-if="percentSale">
          <span>Sale {{ percentSale }}%</span>
          <span>- ${{ formatPrice(priceSale) }}</span>
        </div>
        <div class="wrap-price-tax">
          <span>Tax (10%)</span>
          <span>${{ formatPrice(tax) }} </span>
        </div>
        <div class="wrap-price-total">
          <span>Total Price</span>
          <span>${{ formatPrice(totalPrice) }}</span>
        </div>
        <button class="btn-checkout">CHECK OUT</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Checkout",
  props: {
    products: Object,
  },
  data() {
    return {
      vat: 10,
      promoCodes: [
        { code: "test", percent: 20 },
        { code: "chinhcp", percent: 80 },
      ],
      percentSale: 0,
      promoCode: "",
      alertInput: "",
    };
  },
  computed: {
    subTotal() {
      return this.products.reduce(
        (sumPrice, product) => sumPrice + product.price * product.num,
        0
      );
    },
    priceSale() {
      return (this.subTotal * this.percentSale) / 100;
    },
    tax() {
      return (this.subTotal * this.vat) / 100;
    },
    totalPrice() {
      return this.subTotal - this.priceSale + this.tax;
    },
  },
  watch: {
    promoCode(newState) {
      if (newState === "") {
        this.alertInput = "";
      }
    },
  },
  methods: {
    formatPrice(value) {
      let val = (value / 1).toFixed(2).replace(".", ",");
      return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
    },
    applyPromoCode() {
      let objPromoCode = this.promoCodes.find(
        (element) => element.code.toLowerCase() === this.promoCode.toLowerCase()
      );
      this.percentSale = objPromoCode ? objPromoCode.percent : 0;
      if (this.percentSale) {
        this.alertInput = `You have discount ${this.percentSale}%`;
      } else {
        this.alertInput = "Exprired or invalid code!";
      }
      if (!this.promoCode) {
        this.alertInput = "Invalid code!";
      }
    },
  },
};
</script>

<style scoped>
.wrap-cart-info {
  flex: 30%;
  padding-left: 40px;
  box-sizing: border-box;
  background-color: #f4f4f4;
  position: sticky;
  top: 0;
  padding-right: 40px;
  margin-top: 1rem;
  /* box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px; */
}
.wrap-cart-info > div {
  background-color: #f4f4f4;
  position: sticky;
  top: 0;
}
.hide {
  visibility: hidden;
}
.wrap-promo-code {
  border-bottom: 1px solid rgb(233, 233, 233);
  padding-bottom: 30px;
}
.wrap-promo-code p {
  padding-top: 15px;
}
.wrap-promo-code input {
  width: inherit;
  border-left: 1px solid rgb(145, 145, 145);
  text-align: left;
}

.wrap-checkout {
  padding-top: 20px;
}
.wrap-price-tax,
.wrap-price-subtotal,
.wrap-price-total,
.wrap-price-sale {
  display: flex;
  justify-content: space-between;
  padding: 15px 0;
  font-weight: bold;
}

.wrap-price-total {
  font-size: 1.2rem;
}
.btn-checkout {
  width: 100%;
  padding: 15px 0;
  cursor: pointer;
  outline: none;
  border: none;
  background-color: #1f9aff;
  color: white;
  font-weight: 700;
  font-size: 20px;
  border-radius: 50px;
}
small {
  padding-top: 15px;
}
@media only screen and (max-width: 992px) {
  .wrap-cart-info {
    flex: 100%;
    padding-left: 0;
  }
}
</style>
