import { conf } from "./constant"

export const fetchit = async (type) => {

  if(!conf[type]) return {"error": "no data"};

  const res = await fetch(conf[type]);
  return await res.json();
}

export async function fetchme(url, method = "GET", body = null) {
  const res = await fetch(url, { method, body }, );
  if (res.status !== 200) {
    return {"error": "no data"}
  }
  return await res.json();
}