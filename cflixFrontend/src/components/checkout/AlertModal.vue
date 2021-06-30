<template>
  <div class="wrap-modal">
    <div class="modal">
      <div class="modal-content">Do you want to delete this product?</div>
      <div class="modal-group-action">
        <button class="btn-yes" @click="confirmYes">Confirm</button>
        <button class="btn-no" @click="confirmNo">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "CompModal",
  methods: {
    confirmYes() {
      this.$emit("confirm", true);
    },
    confirmNo() {
      this.$emit("confirm", false);
    },
  },
};
</script>

<style>
.wrap-modal {
  z-index: 2;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.3);
}
.modal {
  width: 500px;
  height: 200px;
  background-color: white;
  border-radius: 10px;
  padding: 0 20px;
}
.btn-yes {
  height: 50px;
  font-size: 20px;
  border-radius: 10px;
}
.btn-no {
  background-color: rgb(41, 41, 41);
  border-top: 1px solid rgb(41, 41, 41);
  border-bottom: 1px solid rgb(41, 41, 41);
  height: 50px;
  font-size: 20px;
  border-radius: 10px;
}
.modal-content {
  padding: 20px 0;
  font-size: 25px;
}
.modal-group-action {
  width: 100%;
  display: flex;
  justify-content: space-between;
  padding-top: 3rem;
  /* margin-top: 5rem;
  margin-bottom: 20px; */
}
.modal-group-action button {
  width: 40%;
}
</style>
