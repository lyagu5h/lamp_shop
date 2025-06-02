import { Order, Product } from "./types";

const API = import.meta.env.VITE_API_URL ?? "http://localhost:8000";

export async function getProducts() {
  const r = await fetch(`${API}/products`);
  if (!r.ok) throw new Error("Не удалось получить товары");
  return (await r.json()) as Product[];
}

export async function getProduct(id: number) {
  const r = await fetch(`${API}/products/${id}`);
  if (!r.ok) throw new Error("Товар не найден");
  return (await r.json()) as Product;
}

// export interface OrderPayload extends Pick<Order, "customer_name" | "customer_email" | "phone" | "address" | "items" | "total_amount"> { [key: string]: never; }

export async function createOrder(data: Pick<Order, "customer_name" | "customer_email" | "customer_phone" | "address" | "items" | "total_amount">) {
  const r = await fetch(`${API}/orders`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (!r.ok) throw new Error("Ошибка при создании заказа");
  return (await r.json()) as Order;
}

export async function getOrder(id: number) {
  const r = await fetch(`${API}/orders/${id}`);
  if (!r.ok) throw new Error("Заказ не найден");
  return (await r.json()) as Order;
}