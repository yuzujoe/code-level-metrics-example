import http from 'k6/http';
import {check, sleep} from 'k6';

export const options = {
    vus: 50,
    duration: '30m',
}

export default function () {
    const res = http.get(`http://localhost:8000/example`);
    check(res, {'success rate': (r) => r.status === 200});
    sleep(0.3)
}
