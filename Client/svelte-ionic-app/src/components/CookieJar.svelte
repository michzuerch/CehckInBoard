{#if showCookieJar}
  <div
    in:fly="{{ y: 200, delay: 7000, duration: 1500, easing: elasticInOut }}"
    out:fly="{{ y: 200, delay: 10, duration: 1000, easing: elasticInOut }}"
    class="card"
  >
    <div class="card-inset">
      <ion-card>
        <ion-card-header>
          <ion-card-header>Cookie monster alert!</ion-card-header>
        </ion-card-header>
        <ion-card-content>
          This app/website is using cookies to track usage. Please click the
          button to continue and agree. <br />
          <ion-button on:click="{userAgrees}" expand="full" color="primary" on>
            I Agree
          </ion-button>
        </ion-card-content>
      </ion-card>
    </div>
  </div>
{/if}

<style>
.card {
  z-index: 1000;
  position: absolute;
  width: 100%;
  bottom: 10px;
  background: transparent;
}

.card-inset {
  margin: auto;
  width: 90%;
}

@media only screen and (min-width: 500px) {
  .card-inset {
    width: 50%;
  }
}
</style>

<script>
import { fly } from "svelte/transition";
import { elasticInOut } from "svelte/easing";

import localforage from "localforage";

// we don't want to show, unless we need to
let showCookieJar = false;
localforage.getItem("cookie-agree").then((value) => {
  console.log("Cookie value", value);
  if (!value) {
    showCookieJar = true;
  }
});

const userAgrees = () => {
  showCookieJar = false;
  console.log("User agreed to Cookie!");
  localforage.setItem("cookie-agree", true);
};
</script>
