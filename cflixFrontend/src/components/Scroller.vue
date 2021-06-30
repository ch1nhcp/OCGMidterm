<template>
  <div class="section-header">{{ category }}</div>
  <carousel :items-to-show="1" class="film-carousel-wrapper">
    <slide v-for="slide in 5" :key="slide">
      <div class="section">
        <div class="hero-slide-item">
          <img src="../assets/images/transformer-banner.jpg" alt="" />
          <div class="overlay"></div>
          <div class="hero-slide-item-content">
            <div class="item-content-wraper">
              <div class="item-content-title">
                Transformer
              </div>
              <div class="movie-infos">
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
              <div class="item-content-description">
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quas,
                possimus eius. Deserunt non odit, cum vero reprehenderit
                laudantium odio vitae autem quam, incidunt molestias ratione
                mollitia accusantium, facere ab suscipit.
              </div>
              <div class="item-action">
                <a href="#" class="btn btn-hover">
                  <i class="bx bxs-right-arrow"></i>
                  <span>Watch now</span>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </slide>
    <template #addons>
      <navigation />
    </template>
  </carousel>
</template>

<script>
import "vue3-carousel/dist/carousel.css";
import { Carousel, Slide, Navigation } from "vue3-carousel";

export default {
  name: "ProductsCarousel",
  components: {
    Carousel,
    Slide,
    Navigation
  },
  props: {
    category: String,
  },
};
</script>

<style>
/* :root {
    --main-color: #c0392b;
    --body-bg: #181616;
    --box-bg: #221f1f;
    --text-color: #ffffff;

    --nav-height: 60px;
    --space-top: 30px;
}
 */



a {
    text-decoration: none;
    color: unset;
}

img {
    max-width: 100%;
}

.main-color {
    color: var(--main-color);
}



.overlay {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    background-color: rgba(0, 0, 0, 0.5);
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
}

.btn-hover::before {
    z-index: 1;
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 50%;
    height: 100%;
    background-color: var(--main-color);
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
    color: var(--text-color);
    display: flex;
}

.item-content-wrapper {
    padding-left: 5rem;
    width: 40%;
    min-width: 30%;
    background-color: rgba(0, 0, 0, 0.8);
    display: flex;
    flex-direction: column;
    justify-content: center;
    position: relative;
}

.item-content-wrapper::before {
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
}

.movie-infos {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    margin-top: calc(var(--space-top) / 2);
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
    color: var(--main-color);
}

.movie-info ~ .movie-info {
    margin-left: 5px;
}

.item-content-description {
    font-size: 1.15rem;
    margin-top: var(--space-top);
    line-height: 1.5rem;
}

.item-action {
    margin-top: var(--space-top);
}

/* .section {
    padding-top: 80px;
} */

.section-header {
    margin-bottom: 30px;
    padding-left: 20px;
    text-transform: uppercase;
    font-size: 1.5rem;
    font-weight: 700;
    border-left: 4px solid var(--main-color);
    display: flex;
    align-items: center;
}


/* RESPONSIVE */

@media only screen and (max-width: 1280px) {
    html {
        font-size: 14px;
    }

    .hero-slide-item {
        padding-top: 50%;
    }

    .item-content-wrapper {
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

    .item-content-wrapper {
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

    .hamburger-menu {
        display: grid;
    }

    .nav-menu {
        /* display: none; */
        position: absolute;
        top: 100%;
        left: -100%;
        background-color: #000000;
        flex-direction: column;
        width: 80%;
        height: 100vh;
        padding: 20px;
        transition: 0.3s ease-in-out;
    }

    .nav-menu li {
        margin: 10px 30px;
    }

    .nav-menu.active {
        left: 0;
    }
}

</style>
