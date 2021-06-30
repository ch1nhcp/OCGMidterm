<template>
  <!-- header -->
  <header>
    <!-- mobile menu -->
    <div class="mobile-menu bg-second">
      <a href="#" class="mb-logo">CFLIX</a>
      <span class="mb-menu-toggle" id="mb-menu-toggle">
        <i class="bx bx-menu"></i>
      </span>
    </div>
    <!-- end mobile menu -->
    <!-- main header -->
    <div class="header-wrapper" id="header-wrapper">
      <span class="mb-menu-toggle mb-menu-close" id="mb-menu-close">
        <i class="bx bx-x"></i>
      </span>
      <!-- top header -->
      <div class="bg-second">
        <div class="top-header container">
          <ul class="devided">
            <li>
              <a href="#">+84862332808</a>
            </li>
            <li>
              <a href="#">cflix@gmail.com</a>
            </li>
          </ul>
        </div>
      </div>
      <!-- end top header -->
      <!-- mid header -->
      <div class="bg-main">
        <div class="mid-header container">
          <a href="/" class="logo">CFLIX</a>
          <div class="search">
            <input type="text" placeholder="Search"/>
            <i class="bx bx-search-alt"></i>
          </div>
          <ul class="user-menu">
            <li>
              <router-link to="/checkout"><i class="bx bx-cart"></i></router-link>
              <!-- <a href="#"><i class="bx bx-cart"></i></a> -->
            </li>
            <li>
              <a href="#"><i class="bx bx-user-circle"></i></a>
            </li>
          </ul>
        </div>
      </div>
      <!-- end mid header -->
      <!-- bottom header -->
      <div class="bg-second">
        <div class="bottom-header container">
          <ul class="main-menu">
            <li><router-link to="/" class="active">Home</router-link></li>
            <li><router-link to="/films" class="active">Films</router-link></li>
            <li><router-link to="/news" class="active">News</router-link></li>
            <li><router-link to="/about" class="active">About</router-link></li>
            <!-- <li><a href="#" style="padding: 0 1rem">Home</a></li>
            <li><a href="#" style="padding: 0 1rem">Films</a></li>
            <li><a href="#" style="padding: 0 1rem">News</a></li>
            <li><a href="#" style="padding: 0 1rem">About</a></li> -->
          </ul>
        </div>
      </div>
      <!-- end bottom header -->
    </div>
    <!-- end main header -->
  </header>
  <!-- end header -->
</template>

<script>
// import "./css/app.css";
import "./css/grid.css";
export default {
  name: "Header",
  mounted() {
    document.querySelectorAll(".dropdown > a").forEach((e) => {
      e.addEventListener("click", (event) => event.preventDefault());
    });

    document.querySelectorAll(".mega-dropdown > a").forEach((e) => {
      e.addEventListener("click", (event) => event.preventDefault());
    });

    document
      .querySelector("#mb-menu-toggle")
      .addEventListener("click", () =>
        document.querySelector("#header-wrapper").classList.add("active")
      );

    document
      .querySelector("#mb-menu-close")
      .addEventListener("click", () =>
        document.querySelector("#header-wrapper").classList.remove("active")
      );
  },
};
</script>

<style scoped src="./css/app.css"></style>
