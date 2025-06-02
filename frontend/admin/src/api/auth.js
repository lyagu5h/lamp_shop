import { authApi } from "./axios";

export async function login(username, password) {
  const response = await authApi.post("/login", { username, password });
  return response.data;
}

export async function register(username, password, role) {
  return authApi.post("/register", { username, password, role });
}