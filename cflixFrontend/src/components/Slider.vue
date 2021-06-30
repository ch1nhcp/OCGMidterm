<template>
  <link rel="preconnect" href="https://fonts.gstatic.com" />
  <link
    href="https://fonts.googleapis.com/css2?family=Cairo:wght@200;300;400;600;700;900&display=swap"
    rel="stylesheet"
  />
  <!-- MAIN HTML CONTENT -->

  <!-- HERO SECTION -->
  <div class="hero-section">
    <!-- HERO SLIDE -->
    <div class="hero-slide">
      <div class="owl-carousel carousel-nav-center" id="hero-carousel">
        <!-- SLIDE ITEM -->
        <div class="hero-slide-item">
          <img src="../assets/images/black-banner.png" alt="" />
          <div class="overlay"></div>
          <div class="hero-slide-item-content">
            <div class="item-content-wraper">
              <div class="item-content-title top-down">
                Black Panther
              </div>
              <div class="movie-infos top-down delay-2">
                <div class="movie-info">
                  <i class="bx bxs-star"></i>
                  <span>9.5</span>
                </div>
                <div class="movie-info">
                  <i class="bx bxs-time"></i>
                  <span>120 mins</span>
                </div>
                <div class="movie-info">
                  <span>HD</span>
                </div>
                <div class="movie-info">
                  <span>16+</span>
                </div>
              </div>
              <div class="item-content-description top-down delay-4">
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quas,
                possimus eius. Deserunt non odit, cum vero reprehenderit
                laudantium odio vitae autem quam, incidunt molestias ratione
                mollitia accusantium, facere ab suscipit.
              </div>
              <div class="item-action top-down delay-6">
                <router-link to="/productsdetail" class="btn btn-hover">
                  <i class="bx bxs-right-arrow"></i>
                  <span>Watch now</span>
                </router-link>
              </div>
            </div>
          </div>
        </div>
        <!-- END SLIDE ITEM -->
        <!-- SLIDE ITEM -->
        <!-- <div class="hero-slide-item">
          <img src="../assets/images/supergirl-banner.jpg" alt="" />
          <div class="overlay"></div>
          <div class="hero-slide-item-content">
            <div class="item-content-wraper">
              <div class="item-content-title top-down">
                Supergirl
              </div>
              <div class="movie-infos top-down delay-2">
                <div class="movie-info">
                  <i class="bx bxs-star"></i>
                  <span>9.5</span>
                </div>
                <div class="movie-info">
                  <i class="bx bxs-time"></i>
                  <span>120 mins</span>
                </div>
                <div class="movie-info">
                  <span>HD</span>
                </div>
                <div class="movie-info">
                  <span>16+</span>
                </div>
              </div>
              <div class="item-content-description top-down delay-4">
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quas,
                possimus eius. Deserunt non odit, cum vero reprehenderit
                laudantium odio vitae autem quam, incidunt molestias ratione
                mollitia accusantium, facere ab suscipit.
              </div>
              <div class="item-action top-down delay-6">
                <a href="#" class="btn btn-hover">
                  <i class="bx bxs-right-arrow"></i>
                  <span>Watch now</span>
                </a>
              </div>
            </div>
          </div>
        </div> -->
        <!-- END SLIDE ITEM -->
        <!-- SLIDE ITEM -->
        <!-- <div class="hero-slide-item">
          <img src="../assets/images/wanda-banner.jpg" alt="" />
          <div class="overlay"></div>
          <div class="hero-slide-item-content">
            <div class="item-content-wraper">
              <div class="item-content-title top-down">
                Wanda Vision
              </div>
              <div class="movie-infos top-down delay-2">
                <div class="movie-info">
                  <i class="bx bxs-star"></i>
                  <span>9.5</span>
                </div>
                <div class="movie-info">
                  <i class="bx bxs-time"></i>
                  <span>120 mins</span>
                </div>
                <div class="movie-info">
                  <span>HD</span>
                </div>
                <div class="movie-info">
                  <span>16+</span>
                </div>
              </div>
              <div class="item-content-description top-down delay-4">
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quas,
                possimus eius. Deserunt non odit, cum vero reprehenderit
                laudantium odio vitae autem quam, incidunt molestias ratione
                mollitia accusantium, facere ab suscipit.
              </div>
              <div class="item-action top-down delay-6">
                <a href="#" class="btn btn-hover">
                  <i class="bx bxs-right-arrow"></i>
                  <span>Watch now</span>
                </a>
              </div>
            </div>
          </div>
        </div> -->
        <!-- END SLIDE ITEM -->
      </div>
    </div>
    <!-- END HERO SLIDE -->
  </div>
  <!-- END HERO SECTION -->
  <!-- END HTML CONTENT -->
</template>

<script>
export default {};
</script>

<style scoped>
/* BUTTON CSS */
a.btn:hover,
a.logo:hover,
a.movie-item:hover {
  color: unset;
}
.btn {
  color: #ffffff;
  padding: 0.25rem 1.5rem;
  text-transform: uppercase;
  font-size: 1.25rem;
  font-weight: 700;
  display: inline-flex;
  position: relative;
  align-items: center;
  font-family: "Cairo", sans-serif;
  text-decoration: none;
}

.btn-hover::before {
  z-index: 1;
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 50%;
  height: 100%;
  background-color: #c0392b;
  transition: 0.3s ease-in-out;
}

.btn-hover:hover::before {
  width: 100%;
}

.btn i,
.btn span {
  z-index: 1;
}

.btn i {
  margin-right: 1rem;
}

/* END BUTTON CSS */
.hero-slide-item {
  padding-top: 40%;
  position: relative;
  overflow: hidden;
}

.hero-slide-item img {
  width: 100%;
  position: absolute;
  top: 0;
  left: 0;
}

.hero-slide-item-content {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  color: #ffffff;
  display: flex;
  font-family: "Cairo", sans-serif;
}

.item-content-wraper {
  padding-left: 5rem;
  width: 40%;
  background-color: rgba(0, 0, 0, 0.8);
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: relative;
}

.item-content-wraper::before {
  content: "";
  position: absolute;
  top: 0;
  left: 100%;
  width: 50%;
  height: 100%;
  background: linear-gradient(to right, rgba(0, 0, 0, 0.8), rgba(0, 0, 0, 0));
}

.item-content-title {
  font-size: 5rem;
  line-height: 5rem;
  font-weight: 900;
  font-family: "Cairo", sans-serif;
}

.movie-infos {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  margin-top: 30px;
}

.movie-info {
  display: flex;
  align-items: center;
  font-size: 0.8rem;
  font-weight: 600;
}

.movie-info span {
  margin-left: 5px;
}

.movie-info i {
  color: #ffffff;
}

.movie-info ~ .movie-info {
  margin-left: 5px;
}

.item-content-description {
  font-size: 1.15rem;
  margin-top: 30px;
  line-height: 1.5rem;
}

.item-action {
  margin-top: 30px;
}

/* RESPONSIVE */

@media only screen and (max-width: 1280px) {
  html {
    font-size: 14px;
  }

  .hero-slide-item {
    padding-top: 50%;
  }

  .item-content-wraper {
    width: 50%;
  }
}

@media only screen and (max-width: 850px) {
  html {
    font-size: 12px;
  }

  .container {
    padding: 0 15px;
  }

  .hero-slide-item {
    height: max-content;
    padding-top: unset;
  }

  .item-content-wraper {
    width: 100%;
    height: 100%;
    padding-bottom: 20px;
    padding-right: 5rem;
    background-color: rgba(0, 0, 0, 0.2);
  }

  .hero-slide-item-content {
    position: relative;
  }

  .hero-slide-item img {
    height: 100%;
  }

  .item-content-title {
    font-size: 3rem;
  }
}
</style>
