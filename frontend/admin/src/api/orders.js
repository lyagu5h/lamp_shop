import { ordersApi } from "./axios";

export async function fetchOrders() {
  const res = await ordersApi.get("/");
  return res.data;
}

export async function fetchOrderById(id) {
  const res = await ordersApi.get(`/${id}`);
  return res.data;
}

export async function createOrder(data) {
  const res = await ordersApi.post("/", data);
  return res.data;
}

export async function updateOrderStatus(id, statusData) {
  const res = await ordersApi.patch(`/${id}/status`, statusData);
  return res.data;
}
