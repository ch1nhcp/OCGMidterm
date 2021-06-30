<template>
  <!-- product-detail content -->
  <!-- <div class="bg-main">
        <div class="container"> -->
          <div v-for="product in products" :key="product.id">
  <div class="row product-row "  v-if="product.Id == $route.params.id" >
    <div class="col-5 col-md-12" >
      <div class="product-img" id="product-img" style="margin:auto">
        <img :src="product.Img" alt="" />
      </div>
      <div class="box">
        <div class="product-img-list">
          <div class="product-img-item">
            <img src="../assets/images/supergirl-banner.jpg" alt="" />
          </div>
          <div class="product-img-item">
            <img src="../assets/images/black-banner.png" alt="" />
          </div>
          <div class="product-img-item">
            <img src="../assets/images/transformer-banner.jpg" alt="" />
          </div>
        </div>
      </div>
    </div>
    <div class="col-7 col-md-12"   >
      <div class="product-info" >
        <div>
        <h1>
          {{product.Name}}
        </h1>
        <div class="product-info-detail">
          <span class="product-info-detail-title">Category: </span>
          <a href="#">Action, </a>
          <a href="#"> Scifi</a>
        </div>
        <div class="product-info-detail">
          <span class="product-info-detail-title">Rated:</span>
          <span class="rating">
            <i class="bx bxs-star"></i>
            <i class="bx bxs-star"></i>
            <i class="bx bxs-star"></i>
            <i class="bx bxs-star"></i>
            <i class="bx bxs-star"></i>
          </span>
        </div>
        <p class="product-description">
          Lorem, ipsum dolor sit amet consectetur adipisicing elit. Quo libero
          alias officiis dolore doloremque eveniet culpa dignissimos, itaque,
          cum animi excepturi sed veritatis asperiores soluta, nisi atque quae
          illum. Ipsum.
        </p>
        <div class="product-info-price">${{product.Saleprice}}</div>
        <div class="product-quantity-wrapper">
          <span class="product-quantity-btn">
            <i class="bx bx-minus"></i>
          </span>
          <span class="product-quantity">1</span>
          <span class="product-quantity-btn">
            <i class="bx bx-plus"></i>
          </span>
        </div>
        <div>
          <button class="btn-flat btn-hover" @click="greet">
            add to cart
          </button>
        </div>
        </div>
      </div>
    </div>
  </div>
  </div>
  <!-- </div> -->
</template>

<script>
// import "./css/app.css";
import "./css/grid.css";
export default {
   data() {
    return {
      products: [],
    };
  },
  methods: {
    greet: function() {
      alert("Add to cart successfully!! ");
    },
  },
    async created() {
    const response = await fetch("http://127.0.0.1:3000/film");
    this.products = await response.json();
    // axios.get("http://127.0.0.1:3000/product").then((response) => {
    //   this.products = Object.values(response.data);
    //   console.log(this.products);
    // });
  },
};  
</script>

<style scoped src="./css/app.css">

</style>
