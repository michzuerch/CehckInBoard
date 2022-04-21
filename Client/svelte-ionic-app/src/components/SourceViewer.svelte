<svelte:head>
  <title>Sourceviewer {name}</title>
</svelte:head>
<ion-header translucent="true">
  <ion-toolbar>
    <ion-buttons slot="end">
      <ion-button
        on:click="{() => {
          if (REPLlink.length > 1) {
            REPLclick();
            window.open(REPLlink, '_blank');
          } else {
            showNoREPLToast();
          }
        }}"
      >
        REPL
      </ion-button>

      <ion-button
        on:click="{() => {
          APIclick();
          window.open(APIlink, '_blank');
        }}"
      >
        API
      </ion-button>

      {#if sourceCode.length > 11}
        <ion-button on:click="{copySource}">COPY</ion-button>
      {/if}
      <ion-button on:click="{closeOverlay}">
        <ion-icon name="close"></ion-icon>
      </ion-button>
    </ion-buttons>
  </ion-toolbar>

  <ion-segment value="{codeLanguage}" on:ionChange="{segmentChange}" scrollable>
    <ion-segment-button value="svelte">
      <ion-label>Svelte</ion-label>
    </ion-segment-button>
    <ion-segment-button value="js">
      <ion-label>Javascript</ion-label>
    </ion-segment-button>
  </ion-segment>
</ion-header>

<ion-content scroll-x="true" style="--padding-start: 15px;--padding-end: 15px;">
  <pre
    style="-webkit-user-select: text; /* Chrome 49+ */
  -moz-user-select: text; /* Firefox 43+ */
  -ms-user-select: text; /* No support yet */
  user-select: text; /* Likely future */"
  >{sourceCode}</pre>
</ion-content>

<style>
pre {
  -webkit-user-select: all; /* Chrome 49+ */
  -moz-user-select: all; /* Firefox 43+ */
  -ms-user-select: all; /* No support yet */
  user-select: all; /* Likely future */
}
</style>

<script lang="ts">
import { Plugins } from "@capacitor/core";
import { fromFetch } from "rxjs/fetch";
import { IonicShowToast } from "../services/IonicControllers";

import localforage from "localforage";

import firebase from "firebase/app";

const { Clipboard } = Plugins;
const { componentProps } = document.querySelector("ion-modal");

const defaultAnalytics = firebase.analytics();
let REPLlink;
let APIlink;
let codeLanguage = "svelte";
let sourceCode = "Loading....";
let sources = {};
let languages = ["svelte", "js"];
languages.forEach((lang) => {
  sources[lang] = "Loading " + lang + "....";
});

localforage.getItem("source-language").then((value: string) => {
  if (value) {
    codeLanguage = value;
    sourceCode = sources[codeLanguage];
  }
});

// we want to measure some stuff
defaultAnalytics.logEvent("page_view", { page_title: "sourceviewer" });

// probably can be done in easier way, but I am lazy
let name = "/";
if (componentProps.name) {
  name = componentProps.name;
}

// generate the name to source file
name = name.replace("ionic/", "");
name = name.substr(1);
name = name.charAt(0).toUpperCase() + name.slice(1);
if (name.length == 0) {
  name = "Splash";
}
if (name.toLowerCase().includes("tabs")) {
  name = "tabs";
}

console.log("Svelte name to load", name);

// cannot use function as Safari request window.open to be called in on:click
// therefor, preload the links
fromFetch("/assets/json/repls.json").subscribe((response) => {
  response.json().then((json) => {
    const url = json[name.toLowerCase()];
    if (url) {
      REPLlink = url;
    } else {
      REPLlink = "";
    }
  });
});

// try to generate the url to the api docs
// if (apiName.charAt(apiName.length - 1) == "s") {
//  apiName = apiName.slice(0, -1);
// }

let apiName = name.toLowerCase();
console.log("Raw APINAME", apiName);

// do a mapping for some exceptions
apiName = apiName.replace("/", "");
fromFetch("/assets/json/api-mappings.json").subscribe(
  (response) => {
    response.json().then((json) => {
      let url = "https://ionicframework.com/docs/api/";
      if (json[apiName]) {
        url += json[apiName];
        // if we found an exception - we replace
        if (json[apiName].toLowerCase().substring(0, 4) == "http") {
          url = json[apiName];
        }
      } else {
        url += apiName;
      }
      APIlink = url;
    });
  },
  (error) => {
    console.error("Error HTTP", error);
  }
);

// and more automated ways to load the other languages
languages.forEach((lang) => {
  if (lang !== "svelte") {
    console.log("Loading code ", lang);
    sources[lang] = "No code available";
    fromFetch(
      "/assets/src/" + lang + "/" + name.toLowerCase() + "/index.html"
    ).subscribe(
      (response) => {
        response.text().then((txt) => {
          sources[lang] = txt;
        });
      },
      (err) => {
        sources[lang] = "No code available";
        console.log("Loaded code error", err);
      },
      () => {
        sourceCode = sources[codeLanguage];
        console.log("Loaded code", sources);
      }
    );
  }
});

// we use a separate way to load the svelte source
// hack the name for tabs
if (name == "tabs") {
  name = "tabs/[tab]";
}
fromFetch("/assets/src/svelte/" + name + ".svelte").subscribe((response) => {
  response.text().then((txt) => {
    if (txt.search("<!DOCTYPE html>") > -1) {
      sources[
        "svelte"
      ] = `No svelte file found for ${name}. Please check github repo.`;
    } else {
      sources["svelte"] = txt;
      sourceCode = sources[codeLanguage];
    }
  });
});

const segmentChange = (value) => {
  codeLanguage = value.detail.value;
  localforage.setItem("source-language", codeLanguage);
  sourceCode = sources[codeLanguage];
  defaultAnalytics.logEvent("page_view", {
    page_title: "source" + codeLanguage,
  });
};

const REPLclick = () => {
  defaultAnalytics.logEvent("page_view", { page_title: "REPL_view" });
};

const APIclick = () => {
  defaultAnalytics.logEvent("page_view", { page_title: "API_view" });
};

const closeOverlay = () => {
  const modal = document.querySelector("ion-modal");
  modal.dismiss({});
};

const showNoREPLToast = () => {
  IonicShowToast({
    color: "danger",
    duration: 2000,
    message: "No REPL link configured",
    showCloseButton: true,
  });
};

// somehow does not fly on iOS?
const copySource = () => {
  Clipboard.write({
    string: sourceCode,
  })
    .then(() => {
      IonicShowToast({
        color: "dark",
        duration: 2000,
        message: "Copied...",
        showCloseButton: true,
      });
    })
    .catch((message) => {
      IonicShowToast({
        color: "danger",
        duration: 2000,
        message,
        showCloseButton: true,
      });
    });

  setTimeout(closeOverlay, 1000);
};
</script>
