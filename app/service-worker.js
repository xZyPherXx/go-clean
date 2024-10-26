const CACHE_OBJ = [

    "./",
    "./manifest.json",
    "./src/style.css",
    "./src/index.js",
]

self.addEventListener('install', event => {

    console.log("Install!");
    console.log(event);

    event.waitUntil(

        caches.open('pwa-cache').then(cache => {

            console.log('Cahce Open!')
            return cache.addAll(CACHE_OBJ)
        })

    );

});

self.addEventListener('fetch', event => {

    console.log(`Fetching : ${event.request.url}`);
    event.respondWith(

        fetch(event.request).then((response) => {

            const responseClone = response.clone();
            caches.open('pwa-cache').then((cache) => {

                cache.put(event.request, responseClone);
            });

            return response;
        }).catch(() => {

            return caches.match(event.request);
        })

    );

});
