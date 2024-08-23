import http from 'k6/http';
import { check, sleep } from 'k6';
import { SharedArray } from 'k6/data';

const BASE_URL = 'http://localhost:8083/get/';

const NUMBERS = new SharedArray('random numbers', function () {
  return Array.from({ length: 1000 }, () => Math.floor(Math.random() * 100) + 1);
});

export const options = {
  stages: [
    { duration: '5s', target: 100 },
    { duration: '2s', target: 30000 },
    { duration: '30s', target: 30000 },
    { duration: '10s', target: 0 }
  ],
};

export default function () {
  const randomNumber = NUMBERS[Math.floor(Math.random() * NUMBERS.length)];

  const url = `${BASE_URL}${randomNumber}`;

  const response = http.get(url);

  check(response, {
    'is status 200': (r) => r.status === 200,
  });

  sleep(1);
}
