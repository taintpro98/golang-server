import http from 'k6/http';

export let options = {
  duration: '20s',
  vus: 1,
};

export function setup() {
  const tokens = [
    'eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNDJjOTU2ZjYtNjI1ZS00YjAyLWJkNjktMGQ1NGYwNTg3ZTE2In0.OMdXGRNhHeIYD90MaC4dfdbHNXvUP47ZuuxuvecOZ3kma_LyYPCwzJckNbArCT14zk3BjGyVrObwuUR83FnBA7KbTL2fZPT4t0KRcMC0tk19Pen3B3eL2bKwZiEZJGBvNO2AaireSQTh02TVALc83KLvh7I1Zi9jhIjLkaLuHF8QQi-kBSw-LNBKpkhOQcKtRA_v2LN-Nz3g_97eshcRCjCLJFBOV4ngWaBjSqoxm-lqQowqsmZBaV8S7A3e20jBn_JBjzHTo8qMvF-UMRq-_XLyWv3d7b4WklpP6V2XdUJOc3WLpHfYUrM7S2L35wMpxidvcXAgw5BrgsyTV9k9PNLAi3iKt17vByNVnwwo8OU-uXbqNQ3_PKVEXD_6iSr4dVG3MDYi3jQIGeQSOJWODXBEEvYL2-CU2bp3Zzd94ghUxb9UE2zGKu1Xe6wrt-kmXn9sP3r0ACrjnuzHCpPcKKQXFhwFuKp78bNgw3yKIQDRyCIMZfnH0FLoEHE4BcQVwu8_0NGMMwq5hb2Wmd1AqZkP6X2evc5Pf0DZnFidyEhyBu1p6pwIQEthBv6ivzPkxMHrN20ABey3Zm_lt7XfCcHvzIHCcMRiNEKbe6bcntrakZEQJ--majfLLPlUTc9mxWWx7-JmptGf6r88mid9uANS4rC9Mv8bqHrIz3duqAk'
  ]
  return { tokens };
}

export default function (data) {
  const slotID = '4bdacdba-16b5-4160-bec2-4df443611f13';
  const randomIndex = Math.floor(Math.random() * data.tokens.length);
  const token = data.tokens[randomIndex];

  const headers = {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${token}`,
  };

  const getResponse = http.get(`http://localhost:5000/v1/public/slots/${slotID}`, { headers });
  if (getResponse.status === 200) {
    const getResponseData = JSON.parse(getResponse.body).data
    const randomSeat = Math.floor(Math.random() * getResponseData.seats.length);
    const postBody = {
      seat_id: randomSeat.seat_id
    }
    const postResponse = http.post(`http://localhost:5000/v1/public/slots/${slotID}`, JSON.stringify(postBody), { headers });
    if (postResponse.status === 200) {
      console.log(`POST response status code: ${postResponse.body}`);
    }
  }
}

// result: running (20.0s), 0/1 VUs, 3435 complete and 0 interrupted iterations