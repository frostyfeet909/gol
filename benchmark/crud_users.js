import http from "k6/http";
import { check, sleep } from "k6";
import { randomString } from "https://jslib.k6.io/k6-utils/1.4.0/index.js";

export const options = {
  vus: 5,              // virtual users
  duration: "30s",     // run time
  thresholds: {
    http_req_failed: ["rate<0.01"],      // <1% errors
    http_req_duration: ["p(95)<300"],    // p95 under 300ms (tweak)
  },
};

__ENV.BASE_URL =  "https://gol-app.fly.dev/"; // uncomment to set default base URL

const BASE_URL = __ENV.BASE_URL || "http://127.0.0.1:8080";
const AUTH = __ENV.AUTH || ""; // e.g. "Bearer <token>"

function headers() {
  const h = { "Content-Type": "application/json", "Accept": "application/json" };
  if (AUTH) h["Authorization"] = AUTH;
  return h;
}

export default function () {
  // --- create 10 users
  const ids = [];
  for (let i = 0; i < 10; i++) {
    const body = JSON.stringify({
      email: `algie_${__VU}_${__ITER}_${i}_${randomString(6)}@example.com`,
      name: `User ${__VU}-${__ITER}-${i}`,
    });

    const res = http.post(`${BASE_URL}/api/v1/users`, body, { headers: headers() });
    check(res, { "create: 201/200": (r) => r.status === 201 || r.status === 200 });

    if (res.status < 200 || res.status >= 300) {
        console.error(`CREATE payload=${body}`);
        console.error(`CREATE status=${res.status}`);
        console.error(`CREATE resp_headers=${JSON.stringify(res.headers)}`);
        console.error(`CREATE body=${res.body}`);
        return;
    }

    // adjust depending on your API response shape:
    // e.g. { id: "..." } or { user: { id: "..." } }
    const json = res.json();
    const id = json.id || (json.user && json.user.id);
    if (id) ids.push(id);
  }

  // --- list users
  {
    const res = http.get(`${BASE_URL}/api/v1/users`, { headers: headers() });
    check(res, { "list: 200": (r) => r.status === 200 });
  }

  // --- get each user individually
  for (const id of ids) {
    const res = http.get(`${BASE_URL}/api/v1/users/${id}`, { headers: headers() });
    check(res, { "get: 200": (r) => r.status === 200 });
  }

  // --- delete each user
  for (const id of ids) {
    const res = http.del(`${BASE_URL}/api/v1/users/${id}`, null, { headers: headers() });
    check(res, { "delete: 200/204": (r) => r.status === 200 || r.status === 204 });
  }

  sleep(0.2);
}
