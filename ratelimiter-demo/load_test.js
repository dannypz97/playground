import { check } from "k6";
import http from "k6/http";

export const options = {
    scenarios: {
      scenario: {
        executor: "ramping-arrival-rate",
        startRate: 15,
        timeUnit: "1s",
        preAllocatedVUs: 5,
        maxVUs: 20,
        exec: "scenario",
        env: {},
        stages: [
            {
                target: 15,
                duration: "2m"
            },
        ]
      }
    }
};

const hostUrl = __ENV.HOST_URL || "http://localhost:8080";


export function scenario() {
    const response = http.get(hostUrl);

    check(response, {
        'token available': (r) => {
            return r.body.trim() === "true";
        }
    });
}