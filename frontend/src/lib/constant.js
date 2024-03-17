const API_URL = import.meta.env.VITE_API_URL ? import.meta.env.VITE_API_URL + '/api' : "/api"

export const conf = {
  fields: `${API_URL}/fields`,
  alldata : `${API_URL}/allcitizen`,
  citizen : `${API_URL}/citizen`,
}