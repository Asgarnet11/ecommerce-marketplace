import axios from "axios";
import { useAuth } from "../stores/auth";

const api = axios.create({ baseURL: "http://localhost:8080" });

api.interceptors.request.use((cfg) => {
  const auth = useAuth.getState ? useAuth.getState() : useAuth(); // support HMR
  const token = auth?.token;
  if (token) cfg.headers.Authorization = `Bearer ${token}`;
  return cfg;
});

export default api;
