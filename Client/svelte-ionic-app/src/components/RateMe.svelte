{#if showRateMe}
  <div
    in:fly="{{ y: 200, duration: 1500, easing: elasticInOut }}"
    out:fly="{{ y: 200, delay: 10, duration: 1000, easing: elasticInOut }}"
    class="card"
  >
    <div class="card-inset">
      <ion-card>
        <ion-card-content>
          <h2>Rate this app!</h2>
          <ion-grid>
            <ion-row>
              {#each stars as star, id}
                <ion-col>
                  <ion-icon
                    name="{star}"
                    id="{id + 1}"
                    on:click="{starClick}"
                  ></ion-icon>
                </ion-col>
              {/each}
            </ion-row>
          </ion-grid>
          <br />
          <h2>{rateTotal} others said...</h2>

          <ion-grid>
            {#each rateLabels as label, id}
              <ion-row>
                <ion-col>{label} ({rateCount[5 - id]})</ion-col>
                <ion-col>
                  <ion-progress-bar
                    value="{rateCount[5 - id] / rateTotal}"
                  ></ion-progress-bar>
                </ion-col>
              </ion-row>
            {/each}
          </ion-grid>

          {#if rate > 0 && rate < 5}
            Help us improve and give feedback. <br />
            {#each checkBoxes as { value, label }}
              <ion-item>
                <ion-checkbox
                  value="{value}"
                  on:ionChange="{checkBoxChange}"
                  color="dark"
                  slot="start"
                ></ion-checkbox>
                <ion-label>{label}</ion-label>
              </ion-item>
            {/each}
            <ion-item>
              <ion-label position="stacked">Tell us more...</ion-label>
              <ion-textarea on:ionChange="{changeValue}"></ion-textarea>
            </ion-item>
          {/if}

          {#if showDone}
            <br />
            <ion-button on:click="{userDone}" expand="full" color="primary" on>
              Close
            </ion-button>
          {:else}
            <ion-button
              on:click="{userCancel}"
              expand="full"
              color="primary"
              on
            >
              No thx - maybe next time
            </ion-button>
          {/if}
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
  bottom: 10%;
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

<script lang="ts">
import { fly } from "svelte/transition";
import { elasticInOut } from "svelte/easing";
import localforage from "localforage";

import { path } from "../services/routestore";

import firebase from "firebase/app";
import "firebase/firestore";

let feedback = {};
let showDone = false;
let rate = 0;
const rateCount = { 1: 0, 2: 0, 3: 0, 4: 0, 5: 0 };
const rateLabels = ["5 stars", "4 stars", "3 stars", "2 stars", "1 star"];
let rateTotal = 0;
let showRateMe = false;

let stars = [
  "star-outline",
  "star-outline",
  "star-outline",
  "star-outline",
  "star-outline",
];

let checkBoxes = [
  { value: "begging", label: "To often asking for rating" },
  { value: "helpfull", label: "Very helpful app" },
  { value: "buggy", label: "A bit buggy" },
  { value: "newfeature", label: "New feature:" },
];

// we don't want to show, unless we need to
localforage.getItem("rate-me-1").then((value) => {
  console.log("RateMe value", value);
  // showRateMe = true; // testing purposes should be on then
  if (!value) {
    setTimeout(() => {
      //   showRateMe = true;  // got complaints from user about rating-begging. Maybe asking too many times and localstorage does not work?
    }, 1000 * 60 * 1);
  }
});

// allow to show RateMe when menu is selected - unfortunately we cannot grap from routify?
path.subscribe((p) => {
  if (p === "/RateMe") {
    showRateMe = true;
    path.set("/"); // reset in case user wants to rate again without having to navigate away
  }
});

// lets try to load all the earlier ratings - if we didn't get them earlier
if (rateTotal === 0) {
  let db = firebase.firestore();
  db.collection("feedback")
    .get()
    .then((querySnapshot) => {
      querySnapshot.forEach((doc) => {
        const docData = doc.data();

        const loadedRate = docData.feedbackText.rate;
        if (loadedRate > 0 && loadedRate < 6) {
          rateCount[loadedRate] += 1;
          rateTotal += 1;

          if (docData.feedbackText.notes) {
            console.log("Notes received", docData.feedbackText.notes);
          }
        }
      });
    });
}

const starClick = (event) => {
  console.log("Starclick", event.srcElement.id, rate);
  if (rate === 0) {
    const id = event.srcElement.id;
    rate = event.srcElement.id;
    stars = stars.map((star, i) => {
      if (i < rate) {
        return "star-sharp";
      } else {
        return "star-outline";
      }
    });

    feedback["rate"] = rate;

    showDone = true;

    if (rate == 5) {
      setTimeout(() => {
        userDone();
      }, 2000);
    }
  }
};

const checkBoxChange = (event) => {
  feedback[event.detail.value] = event.detail.checked;
};

const changeValue = (event) => {
  feedback["notes"] = event.detail.value;
};

const userCancel = () => {
  showRateMe = false;
};

const userDone = () => {
  showRateMe = false;

  // let's track some stuff
  feedback["time"] = Date.now();
  const d = new Date();
  feedback["utc"] = d.toUTCString();

  // we want lots of info! - and some info is too much, so we skip
  [
    // "product",
    "platform",
    //  "userAgent",
    //  "appName",
    "appVersion",
    //  "appCodeName",
  ].forEach((key) => {
    feedback[key] = navigator[key];
    if (typeof feedback[key] === "undefined") {
      feedback[key] = "Unknown";
    }
  });

  // stringify and firebase can give runtime errors
  let feedbackText = "";
  try {
    // we don't want to post to production database while in development
    if (window.location.hostname !== "localhost") {
      let db = firebase.firestore();
      feedbackText = JSON.parse(JSON.stringify(feedback, null, 2)); // convert into a saveable object
      db.collection("feedback").add({ feedbackText });
    }
  } catch (e) {
    console.log("Errors", e);
  }
  console.log("User rateme", feedbackText);
  localforage.setItem("rate-me-1", true);

  // reset for another review - after animation end
  setTimeout(() => {
    rate = 0;
    stars = [
      "star-outline",
      "star-outline",
      "star-outline",
      "star-outline",
      "star-outline",
    ];
  }, 1000);
};
</script>
