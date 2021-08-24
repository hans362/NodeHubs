<template>
  <footer class="footer px-4">
    <div class="row align-items-center justify-content-lg-between">
      <div class="col-lg-6">
        <div class="copyright text-center text-lg-left text-muted">
          © {{ year }}
          <a
            href="https://github.com/hans362/NodeHubs"
            class="font-weight-bold ml-1"
            target="_blank"
            >NodeHubs</a
          >
          Made By
          <a
            href="https://hans362.cn"
            class="font-weight-bold ml-1"
            target="_blank"
            >Hans362</a
          >
        </div>
      </div>
    </div>
  </footer>
</template>
<script>
export default {
  data() {
    return {
      year: new Date().getFullYear(),
    };
  },
};
</script>
<style></style>
