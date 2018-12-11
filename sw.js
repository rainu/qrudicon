importScripts('/qrudicon/_nuxt/workbox.4c4f5ca6.js')



workbox.precaching.precacheAndRoute([
  {
    "url": "/qrudicon/_nuxt/app.0f8a1c8797b823aa979a.js",
    "revision": "c5897fa6abc27c6ea214d62ec8a2bed2"
  },
  {
    "url": "/qrudicon/_nuxt/layouts/default.f77873a261de3827fae5.js",
    "revision": "fb863d2afa8b03e736d9c337ca014cc9"
  },
  {
    "url": "/qrudicon/_nuxt/manifest.cab4f5c6cbc444c5b901.js",
    "revision": "ccd021246989dcd350db9eff16e0e812"
  },
  {
    "url": "/qrudicon/_nuxt/pages/index.123f1ce448a2e96b139a.js",
    "revision": "1cc35cbc2e505c3cf3fa5d61077219bb"
  },
  {
    "url": "/qrudicon/_nuxt/vendor.7de63f0960414e680473.js",
    "revision": "598cc7b718ffe01c932769de53d22afe"
  }
], {
  "cacheId": "qrudicon",
  "directoryIndex": "/",
  "cleanUrls": false
})



workbox.clientsClaim()
workbox.skipWaiting()


workbox.routing.registerRoute(new RegExp('/qrudicon/_nuxt/.*'), workbox.strategies.cacheFirst({}), 'GET')

workbox.routing.registerRoute(new RegExp('/qrudicon/.*'), workbox.strategies.networkFirst({}), 'GET')





