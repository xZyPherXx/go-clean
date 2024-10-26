// console.log('Index!')
// if ('serviceWorker' in navigator) {

//     navigator.serviceWorker.register("/service-worker.js").then(reg => {

//         console.log("Service Worker Registered!", reg);
//     }).catch(err => {

//         console.log("Service Worker Failed!", err);
//     })

// }

async function fetchProducts() {

    // const cache = await caches.open('pwa-cache');
    // if (navigator.onLine) {

        const res = await fetch('localhost:8000/api/products');
        // const res = await fetch('https://pwa.zyph6rx.site:8443/api/products');
        const products = await res.json();
        // cache.put('/api/products', new Response(JSON.stringify(products)));
        return products;
    // }

    // const cachedResponse = await cache.match('/api/products');
    // if (cachedResponse) return await cachedResponse.json();
    // throw new Error("No cached");
}

function createCard(product) {

    const categories = product.description.split(", ").map(part => part.trim()).join(", ");
    const card = document.createElement('div');
    card.className = 'product-card';
    card.innerHTML = `
        <img src="${product.image_url}" alt="${product.name}" />
        <h2>${product.name}</h2>
        <p>${categories}</p>
        <p>Price: $${product.price.toFixed(2)}</p>`;

    return card;
}

function displayCard(products) {

    const productContainer = document.getElementById('productContainer');
    productContainer.innerHTML = '';
    products.forEach(product => {

        const card = createCard(product);
        productContainer.appendChild(card);
    });

}

function filterCard(products) {

    const selected = document.getElementById('categorySelect').value;
    const filtered = selected === 'All'

        ? products
        : products.filter(product => {

            const category = product.description.split(", ").map(part => part.trim());
            return category.includes(selected);
        });

    displayCard(filtered);
}

fetchProducts().then(products => {

    displayCard(products);
    document.getElementById('categorySelect').addEventListener('change', () => {

        filterCard(products);
    });

}).catch(err => {

    console.error(err);
});