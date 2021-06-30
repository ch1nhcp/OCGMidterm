<template>
  <div class="products">
    <div class="container">
      <!-- <h2 style="text-align:center">All items: {{ products.length }} items</h2>
      <h2 style="text-align:center">
        All onsale items: {{ onSaleCount() }} items
      </h2>
      <h2 style="text-align:center">
        Total price: {{ formatPrice(totalPrice()) }} $
      </h2>
      <h2 style="text-align:center">
        Average price: {{ formatPrice(averagePrice()) }} $
      </h2> -->

      <!-- Sort by... -->
      <div style="display: flex">
        <span>
          <p style="margin-left:30px; font-weight: bold">Sort by:</p>
          <select
            style="margin-left:30px"
            v-model="selected"
            @change="onChange"
          >
            <option disabled value="">Choose one</option>
            <option value="default">Default</option>
            <option value="priceAscending">Price, Low -> High </option>
            <option value="priceDecending">Price, High -> Low </option>
          </select>
        </span>
        <span style="flex:wrap;">
          <p style="margin-left:30px; font-weight: bold">Category:</p>
          <input
            style="margin-left:30px"
            type="checkbox"
            id="jack"
            value="Jack"
            v-model="checkedNames"
          />
          <label for="jack">Action</label>
          <input
            style="margin-left:15px"
            type="checkbox"
            id="john"
            value="John"
            v-model="checkedNames"
          />
          <label for="john">Romance</label>
          <input
            style="margin-left:15px"
            type="checkbox"
            id="john"
            value="John"
            v-model="checkedNames"
          />
          <label for="john">Sport</label>
          <input
            style="margin-left:15px"
            type="checkbox"
            id="john"
            value="John"
            v-model="checkedNames"
          />
          <label for="john">Family</label>
          <input
            style="margin-left:15px"
            type="checkbox"
            id="john"
            value="John"
            v-model="checkedNames"
          />
          <label for="john">Horror</label>
          <input
            style="margin-left:15px"
            type="checkbox"
            id="john"
            value="John"
            v-model="checkedNames"
          />
          <label for="john">Cattoon</label>
        </span>
      </div>

      <!-- <span style="margin-left: 50px">Sorted by: {{ selected }}</span> -->

      <!-- <span style="margin-left: 30px"
        >Search:
        <input type="text" v-on:keyup.enter="search" />
        <button style="cursor:pointer; width: 20px;" v-on:click="resetInput">
          x
        </button>
      </span> -->

      <div class="product-items">
        <!-- single product -->
        <div class="product" v-for="product in products" :key="product.id">
          <div class="product-content">
            <div class="product-img">
              <img :src="product.img" :alt="imageAlt" />
            </div>
            <div class="product-btns">
              <button type="button" class="btn-cart">
                add to cart
                <span><i class="fas fa-plus"></i></span>
              </button>
              <button type="button" class="btn-buy">
                buy now
                <span><i class="fas fa-shopping-cart"></i></span>
              </button>
            </div>
          </div>

          <div class="product-info">
            <a href="#" class="product-name">{{ product.name }}</a>
            <p v-show="product.isOnSale" class="product-price">
              $ {{ product.price }}
            </p>
            <p class="product-price">
              $ {{ (product.price * (100 - product.sale)) / 100 }}
            </p>
          </div>

          <div v-show="product.isOnSale" class="off-info">
            <h2 class="sm-title">{{ product.sale }}% off</h2>
          </div>
        </div>
        <!-- end of single product -->
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      products: [
        {
          id: 1,
          img: "https://i.imgur.com/zPHHQrL.png",
          name: "Product 3",
          price: 199,
          sale: 20,
          isOnSale: true,
        },
        {
          id: 2,
          img: "https://i.imgur.com/dBfEWoN.jpg",
          name: "Product 2",
          price: 299,
          sale: 10,
          isOnSale: true,
        },
        {
          id: 3,
          img: "https://i.imgur.com/Ujqq4tk.jpg",
          name: "Product 1",
          price: 169,
          sale: 0,
          isOnSale: false,
        },
        {
          id: 4,
          img: "https://i.imgur.com/3RdqlDa.jpg",
          name: "Product 4",
          price: 799,
          sale: 25,
          isOnSale: true,
        },
        {
          id: 5,
          img: "https://i.imgur.com/zPHHQrL.png",
          name: "Product 2021",
          price: 19,
          sale: 20,
          isOnSale: true,
        },
        {
          id: 6,
          img: "https://i.imgur.com/msp7gU8.jpg",
          name: "Product 6",
          price: 999,
          sale: 5,
          isOnSale: true,
        },
        {
          id: 7,
          img: "https://i.imgur.com/zPHHQrL.png",
          name: "Product 7",
          price: 150,
          sale: 20,
          isOnSale: true,
        },
        {
          id: 8,
          img: "https://i.imgur.com/msp7gU8.jpg",
          name: "Product 8",
          price: 99,
          sale: 0,
          isOnSale: false,
        },
      ],
      selected: "",
      imageAlt: "image",
      productClone: [],
    };
  },
  methods: {
    formatPrice(value) {
      let val = (value / 1).toFixed(2).replace(".", ",");
      return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
    },
    onSaleCount() {
      let count = 0;
      for (let i = 0; i < this.products.length; i++) {
        if (this.products[i].isOnSale == true) {
          count++;
        }
      }
      return count;
    },
    totalPrice() {
      let sum = 0;
      for (let i = 0; i < this.products.length; i++) {
        sum += this.products[i].price;
      }
      return sum;
    },

    averagePrice() {
      return this.totalPrice() / this.products.length;
    },

    onChange(event) {
      console.log(event.target.value);

      switch (event.target.value) {
        case "default":
          return this.products.sort(function(a, b) {
            return a.id - b.id;
          });
        case "priceAscending":
          return this.products.sort(function(a, b) {
            return a.price - b.price;
          });
        case "priceDecending":
          return this.products.sort(function(a, b) {
            return b.price - a.price;
          });
      }
    },

    resetInput() {
      location.reload();
    },

    search() {
      if (event.target.value != "") {
        // tranh truong hop khong nhap gi ma van tim duoc san pham
        let keyword = event.target.value; // gán keyword = giá trị nhập vào ô trống
        console.log(keyword.toLowerCase()); // in keyword ra trong Console

        if (this.productClone.length == 0) {
          this.productClone = this.products; // clone sang array mới
        }
        if (this.products.length < this.productClone.length) {
          this.products = this.productClone; // nếu array cũ mất đi sản phẩm thì gán = array clone
        }

        console.log(
          this.products.filter((product) => product.name.match(keyword)), // in ra mang phu hop voi ket qua tim kiem
          alert(
            "Find " +
              this.products.filter((product) => product.name.match(keyword))
                .length +
              " product !"
          )
        );

        //kiem tra neu khong tim duoc san pham thi Alert! sau do reset mang ve ban dau
        if (
          this.products.filter((product) => product.name.match(keyword))
            .length == 0
        ) {
          this.products = this.productClone;
          // alert("Cannot find this product!");
          event.target.value = "";
          return this.products;
        }

        //gan products = mang cac gia tri moi tim duoc
        return (this.products = this.products.filter((product) =>
          product.name.match(keyword)
        ));
      }
    },
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Quicksand:wght@300;400;500;600;700&family=Roboto:wght@300;400;500;700;900&display=swap");
@import url("https://fonts.googleapis.com/css2?family=Fredoka+One&display=swap");

:root {
  --white-light: rgba(255, 255, 255, 0.5);
  --alice-blue: #fff;
  --blue-main: #439cff;
  --gray: #ededed;
}
/* * {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}
body {
  font-family: "Quicksand", sans-serif;
} */

/* Utility stylings */
img {
  width: 100%;
  display: block;
}
.container {
  width: 88vw;
  margin: 0 auto;
}
.lg-title,
.md-title,
.sm-title {
  font-family: "Roboto", sans-serif;
  padding: 0.6rem 0;
  text-transform: capitalize;
}
.lg-title {
  font-size: 2.5rem;
  font-weight: 500;
  text-align: center;
  padding: 1.3rem 0;
  opacity: 0.9;
  color: #439cff;
  font-family: "Fredoka One", cursive;
}
.md-title {
  font-size: 2rem;
  font-family: "Roboto", sans-serif;
}
.sm-title {
  font-weight: 300;
  font-size: 1rem;
  text-transform: uppercase;
}
.text-light {
  font-size: 1rem;
  font-weight: 600;
  line-height: 1.5;
  opacity: 0.5;
  margin: 0.4rem 0;
}

/* product section */
.products {
  background: var(--alice-blue);
  padding: 3.2rem 0;
}
.products .text-light {
  text-align: center;
  width: 70%;
  margin: 0.9rem auto;
}
.product {
  margin: 2rem;
  position: relative;
  box-shadow: rgba(50, 50, 93, 0.25) 0px 6px 12px -2px,
    rgba(0, 0, 0, 0.3) 0px 3px 7px -3px;
  transition: cubic-bezier(0.165, 0.84, 0.44, 1);
}
.product:hover {
  transform: scale(1.01);
  -webkit-transition: all 0.2s 0s linear;
  -moz-transition: all 0.2s 0s linear;
  -o-transition: all 0.2s 0s linear;
  transition: all 0.2s 0s linear;
}
.product-content {
  background: var(--gray);
  padding: 3rem 0.5rem 2rem 0.5rem;
  cursor: pointer;
  /* border-radius: 20px; */
}
.product-img {
  background: var(--white-light);
  box-shadow: 0 0 20px 10px var(--white-light);
  width: 200px;
  height: 200px;
  margin: 0 auto;
  border-radius: 50%;
  /* transition: background 0.5s ease; */
}
.product-btns {
  display: flex;
  justify-content: center;
  margin-top: 1.4rem;
  opacity: 0;
  transition: opacity 0.6s ease;
}
.btn-cart,
.btn-buy {
  background: transparent;
  border: 1px solid black;
  padding: 0.8rem 0;
  width: 125px;
  font-family: inherit;
  text-transform: uppercase;
  cursor: pointer;
  border: none;
  transition: all 0.6s ease;
}
.btn-cart {
  background: black;
  color: white;
}
.btn-cart:hover {
  background: var(--blue-main);
}
.btn-buy {
  background: white;
}
.btn-buy:hover {
  background: var(--blue-main);
  color: #fff;
}
.product-info {
  background: white;
  padding: 2rem;
}
.product-info-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.product-name {
  color: black;
  display: block;
  text-decoration: none;
  font-size: 1rem;
  text-transform: uppercase;
  font-weight: bold;
}
.product-price {
  padding-top: 0.6rem;
  padding-right: 0.6rem;
  display: inline-block;
}
.product-price:first-of-type {
  text-decoration: line-through;
  color: var(--blue-main);
}
/* .product-img img {
  transition: transform 0.6s ease;
} */
/* .product:hover .product-img img {
  transform: scale(1.1);
} */
.product:hover .product-img {
  background: var(--blue-main);
}
.product:hover .product-btns {
  opacity: 1;
}
.off-info .sm-title {
  background: var(--blue-main);
  color: white;
  display: inline-block;
  padding: 0.5rem;
  position: absolute;
  top: 0;
  left: 0;
  writing-mode: vertical-lr;
  transform: rotate(180deg);
  z-index: 1;
  letter-spacing: 3px;
  cursor: pointer;
}

/* Media Queries */
@media screen and (min-width: 992px) {
  .product-items {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
  }
}
@media screen and (min-width: 1200px) {
  .product-items {
    grid-template-columns: repeat(3, 1fr);
  }
  .product {
    margin-right: 1rem;
    margin-left: 1rem;
  }
  .products .text-light {
    width: 50%;
  }
}

@media screen and (min-width: 1336px) {
  .product-items {
    grid-template-columns: repeat(4, 1fr);
  }
  .product-collection-wrapper {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
  }
  .flex {
    height: 60vh;
  }
}
</style>
