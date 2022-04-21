<Router routes="{routes}" />

<style>
:global(html) {
  /* this will apply to <html doc> */
  --ion-font-family: "Helvetica Neue", "Roboto", sans-serif;
}
</style>

<script>
import { Router } from "@roxi/routify";

export let routes;

// import { FirebaseApp } from "sveltefire";
import firebase from "firebase/app";
import "firebase/firestore";
import "firebase/auth";

// Notice regarding imports below
// https://stackoverflow.com/questions/63835251/firebase-modules-cause-registercomponent-error-in-svelte
import "firebase/performance";
import "firebase/analytics";

let firebaseConfig = {
  apiKey: "AIzaSyAOul3jp5hfGOfRWwqxjrvU4i965DE11Qk",
  authDomain: "ionicsvelte.firebaseapp.com",
  databaseURL: "https://ionicsvelte.firebaseio.com",
  projectId: "ionicsvelte",
  storageBucket: "ionicsvelte.appspot.com",
  messagingSenderId: "512207280782",
  appId: "1:512207280782:web:aa2cc110c3ac12e5ef07df",
  measurementId: "G-NNGFLY1HYH",
};

firebase.initializeApp(firebaseConfig);
const defaultAnalytics = firebase.analytics();
defaultAnalytics.logEvent("page_view", { page_title: "app_start" });
</script>
