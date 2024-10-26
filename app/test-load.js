// k6 script for testing performance
import http from 'k6/http';

export let options = {
    vus: 100,          // concurrent users
    duration: '30s',   // duration of test
};

export default function () {
    http.post('http://localhost:8000/api/products', { name: "Test User",  description : "Tactical, FPS",
        price: 14.00, image_url: "rainbow-six-siege.png"
    });
}
