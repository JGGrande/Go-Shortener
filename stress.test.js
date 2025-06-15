import http from 'k6/http';
import { check, sleep } from 'k6';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";
import { Trend } from 'k6/metrics';

const config = {
  host: 'http://localhost:8080',
  language: 'golang'
}

export let options = {
  stages: [ 
    { duration: '30s', target: 100 },
    { duration: '30s', target: 200 },
    { duration: '30s', target: 300 },
    { duration: '30s', target: 400 },
    { duration: '1m', target: 0 },     
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95% das requisições abaixo de 500ms
    http_req_failed: ['rate<0.01'],   // Taxa de erro abaixo de 1%
  },
};


const BASE_URL = config.host;

const shortenerTrend = new Trend('create_shortener_duration');
const redirectTrend = new Trend('redirect_duration');

export default function () {
  const postRes = http.post(`${BASE_URL}/shortener`, JSON.stringify({
    url: 'https://example.com'
  }), {
    headers: { 'Content-Type': 'application/json' },
  });

  shortenerTrend.add(postRes.timings.duration);

  check(postRes, {
    'POST status 201': (r) => r.status === 201,
    'POST has identifier': (r) => r.json('slug') !== undefined,
  });

  const identifier = postRes.json('slug');

  const getRes = http.get(`${BASE_URL}/r/${identifier}`, {
    redirects: 0
  });

  redirectTrend.add(getRes.timings.duration);

  check(getRes, {
    'GET status is redirect': (r) => r.status === 302,
  });

  sleep(1);  
}

export function handleSummary(data) {
  return {
    [`summary-${config.language}.html`]: htmlReport(data),
  };
}

